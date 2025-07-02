package internalmodel

import "github.com/car-journal/api-backend/lib/uuid"

const CarTableName = "fuel_entries"

type Car struct {
	BaseModel

	UserID                         uuid.UUID `json:"user_id"`
	Brand                          string    `json:"brand"`
	Model                          string    `json:"model"`
	ManufactureYear                *int16    `json:"manufacture_year"`
	CylinderCapacity               *int16    `json:"cylinder_capacity"`
	VehicleIdentityNumber          *string   `json:"vehicle_identity_number"`
	EngineNumber                   *string   `json:"engine_number"`
	Color                          *string   `json:"color"`
	FuelType                       string    `json:"fuel_type"`
	RegistrationYear               *int16    `json:"registration_year"`
	VehicleOwnershipDocumentNumber *string   `json:"vehicle_ownership_document_number"`
}
