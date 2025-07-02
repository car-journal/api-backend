// Package carpayload handles payload for car domain
package carpayload

import (
	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	UserID                         string
	Brand                          string  `json:"brand"`
	Model                          string  `json:"model"`
	ManufactureYear                *int16  `json:"manufacture_year"`
	CylinderCapacity               *int16  `json:"cylinder_capacity"`
	VehicleIdentityNumber          *string `json:"vehicle_identity_number"`
	EngineNumber                   *string `json:"engine_number"`
	Color                          *string `json:"color"`
	FuelType                       string  `json:"fuel_type"`
	RegistrationYear               *int16  `json:"registration_year"`
	VehicleOwnershipDocumentNumber *string `json:"vehicle_ownership_document_number"`
}

type UpdatePayload struct {
	ID                             string          `json:"id"`
	UserID                         string          `json:"user_id"`
	Brand                          *string         `json:"brand"`
	Model                          *string         `json:"model"`
	ManufactureYear                nullable.Int16  `json:"manufacture_year"`
	CylinderCapacity               nullable.Int16  `json:"cylinder_capacity"`
	VehicleIdentityNumber          nullable.String `json:"vehicle_identity_number"`
	EngineNumber                   nullable.String `json:"engine_number"`
	Color                          nullable.String `json:"color"`
	FuelType                       *string         `json:"fuel_type"`
	RegistrationYear               nullable.Int16  `json:"registration_year"`
	VehicleOwnershipDocumentNumber nullable.String `json:"vehicle_ownership_document_number"`
}
