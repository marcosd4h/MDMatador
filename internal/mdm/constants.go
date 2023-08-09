package mdm

// MS-MDE2 and MS-MDM HTTP Endpoints
const (
	// MSMDE2_DiscoveryPath is the HTTP endpoint path that serves the IDiscoveryService functionality.
	// This is the endpoint that process the Discover and DiscoverResponse messages
	// See the section 3.1 on the MS-MDE2 specification for more details:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/2681fd76-1997-4557-8963-cf656ab8d887
	MSMDE2_DiscoveryPath = "/EnrollmentServer/Discovery.svc"

	// MSMDE2_PolicyPath is the HTTP endpoint path that delivers the X.509 Certificate Enrollment Policy (MS-XCEP) functionality.
	// This is the endpoint that process the GetPolicies and GetPoliciesResponse messages
	// See the section 3.3 on the MS-MDE2 specification for more details on this endpoint requirements:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/8a5efdf8-64a9-44fd-ab63-071a26c9f2dc
	// The MS-XCEP specification is available here:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-xcep/08ec4475-32c2-457d-8c27-5a176660a210
	MSMDE2_PolicyPath = "/EnrollmentServer/Policy.svc"

	// MSMDE2_AuthPath is the HTTP endpoint path that delivers the Security Token Servicefunctionality.
	// The MS-MDE2 protocol is agnostic to the token format and value returned by this endpoint.
	// See the section 3.2 on the MS-MDE2 specification for more details:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/27ed8c2c-0140-41ce-b2fa-c3d1a793ab4a
	MSMDE2_AuthPath = "/EnrollmentServer/Auth.svc"

	// MSMDE2_EnrollPath is the HTTP endpoint path that delivers WS-Trust X.509v3 Token Enrollment (MS-WSTEP) functionality.
	// This is the endpoint that process the RequestSecurityToken and RequestSecurityTokenResponseCollection messages
	// See the section 3.4 on the MS-MDE2 specification for more details on this endpoint requirements:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/5b02c625-ced2-4a01-a8e1-da0ae84f5bb7
	// The MS-WSTEP specification is available here:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wstep/4766a85d-0d18-4fa1-a51f-e5cb98b752ea
	MSMDE2_EnrollPath = "/EnrollmentServer/Enrollment.svc"

	// ManagementPath is the HTTP endpoint path that delivers WS-Trust X.509v3 Token Enrollment (MS-WSTEP) functionality.
	// This is the endpoint that process the RequestSecurityToken and RequestSecurityTokenResponseCollection messages
	// See the section 3.4 on the MS-MDE2 specification for more details:
	// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/5b02c625-ced2-4a01-a8e1-da0ae84f5bb7
	MSMDM_ManagementPath = "/ManagementServer/MDM.svc"
)

// Device Enrolled States
const (
	// Device is not yet MDM enrolled
	MDMDeviceStateNotEnrolled = "MDMDeviceEnrolledNotEnrolled"

	// Device is MDM enrolled
	MDMDeviceStateEnrolled = "MDMDeviceEnrolledEnrolled"

	// Device is MDM enrolled and managed
	MDMDeviceStateManaged = "MDMDeviceEnrolledManaged"
)

