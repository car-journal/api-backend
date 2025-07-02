package internalusercontroller

import (
	"context"
	"net/http"

	userprofilepayload "github.com/car-journal/api-backend/internal/domain/userprofile/payload"
	userprofileservice "github.com/car-journal/api-backend/internal/domain/userprofile/service"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	bodyparser "github.com/car-journal/api-backend/lib/parser/body"
)

func Update(userProfileService userprofileservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body userprofilepayload.UpdatePayload
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			me := authz.GetAuthUser(ctx)
			body.UserID = me.ID
			return userProfileService.Update(ctx, body)
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, nil, nil)
	}
}
