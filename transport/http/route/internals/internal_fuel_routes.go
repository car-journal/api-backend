package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalfuelcontroller "github.com/car-journal/api-backend/transport/http/controller/internals/fuel"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Fuel(internalRoute *mux.Router, app container.AppContainer) {
	internalCarRoutes := internalRoute.PathPrefix(routeconst.FuelPrefix).Subrouter()
	internalCarRoutes.HandleFunc("", internalfuelcontroller.List(app.Services.Fuel)).Methods(http.MethodGet).Name("internalfuelcontroller.List")
}
