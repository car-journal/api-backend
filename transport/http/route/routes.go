// Package route handles all routes
package route

import (
	"log"
	"net/http"

	"github.com/car-journal/api-backend/lib/parser"
	"github.com/car-journal/api-backend/transport/container"
	adminroute "github.com/car-journal/api-backend/transport/http/route/admin"
	externalroute "github.com/car-journal/api-backend/transport/http/route/external"
	internalroute "github.com/car-journal/api-backend/transport/http/route/internals"
	"github.com/gorilla/mux"
)

func Route(app container.AppContainer) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("ping")
		parser.Json(writer, nil, nil)
	}).Methods(http.MethodGet)

	v1 := r.PathPrefix("/v1").Subrouter()

	// Middlewares
	v1.Use(
		app.Middleware.PostgresDB(),
	)

	// Admin routes
	adminroute.AdminRoutes(v1, app, r)

	// External routes
	externalroute.External(v1, app)

	// Internal routes
	internalroute.Internal(v1, app)

	return r
}
