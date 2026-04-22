package internalusercontroller

import (
	"context"
	"net/http"

	userprofileservice "github.com/car-journal/api-backend/internal/domain/userprofile/service"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func FindByUserID(userProfileService userprofileservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var userProfile *internalmodel.UserProfile
		userID := mux.Vars(request)["user_id"]

		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			userProfile, err = userProfileService.FindByUserID(ctx, userID)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		parser.JSON(writer, userProfile, nil)
	}
}
