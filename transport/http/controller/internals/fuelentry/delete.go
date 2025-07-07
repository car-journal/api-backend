package internalfuelentrycontroller

import (
	"context"
	"net/http"

	fuelentryservice "github.com/car-journal/api-backend/internal/domain/fuelentries/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func Delete(fuelEntryService fuelentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["fuel_entry_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			return fuelEntryService.Delete(ctx, id)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
