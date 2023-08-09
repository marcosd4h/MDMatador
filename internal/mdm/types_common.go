package mdm

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/marcosd4h/MDMatador/internal/validator"
)

// MS-MDE2 is a client-to-server protocol that consists of a SOAP-based Web service.
// SOAP is a lightweight and XML based protocol that consists of three parts:
// - An envelope that defines a framework for describing what is in a message and how to process it
// - A set of encoding rules for expressing instances of application-defined datatypes
// - And a convention for representing remote procedure calls and responses.

// Below are the XML types used by the SOAP protocol

// SoapResponse is the Soap Envelope Response type for MS-MDE2 responses from the server
// This envelope XML message is composed by a mandatory SOAP envelope, a SOAP header, and a SOAP body
type SoapResponse struct {
	XMLName xml.Name       `xml:"s:Envelope"`
	XMLNSS  string         `xml:"xmlns:s,attr"`
	XMLNSA  string         `xml:"xmlns:a,attr"`
	XMLNSU  *string        `xml:"xmlns:u,attr,omitempty"`
	Header  ResponseHeader `xml:"s:Header"`
	Body    BodyResponse   `xml:"s:Body"`
}

// Error returns soap fault error if present
func (msg *SoapResponse) GetError() error {
	if msg.Body.SoapFault != nil {
		return msg.Body.SoapFault.Error()
	}
	return nil
}

// SoapRequest is the Soap Envelope Request type for MS-MDE2 responses to the server
// This envelope XML message is composed by a mandatory SOAP envelope, a SOAP header, and a SOAP body
type SoapRequest struct {
	XMLName   xml.Name      `xml:"Envelope"`
	XMLNSS    string        `xml:"s,attr"`
	XMLNSA    string        `xml:"a,attr"`
	XMLNSU    *string       `xml:"u,attr,omitempty"`
	XMLNSWsse *string       `xml:"wsse,attr,omitempty"`
	XMLNSWST  *string       `xml:"wst,attr,omitempty"`
	XMLNSAC   *string       `xml:"ac,attr,omitempty"`
	Header    RequestHeader `xml:"Header"`
	Body      BodyRequest   `xml:"Body"`
}

// HTTP request header field used to indicate the intent of the SOAP request, using a URI value
// See section 6.1.1 on SOAP Spec - https://www.w3.org/TR/2000/NOTE-SOAP-20000508/#_Toc478383527
type Action struct {
	Content        string `xml:",chardata"`
	MustUnderstand string `xml:"s:mustUnderstand,attr"`
}

// ActivityId is a unique identifier for the activity
type ActivityId struct {
	Content       string `xml:",chardata"`
	CorrelationId string `xml:"CorrelationId,attr"`
	XMLNS         string `xml:"xmlns,attr"`
}

// Timestamp for certificate authentication
type Timestamp struct {
	ID      string `xml:"u:Id,attr"`
	Created string `xml:"u:Created"`
	Expires string `xml:"u:Expires"`
}

// Security token container
type WsSecurity struct {
	XMLNS          string    `xml:"xmlns:o,attr"`
	MustUnderstand string    `xml:"s:mustUnderstand,attr"`
	Timestamp      Timestamp `xml:"u:Timestamp"`
}

// Security token container for encoded security sensitive data
type BinSecurityToken struct {
	Content  string `xml:",chardata"`
	Value    string `xml:"ValueType,attr"`
	Encoding string `xml:"EncodingType,attr"`
}

// TokenSecurity is the security token container for BinSecurityToken
type TokenSecurity struct {
	MustUnderstand string           `xml:"s:mustUnderstand,attr"`
	Security       BinSecurityToken `xml:"BinarySecurityToken"`
}

// To target endpoint header field
type To struct {
	Content        string `xml:",chardata"`
	MustUnderstand string `xml:"mustUnderstand,attr"`
}

// ReplyTo message correlation header field
type ReplyTo struct {
	Address string `xml:"Address"`
}

