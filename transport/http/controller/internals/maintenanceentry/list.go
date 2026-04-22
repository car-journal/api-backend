package internalmaintenanceentrycontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	maintenanceentrypayload "github.com/car-journal/api-backend/internal/domain/maintenanceentry/payload"
	maintenanceentryservice "github.com/car-journal/api-backend/internal/domain/maintenanceentry/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
	carjournalstrings "github.com/car-journal/api-backend/lib/strings"
	"github.com/gorilla/mux"
)

func List(maintenanceEntryService maintenanceentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var entries []*internalmodel.MaintenanceEntry
		var pageResponse *filter.PageResponse

		queryParams := request.URL.Query()
		name := queryParams.Get("name")
		categoryIDs := carjournalstrings.Split(queryParams.Get("category_ids"), ",")
		sorts := carjournalstrings.Split(queryParams.Get("sorts"), ",")
		carID := mux.Vars(request)["car_id"]
		pageParams := filter.ParsePage(queryParams, config.DefaultLimit)

		payload := maintenanceentrypayload.ListPayload{
			Name:        name,
			CarID:       carID,
			CategoryIDs: categoryIDs,
			Sorts:       sorts,
			PageParams:  pageParams,
		}

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			entries, err = maintenanceEntryService.List(ctx, payload)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, entries)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
