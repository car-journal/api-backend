package securitymiddleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/car-journal/config"
	authservice "github.com/car-journal/internal/domain/auth/service"
	userconst "github.com/car-journal/internal/domain/user/const"
	userservice "github.com/car-journal/internal/domain/user/service"
	"github.com/car-journal/lib/authz"
	"github.com/car-journal/lib/httperror"
	"github.com/car-journal/lib/httperror/const/errortype"
	"github.com/car-journal/lib/parser"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Authentication(authService authservice.Interface, userService userservice.Interface) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			// 1. Get token from header authorization
			authorization := r.Header.Get(config.AUTHORIZATION)
			bearer := strings.Replace(authorization, config.BEARER, "", -1)
			if bearer == "" {
				parser.JSON(w, nil, httperror.New(errortype.UNAUTHORIZED, fmt.Errorf("missing authorization")))
				return
			}

			// 2. Parse token with rsa private key, rsa public key
			jwtPublicKey, err := base64.StdEncoding.DecodeString(config.Get(config.JWT_PUBLIC_KEY))
			if err != nil {
				parser.JSON(w, nil, httperror.New(errortype.INTERNAL_SERVER, err))
				return
			}

			publicKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPublicKey)
			if nil != err {
				parser.JSON(w, nil, httperror.New(errortype.INTERNAL_SERVER, fmt.Errorf("missing pub key")))
				return
			}

			// 3. Decode token and get jwt id
			token, err := new(jwt.Parser).ParseWithClaims(bearer, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			})
			if err != nil {
				parser.JSON(w, nil, httperror.New(errortype.UNAUTHORIZED, err))
				return
			}

			var jti string
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				jti = claims[config.JTI].(string)
			}

			// 4. Get oauth access token by jwt id
			oauthAccessToken, err := authService.FindActiveOauthAccessTokenByID(ctx, jti)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				parser.JSON(w, nil, httperror.New(errortype.UNAUTHORIZED, fmt.Errorf("oauth access token not found")))
				return
			} else if err != nil {
				parser.JSON(w, nil, httperror.New(errortype.INTERNAL_SERVER, err))
				return
			} else if oauthAccessToken == nil {
				parser.JSON(w, nil, httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("user id not found")))
				return
			}
			userID := oauthAccessToken.UserID

			// 5. Get user by id
			userData, err := userService.FindByID(ctx, userID.String(), userconst.WITH_USER_PROFILE)
			if err != nil {
				parser.JSON(w, nil, err)
				return
			}

			// 8. Set me
			var firstName string
			var lastName *string
			if userData.UserProfile != nil {
				firstName = userData.UserProfile.FirstName
				lastName = userData.UserProfile.LastName
			}
			me := authz.Authz{
				ID:                 userData.ID.String(),
				Email:              userData.Email,
				CreatedAt:          &userData.CreatedAt,
				UpdatedAt:          &userData.UpdatedAt,
				FirstName:          firstName,
				LastName:           lastName,
				OauthAccessTokenID: oauthAccessToken.ID.String(),
			}

			ctx = authz.SetAuthUser(ctx, me)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
