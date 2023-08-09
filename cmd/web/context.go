package main

import (
	"context"
	"net/http"

	"github.com/marcosd4h/MDMatador/internal/database"
	"github.com/marcosd4h/MDMatador/internal/mdm"
)

type contextKey string

const (
	authenticatedUserContextKey = contextKey("authenticatedUser")
	mdmDeviceContextKey         = contextKey("device")
)

func contextSetAuthenticatedUser(r *http.Request, user *database.User) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedUserContextKey, user)
	return r.WithContext(ctx)
}

func contextGetAuthenticatedUser(r *http.Request) *database.User {
	user, ok := r.Context().Value(authenticatedUserContextKey).(*database.User)
	if !ok {
		return nil
	}

	return user
}

func contextSetMDMDevice(r *http.Request, device *mdm.MDMWindowsEnrolledDevice) *http.Request {
	ctx := context.WithValue(r.Context(), mdmDeviceContextKey, device)
	return r.WithContext(ctx)
}
