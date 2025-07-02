package fuelentriespayload

import (
	"time"

	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	UserID      string
	Gender      *string
	FirstName   string
	LastName    *string
	DateOfBirth *time.Time
	PictureURL  *string
}

type UpdatePayload struct {
	UserID      string          `json:"user_id"`
	Gender      nullable.String `json:"gender"`
	FirstName   *string         `json:"first_name"`
	LastName    nullable.String `json:"last_name"`
	DateOfBirth nullable.String `json:"date_of_birth"`
	PictureURL  nullable.String `json:"profile_picture_url"`
}
