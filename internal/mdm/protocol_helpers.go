package mdm

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"net/url"
	"path"
	"strconv"
	"text/template"
	"time"

	"github.com/google/uuid"
)

// getUtcTime returns the current timestamp plus the specified number of minutes,
// formatted as "2006-01-02T15:04:05.000Z".
func getUtcTime(minutes int) string {
	// Get the current time and then add the specified number of minutes
	now := time.Now()
	future := now.Add(time.Duration(minutes) * time.Minute)

	// Format and return the future time as a string
	return future.UTC().Format("2006-01-02T15:04:05.000Z")
}

// resolveURL resolves a relative path to a server URL (typically the Fleet
// server's). If cleanQuery is true, the query string part is cleared.
func resolveURL(serverURL, relPath string, cleanQuery bool) (string, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, relPath)
	if cleanQuery {
		u.RawQuery = ""
	}
	return u.String(), nil
}

func ResolveWindowsMDMDiscovery(serverURL string) (string, error) {
	return resolveURL(serverURL, MSMDE2_DiscoveryPath, false)
}

func ResolveWindowsMDMPolicy(serverURL string) (string, error) {
	return resolveURL(serverURL, MSMDE2_PolicyPath, false)
}

func ResolveWindowsMDMEnroll(serverURL string) (string, error) {
	return resolveURL(serverURL, MSMDE2_EnrollPath, false)
}

func ResolveWindowsMDMAuth(serverURL string) (string, error) {
	return resolveURL(serverURL, MSMDE2_AuthPath, false)
}

func ResolveWindowsMDMManagement(serverURL string) (string, error) {
	return resolveURL(serverURL, MSMDM_ManagementPath, false)
}

// NewDiscoverResponse creates a new DiscoverResponse struct based on the auth policy, policy url, and enrollment url
// DiscoverResponse message contains the Uniform Resource Locators (URLs) of service endpoints required for the following steps
func NewDiscoverResponse(authPolicy string, policyUrl string, enrollmentUrl string, authUrl string) (*DiscoverResponse, error) {

	if (len(authPolicy) == 0) || (len(policyUrl) == 0) || (len(enrollmentUrl) == 0) {
		return nil, errors.New("invalid parameters")
	}

	if len(authUrl) == 0 {
		return &DiscoverResponse{
			XMLNS: DiscoverNS,
			DiscoverResult: DiscoverResult{
				AuthPolicy:                 authPolicy,
				EnrollmentVersion:          EnrollmentVersionV4,
				EnrollmentPolicyServiceUrl: policyUrl,
				EnrollmentServiceUrl:       enrollmentUrl,
			},
		}, nil
	}

	return &DiscoverResponse{
		XMLNS: DiscoverNS,
		DiscoverResult: DiscoverResult{
			AuthPolicy:                 authPolicy,
			EnrollmentVersion:          EnrollmentVersionV4,
			EnrollmentPolicyServiceUrl: policyUrl,
			EnrollmentServiceUrl:       enrollmentUrl,
			EnrollmentAuthServiceUrl:   &authUrl,
		},
	}, nil
}

