// Package internalmaintenancecategorycontroller handles controller for maintenance category domain
package internalmaintenancecategorycontroller

import (
	"context"
	"net/http"

	maintenancecategorypayload "github.com/car-journal/api-backend/internal/domain/maintenancecategory/payload"
	maintenancecategoryservice "github.com/car-journal/api-backend/internal/domain/maintenancecategory/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
)

func Create(maintenanceCategoryService maintenancecategoryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body maintenancecategorypayload.CreatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}
		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			return maintenanceCategoryService.Create(ctx, body)
		}); errTrans != nil {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