// ResponseHeader is the header for MDM responses from the server
type ResponseHeader struct {
	Action     Action      `xml:"a:Action"`
	ActivityId *ActivityId `xml:"ActivityId,omitempty"`
	RelatesTo  string      `xml:"a:RelatesTo"`
	Security   *WsSecurity `xml:"o:Security,omitempty"`
}

// RequestHeader is the header for MDM requests to the server
type RequestHeader struct {
	Action    Action         `xml:"Action"`
	MessageID string         `xml:"MessageID"`
	ReplyTo   ReplyTo        `xml:"ReplyTo"`
	To        To             `xml:"To"`
	Security  *TokenSecurity `xml:"Security,omitempty"`
}

// BodyReponse is the body of the MDM SOAP response message
type BodyResponse struct {
	Xsd                                    *string                                 `xml:"xmlns:xsd,attr,omitempty"`
	Xsi                                    *string                                 `xml:"xmlns:xsi,attr,omitempty"`
	DiscoverResponse                       *DiscoverResponse                       `xml:"DiscoverResponse,omitempty"`
	GetPoliciesResponse                    *GetPoliciesResponse                    `xml:"GetPoliciesResponse,omitempty"`
	RequestSecurityTokenResponseCollection *RequestSecurityTokenResponseCollection `xml:"RequestSecurityTokenResponseCollection,omitempty"`
	SoapFault                              *SoapFault                              `xml:"s:fault,omitempty"`
}

// BodyRequest is the body of the MDM SOAP request message
type BodyRequest struct {
	Xsi                  *string               `xml:"xsi,attr,omitempty"`
	Xsd                  *string               `xml:"xsd,attr,omitempty"`
	Discover             *Discover             `xml:"Discover,omitempty"`
	GetPolicies          *GetPolicies          `xml:"GetPolicies,omitempty"`
	RequestSecurityToken *RequestSecurityToken `xml:"RequestSecurityToken,omitempty"`
}

// isValidHeader checks for required fields in the header
func (req *SoapRequest) isValidHeader() error {
	// Check for required fields

	if !validator.HasContent(req.XMLNSS) {
		return errors.New("XMLNSS")
	}

	if !validator.HasContent(req.XMLNSA) {
		return errors.New("XMLNSA")
	}

	if !validator.HasContent(req.Header.MessageID) {
		return errors.New("Header.MessageID")
	}

	if !validator.HasContent(req.Header.Action.Content) {
		return errors.New("Header.Action")
	}

	if !validator.HasContent(req.Header.ReplyTo.Address) {
		return errors.New("Header.ReplyTo")
	}

	if !validator.HasContent(req.Header.To.Content) {
		return errors.New("Header.To")
	}

	return nil
}

// isValidBody checks for the presence of only one message
func (req *SoapRequest) isValidBody() error {
	nonNilCount := 0

	if req.Body.Discover != nil {
		nonNilCount++
	}
	if req.Body.GetPolicies != nil {
		nonNilCount++
	}
	if req.Body.RequestSecurityToken != nil {
		nonNilCount++
	}

	if nonNilCount != 1 {
		return errors.New("invalid SOAP body: Multiple messages or no message")
	}

	return nil
}

// IsValidDiscoveryMsg checks for required fields in the Discover message
func (req *SoapRequest) IsValidDiscoveryMsg() error {

	if err := req.isValidHeader(); err != nil {
		return fmt.Errorf("invalid discover message header: %s", err)
	}

	if err := req.isValidBody(); err != nil {
		return fmt.Errorf("invalid discover message body: %s", err)
	}

	if req.Body.Discover == nil {
		return errors.New("invalid discover message: Discover message not present")
	}

	if !validator.HasContent(req.Body.Discover.XMLNS) {
		return errors.New("invalid discover message: XMLNS")
	}

	// Traverse the AuthPolicies slice and check for valid values
	isInvalidAuth := true
	for _, authPolicy := range req.Body.Discover.Request.AuthPolicies.AuthPolicy {
		if authPolicy == AuthOnPremise {
			isInvalidAuth = false
			break
		}
	}

	if isInvalidAuth {
		return errors.New("invalid discover message: Request.AuthPolicies")
	}

	return nil
}

