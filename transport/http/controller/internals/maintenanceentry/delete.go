package internalmaintenanceentrycontroller

import (
	"context"
	"net/http"

	maintenanceentryservice "github.com/car-journal/api-backend/internal/domain/maintenanceentry/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func Delete(maintenanceEntryService maintenanceentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			return maintenanceEntryService.Delete(ctx, id)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