// MS-MDE2 Config values
const (
	// Minimum supported version
	EnrollmentVersionV4 = "4.0"

	// Maximum supported version
	EnrollmentVersionV5 = "5.0"

	// Minimal Key Length for SHA256WithRSA encryption
	MinKeyLength = "2048"

	// Certificate Validity Period in seconds (365 days)
	CertValidityPeriodInSecs = "31536000"

	// Certificate Renewal Period in seconds (180 days)
	CertRenewalPeriodInSecs = "15552000"

	// Provisioning Doc Certificate Renewal Period (365 days)
	WstepCertRenewalPeriodInDays = "365"

	// Certificate Renewal Period in seconds (180 days)
	PolicyCertRenewalPeriodInSecs = "15552000"

	// Provisioning Doc Server supports ROBO auto certificate renewal
	// TODO: Add renewal support
	WstepROBOSupport = "true"

	// Provisioning Doc Server retry interval
	WstepRenewRetryInterval = "4"

	// The PROVIDER-ID paramer specifies the server identifier for a management server used in the current management session
	C2ProviderID = "MDMatador"

	// The NAME parameter is used in the APPLICATION characteristic to specify a user readable application identity
	DocProvisioningAppName = C2ProviderID

	// The CONNRETRYFREQ parameter is used in the APPLICATION characteristic to specify a user readable application identity
	DocProvisioningAppConnRetryFreq = "6"

	// The INITIALBACKOFFTIME parameter is used to specify the initial wait time in milliseconds when the DM client retries for the first time
	DocProvisioningAppInitialBackoffTime = "30000"

	// The MAXBACKOFFTIME parameter is used to specify the maximum number of milliseconds to sleep after package-sending failure
	DocProvisioningAppMaxBackoffTime = "120000"

	// The DocProvisioningVersion attributes defines the version of the provisioning document format
	DocProvisioningVersion = "1.1"

	// The number of times the DM client should retry to connect to the server when the client is initially configured or enrolled to communicate with the server.
	// If the value is set to 0 and the IntervalForFirstSetOfRetries value isn't 0, then the schedule will be set to repeat an infinite number of times and second set and this set of schedule won't set in this case
	DmClientCSPNumberOfFirstRetries = "0"

	// The waiting time (in minutes) for the initial set of retries as specified by the number of retries in NumberOfFirstRetries
	DmClientCSPIntervalForFirstSetOfRetries = "1"

	// The number of times the DM client should retry a second round of connecting to the server when the client is initially configured/enrolled to communicate with the server
	DmClientCSPNumberOfSecondRetries = "0"

	// The waiting time (in minutes) for the second set of retries as specified by the number of retries in NumberOfSecondRetries
	DmClientCSPIntervalForSecondSetOfRetries = "1"

	// The number of times the DM client should retry connecting to the server when the client is initially configured/enrolled to communicate with the server
	DmClientCSPNumberOfRemainingScheduledRetries = "0"

	// The waiting time (in minutes) for the initial set of retries as specified by the number of retries in NumberOfRemainingScheduledRetries
	DmClientCSPIntervalForRemainingScheduledRetries = "1560"

	// It allows the IT admin to require the device to start a management session on any user login, regardless of if the user has preciously logged in
	DmClientCSPPollOnLogin = "true"

	// It specifies whether the DM client should send out a request pending alert in case the device response to a DM request is too slow.
	DmClientCSPEnableOmaDmKeepAliveMessage = "true"

	// CSR issuer should be verified during enrollment
	EnrollVerifyIssue = true

	// UPN used for both programmatic and user-driven enrollment
	DefaultC2UPN = "infected@demomatador.io"
)

// Soap Error constants
// Details here: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/0a78f419-5fd7-4ddb-bc76-1c0f7e11da23

type SoapError string

const (
	// Message format is bad
	SoapErrorMessageFormat SoapError = "s:messageformat"

	// User not recognized
	SoapErrorAuthentication SoapError = "s:authentication"

	// User not allowed to enroll
	SoapErrorAuthorization SoapError = "s:authorization"

	// Failed to get certificate
	SoapErrorCertificateRequest SoapError = "s:certificaterequest"

	// Generic failure from management server, such as a database access error
	SoapErrorEnrollmentServer SoapError = "s:enrollmentserver"

	// The server hit an unexpected issue
	SoapErrorInternalServiceFault SoapError = "s:internalservicefault"

	// Cannot parse the security header
	SoapErrorInvalidSecurity SoapError = "a:invalidsecurity"
)

