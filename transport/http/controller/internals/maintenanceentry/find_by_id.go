package internalmaintenanceentrycontroller

import (
	"context"
	"net/http"

	maintenanceentryservice "github.com/car-journal/api-backend/internal/domain/maintenanceentry/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByID(maintenanceEntryService maintenanceentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var entry *internalmodel.MaintenanceEntry
		id := mux.Vars(request)["id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			entry, err = maintenanceEntryService.FindByID(ctx, id)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, entry, nil)
	}
}
