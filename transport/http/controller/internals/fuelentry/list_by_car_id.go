package internalfuelentrycontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	fuelentryservice "github.com/car-journal/api-backend/internal/domain/fuelentry/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FuelEntryListByCarID(fuelEntryService fuelentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var fuelEntries []*internalmodel.FuelEntry
		var pageResponse *filter.PageResponse
		carID := mux.Vars(request)["car_id"]
		pageParams := filter.ParsePage(request.URL.Query(), config.DefaultLimit)

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			fuelEntries, err = fuelEntryService.ListByCarID(ctx, carID, me.ID, pageParams)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, fuelEntries)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
