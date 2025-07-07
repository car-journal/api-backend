package internalfuelentrycontroller

import (
	"context"
	"net/http"

	fuelentryservice "github.com/car-journal/api-backend/internal/domain/fuelentries/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByID(fuelEntryService fuelentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var fuelEntries *internalmodel.FuelEntry
		id := mux.Vars(request)["fuel_entry_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			fuelEntries, err = fuelEntryService.FindByID(ctx, id)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, fuelEntries, nil)
	}
}
