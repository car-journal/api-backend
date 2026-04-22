package internalmaintenancecategorycontroller

import (
	"context"
	"net/http"

	maintenancecategorypayload "github.com/car-journal/api-backend/internal/domain/maintenancecategory/payload"
	maintenancecategoryservice "github.com/car-journal/api-backend/internal/domain/maintenancecategory/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
	"github.com/gorilla/mux"
)

func Update(maintenanceCategoryService maintenancecategoryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body maintenancecategorypayload.UpdatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		id := mux.Vars(request)["id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			body.ID = id
			return maintenanceCategoryService.Update(ctx, body)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
