package internalcarcontroller

import (
	"context"
	"net/http"

	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func Delete(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		carID := mux.Vars(request)["car_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			return carService.Delete(ctx, carID)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
