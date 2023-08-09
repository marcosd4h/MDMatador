package request

import (
	"bytes"
	"io"
	"net/http"

	"github.com/marcosd4h/MDMatador/internal/mdm"
)

func getBodyBytes(req *http.Request) (*bytes.Buffer, error) {
	var bodyBytes bytes.Buffer
	_, err := io.Copy(&bodyBytes, req.Body)

	return &bodyBytes, err
}

func SoapRequest(req *http.Request) (*mdm.SoapRequest, error) {

	bodyBytes, err := getBodyBytes(req)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	// Parsing the Soap Request
	soapReq, err := mdm.NewSoapRequest(bodyBytes.Bytes())
	if err != nil {
		return nil, err
	}

	return soapReq, nil
}

func SyncMLRequest(req *http.Request) (*mdm.SyncML, error) {

	bodyBytes, err := getBodyBytes(req)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	// Parsing the SyncML request
	syncMLReq, err := mdm.NewSyncMLFromRequest(bodyBytes.Bytes())
	if err != nil {
		return nil, err
	}

	return syncMLReq, nil
}
