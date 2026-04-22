// Package internalmaintenanceentrycontroller handles controller for maintenance entry domain
package internalmaintenanceentrycontroller

import (
	"context"
	"net/http"

	maintenanceentrypayload "github.com/car-journal/api-backend/internal/domain/maintenanceentry/payload"
	maintenanceentryservice "github.com/car-journal/api-backend/internal/domain/maintenanceentry/service"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
	"github.com/gorilla/mux"
)

func Create(maintenanceEntryService maintenanceentryservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body maintenanceentrypayload.CreatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		carID := mux.Vars(request)["car_id"]
		body.CarID = carID
		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			return maintenanceEntryService.Create(ctx, body)
		}); errTrans != nil {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
