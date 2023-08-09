package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/marcosd4h/MDMatador/internal/version"
)

func (app *application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"AuthenticatedUser": contextGetAuthenticatedUser(r),
		"CSRFToken":         nosurf.Token(r),
		"Version":           version.Get(),
	}

	return data
}

func (app *application) backgroundTask(r *http.Request, fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.reportServerError(r, fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			app.reportServerError(r, err)
		}
	}()
}
