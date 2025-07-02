package internalauthcontroller

import (
	"context"
	"net/http"

	authpayload "github.com/car-journal/api-backend/internal/domain/auth/payload"
	authservice "github.com/car-journal/api-backend/internal/domain/auth/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
)

func UpdatePassword(authService authservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body authpayload.UpdatePassword
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			body.ID = me.ID
			return authService.UpdatePassword(ctx, body)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
