package externalroute

import (
	"github.com/car-journal/transport/container"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func External(v1Routes *mux.Router, app container.AppContainer) {
	externalRoutes := v1Routes.PathPrefix(routeconst.EXTERNAL_PREFIX).Subrouter()

	Auth(externalRoutes, app)
}