// NewGetPoliciesResponse creates a new GetPoliciesResponse struct based on the minimal key length, certificate validity period, and renewal period
func NewGetPoliciesResponse(minimalKeyLength string, certificateValidityPeriodSeconds string, renewalPeriodSeconds string) (*GetPoliciesResponse, error) {

	if (len(minimalKeyLength) == 0) || (len(certificateValidityPeriodSeconds) == 0) || (len(renewalPeriodSeconds) == 0) {
		return nil, errors.New("invalid parameters")
	}

	return &GetPoliciesResponse{
		XMLNS: PolicyNS,
		Response: Response{
			PolicyFriendlyName: ContentAttr{
				Xsi:   DefaultStateXSI,
				XMLNS: EnrollXSI,
			},
			NextUpdateHours: ContentAttr{
				Xsi:   DefaultStateXSI,
				XMLNS: EnrollXSI,
			},
			PoliciesNotChanged: ContentAttr{
				Xsi:   DefaultStateXSI,
				XMLNS: EnrollXSI,
			},
			Policies: Policies{
				Policy: Policy{
					PolicyOIDReference: "0",
					CAs: GenericAttr{
						Xsi: DefaultStateXSI,
					},

					// These are MS-XCEP Attributes defined in section 3.1.4.1.3.1
					// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-xcep/cd22d3a0-f469-4a44-95ed-d10ce4dc2063
					Attributes: Attributes{
						CommonName:                "MDMatadorAttributes",
						PolicySchema:              "3",
						HashAlgorithmOIDReference: "0",
						Revision: Revision{
							MajorRevision: "101",
							MinorRevision: "0",
						},
						CertificateValidity: CertificateValidity{
							ValidityPeriodSeconds: certificateValidityPeriodSeconds,
							RenewalPeriodSeconds:  renewalPeriodSeconds,
						},
						Permission: Permission{
							Enroll:     "true",
							AutoEnroll: "false",
						},
						SupersededPolicies: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						PrivateKeyFlags: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						SubjectNameFlags: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						EnrollmentFlags: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						GeneralFlags: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						RARequirements: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						KeyArchivalAttributes: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						Extensions: GenericAttr{
							Xsi: DefaultStateXSI,
						},
						PrivateKeyAttributes: PrivateKeyAttributes{
							MinimalKeyLength: minimalKeyLength,
							KeySpec: GenericAttr{
								Xsi: DefaultStateXSI,
							},
							KeyUsageProperty: GenericAttr{
								Xsi: DefaultStateXSI,
							},
							Permissions: GenericAttr{
								Xsi: DefaultStateXSI,
							},
							AlgorithmOIDReference: GenericAttr{
								Xsi:     DefaultStateXSI,
								Content: "1",
							},
							CryptoProviders: []ProviderAttr{
								{Content: "Microsoft Platform Crypto Provider"},
								{Content: "Microsoft Software Key Storage Provider"},
							},
						},
					},
				},
			},
		},

		// These are MS-XCEP OIDs defined in section 3.1.4.1.3.16
		// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-xcep/161aab9f-d159-4df3-85c9-f732ed2a8445
		OIDs: OIDs{
			OID: []OID{
				{
					// SHA256WithRSA OID
					// https://oidref.com/2.16.840.1.101.3.4.2.1
					Value:          "2.16.840.1.101.3.4.2.1",
					Group:          "4",
					OIDReferenceID: "0",
					DefaultName:    "szOID_NIST_sha256",
				},
				{
					// RSA OID
					// https://oidref.com/1.2.840.113549.1.1.1
					Value:          "1.2.840.113549.1.1.1",
					Group:          "3",
					OIDReferenceID: "1",
					DefaultName:    "szOID_RSA_RSA",
				},
			},
		},
	}, nil
}

// NewRequestSecurityTokenResponseCollection creates a new RequestSecurityTokenResponseCollection struct based on the provisioned token
func NewRequestSecurityTokenResponseCollection(provisionedToken string) (*RequestSecurityTokenResponseCollection, error) {

	if len(provisionedToken) == 0 {
		return nil, errors.New("invalid parameters")
	}

	enrollSecExtVal := EnrollSecExt
	return &RequestSecurityTokenResponseCollection{
		XMLNS: EnrollWSTrust,
		RequestSecurityTokenResponse: RequestSecurityTokenResponse{
			TokenType: EnrollTokenType,
			DispositionMessage: SecAttr{
				Content: "",
				XmlNS:   EnrollReq,
			},
			RequestID: SecAttr{
				Content: "0",
				XmlNS:   EnrollReq,
			},
			RequestedSecurityToken: RequestedSecurityToken{
				BinarySecurityToken: BinarySecurityToken{
					Content:      provisionedToken,
					XMLNS:        &enrollSecExtVal,
					ValueType:    EnrollPDoc,
					EncodingType: EnrollEncode,
				},
			},
		},
	}, nil
}

// NewSoapFault creates a new SoapFault struct based on the error type, original message type, and error message
func NewSoapFault(errorType SoapError, origMessage MDEMessageType, errorMessage error) *SoapFault {
	return &SoapFault{
		OriginalMessageType: origMessage,
		Code: Code{
			Value: SoapFaultRecv,
			Subcode: Subcode{
				Value: string(errorType),
			},
		},
		Reason: Reason{
			Text: ReasonText{
				Content: errorMessage.Error(),
				Lang:    SoapFaultLocale,
			},
		},
	}
}

