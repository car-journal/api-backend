package adminroute

import (
	"net/http"

	"github.com/car-journal/transport/container"
	admincontroller "github.com/car-journal/transport/http/controller/admin"
	routeconst "github.com/car-journal/transport/http/route/const"
	"github.com/gorilla/mux"
)

func AdminRoutes(v1Routes *mux.Router, app container.AppContainer, baseRoute *mux.Router) {
	adminRoutes := v1Routes.PathPrefix(routeconst.ADMIN_PREFIX).Subrouter()
	// adminRoutes.Use(
	// 	app.Middleware.Authentication(app.Services.Auth, app.Services.User),
	// )
	adminRoutes.HandleFunc("/get-all-routes", admincontroller.GetAllRoutes(baseRoute)).Methods(http.MethodGet)
}
