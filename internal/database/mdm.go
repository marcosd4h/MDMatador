package database

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"errors"
	"time"

	"github.com/marcosd4h/MDMatador/internal/common"
	"github.com/marcosd4h/MDMatador/internal/mdm"
)

/////////////////////////////////////////////////////////////////
/// mdm_enrollments table
/// This table is used to store the enrolled devices

func (db *DB) MDMInsertEnrolledDevice(device *mdm.MDMWindowsEnrolledDevice) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_enrollments(device_id, hardware_id, device_name, device_type, os_locale, os_edition, os_version, client_version, last_seen, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := db.ExecContext(ctx, query, device.ID, device.HWID, device.Name, device.Type, device.OSLocale, device.OSEdition, device.OSVersion, device.ClientVersion, time.Now(), time.Now(), time.Now())
	return err
}

func (db *DB) MDMGetEnrolledDevice(mdmDeviceID string) (*mdm.MDMWindowsEnrolledDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var enrolledDevice mdm.MDMWindowsEnrolledDevice

	query := `SELECT * FROM mdm_enrollments WHERE device_id = $1`

	err := db.GetContext(ctx, &enrolledDevice, query, mdmDeviceID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &enrolledDevice, err
}

func (db *DB) MDMGetEnrolledDeviceByHWID(mdmDeviceHWID string) (*mdm.MDMWindowsEnrolledDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var enrolledDevice mdm.MDMWindowsEnrolledDevice

	query := `SELECT * FROM mdm_enrollments WHERE hardware_id = $1`

	err := db.GetContext(ctx, &enrolledDevice, query, mdmDeviceHWID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &enrolledDevice, err
}

func (db *DB) MDMGetEnrolledDevices() ([]*mdm.MDMWindowsEnrolledDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var enrolledDevices []*mdm.MDMWindowsEnrolledDevice

	query := `SELECT * FROM mdm_enrollments`

	err := db.SelectContext(ctx, &enrolledDevices, query)
	if errors.Is(err, sql.ErrNoRows) {
		return []*mdm.MDMWindowsEnrolledDevice{}, nil
	}

	return enrolledDevices, err
}

func (db *DB) MDMDeleteEnrolledDevice(deviceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_enrollments WHERE device_id = $1`

	_, err := db.ExecContext(ctx, query, deviceID)
	return err
}

func (db *DB) MDMDeleteEnrolledDeviceByHWID(deviceHWID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_enrollments WHERE hardware_id = $1`

	_, err := db.ExecContext(ctx, query, deviceHWID)
	return err
}

func (db *DB) MDMIsValidDeviceID(deviceID string) bool {

	// Checking the storage to see if the device is already enrolled
	device, err := db.MDMGetEnrolledDevice(deviceID)
	if err != nil || device == nil {
		return false
	}

	return true
}

func (db *DB) MDMUpdateLastSeen(deviceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		UPDATE mdm_enrollments 
		SET last_seen = $1 
		WHERE device_id = $3
	`
	_, err := db.ExecContext(ctx, query, time.Now().Format(time.RFC822), deviceID)
	return err
}

/////////////////////////////////////////////////////////////////
/// mdm_certificates table
/// This table is used to store the server identity certificate

func (db *DB) InsertMdmCertificate(cert *mdm.MDMCertificate) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_certificates(identity_cert, identity_key, created_at)
		VALUES (?, ?, ?)`

	_, err := db.ExecContext(ctx, query, cert.IdentityCert, cert.IdentityKey, time.Now())
	return err
}

func (db *DB) GetMDMCertificate(id int) (*mdm.MDMCertificate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var cert mdm.MDMCertificate

	query := `
		SELECT * FROM mdm_certificates
		WHERE id = ?`

	err := db.GetContext(ctx, &cert, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &cert, err
}

func (db *DB) DeleteMDMCertificate(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_certificates WHERE id = ?`

	_, err := db.ExecContext(ctx, query, id)
	return err
}

func (db *DB) GetIdentityCert() (*x509.Certificate, *rsa.PrivateKey, error) {
	cert, err := db.GetMDMCertificate(1) // Try to get the first certificate
	if err != nil {
		return nil, nil, err
	}

	// If the certificate exists, we just return it
	if cert != nil {
		certificate, privateKey, err := common.ParseX509Keypair(cert.IdentityCert, cert.IdentityKey)
		if err != nil {
			return nil, nil, err
		}

		return certificate, privateKey, nil
	}

	// No certificate in the database, so a new one is generated
	certDataContainer, err := common.GetNewMDMCertificate()
	if err != nil {
		return nil, nil, err
	}

	// The new certificate is stored in the database
	newCert := &mdm.MDMCertificate{
		IdentityCert: certDataContainer.CertData,
		IdentityKey:  x509.MarshalPKCS1PrivateKey(certDataContainer.PrivateKey),
		CreatedAt:    time.Now(),
	}

	if err := db.InsertMdmCertificate(newCert); err != nil {
		return nil, nil, err
	}

	return certDataContainer.Certificate, certDataContainer.PrivateKey, nil
}

/////////////////////////////////////////////////////////////////
/// mdm_device_settings table
/// This table is used to store the device settings

func (db *DB) InsertDeviceSetting(setting *mdm.KSDeviceSetting) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_device_settings(device_id, setting_uri, setting_value, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.ExecContext(ctx, query, setting.DeviceID, setting.SettingURI, setting.SettingValue, time.Now(), time.Now())
	return err
}

func (db *DB) GetDeviceSettings(deviceID string) ([]*mdm.KSDeviceSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var settings []*mdm.KSDeviceSetting

	query := `SELECT * FROM mdm_device_settings WHERE device_id = $1`

	err := db.SelectContext(ctx, &settings, query, deviceID)
	if errors.Is(err, sql.ErrNoRows) {
		return []*mdm.KSDeviceSetting{}, nil
	}

	return settings, err
}

func (db *DB) DeleteDeviceSetting(deviceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_device_settings WHERE device_id = $1`

	_, err := db.ExecContext(ctx, query, deviceID)
	return err
}

func (db *DB) DeleteCustomDeviceSetting(deviceID string, settingURI string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_device_settings WHERE device_id = $1 AND setting_uri = $2`

	_, err := db.ExecContext(ctx, query, deviceID, settingURI)
	return err
}

func (db *DB) GetCustomDeviceSetting(deviceID string, settingURI string) (*mdm.KSDeviceSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
        SELECT device_id, setting_uri, setting_value, created_at, updated_at
        FROM mdm_device_settings
        WHERE device_id = $1 AND setting_uri = $2`

	setting := &mdm.KSDeviceSetting{}
	err := db.GetContext(ctx, setting, query, deviceID, settingURI)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return setting, err
}

func (db *DB) UpdateDeviceSettingWithType(setting *mdm.KSDeviceSetting) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_device_settings(device_id, setting_uri, setting_value, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (device_id, setting_uri)
		DO UPDATE SET setting_value = excluded.setting_value, updated_at = excluded.updated_at`

	_, err := db.ExecContext(ctx, query, setting.DeviceID, setting.SettingURI, setting.SettingValue, time.Now(), time.Now())
	return err
}

func (db *DB) UpdateDeviceSetting(deviceID string, settingURI string, settingValue string) error {
	deviceSetting := mdm.NewDeviceSetting(deviceID, settingURI, settingValue)
	return db.UpdateDeviceSettingWithType(deviceSetting)
}

/////////////////////////////////////////////////////////////////
/// mdm_pending_device_operations table
/// This table is used to store the pending device operations

func (db *DB) InsertPendingDeviceOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_pending_device_operations(device_id, cmd_verb, setting_uri, setting_value)
		VALUES ($1, $2, $3, $4)`

	_, err := db.ExecContext(ctx, query, deviceID, cmdVerb, settingURI, settingValue)
	return err
}

func (db *DB) GetPendingDeviceOperations(deviceID string) ([]*mdm.PendingDeviceOperation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var operations []*mdm.PendingDeviceOperation

	query := `SELECT * FROM mdm_pending_device_operations WHERE device_id = $1`

	err := db.SelectContext(ctx, &operations, query, deviceID)
	if errors.Is(err, sql.ErrNoRows) {
		return []*mdm.PendingDeviceOperation{}, nil
	}

	return operations, err
}

func (db *DB) DeletePendingDeviceOperations(deviceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_pending_device_operations WHERE device_id = $1`

	_, err := db.ExecContext(ctx, query, deviceID)
	return err
}

func (db *DB) DeletePendingDeviceOperation(deviceID string, cmdVerb string, settingURI string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `DELETE FROM mdm_pending_device_operations WHERE device_id = $1 AND cmd_verb = $2 AND setting_uri = $3`

	_, err := db.ExecContext(ctx, query, deviceID, cmdVerb, settingURI)
	return err
}

func (db *DB) QueuePendingDeviceOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO mdm_pending_device_operations(device_id, cmd_verb, setting_uri, setting_value)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (device_id, cmd_verb, setting_uri, setting_value)
		DO UPDATE SET device_id = excluded.device_id, cmd_verb = excluded.cmd_verb, setting_uri = excluded.setting_uri, setting_value = excluded.setting_value`

	_, err := db.ExecContext(ctx, query, deviceID, cmdVerb, settingURI, settingValue)

	return err
}

func (db *DB) GetAllPendingOperations(deviceID string) ([]*mdm.PendingDeviceOperation, error) {

	operations, err := db.GetPendingDeviceOperations(deviceID)
	if err != nil {
		return nil, err
	}

	if (operations != nil) && (len(operations) != 0) {
		err = db.DeletePendingDeviceOperations(deviceID)
		if err != nil {
			return nil, err
		}
	}

	return operations, nil
}

func (db *DB) GetCustomPendingOperations(deviceID string, targetOp string) (*mdm.PendingDeviceOperation, error) {

	operations, err := db.GetPendingDeviceOperations(deviceID)
	if err != nil {
		return nil, err
	}

	if (operations != nil) && (len(operations) != 0) {
		for _, operation := range operations {
			if operation.SettingURI == targetOp {
				return operation, nil
			}
		}
	}

	return nil, nil
}

// QueueProtoCmdOperation will queue an asynchronous MDM command to the device and return immediately
// The command is queued and will be sent to the device when it connects to the server
func (db *DB) QueueProtoCmdOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error {

	err := db.QueuePendingDeviceOperation(deviceID, cmdVerb, settingURI, settingValue)
	if err != nil {
		return err
	}

	return nil
}