// GetAuthSTSResponse returns the HTML content that will be rendered in the webview2 container during user-driven enrollment
func GetAuthSTSResponse(appru string) ([]byte, error) {

	// This could be used to pass an JWT auth token to authenticate the policy and enrollment endpoints
	encodedBST := "dummy"

	// STS Auth Endpoint returns HTML content that gets render in a webview container
	// The webview container expect a POST request to the appru URL with the wresult parameter set to the auth token
	// The security token in wresult is later passed back in <wsse:BinarySecurityToken>
	// This string is opaque to the enrollment client; the client does not interpret the string.
	// The returned HTML content contains a JS script that will perform a POST request to the appru URL automatically
	// This will set the wresult parameter to the value of auth token
	tmpl, err := template.New("").Parse(`
				<h3>MDM Federated Login</h3>
								
				<script>
				function performPost() {
				  var form = document.createElement('form');
				  form.method = 'POST';
				  form.action = "{{.ActionURL}}"

				  // Add any form fields or data you want to send
				  var input1 = document.createElement('input');
				  input1.type = 'hidden';
				  input1.name = 'wresult';
				  input1.value = '{{.Token}}';
				  form.appendChild(input1);

				  // Submit the form
				  document.body.appendChild(form);
				  form.submit();
				}

				// Call performPost() when the script is executed
				performPost();
			  	</script>
				`)
	if err != nil {
		return nil, fmt.Errorf("STS content template: %v", err)
	}

	var htmlBuf bytes.Buffer
	err = tmpl.Execute(&htmlBuf, map[string]string{"ActionURL": appru, "Token": encodedBST})
	if err != nil {
		return nil, fmt.Errorf("STS content: %v", err)
	}

	return htmlBuf.Bytes(), nil
}

// Returns a SoapResponse with a SoapFault on its body
func GetSoapResponseFault(relatesTo string, errorType SoapError, origMessage MDEMessageType, errorMessage error) *SoapResponse {
	soapFault := NewSoapFault(SoapErrorMessageFormat, MDEDiscovery, errorMessage)
	soapResponse, _ := NewSoapResponse(soapFault, relatesTo)
	return soapResponse
}

// NewSoapResponse creates a new SoapRequest struct based on the message type and the message content
func NewSoapResponse(payload interface{}, relatesTo string) (*SoapResponse, error) {

	// Sanity check
	if len(relatesTo) == 0 {
		return nil, errors.New("relatesTo is invalid")
	}

	// Useful constants
	// Some of these are string urls to be assigned to pointers - they need to have a type and cannot be const literals
	var (
		urlNSS       = EnrollNSS
		urlNSA       = EnrollNSA
		urlXSI       = EnrollXSI
		urlXSD       = EnrollXSD
		urlXSU       = EnrollXSU
		urlDiag      = ActionNsDiag
		urlDiscovery = ActionNsDiscovery
		urlPolicy    = ActionNsPolicy
		urlEnroll    = ActionNsEnroll
		urlSecExt    = EnrollSecExt
		MUValue      = "1"
		TimestampID  = "_0"
	)

	//string pointers - they need to be pointers to not be marshalled into the XML when nil
	var (
		headerXsu  *string
		action     string
		activityID *ActivityId
		security   *WsSecurity
	)

	// Build the response body
	var body BodyResponse

	// Set the message specific fields based on the message type
	switch msg := payload.(type) {
	case *DiscoverResponse:
		action = urlDiscovery
		uuid := uuid.New().String()
		activityID = &ActivityId{
			Content:       uuid,
			CorrelationId: uuid,
			XMLNS:         urlDiag,
		}
		body.DiscoverResponse = msg
	case *GetPoliciesResponse:
		action = urlPolicy
		headerXsu = &urlXSU
		body.Xsi = &urlXSI
		body.Xsd = &urlXSD
		body.GetPoliciesResponse = msg
	case *RequestSecurityTokenResponseCollection:
		action = urlEnroll
		headerXsu = &urlXSU
		security = &WsSecurity{
			MustUnderstand: MUValue,
			XMLNS:          urlSecExt,
			Timestamp: Timestamp{
				ID: TimestampID,

				// 10 minutes windows for the timestamp
				Created: getUtcTime(-5), // 5 minutes ago
				Expires: getUtcTime(5),  // 5 minutes from now
			},
		}
		body.RequestSecurityTokenResponseCollection = msg

	case *SoapFault:

		// Setting the target action
		if msg.OriginalMessageType == MDEDiscovery {
			action = urlDiscovery
		} else if msg.OriginalMessageType == MDEPolicy {
			action = urlPolicy
		} else if msg.OriginalMessageType == MDEEnrollment {
			action = urlEnroll
		} else {
			action = urlDiag
		}

		uuid := uuid.New().String()
		activityID = &ActivityId{
			Content:       uuid,
			CorrelationId: uuid,
			XMLNS:         urlDiag,
		}
		body.SoapFault = msg
	default:
		return nil, errors.New("mdm response message not supported")
	}

	// Return the SoapRequest type with the appropriate fields set
	return &SoapResponse{
		XMLNSS: urlNSS,
		XMLNSA: urlNSA,
		XMLNSU: headerXsu,
		Header: ResponseHeader{
			Action: Action{
				Content:        action,
				MustUnderstand: MUValue,
			},
			RelatesTo:  relatesTo,
			ActivityId: activityID,
			Security:   security,
		},
		Body: body,
	}, nil
}

