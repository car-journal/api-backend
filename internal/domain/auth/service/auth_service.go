// Package authservice handles services for auth domain
package authservice

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/car-journal/api-backend/config"
	authdto "github.com/car-journal/api-backend/internal/domain/auth/dto"
	authpayload "github.com/car-journal/api-backend/internal/domain/auth/payload"
	authrepository "github.com/car-journal/api-backend/internal/domain/auth/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	userprofilerepository "github.com/car-journal/api-backend/internal/domain/userprofile/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/authz"
	"github.com/car-journal/api-backend/lib/hash"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/logger"
	"github.com/car-journal/api-backend/lib/uuid"

	"github.com/golang-jwt/jwt/v4"
)

type Interface interface {
	Register(ctx context.Context, payload authpayload.RegisterPayload) error
	FindActiveOauthAccessTokenByID(ctx context.Context, id string) (*internalmodel.OauthAccessToken, error)
	Me(ctx context.Context) (authdto.Me, error)
	Login(ctx context.Context, payload authpayload.Login) (*authdto.LoginResponse, error)
	UpdatePassword(ctx context.Context, payload authpayload.UpdatePassword) error
}

type service struct {
	authRepository        authrepository.Interface
	hashLib               hash.HashInterface
	userRepository        userrepository.Interface
	userProfileRepository userprofilerepository.Interface
	uuidLib               uuid.UUIDInterface
}

func Service(
	authRepository authrepository.Interface,
	hashLib hash.HashInterface,
	userRepository userrepository.Interface,
	userProfileRepository userprofilerepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		authRepository:        authRepository,
		hashLib:               hashLib,
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		uuidLib:               uuidLib,
	}
}

func (s service) Register(ctx context.Context, payload authpayload.RegisterPayload) error {
	if payload.ClientSecret != config.Get(config.ClientSecret) {
		return httperror.New(errortype.Forbidden, fmt.Errorf("secret not validated"))
	}

	if isEmailExists := s.userRepository.IsEmailExists(ctx, payload.Email); isEmailExists {
		return httperror.New(errortype.AlreadyRegistered, fmt.Errorf("email already exists"))
	}

	id := s.uuidLib.GenerateNewUUID()
	password, err := s.hashLib.Make(payload.Password)
	if err != nil {
		return err
	}

	if err := s.userRepository.Create(ctx, &internalmodel.User{
		BaseModel: internalmodel.BaseModel{
			ID: id,
		},
		Email:    payload.Email,
		Password: password,
	}); err != nil {
		return err
	}

	return s.userProfileRepository.Save(ctx, internalmodel.UserProfile{
		UserID:      id,
		Gender:      payload.Gender,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		DateOfBirth: payload.DateOfBirth,
	})
}

func (s service) FindActiveOauthAccessTokenByID(ctx context.Context, id string) (*internalmodel.OauthAccessToken, error) {
	return s.authRepository.FindActiveOauthAccessTokenByID(ctx, id)
}

func (s service) Me(ctx context.Context) (authdto.Me, error) {
	me := authz.GetAuthUser(ctx)
	return authdto.Me{
		ID:         me.ID,
		Email:      me.Email,
		CreatedAt:  me.CreatedAt,
		UpdatedAt:  me.UpdatedAt,
		FirstName:  me.FirstName,
		LastName:   me.LastName,
		PictureURL: me.PictureURL,
	}, nil
}

func (s service) Login(ctx context.Context, payload authpayload.Login) (*authdto.LoginResponse, error) {
	if payload.ClientSecret != config.Get(config.ClientSecret) {
		return nil, httperror.New(errortype.Forbidden, fmt.Errorf("secret not validated"))
	}
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if nil != err {
		logger.LoggerInterface.Log(err.Error())
		return nil, err
	}
	if user == nil {
		logger.LoggerInterface.Log("user does not exists")
		return nil, httperror.New(errortype.Unauthorized, fmt.Errorf("invalid credential"))
	}

	if err := s.hashLib.Check(user.Password, payload.Password); nil != err {
		logger.LoggerInterface.Log("invalid credential")
		return nil, httperror.New(errortype.Unauthorized, fmt.Errorf("invalid credential"))
	}

	return s.createJWTToken(ctx, *user)
}

func (s service) UpdatePassword(ctx context.Context, payload authpayload.UpdatePassword) error {
	user, err := s.userRepository.FindByID(ctx, payload.ID)
	if nil != err {
		return err
	}

	if err := s.hashLib.Check(user.Password, payload.OldPassword); nil != err {
		return httperror.New(errortype.Unauthorized, fmt.Errorf("invalid credential"))
	}

	password, err := s.hashLib.Make(payload.NewPassword)
	if err != nil {
		return err
	}

	user.Password = password

	if err := s.userRepository.UpdatePassword(ctx, user.ToModel()); nil != err {
		return err
	}

	return s.authRepository.RevokeOauthAccessTokensByUserID(ctx, payload.ID)
}

func (s service) createJWTToken(ctx context.Context, userData internalmodel.User) (*authdto.LoginResponse, error) {
	jti := s.uuidLib.GenerateNewUUID()
	tokenExp := time.Now().Add(time.Hour * 24 * 7)
	tokenExpUnix := tokenExp.Unix()

	jwtPrivateKey, err := base64.StdEncoding.DecodeString(config.Get(config.JwtPrivateKey))
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKey)
	if nil != err {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = jti
	claims["exp"] = tokenExpUnix
	t, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	if _, err := s.authRepository.CreateOauthAccessToken(ctx, &internalmodel.OauthAccessToken{
		ID:        jti,
		UserID:    userData.ID,
		Revoked:   false,
		ExpiresAt: &tokenExp,
	}); nil != err {
		return nil, err
	}

	return &authdto.LoginResponse{
		AccessToken: t,
		ExpiresIn:   tokenExpUnix,
		TokenType:   "Bearer",
	}, nil
}
