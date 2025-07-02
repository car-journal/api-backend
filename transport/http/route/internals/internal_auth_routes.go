package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalauthcontroller "github.com/car-journal/api-backend/transport/http/controller/internals/auth"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Auth(internalRoute *mux.Router, app container.AppContainer) {
	internalAuthRoutes := internalRoute.PathPrefix(routeconst.AuthPrefix).Subrouter()
	internalAuthRoutes.HandleFunc("/update-password", internalauthcontroller.UpdatePassword(app.Services.Auth)).Methods(http.MethodPatch).Name("internalauthcontroller.UpdatePassword")
	internalAuthRoutes.HandleFunc("/me", internalauthcontroller.Me(app.Services.Auth)).Methods(http.MethodGet).Name("internalauthcontroller.Me")
}