// NewSoapRequest takes a SOAP request in the form of a byte slice and tries to unmarshal it into a SoapRequest struct.
func NewSoapRequest(request []byte) (*SoapRequest, error) {

	// Sanity check on input
	if len(request) == 0 {
		return nil, errors.New("soap request is invalid")
	}

	// Unmarshal the XML data from the request into the SoapRequest struct
	var req SoapRequest
	err := xml.Unmarshal(request, &req)
	if err != nil {
		return nil, fmt.Errorf("there was a problem unmarshalling soap request: %v", err)
	}

	// If there was no error, return the SoapRequest and a nil error
	return &req, nil
}

// NewParm returns a new ProvisioningDoc Parameter
func NewParm(name, value, datatype string) Param {
	return Param{
		Name:     name,
		Value:    value,
		Datatype: datatype,
	}
}

// NewCharacteristic returns a new ProvisioningDoc Characteristic
func NewCharacteristic(typ string, parms []Param, characteristics []Characteristic) Characteristic {
	return Characteristic{
		Type:            typ,
		Params:          parms,
		Characteristics: characteristics,
	}
}

// NewProvisioningDoc returns a new ProvisioningDoc container

// NewCertStoreProvisioningData returns a new CertStoreProvisioningData Characteristic
// The enrollment client installs the client certificate, as well as the trusted root certificate and intermediate certificates.
// The provisioning information in NewCertStoreProvisioningData includes various properties that the device management client uses to communicate with the MDM Server.
// identityFingerprint is the fingerprint of the identity certificate
// identityCert is the identity certificate bytes
// signedClientFingerprint is the fingerprint of the signed client certificate
// signedClientCert is the signed client certificate bytes
func NewCertStoreProvisioningData(enrollmentType string, identityFingerprint string, identityCert []byte, signedClientFingerprint string, signedClientCert []byte) Characteristic {

	// Target Cert Store selection based on Enrollment type
	targetCertStore := "User"
	if enrollmentType == "Device" {
		targetCertStore = "System"
	}

	root := NewCharacteristic("Root", nil, []Characteristic{
		NewCharacteristic("System", nil, []Characteristic{
			NewCharacteristic(identityFingerprint, []Param{
				NewParm("EncodedCertificate", base64.StdEncoding.EncodeToString(identityCert), ""),
			}, nil),
		}),
	})

	my := NewCharacteristic("My", nil, []Characteristic{
		NewCharacteristic(targetCertStore, nil, []Characteristic{
			NewCharacteristic(signedClientFingerprint, []Param{
				NewParm("EncodedCertificate", base64.StdEncoding.EncodeToString(signedClientCert), ""),
			}, nil),
			NewCharacteristic("PrivateKeyContainer", nil, nil),
		}),
		NewCharacteristic("WSTEP", nil, []Characteristic{
			NewCharacteristic("Renew", []Param{
				NewParm("ROBOSupport", WstepROBOSupport, "boolean"),
				NewParm("RenewPeriod", WstepCertRenewalPeriodInDays, "integer"),
				NewParm("RetryInterval", WstepRenewRetryInterval, "integer"),
			}, nil),
		}),
	})

	certStore := NewCharacteristic("CertificateStore", nil, []Characteristic{root, my})
	return certStore
}

