package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalmaintenanceentrycontroller "github.com/car-journal/api-backend/transport/http/controller/internals/maintenanceentry"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func MaintenanceEntry(internalRoute *mux.Router, app container.AppContainer) {
	maintenanceEntryRoutes := internalRoute.PathPrefix(routeconst.MaintenanceEntryPrefix).Subrouter()
	maintenanceEntryRoutes.HandleFunc("/{car_id}", internalmaintenanceentrycontroller.Create(app.Services.MaintenanceEntry)).Methods(http.MethodPost).Name("internalmaintenanceentrycontroller.Create")
	maintenanceEntryRoutes.HandleFunc("/{car_id}", internalmaintenanceentrycontroller.List(app.Services.MaintenanceEntry)).Methods(http.MethodGet).Name("internalmaintenanceentrycontroller.List")
	maintenanceEntryRoutes.HandleFunc("/{id}", internalmaintenanceentrycontroller.FindByID(app.Services.MaintenanceEntry)).Methods(http.MethodGet).Name("internalmaintenanceentrycontroller.FindByID")
	maintenanceEntryRoutes.HandleFunc("/{id}", internalmaintenanceentrycontroller.Update(app.Services.MaintenanceEntry)).Methods(http.MethodPatch).Name("internalmaintenanceentrycontroller.Update")
	maintenanceEntryRoutes.HandleFunc("/{id}", internalmaintenanceentrycontroller.Delete(app.Services.MaintenanceEntry)).Methods(http.MethodDelete).Name("internalmaintenanceentrycontroller.Delete")
}
