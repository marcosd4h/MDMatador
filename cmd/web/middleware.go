package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/justinas/nosurf"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) preventCSRF(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})

	return csrfHandler
}

func (app *application) authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := app.sessionStore.Get(r, "session")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		userID, ok := session.Values["userID"].(int)
		if ok {
			user, err := app.db.GetUser(userID)
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			if user != nil {
				r = contextSetAuthenticatedUser(r, user)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticatedUser := contextGetAuthenticatedUser(r)

		if authenticatedUser == nil {
			session, err := app.sessionStore.Get(r, "session")
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			session.Values["redirectPathAfterLogin"] = r.URL.Path

			err = session.Save(r, w)
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAnonymousUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticatedUser := contextGetAuthenticatedUser(r)

		if authenticatedUser != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) MDMDeviceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceID := chi.URLParam(r, "DeviceID")
		if deviceID == "" {
			app.serverError(w, r, errors.New("mdmDeviceInfo middleware: device not found"))
			return
		}

		device, err := app.db.MDMGetEnrolledDevice(deviceID)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		r = contextSetMDMDevice(r, device)
		next.ServeHTTP(w, r)
	})
}

// drainBody reads all of bytes to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
//
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, body []byte, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, nil, err
	}
	if err = b.Close(); err != nil {
		return nil, b, nil, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), buf.Bytes(), nil
}

func (app *application) customLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		shouldLog := strings.HasPrefix(r.URL.Path, "/EnrollmentServer") || strings.HasPrefix(r.URL.Path, "/ManagementServer")

		if !shouldLog {
			// Skip logging, call next handler
			next.ServeHTTP(w, r)
			return
		}

		// grabbing Input Header and Body
		reqHeader, err := httputil.DumpRequest(r, false)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		var bodyBytes []byte
		reqBodySave := r.Body
		if r.Body != nil {
			reqBodySave, r.Body, bodyBytes, err = drainBody(r.Body)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		}
		r.Body = reqBodySave

		var reqBody string
		if len(bodyBytes) > 0 {
			reqBody = xmlfmt.FormatXML(string(bodyBytes), " ", "  ")
		}

		fmt.Printf("\n\n============================= Input Request =============================\n")
		fmt.Printf("Timestamp: %s\n", time.Now().Format(time.RFC822))
		fmt.Printf("%s\n", reqHeader)
		fmt.Printf("%s\n", reqBody)
		fmt.Printf("=========================================================================\n")

		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		// grabbing Output Header and Body
		rawResBody := rec.Body.Bytes()

		var resBody string
		if len(rawResBody) > 0 {
			resBody = xmlfmt.FormatXML(string(rawResBody), " ", "  ")
		}

		resHeader, err := httputil.DumpResponse(rec.Result(), false)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		fmt.Printf("\n\n============================= Output Response =============================\n")
		fmt.Printf("Timestamp: %s\n", time.Now().Format(time.RFC822))
		fmt.Printf("%s\n", resHeader)
		fmt.Printf("%s\n", resBody)
		fmt.Printf("=========================================================================\n")

		// we copy the captured response headers to our new response
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		w.Write(rec.Body.Bytes())
	})
}
