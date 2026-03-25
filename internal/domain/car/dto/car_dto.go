// Package cardto handles const for car domain
package cardto

import internalmodel "github.com/car-journal/api-backend/internal/model"

type CarWithAverageFuelConsumptionRate struct {
	internalmodel.BaseModel

	Brand                      string  `json:"brand"`
	Model                      string  `json:"model"`
	ManufactureYear            *int16  `json:"manufacture_year"`
	CylinderCapacity           *int16  `json:"cylinder_capacity"`
	Color                      string  `json:"color"`
	FuelType                   string  `json:"fuel_type"`
	AverageFuelConsumptionRate float64 `json:"average_fuel_consumption_rate"`
}
