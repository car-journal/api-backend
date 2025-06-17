package internalauthcontroller

import (
	"context"
	"net/http"

	authdto "github.com/car-journal/internal/domain/auth/dto"
	authservice "github.com/car-journal/internal/domain/auth/service"
	"github.com/car-journal/lib/database"
	"github.com/car-journal/lib/parser"
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
			parser.Json(writer, nil, errTrans)
			return
		}
		parser.Json(writer, response, nil)
	}
}