func (req *SoapRequest) IsProgrammaticDiscovery() (bool, error) {

	if err := req.isValidHeader(); err != nil {
		return false, fmt.Errorf("invalid discover message header: %s", err)
	}

	if err := req.isValidBody(); err != nil {
		return false, fmt.Errorf("invalid discover message body: %s", err)
	}

	if validator.IsEmail(req.Body.Discover.Request.EmailAddress) {
		return false, nil
	}

	return true, nil
}

// IsValidGetPolicyMsg checks for required fields in the GetPolicies message
func (req *SoapRequest) IsValidGetPolicyMsg() error {
	if err := req.isValidHeader(); err != nil {
		return fmt.Errorf("invalid getpolicies message:  %s", err)
	}

	if err := req.isValidBody(); err != nil {
		return fmt.Errorf("invalid getpolicies message: %s", err)
	}

	if req.Body.GetPolicies == nil {
		return errors.New("invalid getpolicies message:  GetPolicies message not present")
	}

	if !validator.HasContent(req.Body.GetPolicies.XMLNS) {
		return errors.New("invalid getpolicies message: XMLNS")
	}

	return nil
}

// IsValidRequestSecurityTokenMsg checks for required fields in the RequestSecurityToken message
func (req *SoapRequest) IsValidRequestSecurityTokenMsg() error {

	if err := req.isValidHeader(); err != nil {
		return fmt.Errorf("invalid requestsecuritytoken message:: %s", err)
	}

	if err := req.isValidBody(); err != nil {
		return fmt.Errorf("invalid requestsecuritytoken message: %s", err)
	}

	if req.Body.RequestSecurityToken == nil {
		return errors.New("invalid requestsecuritytoken message: RequestSecurityToken message not present")
	}

	if !validator.HasContent(req.Body.RequestSecurityToken.TokenType) {
		return errors.New("invalid requestsecuritytoken message: TokenType")
	}

	if !validator.HasContent(req.Body.RequestSecurityToken.RequestType) {
		return errors.New("invalid requestsecuritytoken message: RequestType")
	}

	if !validator.HasContent(req.Body.RequestSecurityToken.BinarySecurityToken.ValueType) {
		return errors.New("invalid requestsecuritytoken message: BinarySecurityToken.ValueType")
	}

	if req.Body.RequestSecurityToken.BinarySecurityToken.ValueType != EnrollReqTypePKCS10 &&
		req.Body.RequestSecurityToken.BinarySecurityToken.ValueType != EnrollReqTypePKCS7 {
		return errors.New("invalid requestsecuritytoken message: BinarySecurityToken.EncodingType  not supported")
	}

	if !validator.HasContent(req.Body.RequestSecurityToken.BinarySecurityToken.Content) {
		return errors.New("invalid requestsecuritytoken message: BinarySecurityToken.Content")
	}

	if len(req.Body.RequestSecurityToken.AdditionalContext.ContextItems) == 0 {
		return errors.New("invalid requestsecuritytoken message: AdditionalContext.ContextItems missing")
	}

	reqVersion, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemRequestVersion)
	if err != nil || (reqVersion != EnrollmentVersionV5 && reqVersion != EnrollmentVersionV4) {
		reqVersion = EnrollmentVersionV5
	}

	reqEnrollType, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemEnrollmentType)
	if err != nil || reqEnrollType != ReqSecTokenEnrollType {
		return fmt.Errorf("invalid requestsecuritytoken message %s: %s - %v", ReqSecTokenContextItemEnrollmentType, reqEnrollType, err)
	}

	reqDeviceID, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemDeviceID)
	if err != nil || !validator.HasContent(reqDeviceID) {
		return fmt.Errorf("invalid requestsecuritytoken message %s: %s - %v", ReqSecTokenContextItemDeviceID, reqDeviceID, err)
	}

	reqHwDeviceID, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemHWDevID)
	if err != nil || !validator.HasContent(reqHwDeviceID) {
		return fmt.Errorf("invalid requestsecuritytoken message %s: %s - %v", ReqSecTokenContextItemHWDevID, reqHwDeviceID, err)
	}

	reqOSEdition, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemOSEdition)
	if err != nil || !validator.HasContent(reqOSEdition) {
		return fmt.Errorf("invalid requestsecuritytoken message %s: %s - %v", ReqSecTokenContextItemOSEdition, reqOSEdition, err)
	}

	reqOSVersion, err := req.Body.RequestSecurityToken.GetContextItem(ReqSecTokenContextItemOSVersion)
	if err != nil || !validator.HasContent(reqOSEdition) {
		return fmt.Errorf("invalid requestsecuritytoken message %s: %s - %v", ReqSecTokenContextItemOSVersion, reqOSVersion, err)
	}

	return nil
}

