package userprofileservice

import (
	"context"
	"time"

	userprofileconst "github.com/car-journal/internal/domain/userprofile/const"
	userprofilepayload "github.com/car-journal/internal/domain/userprofile/payload"
	userprofilerepository "github.com/car-journal/internal/domain/userprofile/repository"
	internalmodel "github.com/car-journal/internal/model"
	"github.com/car-journal/lib/uuid"
)

type Interface interface {
	Update(ctx context.Context, payload userprofilepayload.UpdatePayload) error
}

type service struct {
	userProfileRepository userprofilerepository.Interface
	uuidLib               uuid.UUIDInterface
}

func Service(
	userProfileRepository userprofilerepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		userProfileRepository: userProfileRepository,
		uuidLib:               uuidLib,
	}
}

func (s service) Update(ctx context.Context, payload userprofilepayload.UpdatePayload) error {
	userProfile, err := s.userProfileRepository.FindByUserID(ctx, payload.UserID)
	if err != nil {
		return err
	}

	// Create user profile it's not exist yet
	if userProfile == nil {
		userID, err := uuid.StringToUUID(payload.UserID)
		if err != nil {
			return err
		}
		userProfile = &internalmodel.UserProfile{
			UserID: userID,
		}
	}

	if payload.Gender.Valid {
		userProfile.Gender = userprofileconst.GetGenderDB(payload.Gender.Value)
	}
	if payload.FirstName != nil {
		userProfile.FirstName = *payload.FirstName
	}
	if payload.LastName.Valid {
		userProfile.LastName = payload.LastName.Value
	}

	if payload.DateOfBirth.Valid {
		var dateOfBirth *time.Time
		if payload.DateOfBirth.Value != nil {
			parsedDateOfBirth, err := time.Parse(time.RFC3339, *payload.DateOfBirth.Value)
			if err != nil {
				return err
			}
			dateOfBirth = &parsedDateOfBirth
		}
		userProfile.DateOfBirth = dateOfBirth
	}

	if payload.PictureURL.Valid {
		userProfile.PictureURL = payload.PictureURL.Value
	}

	return s.userProfileRepository.Save(ctx, *userProfile)
}
