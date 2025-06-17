package container

import (
	authservice "github.com/car-journal/internal/domain/auth/service"
	userservice "github.com/car-journal/internal/domain/user/service"
	dbmiddleware "github.com/car-journal/transport/http/middleware/db"
	securitymiddleware "github.com/car-journal/transport/http/middleware/security"
	"github.com/gorilla/mux"
)

// MiddlewareContainer handle all middleware used in project
type MiddlewareContainer struct {
	Authentication func(authService authservice.Interface, userService userservice.Interface) mux.MiddlewareFunc
	PostgresDB     func() mux.MiddlewareFunc
}

// CreateMiddlewareContainer construct all middlewares used in the app
func CreateMiddlewareContainer() MiddlewareContainer {
	return MiddlewareContainer{
		Authentication: securitymiddleware.Authentication,
		PostgresDB:     dbmiddleware.PostgresDB,
	}
}