// MessageID returns the message ID from the header
func (req *SoapRequest) MessageID() string {
	return req.Header.MessageID
}

// Get Discover MDM Message from the body
func (req *SoapRequest) GetDiscoverMessage() (*Discover, error) {

	if req.Body.Discover == nil {
		return nil, errors.New("invalid body: Discover message not present")
	}

	return req.Body.Discover, nil
}

// Get GetPolicies MDM Message from the body
func (req *SoapRequest) GetPoliciesMessage() (*GetPolicies, error) {

	if req.Body.GetPolicies == nil {
		return nil, errors.New("invalid body: GetPolicies message not present")
	}

	return req.Body.GetPolicies, nil
}

// Get RequestSecurityToken MDM Message from the body
func (req *SoapRequest) GetRequestSecurityTokenMessage() (*RequestSecurityToken, error) {

	if req.Body.RequestSecurityToken == nil {
		return nil, errors.New("invalid body: RequestSecurityToken message not present")
	}

	return req.Body.RequestSecurityToken, nil
}

// MS-MDM is a client-to-server protocol that consists of a SOAP-based Web service.
// MDM is based on the OMA-DM protocol. Messages are issued by a requester and results and status are returned by a responder as a SynCML message.
// A SyncML message is a well-formed XML document that adheres to the document type definition (DTD), but which does not require validation.
// The XML document is identified by a SyncML document (or root) element type that serves as a parent container for the SyncML message.
// The SyncML message consists of a header specified by the SyncHdr  element type and a body specified by the SyncBody element type.
// The SyncML header identifies the routing and versioning information about the SyncML message.
// The SyncML body functions as a container for one or more SyncML commands.
// A SyncML command is specified by individual element types that provide specific details about the command, including any data or meta-information.
// MS-MDM uses a subset of the SyncML message definition specified in OMA-SyncMLRP spec. MDM-specific SyncML xml message format is defined in OMA-DMRP.

type SyncML struct {
	XMLName  xml.Name `xml:"SyncML"`
	Xmlns    string   `xml:"xmlns,attr"`
	SyncHdr  SyncHdr  `xml:"SyncHdr"`
	SyncBody SyncBody `xml:"SyncBody"`
}

type SyncHdr struct {
	VerDTD    string   `xml:"VerDTD"`
	VerProto  string   `xml:"VerProto"`
	SessionID string   `xml:"SessionID"`
	MsgID     string   `xml:"MsgID"`
	Target    *LocURI  `xml:"Target,omitempty"`
	Source    *LocURI  `xml:"Source,omitempty"`
	Meta      *MetaHdr `xml:"Meta,omitempty"`
}

type MetaHdr struct {
	MaxMsgSize *string `xml:"MaxMsgSize,omitempty"`
}

// ProtoCmds contains a slice of SyncML protocol commands
type ProtoCmds *[]SyncMLCmd

// See supported Commands in section 2.2.7.1
type SyncBody struct {
	Final *string `xml:"Final,omitempty"`

	//Request Protocol Commands
	Add     ProtoCmds `xml:"Add,omitempty"`
	Alert   ProtoCmds `xml:"Alert,omitempty"`
	Atomic  ProtoCmds `xml:"Atomic,omitempty"`
	Delete  ProtoCmds `xml:"Delete,omitempty"`
	Exec    ProtoCmds `xml:"Exec,omitempty"`
	Get     ProtoCmds `xml:"Get,omitempty"`
	Replace ProtoCmds `xml:"Replace,omitempty"`

	//Response Protocol Commands
	Results ProtoCmds `xml:"Results,omitempty"`
	Status  ProtoCmds `xml:"Status,omitempty"`

	//Raw container
	Raw ProtoCmds `xml:",omitempty"`
}