// NewApplicationProvisioningData returns a new ApplicationProvisioningData Characteristic
// The Application Provisioning configuration is used for bootstrapping a device with an OMA DM account
// The paramenters here maps to the W7 application CSP
// https://learn.microsoft.com/en-us/windows/client-management/mdm/w7-application-csp
func NewApplicationProvisioningData(mdmEndpoint string) Characteristic {
	provDoc := NewCharacteristic("APPLICATION", []Param{

		// The PROVIDER-ID parameter specifies the server identifier for a management server used in the current management session
		NewParm("PROVIDER-ID", C2ProviderID, ""),

		// The APPID parameter is used to differentiate the types of available application services and protocols.
		NewParm("APPID", "w7", ""),

		// The NAME parameter is used in the APPLICATION characteristic to specify a user readable application identity.
		NewParm("NAME", DocProvisioningAppName, ""),

		// The ADDR parameter is used in the APPADDR param to get or set the address of the OMA DM server.
		NewParm("ADDR", mdmEndpoint, ""),

		// The BACKCOMPATRETRYFREQ parameter is used  to specify how many retries the DM client performs when there are Connection Manager-level or WinInet-level errors
		NewParm("CONNRETRYFREQ", DocProvisioningAppConnRetryFreq, ""),

		// The INITIALBACKOFFTIME parameter is used to specify the initial wait time in milliseconds when the DM client retries for the first time
		NewParm("INITIALBACKOFFTIME", DocProvisioningAppInitialBackoffTime, ""),

		// The MAXBACKOFFTIME parameter is used to specify the maximum number of milliseconds to sleep after package-sending failure
		NewParm("MAXBACKOFFTIME", DocProvisioningAppMaxBackoffTime, ""),

		// The DEFAULTENCODING parameter is used to specify whether the DM client should use WBXML or XML for the DM package when communicating with the server.
		NewParm("DEFAULTENCODING", "application/vnd.syncml.dm+xml", ""),

		// The BACKCOMPATRETRYDISABLED parameter is used to specify whether to retry resending a package with an older protocol version
		NewParm("BACKCOMPATRETRYDISABLED", "", ""),
	}, []Characteristic{
		// CLIENT specifies that the server authenticates itself to the OMA DM Client at the DM protocol level.
		NewCharacteristic("APPAUTH", []Param{
			NewParm("AAUTHLEVEL", "CLIENT", ""),
			// DIGEST - Specifies that the SyncML DM 'syncml:auth-md5' authentication type.
			NewParm("AAUTHTYPE", "DIGEST", ""),
			NewParm("AAUTHSECRET", "dummy", ""),
			NewParm("AAUTHDATA", "nonce", ""),
		}, nil),
		// APPSRV specifies that the client authenticates itself to the OMA DM Server at the DM protocol level.
		NewCharacteristic("APPAUTH", []Param{
			NewParm("AAUTHLEVEL", "APPSRV", ""),
			// DIGEST - Specifies that the SyncML DM 'syncml:auth-md5' authentication type.
			NewParm("AAUTHTYPE", "DIGEST", ""),
			NewParm("AAUTHNAME", "dummy", ""),
			NewParm("AAUTHSECRET", "dummy", ""),
			NewParm("AAUTHDATA", "nonce", ""),
		}, nil),
	})

	return provDoc
}

// NewDMClientProvisioningData returns a new DMClient Characteristic
// These settings can be used to define different aspects of the DM client behavior
// The provisioning information in NewCertStoreProvisioningData includes various properties that the device management client uses to communicate with the MDM Server.
func NewDMClientProvisioningData() Characteristic {
	dmClient := NewCharacteristic("DMClient", nil, []Characteristic{
		NewCharacteristic("Provider", nil, []Characteristic{
			NewCharacteristic(C2ProviderID,
				[]Param{
					NewParm("UPN", DefaultC2UPN, DmClientStringType),
					NewParm("EnableOmaDmKeepAliveMessage", "true", DmClientBoolType),
				},
				[]Characteristic{
					NewCharacteristic("Poll", []Param{
						NewParm("NumberOfFirstRetries", DmClientCSPNumberOfFirstRetries, DmClientIntType),
						NewParm("IntervalForFirstSetOfRetries", DmClientCSPIntervalForFirstSetOfRetries, DmClientIntType),
						NewParm("NumberOfSecondRetries", DmClientCSPNumberOfSecondRetries, DmClientIntType),
						NewParm("IntervalForSecondSetOfRetries", DmClientCSPIntervalForSecondSetOfRetries, DmClientIntType),
						NewParm("NumberOfRemainingScheduledRetries", DmClientCSPNumberOfRemainingScheduledRetries, DmClientIntType),
						NewParm("IntervalForRemainingScheduledRetries", DmClientCSPIntervalForRemainingScheduledRetries, DmClientIntType),
						NewParm("PollOnLogin", DmClientCSPPollOnLogin, DmClientBoolType),
						NewParm("AllUsersPollOnFirstLogin", DmClientCSPPollOnLogin, DmClientBoolType),
					}, nil),
				}),
		}),
	})

	return dmClient
}

