// Package internalodometerentrycontroller handles controller for odometer entry domain
package internalodometerentrycontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	odometerentrydto "github.com/car-journal/api-backend/internal/domain/odometerentry/dto"
	odometerentryservice "github.com/car-journal/api-backend/internal/domain/odometerentry/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FuelEntryListByCarID(odometerEntry odometerentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var odometerEntries []*odometerentrydto.OdometerEntrySummary
		var pageResponse *filter.PageResponse
		carID := mux.Vars(request)["car_id"]
		pageParams := filter.ParsePage(request.URL.Query(), config.DefaultLimit)

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			odometerEntries, err = odometerEntry.ListByCarID(ctx, carID, me.ID)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, odometerEntries)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
