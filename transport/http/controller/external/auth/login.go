package externalauthcontroller

import (
	"context"
	"fmt"
	"net/http"

	authdto "github.com/car-journal/internal/domain/auth/dto"
	authpayload "github.com/car-journal/internal/domain/auth/payload"
	authservice "github.com/car-journal/internal/domain/auth/service"
	"github.com/car-journal/lib/database"
	"github.com/car-journal/lib/parser"
	bodyparser "github.com/car-journal/lib/parser/body"
)

func Login(authService authservice.Interface) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var body authpayload.Login
		if err := bodyparser.Parse(request, &body); nil != err {
			parser.JSON(writer, nil, err)
			return
		}

		var response *authdto.LoginResponse
		if errTrans := database.Run(request.Context(), func(ctx context.Context) (err error) {
			response, err = authService.Login(ctx, body)
			return err
		}); nil != errTrans {
			parser.JSON(writer, nil, errTrans)
			return
		}
		fmt.Println(32)
		parser.JSON(writer, response, nil)
	}
}
