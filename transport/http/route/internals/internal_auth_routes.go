package internalroute

import (
	"net/http"

	"github.com/car-journal/transport/container"
	internalauthcontroller "github.com/car-journal/transport/http/controller/internals/auth"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Auth(internalRoute *mux.Router, app container.AppContainer) {
	internalAuthRoutes := internalRoute.PathPrefix(routeconst.AUTH_PREFIX).Subrouter()
	internalAuthRoutes.HandleFunc("/update-password", internalauthcontroller.UpdatePassword(app.Services.Auth)).Methods(http.MethodPatch)
	internalAuthRoutes.HandleFunc("/me", internalauthcontroller.Me(app.Services.Auth)).Methods(http.MethodGet)
}
