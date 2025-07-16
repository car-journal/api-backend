// Package dbmiddleware handles db context creation
package dbmiddleware

import (
	"net/http"

	"github.com/car-journal/api-backend/config"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/gorilla/mux"
)

func PostgresDB() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := database.Set(r.Context(), config.ConnectGormPostgres())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
