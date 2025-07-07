package internalcarcontroller

import (
	"context"
	"fmt"
	"net/http"

	carpayload "github.com/car-journal/api-backend/internal/domain/car/payload"
	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
	"github.com/gorilla/mux"
)

func Update(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body carpayload.UpdatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			fmt.Println("err", err)
			parser.JSON(writer, nil, err)
			return
		}

		carID := mux.Vars(request)["car_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			body.UserID = me.ID
			body.ID = carID

			return carService.Update(ctx, body)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
