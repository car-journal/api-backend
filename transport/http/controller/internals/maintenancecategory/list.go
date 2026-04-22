package internalmaintenancecategorycontroller

import (
	"context"
	"net/http"

	"github.com/car-journal/api-backend/config"
	maintenancecategorypayload "github.com/car-journal/api-backend/internal/domain/maintenancecategory/payload"
	maintenancecategoryservice "github.com/car-journal/api-backend/internal/domain/maintenancecategory/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/parser"
)

func List(maintenanceCategoryService maintenancecategoryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var categories []*internalmodel.MaintenanceCategory
		var pageResponse *filter.PageResponse

		queryParams := request.URL.Query()
		pageParams := filter.ParsePage(queryParams, config.DefaultLimit)

		payload := maintenancecategorypayload.ListPayload{
			Name:        queryParams.Get("name"),
			Description: queryParams.Get("description"),
			Sorts:       queryParams["sorts"],
			PageParams:  *pageParams,
		}

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			categories, err = maintenanceCategoryService.List(ctx, payload)
			if err != nil {
				return err
			}

			pageResponse, err = filter.BuildPageResponse(ctx, pageParams, categories)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, pageResponse, nil)
	}
}
