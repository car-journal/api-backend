package externalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	externalauthcontroller "github.com/car-journal/api-backend/transport/http/controller/external/auth"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func Auth(externalRoute *mux.Router, app container.AppContainer) {
	externalAuthRoutes := externalRoute.PathPrefix(routeconst.AuthPrefix).Subrouter()
	externalAuthRoutes.HandleFunc("/register", externalauthcontroller.Register(app.Services.Auth)).Methods(http.MethodPost).Name("externalauthcontroller.Register")
	externalAuthRoutes.HandleFunc("/login", externalauthcontroller.Login(app.Services.Auth)).Methods(http.MethodPost).Name("externalauthcontroller.Login")
}
