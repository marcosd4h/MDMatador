package mdm

import (
	"time"
)

///////////////////////////////////////////////////////////////
/// MDMWindowsEnrolledDevice type
/// Contains the information of the enrolled Windows host

type MDMWindowsEnrolledDevice struct {
	ID            string    `db:"device_id"`
	HWID          string    `db:"hardware_id"`
	Name          string    `db:"device_name"`
	Type          string    `db:"device_type"`
	OSLocale      string    `db:"os_locale"`
	OSEdition     string    `db:"os_edition"`
	OSVersion     string    `db:"os_version"`
	ClientVersion string    `db:"client_version"`
	LastSeen      string    `db:"last_seen"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (e MDMWindowsEnrolledDevice) AuthzType() string {
	return "mdm_windows"
}

///////////////////////////////////////////////////////////////
/// MDMCertificate type
/// Contains the information of the server identity certificate

type MDMCertificate struct {
	ID           int       `db:"id"`
	IdentityCert []byte    `db:"identity_cert"`
	IdentityKey  []byte    `db:"identity_key"`
	CreatedAt    time.Time `db:"created_at"`
}

///////////////////////////////////////////////////////////////
/// DeviceSetting type
/// Contains the information of the device settings

type KSDeviceSetting struct {
	DeviceID     string    `db:"device_id"`
	SettingURI   string    `db:"setting_uri"`
	SettingValue string    `db:"setting_value"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func NewDeviceSetting(deviceID string, settingURI string, settingValue string) *KSDeviceSetting {
	return &KSDeviceSetting{
		DeviceID:     deviceID,
		SettingURI:   settingURI,
		SettingValue: settingValue,
	}
}

///////////////////////////////////////////////////////////////
/// PendingDeviceOperation type
/// Contains the information of the pending device operations

type PendingDeviceOperation struct {
	DeviceID     string `json:"device_id" db:"device_id"`
	CmdVerb      string `json:"cmd_verb" db:"cmd_verb"`
	SettingURI   string `json:"setting_uri" db:"setting_uri"`
	SettingValue string `json:"setting_value" db:"setting_value"`
}

///////////////////////////////////////////////////////////////
/// WebSocketCmd type
/// Contains the information of the pending device operations

type WebSocketCmd struct {
	DeviceID string `json:"deviceId"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
}

///////////////////////////////////////////////////////////////
/// TemplateDeviceSetting type
/// Contains the information of the template device settings
/// This is not stored in the database

type TemplateDeviceSetting struct {
	DNSComputerName            string
	DeviceID                   string
	HWDevID                    string
	SMBIOS                     string
	DeviceName                 string
	WindowsEdition             string
	WindowsVersion             string
	OSLocale                   string
	DeviceManufacturer         string
	DeviceModel                string
	Localtime                  string
	FirmwareVersion            string
	HardwareVersion            string
	BIOSVersion                string
	AntivirusStatus            string
	AntivirusSignatureStatus   string
	HVCIStatus                 string
	DeviceGuardStatus          string
	CredentialGuardStatus      string
	SystemGuardStatus          string
	EncryptionComplianceStatus string
	SecureBootStatus           string
	FirewallStatus             string
	CDiskSize                  string
	CDiskFreeSpace             string
	CDiskSystemType            string
	TotalRAM                   string
	ControlFirewall            string
	AVExclusions               string
	AVRTMonitoring             string
	WDAG                       string
	WindowsUpdates             string
	BackgroundImage            string
	StaticContentURL           string
	WindowsVBS                 string
}
