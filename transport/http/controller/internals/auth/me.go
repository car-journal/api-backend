// Package internalauthcontroller handles controller for auth domain
package internalauthcontroller

import (
	"context"
	"net/http"

	authdto "github.com/car-journal/api-backend/internal/domain/auth/dto"
	authservice "github.com/car-journal/api-backend/internal/domain/auth/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
)

func Me(authService authservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var response authdto.Me
		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			response, err = authService.Me(ctx)
			if err != nil {
				return err
			}
			return nil
		}); errTrans != nil {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, response, nil)
	}
}
