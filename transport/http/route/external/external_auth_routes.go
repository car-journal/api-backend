package externalroute

import (
	"net/http"

	"github.com/car-journal/transport/container"
	externalauthcontroller "github.com/car-journal/transport/http/controller/external/auth"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Auth(externalRoute *mux.Router, app container.AppContainer) {
	externalAuthRoutes := externalRoute.PathPrefix(routeconst.AUTH_PREFIX).Subrouter()
	externalAuthRoutes.HandleFunc("/login", externalauthcontroller.Login(app.Services.Auth)).Methods(http.MethodPost)
}
