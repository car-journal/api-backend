package internalfuelentrycontroller

import (
	"context"
	"net/http"

	fuelentrypayload "github.com/car-journal/api-backend/internal/domain/fuelentry/payload"
	fuelentryservice "github.com/car-journal/api-backend/internal/domain/fuelentry/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
	"github.com/gorilla/mux"
)

func Update(fuelEntryService fuelentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body fuelentrypayload.UpdatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		fuelEntryID := mux.Vars(request)["fuel_entry_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			body.UserID = me.ID
			body.ID = fuelEntryID

			return fuelEntryService.Update(ctx, body)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
