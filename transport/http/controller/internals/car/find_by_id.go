package internalcarcontroller

import (
	"context"
	"net/http"

	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByID(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var cars *internalmodel.Car
		carID := mux.Vars(request)["car_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			cars, err = carService.FindByID(ctx, carID, me.ID)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, cars, nil)
	}
}
