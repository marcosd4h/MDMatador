package mdm

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

// MS-MDE2 Message request types
type MDEMessageType int

const (
	MDEDiscovery MDEMessageType = iota
	MDEPolicy
	MDESTSAuth
	MDEEnrollment
	MDEFault
)

// MS-MDE2 Binary Security Token Types
type BinSecTokenType int

const (
	MDETokenPKCS7 = iota
	MDETokenPKCS10
	MDETokenPKCSInvalid
)

///////////////////////////////////////////////////////////////
/// MS-MDE2 Discover message request type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/36e33def-59ab-484f-b0bc-701496346925

type Discover struct {
	XMLNS   string          `xml:"xmlns,attr"`
	Request DiscoverRequest `xml:"request"`
}

type AuthPolicies struct {
	AuthPolicy []string `xml:"AuthPolicy"`
}

type DiscoverRequest struct {
	XMLNS              string       `xml:"i,attr"`
	EmailAddress       string       `xml:"EmailAddress"`
	RequestVersion     string       `xml:"RequestVersion"`
	DeviceType         string       `xml:"DeviceType"`
	ApplicationVersion string       `xml:"ApplicationVersion"`
	OSEdition          string       `xml:"OSEdition"`
	AuthPolicies       AuthPolicies `xml:"AuthPolicies"`
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 GetPolicies message request type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/02b080e4-d1d8-4e0c-af14-b77931cec404

type GetPolicies struct {
	XMLNS         string        `xml:"xmlns,attr"`
	Client        Client        `xml:"client"`
	RequestFilter RequestFilter `xml:"requestFilter"`
}

type ClientContent struct {
	Content string `xml:",chardata"`
	Xsi     string `xml:"nil,attr"`
}

type Client struct {
	LastUpdate        ClientContent `xml:"lastUpdate"`
	PreferredLanguage ClientContent `xml:"preferredLanguage"`
}

type RequestFilter struct {
	Xsi string `xml:"nil,attr"`
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 RequestSecurityToken message request type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/6ba9c509-8bce-4899-85b2-8c3d41f8f845

type RequestSecurityToken struct {
	TokenType           string                 `xml:"TokenType"`
	RequestType         string                 `xml:"RequestType"`
	BinarySecurityToken BinarySecurityToken    `xml:"BinarySecurityToken"`
	AdditionalContext   AdditionalContext      `xml:"AdditionalContext"`
	MapContextItems     map[string]ContextItem `xml:"-"`
}

// BinarySecurityToken contains the base64 encoding representation of the security token
// The token format is defined by the WS-Trust X509v3 Enrollment Extensions [MS-WSTEP] specification
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/5b02c625-ced2-4a01-a8e1-da0ae84f5bb7
type BinarySecurityToken struct {
	Content      string  `xml:",chardata"`
	XMLNS        *string `xml:"xmlns,attr"`
	ValueType    string  `xml:"ValueType,attr"`
	EncodingType string  `xml:"EncodingType,attr"`
}

type ContextItem struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value"`
}

type AdditionalContext struct {
	XMLNS        string        `xml:"xmlns,attr"`
	ContextItems []ContextItem `xml:"ContextItem"`
}

// Get Binary Security Token
func (msg RequestSecurityToken) GetBinarySecurityTokenData() (string, error) {

	if len(msg.BinarySecurityToken.Content) == 0 {
		return "", errors.New("BinarySecurityToken is empty")
	}

	return msg.BinarySecurityToken.Content, nil
}

// Get Binary Security Token Type
func (msg RequestSecurityToken) GetBinarySecurityTokenType() (BinSecTokenType, error) {

	if msg.BinarySecurityToken.ValueType == EnrollReqTypePKCS10 {
		return MDETokenPKCS10, nil
	}

	if msg.BinarySecurityToken.ValueType == EnrollReqTypePKCS7 {
		return MDETokenPKCS7, nil
	}

	return MDETokenPKCSInvalid, errors.New("BinarySecurityToken is invalid")
}

// Get SecurityToken Context Item
func (msg *RequestSecurityToken) GetContextItem(item string) (string, error) {

	if len(msg.AdditionalContext.ContextItems) == 0 {
		return "", errors.New("ContextItems is empty")
	}

	// Generate map of ContextItems if not there
	if msg.MapContextItems == nil {
		contextMap := make(map[string]ContextItem)
		for _, item := range msg.AdditionalContext.ContextItems {
			contextMap[item.Name] = item
		}
		msg.MapContextItems = contextMap
	}

	itemVal, valueFound := (msg.MapContextItems)[item]
	if !valueFound {
		return "", fmt.Errorf("ContextItem item %s is not present", item)
	}

	return itemVal.Value, nil
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 DiscoverResponse message response type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/aa198049-e691-41f9-a45a-b973b9089be7

type DiscoverResponse struct {
	XMLName        xml.Name       `xml:"DiscoverResponse"`
	XMLNS          string         `xml:"xmlns,attr"`
	DiscoverResult DiscoverResult `xml:"DiscoverResult"`
}

type DiscoverResult struct {
	AuthPolicy                 string  `xml:"AuthPolicy"`
	EnrollmentVersion          string  `xml:"EnrollmentVersion"`
	EnrollmentPolicyServiceUrl string  `xml:"EnrollmentPolicyServiceUrl"`
	EnrollmentServiceUrl       string  `xml:"EnrollmentServiceUrl"`
	EnrollmentAuthServiceUrl   *string `xml:"AuthenticationServiceUrl"`
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 GetPoliciesResponse message response type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/6e74dcdb-c3d9-4044-af10-536224904e72

type GetPoliciesResponse struct {
	XMLName  xml.Name `xml:"GetPoliciesResponse"`
	XMLNS    string   `xml:"xmlns,attr"`
	Response Response `xml:"response"`
	OIDs     OIDs     `xml:"oIDs"`
}

type ContentAttr struct {
	Content string `xml:",chardata"`
	Xsi     string `xml:"xsi:nil,attr"`
	XMLNS   string `xml:"xmlns:xsi,attr"`
}

type GenericAttr struct {
	Content string `xml:",chardata"`
	Xsi     string `xml:"xsi:nil,attr"`
}

type CertificateValidity struct {
	ValidityPeriodSeconds string `xml:"validityPeriodSeconds"`
	RenewalPeriodSeconds  string `xml:"renewalPeriodSeconds"`
}

type Permission struct {
	Enroll     string `xml:"enroll"`
	AutoEnroll string `xml:"autoEnroll"`
}

type ProviderAttr struct {
	Content string `xml:",chardata"`
}

type PrivateKeyAttributes struct {
	MinimalKeyLength      string         `xml:"minimalKeyLength"`
	KeySpec               GenericAttr    `xml:"keySpec"`
	KeyUsageProperty      GenericAttr    `xml:"keyUsageProperty"`
	Permissions           GenericAttr    `xml:"permissions"`
	AlgorithmOIDReference GenericAttr    `xml:"algorithmOIDReference"`
	CryptoProviders       []ProviderAttr `xml:"provider"`
}
type Revision struct {
	MajorRevision string `xml:"majorRevision"`
	MinorRevision string `xml:"minorRevision"`
}

type Attributes struct {
	CommonName                string               `xml:"commonName"`
	PolicySchema              string               `xml:"policySchema"`
	CertificateValidity       CertificateValidity  `xml:"certificateValidity"`
	Permission                Permission           `xml:"permission"`
	PrivateKeyAttributes      PrivateKeyAttributes `xml:"privateKeyAttributes"`
	Revision                  Revision             `xml:"revision"`
	SupersededPolicies        GenericAttr          `xml:"supersededPolicies"`
	PrivateKeyFlags           GenericAttr          `xml:"privateKeyFlags"`
	SubjectNameFlags          GenericAttr          `xml:"subjectNameFlags"`
	EnrollmentFlags           GenericAttr          `xml:"enrollmentFlags"`
	GeneralFlags              GenericAttr          `xml:"generalFlags"`
	HashAlgorithmOIDReference string               `xml:"hashAlgorithmOIDReference"`
	RARequirements            GenericAttr          `xml:"rARequirements"`
	KeyArchivalAttributes     GenericAttr          `xml:"keyArchivalAttributes"`
	Extensions                GenericAttr          `xml:"extensions"`
}

type Policy struct {
	PolicyOIDReference string      `xml:"policyOIDReference"`
	CAs                GenericAttr `xml:"cAs"`
	Attributes         Attributes  `xml:"attributes"`
}

type Policies struct {
	Policy Policy `xml:"policy"`
}

type Response struct {
	PolicyID           string      `xml:"policyID"`
	PolicyFriendlyName ContentAttr `xml:"policyFriendlyName"`
	NextUpdateHours    ContentAttr `xml:"nextUpdateHours"`
	PoliciesNotChanged ContentAttr `xml:"policiesNotChanged"`
	Policies           Policies    `xml:"policies"`
}

type OID struct {
	Value          string `xml:"value"`
	Group          string `xml:"group"`
	OIDReferenceID string `xml:"oIDReferenceID"`
	DefaultName    string `xml:"defaultName"`
}

type OIDs struct {
	Content string `xml:",chardata"`
	OID     []OID  `xml:"oID"`
}

///////////////////////////////////////////////////////////////
/// RequestSecurityTokenResponseCollection MS-MDE2 Message response type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/3452fe7d-2441-49f7-8801-c056b58edb6a

type RequestSecurityTokenResponseCollection struct {
	XMLName                      xml.Name                     `xml:"RequestSecurityTokenResponseCollection"`
	XMLNS                        string                       `xml:"xmlns,attr"`
	RequestSecurityTokenResponse RequestSecurityTokenResponse `xml:"RequestSecurityTokenResponse"`
}

type SecAttr struct {
	Content string `xml:",chardata"`
	XmlNS   string `xml:"xmlns,attr"`
}

type RequestedSecurityToken struct {
	BinarySecurityToken BinarySecurityToken `xml:"BinarySecurityToken"`
}

type RequestSecurityTokenResponse struct {
	TokenType              string                 `xml:"TokenType"`
	DispositionMessage     SecAttr                `xml:"DispositionMessage"`
	RequestedSecurityToken RequestedSecurityToken `xml:"RequestedSecurityToken"`
	RequestID              SecAttr                `xml:"RequestID"`
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 SoapFault message response type
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/0a78f419-5fd7-4ddb-bc76-1c0f7e11da23

type Subcode struct {
	Value string `xml:"s:value"`
}

type Code struct {
	Value   string  `xml:"s:value"`
	Subcode Subcode `xml:"s:subcode"`
}

type ReasonText struct {
	Content string `xml:",chardata"`
	Lang    string `xml:"xml:lang,attr"`
}

type Reason struct {
	Text ReasonText `xml:"s:text"`
}

type SoapFault struct {
	XMLName             xml.Name       `xml:"s:fault"`
	Code                Code           `xml:"s:code"`
	Reason              Reason         `xml:"s:reason"`
	OriginalMessageType MDEMessageType `xml:"-"`
}

// Error returns the soap fault as an Error
func (msg SoapFault) Error() error {
	return fmt.Errorf("soap fault: %s", msg.Reason.Text.Content)
}

///////////////////////////////////////////////////////////////
/// MS-MDE2 ProvisioningDoc (XML Provisioning Schema) message type
/// Section 2.2.9.1 on the specification
/// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mde2/35e1aca6-1b8a-48ba-bbc0-23af5d46907a

type Param struct {
	Name     string `xml:"name,attr,omitempty"`
	Value    string `xml:"value,attr,omitempty"`
	Datatype string `xml:"datatype,attr,omitempty"`
}

type Characteristic struct {
	Type            string           `xml:"type,attr"`
	Params          []Param          `xml:"parm"`
	Characteristics []Characteristic `xml:"characteristic,omitempty"`
}

type WapProvisioningDoc struct {
	XMLName         xml.Name         `xml:"wap-provisioningdoc"`
	Version         string           `xml:"version,attr"`
	Characteristics []Characteristic `xml:"characteristic"`
}

// Add Characteristic to the Characteristic container
func (msg *Characteristic) AddCharacteristic(c Characteristic) {
	msg.Characteristics = append(msg.Characteristics, c)
}

// Add Param to the Params container
func (msg *Characteristic) AddParam(name string, value string, dataType string) {
	param := Param{Name: name, Value: value, Datatype: dataType}
	msg.Params = append(msg.Params, param)
}

// Add Characteristic to the WapProvisioningDoc
func (msg *WapProvisioningDoc) AddCharacteristic(c Characteristic) {
	msg.Characteristics = append(msg.Characteristics, c)
}

// GetEncodedB64Representation returns encoded WapProvisioningDoc representation
func (msg WapProvisioningDoc) GetEncodedB64Representation() (string, error) {
	rawXML, err := xml.MarshalIndent(msg, "", "  ")
	if err != nil {
		return "", err
	}

	//Appending the XML header beforing encoding it
	xmlContent := append([]byte(xml.Header), rawXML...)

	// Create a replacer to replace both "\n" and "\t"
	replacer := strings.NewReplacer("\n", "", "\t", "")

	// Use the replacer on the string representation of xmlContent
	xmlStripContent := []byte(replacer.Replace(string(xmlContent)))

	return base64.StdEncoding.EncodeToString(xmlStripContent), nil
}

///////////////////////////////////////////////////////////////
/// MS-MDM Message types

// Possible Session States
type SessionState int

const (
	MDMSessionInitiated SessionState = iota
	MDMSessionFinalized
)

// SessionOrigin represents the origin of the session
type SessionOrigin int

const (
	MDMSessionServerInitiated SessionOrigin = iota
	MDMSessionClientInitiated
)

// UserAgentOrigin represents the origin of the MDM session
type UserAgentOrigin int

const (
	SesOriginUserTriggered  UserAgentOrigin = 1
	SesOriginServer         UserAgentOrigin = 2
	SesOriginWNS            UserAgentOrigin = 3
	SesOriginUserLogin      UserAgentOrigin = 4
	SesOriginPostEnrollment UserAgentOrigin = 5
	SesOriginCSPTriggered   UserAgentOrigin = 6
	SesOriginPublicApi      UserAgentOrigin = 7
)
