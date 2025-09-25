package internalcarcontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	cardto "github.com/car-journal/api-backend/internal/domain/car/dto"
	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
)

func ListCarsWithAverageFuelConsumptionRate(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var cars []*cardto.CarWithAverageFuelConsumptionRate
		var pageResponse *filter.PageResponse
		pageParams := filter.ParsePage(request.URL.Query(), config.DefaultLimit)

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			cars, err = carService.ListCarsWithAverageFuelConsumptionRate(ctx, me.ID)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, cars)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
