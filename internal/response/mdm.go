package response

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/marcosd4h/MDMatador/internal/mdm"
)

func SOAPResponse(w http.ResponseWriter, status int, relatesTo string, data any) error {
	return SOAPResponseWithHeaders(w, status, relatesTo, data, nil)
}

func SOAPResponseWithHeaders(w http.ResponseWriter, status int, relatesTo string, data any, headers http.Header) error {

	res, err := mdm.NewSoapResponse(data, relatesTo)
	if err != nil {
		return err
	}

	xmlRes, err := xml.MarshalIndent(res, "", "\t")
	if err != nil {
		return err
	}

	xmlRes = append(xmlRes, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", mdm.SoapMsgContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(xmlRes)))
	w.WriteHeader(status)
	w.Write(xmlRes)

	return nil
}

func SyncMLResponse(w http.ResponseWriter, status int, msg *mdm.SyncML) error {
	return SyncMLResponseWithHeaders(w, status, msg, nil)
}

func SyncMLResponseWithHeaders(w http.ResponseWriter, status int, msg *mdm.SyncML, headers http.Header) error {

	if msg == nil {
		return errors.New("syncmlresponse: message is nil")
	}

	err := msg.IsValidMsg()
	if err != nil {
		return fmt.Errorf("syncmlresponse: invald msg %v", err)
	}

	rawMsg, err := xml.MarshalIndent(msg, " ", "  ")
	if err != nil {
		return fmt.Errorf("syncmlresponse: message marshalling error %v", err)
	}

	xmlRes := xml.Header + string(rawMsg) + "\n"

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", mdm.SyncMLMsgContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(xmlRes)))
	w.WriteHeader(status)
	w.Write([]byte(xmlRes))

	return nil
}

func DumpXML(v any) {
	rawMsg, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("syncmlresponse: message marshalling error %v", err)
	}

	fmt.Println(string(rawMsg) + "\n")
}
