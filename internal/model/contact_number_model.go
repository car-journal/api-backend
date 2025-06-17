package internalmodel

import "github.com/car-journal/lib/uuid"

const ContactNumberTableName = "contact_numbers"

type ContactNumber struct {
	BaseModel

	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
	Number      string    `json:"number"`
	Notes       *string   `json:"notes"`
	IsMain      bool      `json:"is_main"`
}
