// Package odometerentrydto handles const for odometer entry domain
package odometerentrydto

import (
	fuelentrydto "github.com/car-journal/api-backend/internal/domain/fuelentry/dto"
	internalmodel "github.com/car-journal/api-backend/internal/model"
)

type OdometerEntrySummary struct {
	*internalmodel.OdometerEntry

	FuelEntries *fuelentrydto.FuelEntry `json:"fuel_entries"`
}
