// Package internalcarcontroller handles controller for car domain
package internalcarcontroller

import (
	"context"
	"net/http"

	carpayload "github.com/car-journal/api-backend/internal/domain/car/payload"
	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
)

func Create(carService carservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body carpayload.CreatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}
		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			body.UserID = me.ID
			err = carService.Create(ctx, body)
			if err != nil {
				return err
			}
			return nil
		}); errTrans != nil {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
