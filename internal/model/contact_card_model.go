package internalmodel

import "github.com/car-journal/lib/uuid"

const ContactCardTableName = "contact_cards"

type ContactCard struct {
	BaseModel

	UserID          uuid.UUID `json:"user_id"`
	Name            string    `json:"name"`
	ContactNumberID uuid.UUID `json:"contact_number_id"`
	AddressID       uuid.UUID `json:"address_id"`
}