// MS-MDE2 Message constants
const (
	// xsi:nil indicates value is not present
	DefaultStateXSI = "true"

	// Supported authentication types
	AuthOnPremise = "OnPremise"
	AuthFederated = "Federated"

	// SOAP Fault codes
	SoapFaultRecv = "s:receiver"

	// SOAP Fault default error locale
	SoapFaultLocale = "en-us"

	// String type used by the DM client configuration
	DmClientStringType = "string"

	// Int type used by the DM client configuration
	DmClientIntType = "integer"

	// Bool type used by the DM client configuration
	DmClientBoolType = "boolean"

	// Supported Enroll Type
	ReqSecTokenEnrollType = "Full"

	// SOAP Message Content Type
	SoapMsgContentType = "application/soap+xml; charset=utf-8"

	// HTTP Content Type for SyncML MDM responses
	SyncMLContentType = "application/vnd.syncml.dm+xml"

	// HTTP Content Type for Webcontainer responses
	WebContainerContentType = "text/html; charset=UTF-8"

	// Additional Context items present on the RequestSecurityToken token message
	ReqSecTokenContextItemUXInitiated          = "UXInitiated"
	ReqSecTokenContextItemHWDevID              = "HWDevID"
	ReqSecTokenContextItemLocale               = "Locale"
	ReqSecTokenContextItemTargetedUserLoggedIn = "TargetedUserLoggedIn"
	ReqSecTokenContextItemOSEdition            = "OSEdition"
	ReqSecTokenContextItemDeviceName           = "DeviceName"
	ReqSecTokenContextItemDeviceID             = "DeviceID"
	ReqSecTokenContextItemEnrollmentType       = "EnrollmentType"
	ReqSecTokenContextItemDeviceType           = "DeviceType"
	ReqSecTokenContextItemOSVersion            = "OSVersion"
	ReqSecTokenContextItemApplicationVersion   = "ApplicationVersion"
	ReqSecTokenContextItemNotInOobe            = "NotInOobe"
	ReqSecTokenContextItemRequestVersion       = "RequestVersion"

	// APPRU query param expected by STS Auth endpoint
	STSAuthAppRu = "appru"

	// Login related query param expected by STS Auth endpoint
	STSLoginHint = "login_hint"
)

// XML Namespaces used by the Microsoft Device Enrollment v2 protocol (MS-MDE2)
const (
	DiscoverNS          = "http://schemas.microsoft.com/windows/management/2012/01/enrollment"
	PolicyNS            = "http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy"
	EnrollWSTrust       = "http://docs.oasis-open.org/ws-sx/ws-trust/200512"
	EnrollSecExt        = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	EnrollTokenType     = "http://schemas.microsoft.com/5.0.0.0/ConfigurationManager/Enrollment/DeviceEnrollmentToken"
	EnrollPDoc          = "http://schemas.microsoft.com/5.0.0.0/ConfigurationManager/Enrollment/DeviceEnrollmentProvisionDoc"
	EnrollEncode        = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd#base64binary"
	EnrollReq           = "http://schemas.microsoft.com/windows/pki/2009/01/enrollment"
	EnrollNSS           = "http://www.w3.org/2003/05/soap-envelope"
	EnrollNSA           = "http://www.w3.org/2005/08/addressing"
	EnrollXSI           = "http://www.w3.org/2001/XMLSchema-instance"
	EnrollXSD           = "http://www.w3.org/2001/XMLSchema"
	EnrollXSU           = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	ActionNsDiag        = "http://schemas.microsoft.com/2004/09/ServiceModel/Diagnostics"
	ActionNsDiscovery   = "http://schemas.microsoft.com/windows/management/2012/01/enrollment/IDiscoveryService/DiscoverResponse"
	ActionNsPolicy      = "http://schemas.microsoft.com/windows/pki/2009/01/enrollmentpolicy/IPolicy/GetPoliciesResponse"
	ActionNsEnroll      = EnrollReq + "/RSTRC/wstep"
	EnrollReqTypePKCS10 = EnrollReq + "#PKCS10"
	EnrollReqTypePKCS7  = EnrollReq + "#PKCS7"
)

