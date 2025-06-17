package userservice

import (
	"context"
	"fmt"

	userdto "github.com/car-journal/internal/domain/user/dto"
	userpayload "github.com/car-journal/internal/domain/user/payload"
	userrepository "github.com/car-journal/internal/domain/user/repository"
	userprofileconst "github.com/car-journal/internal/domain/userprofile/const"
	userprofilepayload "github.com/car-journal/internal/domain/userprofile/payload"
	userprofilerepository "github.com/car-journal/internal/domain/userprofile/repository"
	internalmodel "github.com/car-journal/internal/model"
	"github.com/car-journal/lib/hash"
	"github.com/car-journal/lib/httperror"
	"github.com/car-journal/lib/httperror/const/errortype"
	"github.com/car-journal/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, userPayload userpayload.CreatePayload, userProfilePayload userprofilepayload.CreatePayload) error
	FindByID(ctx context.Context, id string, preloads ...string) (*userdto.UserWithUserProfile, error)
	FindByEmail(ctx context.Context, email string) (*internalmodel.User, error)
}

type service struct {
	hashLib               hash.HashInterface
	userRepository        userrepository.Interface
	userProfileRepository userprofilerepository.Interface
	uuidLib               uuid.UUIDInterface
}

func Service(
	hashLib hash.HashInterface,
	userRepository userrepository.Interface,
	userProfileRepository userprofilerepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		hashLib:               hashLib,
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		uuidLib:               uuidLib,
	}
}

func (s service) Create(ctx context.Context, userPayload userpayload.CreatePayload, userProfilePayload userprofilepayload.CreatePayload) error {
	if isExist := s.isExists(ctx, userPayload.Email); isExist {
		return httperror.New(errortype.ALREADY_REGISTERED, fmt.Errorf("email %s already exists", userPayload.Email))
	}
	id := s.uuidLib.GenerateNewUUID()

	password, err := s.hashLib.Make(userPayload.Password)
	if err != nil {
		return err
	}
	err = s.userRepository.Create(ctx, &internalmodel.User{
		BaseModel: internalmodel.BaseModel{
			ID: id,
		},
		Email:    userPayload.Email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return s.userProfileRepository.Save(ctx, internalmodel.UserProfile{
		UserID:      id,
		Gender:      userprofileconst.GetGenderDB(userProfilePayload.Gender),
		FirstName:   userProfilePayload.FirstName,
		LastName:    userProfilePayload.LastName,
		DateOfBirth: userProfilePayload.DateOfBirth,
		PictureURL:  userProfilePayload.PictureURL,
	})
}

func (s service) FindByID(ctx context.Context, id string, preloads ...string) (*userdto.UserWithUserProfile, error) {
	return s.findByID(ctx, id, preloads...)
}

func (s service) findByID(ctx context.Context, id string, preloads ...string) (*userdto.UserWithUserProfile, error) {
	user, err := s.userRepository.FindByID(ctx, id, preloads...)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("user not found"))
	}

	return user, nil
}

func (s service) FindByEmail(ctx context.Context, email string) (*internalmodel.User, error) {
	return s.userRepository.FindByEmail(ctx, email)
}

func (s service) isExists(ctx context.Context, email string) bool {
	user, _ := s.FindByEmail(ctx, email)
	return user != nil
}
