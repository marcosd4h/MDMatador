package main

import (
	"net/http"

	"github.com/marcosd4h/MDMatador/assets"
	"github.com/marcosd4h/MDMatador/internal/mdm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.NotFound(app.notFound)

	// Middlewares
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(app.recoverPanic)
	mux.Use(app.securityHeaders)

	if app.config.debug.http {
		mux.Use(app.customLogger)
	}

	// Handling static files
	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handle("/static/*", fileServer)

	// Handling Webapp routes
	mux.Group(func(mux chi.Router) {

		//TODO uncomment this when release is ready
		//mux.Use(app.preventCSRF)

		mux.Use(app.authenticateMiddleware)

		mux.Group(func(mux chi.Router) {
			mux.Use(app.requireAnonymousUser)

			// User Management Routes
			mux.Group(func(mux chi.Router) {
				mux.Get("/signup", app.signupHandler)
				mux.Post("/signup", app.signupHandler)
				mux.Get("/login", app.loginHandler)
				mux.Post("/login", app.loginHandler)
			})

			//MS-MDE2 and MS-MDM Protocols Routes
			mux.Group(func(mux chi.Router) {
				mux.Get(mdm.MSMDE2_DiscoveryPath, app.discoveryHandler)
				mux.Post(mdm.MSMDE2_DiscoveryPath, app.discoveryHandler)
				mux.Post(mdm.MSMDE2_PolicyPath, app.policyHandler)
				mux.Get(mdm.MSMDE2_AuthPath, app.authHandler)
				mux.Post(mdm.MSMDE2_EnrollPath, app.enrollHandler)
				mux.Post(mdm.MSMDM_ManagementPath, app.managementHandler)
			})

			// Web Socket Routes
			mux.Get("/ws", app.websocketHandler)
		})

		mux.Group(func(mux chi.Router) {
			mux.Use(app.requireAuthenticatedUser)

			// General Routes
			mux.Get("/", app.dashboardHandler)
			mux.Post("/logout", app.logoutHandler)

			// MDM Routes
			mux.Group(func(mux chi.Router) {
				mux.Use(app.MDMDeviceID)

				// Dashboard Routes
				mux.Get("/mdm/device/{DeviceID}", app.getDeviceInfoHandler)
				mux.Get("/mdm/terminal/{DeviceID}", app.getDeviceTerminalHandler)

				// API Routes
				mux.Route("/api", func(r chi.Router) {

					// MDM Device APIs
					mux.Delete("/api/mdm/device/{DeviceID}", app.deviceManagementHandler)
					mux.Post("/api/mdm/device/{DeviceID}", app.deviceManagementHandler)
				})
			})
		})
	})

	return mux
}
