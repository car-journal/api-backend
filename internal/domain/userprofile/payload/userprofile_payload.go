package userprofilepayload

import (
	"time"

	"github.com/car-journal/lib/nullable"
)

type CreatePayload struct {
	UserID      string
	Gender      *string
	FirstName   string
	LastName    *string
	DateOfBirth *time.Time
	PictureURL  *string
}

type Addresses struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"` // id || residential
	IsMain  bool   `json:"is_main"`
}

type ContactNumbers struct {
	Name     string `json:"name"`
	CityCode string `json:"city_code"`
	Number   string `json:"number"`
	IsMain   bool   `json:"is_main"`
}

type UpdatePayload struct {
	UserID      string          `json:"user_id"`
	Gender      nullable.String `json:"gender"`
	FirstName   *string         `json:"first_name"`
	LastName    nullable.String `json:"last_name"`
	DateOfBirth nullable.String `json:"date_of_birth"`
	PictureURL  nullable.String `json:"profile_picture_url"`
}