// MS-MDM Message constants
const (
	// SyncML Message Content Type
	SyncMLMsgContentType = "application/vnd.syncml.dm+xml"

	// SyncML Message Meta Namespace
	SyncMLMetaNamespace = "syncml:metinf"

	// SyncML Cmd Namespace
	SyncCmdNamespace = "SYNCML:SYNCML1.2"

	// SyncML Message Header Name
	SyncMLHdrName = "SyncHdr"

	// Min Supported SyncML version
	SyncMLMinSupportedVersion = "1.2"

	// Max Supported SyncML version
	SyncMLMaxSupportedVersion = "1.2"

	// SyncML ver protocol version
	SyncMLVerProto = "DM/" + SyncMLMinSupportedVersion

	// Wait time for remote shell sessions in seconds
	ShellSessionWaitTime = 70
)

// MS-MDM Status Code constants
// Details here: https://learn.microsoft.com/en-us/windows/client-management/oma-dm-protocol-support

const (
	// The SyncML command completed successfully
	CmdStatusCode200 = "200"

	// 	Accepted for processing
	// This code denotes an asynchronous operation, such as a request to run a remote execution of an application
	CmdStatusCode202 = "202"

	// Authentication accepted
	// Normally you'll only see this code in response to the SyncHdr element (used for authentication in the OMA-DM standard)
	// You may see this code if you look at OMA DM logs, but CSPs don't typically generate this code.
	CmdStatusCode212 = "212"

	// Operation canceled
	// The SyncML command completed successfully, but no more commands will be processed within the session.
	CmdStatusCode214 = "214"

	// Not executed
	// A command wasn't executed as a result of user interaction to cancel the command.
	CmdStatusCode215 = "215"

	// Atomic roll back OK
	// A command was inside an Atomic element and Atomic failed, thhis command was rolled back successfully
	CCmdStatusCode216 = "216"

	// Bad request. The requested command couldn't be performed because of malformed syntax.
	// CSPs don't usually generate this error, however you might see it if your SyncML is malformed.
	CmdStatusCode400 = "400"

	// 	Invalid credentials
	// The requested command failed because the requestor must provide proper authentication. CSPs don't usually generate this error
	CmdStatusCode401 = "401"

	// Forbidden
	// The requested command failed, but the recipient understood the requested command
	CmdStatusCode403 = "403"

	// Not found
	// The requested target wasn't found. This code will be generated if you query a node that doesn't exist
	CmdStatusCode404 = "404"

	// Command not allowed
	// This respond code will be generated if you try to write to a read-only node
	CmdStatusCode405 = "405"

	// Optional feature not supported
	// This response code will be generated if you try to access a property that the CSP doesn't support
	CmdStatusCode406 = "406"

	// Unsupported type or format
	// This response code can result from XML parsing or formatting errors
	CmdStatusCode415 = "415"

	// Already exists
	// This response code occurs if you attempt to add a node that already exists
	CmdStatusCode418 = "418"

	// Permission Denied
	// The requested command failed because the sender doesn't have adequate access control permissions (ACL) on the recipient.
	// An "Access denied" errors usually get translated to this response code.
	CmdStatusCode425 = "425"

	// Command failed. Generic failure.
	// The recipient encountered an unexpected condition, which prevented it from fulfilling the request
	// This response code will occur when the SyncML DPU can't map the originating error code
	CmdStatusCode500 = "500"

	// Atomic failed
	// One of the operations in an Atomic block failed
	CmdStatusCode507 = "507"

	// Atomic roll back failed
	// An Atomic operation failed and the command wasn't rolled back successfully.
	CmdStatusCode516 = "516"
)

// MS-MDM Supported Alerts
// Details on MS-MDM 2.2.7.2: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mdm/72c6ea01-121c-48f9-85da-a26bb12aad51

const (
	// SERVER-INITIATED MGMT
	// Server-initiated device management session
	CmdAlert1200 = "1200"

	// CLIENT-INITIATED MGMT
	// Client-initiated device management session
	CmdAlert1201 = "1201"

	// NEXT MESSAGE
	// Request for the next message of a large object package
	CmdAlert1222 = "1222"

	// SESSION ABORT
	// Informs recipient that the sender wishes to abort the DM session
	CmdAlert1223 = "1223"

	// CLIENT EVENT
	// Informs server that an event has occurred on the client
	CmdAlert1224 = "1224"

	// NO END OF DATA
	// End of Data for chunked object not received.
	CmdAlert1225 = "1225"

	// GENERIC ALERT
	// Generic client generated alert with or without a reference to a Management
	CmdAlert1226 = "1226"
)