// ProtoCmdState is the state of the SyncML protocol commands
type ProtoCmdState int

const (
	Received           ProtoCmdState = iota //Protocol Command was received
	Pending                                 //Protocol Command is on the pending queue and has not been sent yet
	Sent                                    //Protocol Command has been sent
	ResponseProcessing                      //Protocol Command was acknowledged and is being processed
	ResponseAck                             //Protocol Command was acknowledged and processed
)

// Supported protocol command verbs
const (
	CmdAdd     = "Add"     //Protocol Command verb Add
	CmdAlert   = "Alert"   //Protocol Command verb Alert
	CmdAtomic  = "Atomic"  //Protocol Command verb Atomic
	CmdDelete  = "Delete"  //Protocol Command verb Delete
	CmdExec    = "Exec"    //Protocol Command verb Exec
	CmdGet     = "Get"     //Protocol Command verb Get
	CmdReplace = "Replace" //Protocol Command verb Replace
	CmdResults = "Results" //Protocol Command verb Results
	CmdStatus  = "Status"  //Protocol Command verb Status
)

// ProtoCmdOperation is the abstraction to represent a SyncML Protocol Command
type ProtoCmdOperation struct {
	Verb string    `db:"verb"`
	Cmd  SyncMLCmd `db:"cmd"`
}

// Protocol Command
type SyncMLCmd struct {
	XMLName xml.Name   `xml:",omitempty"`
	CmdID   string     `xml:"CmdID"`
	MsgRef  *string    `xml:"MsgRef,omitempty"`
	CmdRef  *string    `xml:"CmdRef,omitempty"`
	Cmd     *string    `xml:"Cmd,omitempty"`
	Data    *string    `xml:"Data,omitempty"`
	Items   *[]CmdItem `xml:"Item,omitempty"`
}

type CmdItem struct {
	Source *string `xml:"Source>LocURI,omitempty"`
	Target *string `xml:"Target>LocURI,omitempty"`
	Meta   *Meta   `xml:"Meta,omitempty"`
	Data   *string `xml:"Data"`
}

type Meta struct {
	Type   *MetaAttr `xml:"Type,omitempty"`
	Format *MetaAttr `xml:"Format,omitempty"`
}

type MetaAttr struct {
	XMLNS   string  `xml:"xmlns,attr"`
	Content *string `xml:",chardata"`
}

type LocURI struct {
	LocURI *string `xml:",omitempty"`
}

func (msg *SyncML) isValidHeader() error {

	if !validator.HasContent(msg.Xmlns) {
		return errors.New("msg namespace")
	}

	//SyncML DTD version check
	if msg.SyncHdr.VerDTD != SyncMLMinSupportedVersion && msg.SyncHdr.VerDTD != SyncMLMaxSupportedVersion {
		return errors.New("unsupported DTD version")
	}

	//SyncML Proto version check
	if msg.SyncHdr.VerDTD != SyncMLMinSupportedVersion && msg.SyncHdr.VerDTD != SyncMLMaxSupportedVersion {
		return errors.New("unsupported proto version")
	}

	//SyncML SessionID check
	if !validator.MinRunes(msg.SyncHdr.SessionID, 1) {
		return errors.New("sessionID")
	}

	//SyncML MsgID check
	if !validator.MinRunes(msg.SyncHdr.MsgID, 1) {
		return errors.New("MsgID")
	}

	//Target LocURI check
	if !validator.HasContent(*msg.SyncHdr.Target.LocURI) {
		return errors.New("Target.LocURI")
	}

	//Device ID check
	if !validator.HasContent(*msg.SyncHdr.Source.LocURI) {
		return errors.New("Source.LocURI")
	}

	return nil

}

