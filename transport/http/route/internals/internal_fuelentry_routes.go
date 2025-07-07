package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalfuelentrycontroller "github.com/car-journal/api-backend/transport/http/controller/internals/fuelentry"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func FuelEntry(internalRoute *mux.Router, app container.AppContainer) {
	internalCarRoutes := internalRoute.PathPrefix(routeconst.FuelEntryPrefix).Subrouter()
	internalCarRoutes.HandleFunc("", internalfuelentrycontroller.Create(app.Services.FuelEntry)).Methods(http.MethodPost).Name("internalfuelentrycontroller.Create")
	internalCarRoutes.HandleFunc("/{fuel_entry_id}", internalfuelentrycontroller.FindByID(app.Services.FuelEntry)).Methods(http.MethodGet).Name("internalfuelentrycontroller.FindByID")
	internalCarRoutes.HandleFunc("/{fuel_entry_id}", internalfuelentrycontroller.Update(app.Services.FuelEntry)).Methods(http.MethodPatch).Name("internalfuelentrycontroller.Update")
	internalCarRoutes.HandleFunc("/{fuel_entry_id}", internalfuelentrycontroller.Delete(app.Services.FuelEntry)).Methods(http.MethodDelete).Name("internalfuelentrycontroller.Delete")
}