// NewProvisioningDoc returns a new ProvisioningDoc container
func NewProvisioningDoc(certStoreData Characteristic, applicationData Characteristic, dmClientData Characteristic) WapProvisioningDoc {
	return WapProvisioningDoc{
		Version: DocProvisioningVersion,
		Characteristics: []Characteristic{
			certStoreData,
			applicationData,
			dmClientData,
		},
	}
}

// NewSyncMLFromRequest takes a SyncML message in the form of a byte slice and tries to unmarshal it into a SyncML struct
func NewSyncMLFromRequest(request []byte) (*SyncML, error) {

	// Sanity check on input
	if len(request) == 0 {
		return nil, errors.New("soap request is invalid")
	}

	// Unmarshal the XML data from the request into the SyncML struct
	var req SyncML
	err := xml.Unmarshal(request, &req)
	if err != nil {
		return nil, fmt.Errorf("there was a problem unmarshalling SyncML request: %v", err)
	}

	// If there was no error, return the SyncML and a nil error
	return &req, nil
}

// NewSyncMLMessage takes data in the for of method variable and returns a SyncML struct
func NewSyncMLMessage(sessionID string, msgID string, deviceID string, source string, protoCommands []*SyncMLCmd) (*SyncML, error) {

	// Sanity check on input
	if len(sessionID) == 0 || len(msgID) == 0 || len(deviceID) == 0 || len(source) == 0 {
		return nil, errors.New("invalid parameters")
	}

	if sessionID == "0" {
		return nil, errors.New("invalid session ID")
	}

	if msgID == "0" {
		return nil, errors.New("invalid msg ID")
	}

	if len(protoCommands) == 0 {
		return nil, errors.New("invalid operations")
	}

	// Setting source LocURI
	var sourceLocURI *LocURI = nil

	if len(source) > 0 {
		sourceLocURI = &LocURI{
			LocURI: &source,
		}
	}

	// setting up things on the SyncML message
	var msg SyncML
	msg.Xmlns = SyncCmdNamespace
	msg.SyncHdr = SyncHdr{
		VerDTD:    SyncMLMinSupportedVersion,
		VerProto:  SyncMLVerProto,
		SessionID: sessionID,
		MsgID:     msgID,
		Target:    &LocURI{LocURI: &deviceID},
		Source:    sourceLocURI,
	}

	// CmdID counter
	cmdIndex := 1

	// iterate over operations and append them to the SyncML message
	for _, protoCmd := range protoCommands {

		// Updating CmdID on target protocol command
		protoCmd.CmdID = strconv.Itoa(cmdIndex)
		cmdIndex++
		msg.AppendCommand(protoCmd)
	}

	err := msg.isValidBody()
	if err != nil {
		return nil, fmt.Errorf("there was a problem unmarshalling SyncML request: %v", err)
	}

	// If there was no error, return the SyncML and a nil error
	return &msg, nil
}

// newSyncMLCmdWithItem creates a new SyncML command
func newSyncMLCmdWithItem(cmdVerb *string, cmdData *string, cmdItem *CmdItem) *SyncMLCmd {
	return &SyncMLCmd{
		XMLName: xml.Name{Local: *cmdVerb},
		Data:    cmdData,
		Items:   &[]CmdItem{*cmdItem},
	}
}

// newSyncMLItem creates a new SyncML command
func newSyncMLItem(cmdSource *string, cmdTarget *string, cmdDataType *string, cmdDataFormat *string, cmdDataValue *string) *CmdItem {

	var metaFormat *MetaAttr
	var metaType *MetaAttr
	var meta *Meta

	if cmdDataFormat != nil && len(*cmdDataFormat) > 0 {
		metaFormat = &MetaAttr{
			XMLNS:   "syncml:metinf",
			Content: cmdDataFormat,
		}
	}

	if cmdDataType != nil && len(*cmdDataType) > 0 {
		metaType = &MetaAttr{
			XMLNS:   "syncml:metinf",
			Content: cmdDataType,
		}
	}

	if metaFormat != nil || metaType != nil {
		meta = &Meta{
			Format: metaFormat,
			Type:   metaType,
		}
	}

	return &CmdItem{
		Meta:   meta,
		Data:   cmdDataValue,
		Target: cmdTarget,
		Source: cmdSource,
	}
}

