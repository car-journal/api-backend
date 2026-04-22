package internalmaintenancecategorycontroller

import (
	"context"
	"net/http"

	maintenancecategoryservice "github.com/car-journal/api-backend/internal/domain/maintenancecategory/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByID(maintenanceCategoryService maintenancecategoryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var category *internalmodel.MaintenanceCategory
		id := mux.Vars(request)["id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			category, err = maintenanceCategoryService.FindByID(ctx, id)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, category, nil)
	}
}
