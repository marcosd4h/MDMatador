package response

import (
	"net/http"
	"strconv"
)

func RawResponse(w http.ResponseWriter, status int, headers http.Header, contentType string, rawBytes []byte) error {

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(rawBytes)))
	w.WriteHeader(status)
	_, err := w.Write(rawBytes)
	if err != nil {
		return err
	}

	return nil
}
