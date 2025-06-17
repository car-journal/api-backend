package internalroute

import (
	"github.com/car-journal/transport/container"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Internal(v1Routes *mux.Router, app container.AppContainer) {
	internalRoutes := v1Routes.PathPrefix(routeconst.INTERNAL_PREFIX).Subrouter()
	internalRoutes.Use(
		app.Middleware.Authentication(app.Services.Auth, app.Services.User),
	)

	Auth(internalRoutes, app)
	UserProfile(internalRoutes, app)
}