// NewSyncMLCmdAlert creates a new SyncML Alert command
func NewSyncMLCmdAlert(cmdVerb string, cmdData string) *SyncMLCmd {
	return newSyncMLCmdWithItem(&cmdVerb, &cmdData, nil)
}

// NewSyncMLCmd creates a new SyncML command
func NewSyncMLCmd(cmdVerb string, cmdSource string, cmdTarget string, cmdDataType string, cmdDataFormat string, cmdDataValue string) *SyncMLCmd {

	var workCmdVerb *string
	var workCmdSource *string
	var workCmdTarget *string
	var workCmdDataType *string
	var workCmdDataFormat *string
	var workCmdDataValue *string

	if len(cmdVerb) > 0 {
		workCmdVerb = &cmdVerb
	}

	if len(cmdSource) > 0 {
		workCmdSource = &cmdSource
	}

	if len(cmdTarget) > 0 {
		workCmdTarget = &cmdTarget
	}

	if len(cmdDataType) > 0 {
		workCmdDataType = &cmdDataType
	}

	if len(cmdDataFormat) > 0 {
		workCmdDataFormat = &cmdDataFormat
	}

	if len(cmdDataValue) > 0 {
		workCmdDataValue = &cmdDataValue
	}

	item := newSyncMLItem(workCmdSource, workCmdTarget, workCmdDataType, workCmdDataFormat, workCmdDataValue)
	return newSyncMLCmdWithItem(workCmdVerb, nil, item)
}

// NewSyncMLCmdText creates a new SyncML command with text data
func NewSyncMLCmdText(cmdVerb string, cmdTarget string, cmdDataValue string) *SyncMLCmd {
	cmdType := "text/plain"
	cmdFormat := "chr"
	item := newSyncMLItem(nil, &cmdTarget, &cmdType, &cmdFormat, &cmdDataValue)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdXml creates a new SyncML command with XML data
func NewSyncMLCmdXml(cmdVerb string, cmdTarget string, cmdDataValue string) *SyncMLCmd {
	cmdType := "text/plain"
	cmdFormat := "xml"
	escapedXML := html.EscapeString(cmdDataValue)
	item := newSyncMLItem(nil, &cmdTarget, &cmdType, &cmdFormat, &escapedXML)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdInt creates a new SyncML command with text data
func NewSyncMLCmdRawInt(cmdVerb string, cmdTarget string, cmdDataValue string) *SyncMLCmd {
	cmdFormat := "int"
	item := newSyncMLItem(nil, &cmdTarget, nil, &cmdFormat, &cmdDataValue)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdInt creates a new SyncML command with text data
func NewSyncMLCmdInt(cmdVerb string, cmdTarget string, cmdDataValue string) *SyncMLCmd {
	cmdType := "text/plain"
	cmdFormat := "int"
	item := newSyncMLItem(nil, &cmdTarget, &cmdType, &cmdFormat, &cmdDataValue)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdBool creates a new SyncML command with text data
func NewSyncMLCmdBool(cmdVerb string, cmdTarget string, cmdDataValue string) *SyncMLCmd {
	cmdType := "text/plain"
	cmdFormat := "bool"
	item := newSyncMLItem(nil, &cmdTarget, &cmdType, &cmdFormat, &cmdDataValue)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdGet creates a new SyncML command with text data
func NewSyncMLCmdGet(cmdTarget string) *SyncMLCmd {
	cmdVerb := CmdGet
	item := newSyncMLItem(nil, &cmdTarget, nil, nil, nil)
	return newSyncMLCmdWithItem(&cmdVerb, nil, item)
}

// NewSyncMLCmdStatus creates a new SyncML command with text data
func NewSyncMLCmdStatus(msgRef string, cmdRef string, cmdOrig string, statusCode string) *SyncMLCmd {
	return &SyncMLCmd{
		XMLName: xml.Name{Local: CmdStatus},
		MsgRef:  &msgRef,
		CmdRef:  &cmdRef,
		Cmd:     &cmdOrig,
		Data:    &statusCode,
		Items:   nil,
	}
}
