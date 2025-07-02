// Package internalroute handles internal routes
package internalroute

import (
	"github.com/car-journal/api-backend/transport/container"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Internal(v1Routes *mux.Router, app container.AppContainer) {
	internalRoutes := v1Routes.PathPrefix(routeconst.InternalPrefix).Subrouter()
	internalRoutes.Use(
		app.Middleware.Authentication(app.Services.Auth, app.Services.User),
	)

	Auth(internalRoutes, app)
	UserProfile(internalRoutes, app)
	Car(internalRoutes, app)
}