func (msg *SyncML) isValidBody() error {
	nonNilCount := 0

	if msg.SyncBody.Add != nil {
		nonNilCount++
	}

	if msg.SyncBody.Alert != nil {
		nonNilCount++
	}

	if msg.SyncBody.Atomic != nil {
		nonNilCount++
	}

	if msg.SyncBody.Delete != nil {
		nonNilCount++
	}

	if msg.SyncBody.Exec != nil {
		nonNilCount++
	}

	if msg.SyncBody.Get != nil {
		nonNilCount++
	}

	if msg.SyncBody.Replace != nil {
		nonNilCount++
	}

	if msg.SyncBody.Status != nil {
		nonNilCount++
	}

	if msg.SyncBody.Results != nil {
		nonNilCount++
	}

	if msg.SyncBody.Raw != nil {
		nonNilCount++
	}

	if nonNilCount == 0 {
		return errors.New("no SyncML protocol commands")
	}

	return nil
}

// IsValidMsg checks for required fields in the SyncML message
func (msg *SyncML) IsValidMsg() error {

	if err := msg.isValidHeader(); err != nil {
		return fmt.Errorf("invalid SyncML header: %s", err)
	}

	if err := msg.isValidBody(); err != nil {
		return fmt.Errorf("invalid SyncML body: %s", err)
	}

	return nil
}

func (msg *SyncML) IsFinal() bool {
	if (msg.SyncBody.Final != nil) && !validator.HasContent(*msg.SyncBody.Final) {
		return true
	}

	return false
}

func (msg *SyncML) GetMessageID() (string, error) {
	if validator.HasContent(msg.SyncHdr.MsgID) {
		return msg.SyncHdr.MsgID, nil
	}

	return "", errors.New("message id is empty")
}

func (msg *SyncML) GetSessionID() (string, error) {
	if validator.HasContent(msg.SyncHdr.SessionID) {
		return msg.SyncHdr.SessionID, nil
	}

	return "", errors.New("session id is empty")
}

func (msg *SyncML) GetSource() (string, error) {
	if (msg.SyncHdr.Source != nil) &&
		(msg.SyncHdr.Source.LocURI != nil) &&
		validator.HasContent(*msg.SyncHdr.Source.LocURI) {

		return *msg.SyncHdr.Source.LocURI, nil
	}

	return "", errors.New("message source is empty")
}

func (msg *SyncML) GetTarget() (string, error) {
	if (msg.SyncHdr.Target != nil) &&
		(msg.SyncHdr.Target.LocURI != nil) &&
		validator.HasContent(*msg.SyncHdr.Target.LocURI) {

		return *msg.SyncHdr.Target.LocURI, nil
	}

	return "", errors.New("message target is empty")
}

// AppendAddCommand appends a SyncML command to the Raw command list
func (msg *SyncML) AppendCommand(cmd *SyncMLCmd) {
	if msg.SyncBody.Raw == nil {
		msg.SyncBody.Raw = &[]SyncMLCmd{}
	}

	if cmd != nil {
		*msg.SyncBody.Raw = append(*msg.SyncBody.Raw, *cmd)
	}
}

// AppendAddCommand appends a SyncML command to the Add command list
func (msg *SyncML) AppendAddCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Add == nil {
		msg.SyncBody.Add = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Add = append(*msg.SyncBody.Add, cmd)
}

// AppendAlertCommand appends a SyncML command to the Alert command list
func (msg *SyncML) AppendAlertCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Alert == nil {
		msg.SyncBody.Alert = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Alert = append(*msg.SyncBody.Alert, cmd)
}

// AppendAtomicCommand appends a SyncML command to the Atomic command list
func (msg *SyncML) AppendAtomicCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Atomic == nil {
		msg.SyncBody.Atomic = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Atomic = append(*msg.SyncBody.Atomic, cmd)
}

// AppendDeleteCommand appends a SyncML command to the Delete command list
func (msg *SyncML) AppendDeleteCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Delete == nil {
		msg.SyncBody.Delete = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Delete = append(*msg.SyncBody.Delete, cmd)
}

