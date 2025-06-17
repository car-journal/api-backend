package internalroute

import (
	"net/http"

	"github.com/car-journal/transport/container"
	internalusercontroller "github.com/car-journal/transport/http/controller/internals/user"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func UserProfile(internalRoute *mux.Router, app container.AppContainer) {
	internalUserProfileRoutes := internalRoute.PathPrefix(routeconst.USER_PROFILES_PREFIX).Subrouter()
	internalUserProfileRoutes.HandleFunc("", internalusercontroller.Update(app.Services.UserProfile)).Methods(http.MethodPatch)
}