// GetFriendlyOSEdition returns the friendly name for the Windows OS edition
func GetFriendlyOSEdition(id string) string {
	var cmdhash = map[string]string{
		"0":   "UNDEFINED",
		"1":   "ULTIMATE",
		"2":   "HOME_BASIC",
		"3":   "HOME_PREMIUM",
		"4":   "ENTERPRISE",
		"5":   "HOME_BASIC_N",
		"6":   "BUSINESS",
		"7":   "STANDARD_SERVER",
		"8":   "DATACENTER_SERVER",
		"9":   "SMALLBUSINESS_SERVER",
		"10":  "ENTERPRISE_SERVER",
		"11":  "STARTER",
		"12":  "DATACENTER_SERVER_CORE",
		"13":  "STANDARD_SERVER_CORE",
		"14":  "ENTERPRISE_SERVER_CORE",
		"15":  "ENTERPRISE_SERVER_IA64",
		"16":  "BUSINESS_N",
		"17":  "WEB_SERVER",
		"18":  "CLUSTER_SERVER",
		"19":  "HOME_SERVER",
		"20":  "STORAGE_EXPRESS_SERVER",
		"21":  "STORAGE_STANDARD_SERVER",
		"22":  "STORAGE_WORKGROUP_SERVER",
		"23":  "STORAGE_ENTERPRISE_SERVER",
		"24":  "SERVER_FOR_SMALLBUSINESS",
		"25":  "SMALLBUSINESS_SERVER_PREMIUM",
		"26":  "HOME_PREMIUM_N",
		"27":  "ENTERPRISE_N",
		"28":  "ULTIMATE_N",
		"29":  "WEB_SERVER_CORE",
		"30":  "MEDIUMBUSINESS_SERVER_MANAGEMENT",
		"31":  "MEDIUMBUSINESS_SERVER_SECURITY",
		"32":  "MEDIUMBUSINESS_SERVER_MESSAGING",
		"33":  "SERVER_FOUNDATION",
		"34":  "HOME_PREMIUM_SERVER",
		"35":  "SERVER_FOR_SMALLBUSINESS_V",
		"36":  "STANDARD_SERVER_V",
		"37":  "DATACENTER_SERVER_V",
		"38":  "ENTERPRISE_SERVER_V",
		"39":  "DATACENTER_SERVER_CORE_V",
		"40":  "STANDARD_SERVER_CORE_V",
		"41":  "ENTERPRISE_SERVER_CORE_V",
		"42":  "HYPERV",
		"43":  "STORAGE_EXPRESS_SERVER_CORE",
		"44":  "STORAGE_STANDARD_SERVER_CORE",
		"45":  "STORAGE_WORKGROUP_SERVER_CORE",
		"46":  "STORAGE_ENTERPRISE_SERVER_CORE",
		"47":  "STARTER_N",
		"48":  "PROFESSIONAL",
		"49":  "PROFESSIONAL_N",
		"50":  "SB_SOLUTION_SERVER",
		"51":  "SERVER_FOR_SB_SOLUTIONS",
		"52":  "STANDARD_SERVER_SOLUTIONS",
		"53":  "STANDARD_SERVER_SOLUTIONS_CORE",
		"54":  "SB_SOLUTION_SERVER_EM",
		"55":  "SERVER_FOR_SB_SOLUTIONS_EM",
		"56":  "SOLUTION_EMBEDDEDSERVER",
		"57":  "SOLUTION_EMBEDDEDSERVER_CORE",
		"59":  "ESSENTIALBUSINESS_SERVER_MGMT",
		"60":  "ESSENTIALBUSINESS_SERVER_ADDL",
		"61":  "ESSENTIALBUSINESS_SERVER_MGMTSVC",
		"62":  "ESSENTIALBUSINESS_SERVER_ADDLSVC",
		"63":  "SMALLBUSINESS_SERVER_PREMIUM_CORE",
		"64":  "CLUSTER_SERVER_V",
		"65":  "EMBEDDED",
		"66":  "STARTER_E",
		"67":  "HOME_BASIC_E",
		"68":  "HOME_PREMIUM_E",
		"69":  "PROFESSIONAL_E",
		"70":  "ENTERPRISE_E",
		"71":  "ULTIMATE_E",
		"72":  "ENTERPRISE_EVALUATION",
		"76":  "MULTIPOINT_STANDARD_SERVER",
		"77":  "MULTIPOINT_PREMIUM_SERVER",
		"79":  "STANDARD_EVALUATION_SERVER",
		"80":  "DATACENTER_EVALUATION_SERVER",
		"84":  "ENTERPRISE_N_EVALUATION",
		"85":  "EMBEDDED_AUTOMOTIVE",
		"86":  "EMBEDDED_INDUSTRY_A",
		"87":  "THINPC",
		"88":  "EMBEDDED_A",
		"89":  "EMBEDDED_INDUSTRY",
		"90":  "EMBEDDED_E",
		"91":  "EMBEDDED_INDUSTRY_E",
		"92":  "EMBEDDED_INDUSTRY_A_E",
		"95":  "STORAGE_WORKGROUP_EVALUATION_SERVER",
		"96":  "STORAGE_STANDARD_EVALUATION_SERVER",
		"97":  "CORE_ARM",
		"98":  "CORE_N",
		"99":  "CORE_COUNTRYSPECIFIC",
		"100": "CORE_SINGLELANGUAGE",
		"101": "CORE",
		"103": "PROFESSIONAL_WMC",
		"104": "MOBILE_CORE",
		"105": "EMBEDDED_INDUSTRY_EVAL",
		"106": "EMBEDDED_INDUSTRY_E_EVAL",
		"107": "EMBEDDED_EVAL",
		"108": "EMBEDDED_E_EVAL",
		"109": "NANO_SERVER",
		"110": "CLOUD_STORAGE_SERVER",
		"111": "CORE_CONNECTED",
		"112": "PROFESSIONAL_STUDENT",
		"113": "CORE_CONNECTED_N",
		"114": "PROFESSIONAL_STUDENT_N",
		"115": "CORE_CONNECTED_SINGLELANGUAGE",
		"116": "CORE_CONNECTED_COUNTRYSPECIFIC",
		"117": "CONNECTED_CAR",
		"118": "INDUSTRY_HANDHELD",
		"119": "PPI_PRO",
		"120": "ARM64_SERVER",
		"121": "EDUCATION",
		"122": "EDUCATION_N",
		"123": "IOTUAP",
		"124": "CLOUD_HOST_INFRASTRUCTURE_SERVER",
		"125": "ENTERPRISE_S",
		"126": "ENTERPRISE_S_N",
		"127": "PROFESSIONAL_S",
		"128": "PROFESSIONAL_S_N",
		"129": "ENTERPRISE_S_EVALUATION",
		"130": "ENTERPRISE_S_N_EVALUATION",
		"135": "HOLOGRAPHIC",
		"136": "HOLOGRAPHIC_BUSINESS",
		"175": "SERVERRDSH",
	}

	value, exists := cmdhash[id]
	if !exists {
		return "UnknownVersion"
	}

	return value
}

