package internalcarcontroller

import (
	"context"
	"net/http"

	cardto "github.com/car-journal/api-backend/internal/domain/car/dto"
	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByIDWithFuelSummary(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var cars *cardto.CarWithFuelSummary
		carID := mux.Vars(request)["car_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			cars, err = carService.FindByIDWithFuelSummary(ctx, carID, me.ID)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, cars, nil)
	}
}