// AppendExecCommand appends a SyncML command to the Exec command list
func (msg *SyncML) AppendExecCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Exec == nil {
		msg.SyncBody.Exec = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Exec = append(*msg.SyncBody.Exec, cmd)
}

// AppendGetCommand appends a SyncML command to the Get command list
func (msg *SyncML) AppendGetCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Get == nil {
		msg.SyncBody.Get = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Get = append(*msg.SyncBody.Get, cmd)
}

// AppendReplaceCommand appends a SyncML command to the Replace command list
func (msg *SyncML) AppendReplaceCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Replace == nil {
		msg.SyncBody.Replace = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Replace = append(*msg.SyncBody.Replace, cmd)
}

// AppendResultsCommand appends a SyncML command to the Results command list
func (msg *SyncML) AppendResultsCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Results == nil {
		msg.SyncBody.Results = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Results = append(*msg.SyncBody.Results, cmd)
}

// AppendStatusCommand appends a SyncML command to the Status command list
func (msg *SyncML) AppendStatusCommand(cmd SyncMLCmd) {
	if msg.SyncBody.Status == nil {
		msg.SyncBody.Status = &[]SyncMLCmd{}
	}

	*msg.SyncBody.Status = append(*msg.SyncBody.Status, cmd)
}

func (msg *SyncML) SetID(cmdID int) {
	msg.SyncHdr.MsgID = strconv.Itoa(cmdID)
}

func (msg *SyncML) GetOrderedCmds() []ProtoCmdOperation {

	// Returns the commands in the order they are defined in the message
	var cmds []ProtoCmdOperation

	// Hash with index as key and ProtoCmdOperation as value
	// Index is the order of the commands in the spec
	// Value is the actual ProtoCmdOperation
	var cmdHash = make(map[string]ProtoCmdOperation)

	if msg.SyncBody.Add != nil {
		for _, cmd := range *msg.SyncBody.Add {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdAdd, Cmd: cmd}
		}
	}

	if msg.SyncBody.Alert != nil {
		for _, cmd := range *msg.SyncBody.Alert {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdAlert, Cmd: cmd}
		}
	}

	if msg.SyncBody.Atomic != nil {
		for _, cmd := range *msg.SyncBody.Atomic {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdAtomic, Cmd: cmd}
		}
	}

	if msg.SyncBody.Delete != nil {
		for _, cmd := range *msg.SyncBody.Delete {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdDelete, Cmd: cmd}
		}
	}

	if msg.SyncBody.Exec != nil {
		for _, cmd := range *msg.SyncBody.Exec {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdExec, Cmd: cmd}
		}
	}

	if msg.SyncBody.Get != nil {
		for _, cmd := range *msg.SyncBody.Get {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdGet, Cmd: cmd}
		}
	}

	if msg.SyncBody.Replace != nil {
		for _, cmd := range *msg.SyncBody.Replace {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdReplace, Cmd: cmd}
		}
	}

	if msg.SyncBody.Results != nil {
		for _, cmd := range *msg.SyncBody.Results {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdResults, Cmd: cmd}
		}
	}

	if msg.SyncBody.Status != nil {
		for _, cmd := range *msg.SyncBody.Status {
			cmdHash[cmd.CmdID] = ProtoCmdOperation{Verb: CmdStatus, Cmd: cmd}
		}
	}

	// Get minimum key value from cmdHash
	minIndex := 10
	for k := range cmdHash {
		i, err := strconv.Atoi(k)
		if err != nil {
			continue
		}

		if minIndex > i {
			minIndex = i
		}
	}

	// Traverse the hash and return the commands in the order they are defined in the spec
	for i := minIndex; i < minIndex+len(cmdHash); i++ {
		val, ok := cmdHash[strconv.Itoa(i)]
		if ok {
			cmds = append(cmds, val)
		}
	}

	return cmds
}

func (cmd *SyncMLCmd) IsValid() bool {

	if ((cmd.Items == nil) || (len(*cmd.Items) == 0)) && cmd.Data == nil {
		return false
	}

	return true
}
