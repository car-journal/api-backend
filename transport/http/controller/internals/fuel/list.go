// Package internalfuelcontroller handles controller for fuel domain
package internalfuelcontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	fuelservice "github.com/car-journal/api-backend/internal/domain/fuel/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
)

func List(fuelService fuelservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var fuels []*internalmodel.Fuel
		var pageResponse *filter.PageResponse
		name := request.URL.Query().Get("name")
		pageParams := filter.ParsePage(request.URL.Query(), config.DefaultLimit)

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			fuels, err = fuelService.List(ctx, name)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, fuels)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
