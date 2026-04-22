// Package maintenancecategorypayload handles payload for maintenance category domain
package maintenancecategorypayload

import (
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type ListPayload struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Sorts       []string    `json:"sorts"`
	PageParams  filter.Page `json:"page_params"`
}

type UpdatePayload struct {
	ID          string          `json:"id"`
	Name        *string         `json:"name"`
	Description nullable.String `json:"description"`
}
