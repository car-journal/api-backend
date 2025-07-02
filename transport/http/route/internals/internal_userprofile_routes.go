package internalroute

import (
	"net/http"

	"github.com/car-journal/api-backend/transport/container"
	internalusercontroller "github.com/car-journal/api-backend/transport/http/controller/internals/user"
	routeconst "github.com/car-journal/api-backend/transport/http/route/const"
	"github.com/gorilla/mux"
)

func UserProfile(internalRoute *mux.Router, app container.AppContainer) {
	internalUserProfileRoutes := internalRoute.PathPrefix(routeconst.UserProfilePrefix).Subrouter()
	internalUserProfileRoutes.HandleFunc("", internalusercontroller.Update(app.Services.UserProfile)).Methods(http.MethodPatch).Name("internalusercontroller.Update")
}
