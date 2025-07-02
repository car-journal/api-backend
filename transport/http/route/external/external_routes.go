// Package externalroute handles external routes
package externalroute

import (
	"github.com/car-journal/api-backend/transport/container"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func External(v1Routes *mux.Router, app container.AppContainer) {
	externalRoutes := v1Routes.PathPrefix(routeconst.ExternalPrefix).Subrouter()

	Auth(externalRoutes, app)
}
