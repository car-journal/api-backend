package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalcarcontroller "github.com/car-journal/api-backend/transport/http/controller/internals/car"
	internalfuelentrycontroller "github.com/car-journal/api-backend/transport/http/controller/internals/fuelentry"
	internalodometerentrycontroller "github.com/car-journal/api-backend/transport/http/controller/internals/odometerentry"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Car(internalRoute *mux.Router, app container.AppContainer) {
	internalCarRoutes := internalRoute.PathPrefix(routeconst.CarPrefix).Subrouter()
	internalCarRoutes.HandleFunc("", internalcarcontroller.Create(app.Services.Car)).Methods(http.MethodPost).Name("internalcarcontroller.Create")
	internalCarRoutes.HandleFunc("", internalcarcontroller.List(app.Services.Car)).Methods(http.MethodGet).Name("internalcarcontroller.List")
	internalCarRoutes.HandleFunc("/{car_id}", internalcarcontroller.FindByID(app.Services.Car)).Methods(http.MethodGet).Name("internalcarcontroller.FindByID")
	internalCarRoutes.HandleFunc("/{car_id}/fuel-entries", internalfuelentrycontroller.FuelEntryListByCarID(app.Services.FuelEntry)).Methods(http.MethodGet).Name("internalfuelentrycontroller.FuelEntryListByCarID")
	internalCarRoutes.HandleFunc("/{car_id}/odometer-entries", internalodometerentrycontroller.FuelEntryListByCarID(app.Services.OdometerEntry)).Methods(http.MethodGet).Name("internalfuelentrycontroller.FuelEntryListByCarID")
	internalCarRoutes.HandleFunc("/{car_id}", internalcarcontroller.Update(app.Services.Car)).Methods(http.MethodPatch).Name("internalcarcontroller.Update")
	internalCarRoutes.HandleFunc("/{car_id}", internalcarcontroller.Delete(app.Services.Car)).Methods(http.MethodDelete).Name("internalcarcontroller.Delete")
}