// Supported CSP properties
const (
	CSPDefaultValue                = "Not Present"
	CSPAlert1201                   = "./Alert/1201"
	CSPDeviceID                    = "./DevInfo/DevId"
	CSPHWDevID                     = "./Vendor/MSFT/DMClient/HWDevID"
	CSPSMBIOS                      = "./DevDetail/Ext/Microsoft/SMBIOSSerialNumber"
	CSPDeviceName                  = "./DevDetail/Ext/Microsoft/DeviceName"
	CSPDNSComputerName             = "./DevDetail/Ext/Microsoft/DNSComputerName"
	CSPWindowsEdition              = "./DevDetail/Ext/Microsoft/OSPlatform"
	CSPWindowsVersion              = "./DevDetail/SwV"
	CSPOSLocale                    = "./DevInfo/Lang"
	CSPDeviceManufacturer          = "./DevInfo/Man"
	CSPDeviceModel                 = "./DevInfo/Mod"
	CSPLocaltime                   = "./DevDetail/Ext/Microsoft/LocalTime"
	CSPFirmwareVersion             = "./DevDetail/FwV"
	CSPHardwareVersion             = "./DevDetail/HwV"
	CSPBIOSVersion                 = "./DevDetail/Ext/Microsoft/SMBIOSVersion"
	CSPAntivirusStatus             = "./Vendor/MSFT/DeviceStatus/Antivirus/Status"
	CSPAntivirusSignatureStatus    = "./Vendor/MSFT/DeviceStatus/Antivirus/SignatureStatus"
	CSPHVCIStatus                  = "./Vendor/MSFT/DeviceStatus/DeviceGuard/HypervisorEnforcedCodeIntegrityStatus"
	CSPDeviceGuardStatus           = "./Vendor/MSFT/DeviceStatus/DeviceGuard/VirtualizationBasedSecurityStatus"
	CSPCredentialGuardStatus       = "./Vendor/MSFT/DeviceStatus/DeviceGuard/LsaCfgCredGuardStatus"
	CSPSystemGuardStatus           = "./Vendor/MSFT/DeviceStatus/DeviceGuard/SystemGuardStatus"
	CSPEncryptionComplianceStatus  = "./Vendor/MSFT/DeviceStatus/Compliance/EncryptionCompliance"
	CSPSecureBootStatus            = "./Vendor/MSFT/DeviceStatus/SecureBootState"
	CSPFirewallStatus              = "./Vendor/MSFT/DeviceStatus/Firewall/Status"
	CSPCDiskSize                   = "./cimV2/Win32_LogicalDisk/Win32_LogicalDisk.DeviceID='C:'/Size"
	CSPCDiskFreeSpace              = "./cimv2/Win32_LogicalDisk/Win32_LogicalDisk.DeviceID='C:'/FreeSpace"
	CSPCDiskSystemType             = "./cimV2/Win32_LogicalDisk/Win32_LogicalDisk.DeviceID='C:'/FileSystem"
	CSPTotalRAM                    = "./cimV2/Win32_PhysicalMemory/Win32_PhysicalMemory.Tag='Physical%20Memory%200'/Capacity"
	CSPDomainProfileFirewall       = "./Vendor/MSFT/Firewall/MdmStore/DomainProfile/EnableFirewall"
	CSPPrivateProfileFirewall      = "./Vendor/MSFT/Firewall/MdmStore/PrivateProfile/EnableFirewall"
	CSPPublicProfileFirewall       = "./Vendor/MSFT/Firewall/MdmStore/PublicProfile/EnableFirewall"
	CSPControlFirewall             = "./Vendor/MSFT/Firewall/CSPControlFirewall"
	CSPPolicyDefenderExcludedPaths = "./Device/Vendor/MSFT/Policy/Config/Defender/ExcludedPaths"
	CSPPolicyDefenderAV            = "./Device/Vendor/MSFT/Policy/Config/Defender/AllowRealtimeMonitoring"
	CSPWDAGAllowSetting            = "./Device/Vendor/MSFT/WindowsDefenderApplicationGuard/Settings/AllowWindowsDefenderApplicationGuard"
	CSPWindowsUpdates              = "./Device/Vendor/MSFT/Policy/Config/Update/AllowAutoUpdate"
	CSPPersonalizationDesktopURL   = "./Vendor/MSFT/Personalization/DesktopImageUrl"
	CSPPolicyWindowsVBS            = "./Device/Vendor/MSFT/Policy/Config/DeviceGuard/EnableVirtualizationBasedSecurity"
	CSPC2runchCSP                  = "./Device/Vendor/OEM/CrunchCSP/Shell"
	CSPC2runchCSPWait              = "./Device/Vendor/OEM/CrunchCSP/Wait"
)
