package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalmaintenancecategorycontroller "github.com/car-journal/api-backend/transport/http/controller/internals/maintenancecategory"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func MaintenanceCategory(internalRoute *mux.Router, app container.AppContainer) {
	maintenanceCategoryRoutes := internalRoute.PathPrefix(routeconst.MaintenanceCategoryPrefix).Subrouter()
	maintenanceCategoryRoutes.HandleFunc("", internalmaintenancecategorycontroller.Create(app.Services.MaintenanceCategory)).Methods(http.MethodPost).Name("internalmaintenancecategorycontroller.Create")
	maintenanceCategoryRoutes.HandleFunc("", internalmaintenancecategorycontroller.List(app.Services.MaintenanceCategory)).Methods(http.MethodGet).Name("internalmaintenancecategorycontroller.List")
	maintenanceCategoryRoutes.HandleFunc("/{id}", internalmaintenancecategorycontroller.FindByID(app.Services.MaintenanceCategory)).Methods(http.MethodGet).Name("internalmaintenancecategorycontroller.FindByID")
	maintenanceCategoryRoutes.HandleFunc("/{id}", internalmaintenancecategorycontroller.Update(app.Services.MaintenanceCategory)).Methods(http.MethodPatch).Name("internalmaintenancecategorycontroller.Update")
	maintenanceCategoryRoutes.HandleFunc("/{id}", internalmaintenancecategorycontroller.Delete(app.Services.MaintenanceCategory)).Methods(http.MethodDelete).Name("internalmaintenancecategorycontroller.Delete")
}
