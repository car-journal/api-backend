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

type CarWithFuelAndMaintenanceSummary struct {
	*internalmodel.Car

	FuelSummary        `json:"fuel_summary"`
	MaintenanceSummary `json:"maintenance_summary"`
}

type FuelSummary struct {
	AverageFuelConsumptionRate  float64                    `json:"average_fuel_consumption_rate"`
	FuelConsumptionRateIncrease float64                    `json:"fuel_consumption_rate_increase,omitempty"`
	FuelConsumptionRateTrend    float64                    `json:"fuel_consumption_rate_trend,omitempty"`
	TotalFuelCost               float64                    `json:"total_fuel_cost"`
	RecentFuelEntries           []*internalmodel.FuelEntry `json:"recent_fuel_entries"`
}

type MaintenanceSummary struct {
	TotalMaintenanceCost     float64                           `json:"total_maintenance_cost"`
	RecentMaintenanceEntries []*internalmodel.MaintenanceEntry `json:"recent_maintenance_entries"`
}
