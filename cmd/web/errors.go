package main

import (
	"net/http"

	"github.com/marcosd4h/MDMatador/internal/mdm"
	"github.com/marcosd4h/MDMatador/internal/response"

	"golang.org/x/exp/slog"
)

func (app *application) reportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()

		//TODO revert this when needed
		//trace   = string(debug.Stack())
		trace = ""
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	http.Error(w, message, http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	http.Error(w, message, http.StatusNotFound)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func (app *application) serverSoapError(w http.ResponseWriter, errorType mdm.SoapError, origMessage mdm.MDEMessageType, messageID string, err error) {
	soapFault := mdm.NewSoapFault(errorType, origMessage, err)
	err = response.SOAPResponse(w, http.StatusInternalServerError, messageID, soapFault)
	if err != nil {
		// TODO: check if logging is needed
		//app.reportError(err)
		//app.rawResponse(w, r, err)

		return
	}
}

func (app *application) rawResponse(w http.ResponseWriter, r *http.Request, rawData string) {
	http.Error(w, rawData, http.StatusOK)
}

func (app *application) serverSyncMLError(w http.ResponseWriter, err error) {
	//app.reportError(err)

	//TODO: check if this is the correct way to handle errors
	/*
		err = response.SyncMLResponse(w, http.StatusInternalServerError, "error")
		if err != nil {
			app.reportError(err)
			return
		}
	*/
}
