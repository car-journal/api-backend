package internalmodel

import (
	"time"

	"github.com/car-journal/lib/uuid"
)

const UserProfileTableName = "user_profiles"

type UserProfile struct {
	BaseModelNoID

	UserID      uuid.UUID  `json:"user_id"`
	Gender      *bool      `json:"gender"`
	FirstName   string     `json:"first_name"`
	LastName    *string    `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	PictureURL  *string    `json:"picture_url"`
}
