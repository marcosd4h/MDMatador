package main

import (
	"net/http"

	"github.com/marcosd4h/MDMatador/internal/password"
	"github.com/marcosd4h/MDMatador/internal/request"
	"github.com/marcosd4h/MDMatador/internal/response"
	"github.com/marcosd4h/MDMatador/internal/validator"
)

func (app *application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	// Getting MDM enrolled devices
	devices, err := app.db.MDMGetEnrolledDevices()
	if err != nil {
		app.serverError(w, r, err)
	}
	data["Devices"] = devices

	err = response.Page(w, http.StatusOK, data, "pages/dashboard.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var form struct {
		Email     string              `form:"Email"`
		Password  string              `form:"Password"`
		Validator validator.Validator `form:"-"`
	}

	switch r.Method {
	case http.MethodGet:
		data := app.newTemplateData(r)
		data["Form"] = form

		err := response.RawPage(w, http.StatusOK, data, "signup", "pages/signup.tmpl")
		if err != nil {
			app.serverError(w, r, err)
		}

	case http.MethodPost:
		err := request.DecodePostForm(r, &form)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		existingUser, err := app.db.GetUserByEmail(form.Email)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		form.Validator.CheckField(form.Email != "", "Email", "Email is required")
		form.Validator.CheckField(validator.Matches(form.Email, validator.RgxEmail), "Email", "Must be a valid email address")
		form.Validator.CheckField(existingUser == nil, "Email", "Email is already in use")

		form.Validator.CheckField(form.Password != "", "Password", "Password is required")
		form.Validator.CheckField(len(form.Password) >= 8, "Password", "Password is too short")
		form.Validator.CheckField(len(form.Password) <= 72, "Password", "Password is too long")

		if form.Validator.HasErrors() {
			data := app.newTemplateData(r)
			data["Form"] = form

			err := response.RawPage(w, http.StatusUnprocessableEntity, data, "signup", "pages/signup.tmpl")
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		}

		hashedPassword, err := password.Hash(form.Password)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		id, err := app.db.InsertUser(form.Email, hashedPassword)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		session, err := app.sessionStore.Get(r, "session")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		session.Values["userID"] = id

		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var form struct {
		Email     string              `form:"Email"`
		Password  string              `form:"Password"`
		Validator validator.Validator `form:"-"`
	}

	switch r.Method {
	case http.MethodGet:
		data := app.newTemplateData(r)
		data["Form"] = form

		err := response.RawPage(w, http.StatusOK, data, "login", "pages/login.tmpl")

		if err != nil {
			app.serverError(w, r, err)
		}

	case http.MethodPost:
		err := request.DecodePostForm(r, &form)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		user, err := app.db.GetUserByEmail(form.Email)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		form.Validator.CheckField(form.Email != "", "Email", "Email is required")
		form.Validator.CheckField(user != nil, "Email", "Email address could not be found")

		if user != nil {
			passwordMatches, err := password.Matches(form.Password, user.HashedPassword)
			if err != nil {
				app.serverError(w, r, err)
				return
			}

			form.Validator.CheckField(form.Password != "", "Password", "Password is required")
			form.Validator.CheckField(passwordMatches, "Password", "Password is incorrect")
		}

		if form.Validator.HasErrors() {
			data := app.newTemplateData(r)
			data["Form"] = form

			err := response.RawPage(w, http.StatusUnprocessableEntity, data, "login", "pages/login.tmpl")
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		}

		session, err := app.sessionStore.Get(r, "session")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		session.Values["userID"] = user.ID

		redirectPath, ok := session.Values["redirectPathAfterLogin"].(string)
		if ok {
			delete(session.Values, "redirectPathAfterLogin")
		} else {
			redirectPath = "/"
		}

		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		http.Redirect(w, r, redirectPath, http.StatusSeeOther)
	}
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	delete(session.Values, "userID")

	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
