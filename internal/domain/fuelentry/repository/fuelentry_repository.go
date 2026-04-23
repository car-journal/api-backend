// Package fuelentryrepository handles db queries for fuel entry domain
package fuelentryrepository

import (
	"context"
	"errors"

	fuelentrydto "github.com/car-journal/api-backend/internal/domain/fuelentry/dto"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.FuelEntry) error
	ListByIDs(ctx context.Context, ids []string) ([]*internalmodel.FuelEntry, error)
	ListByCarID(ctx context.Context, carID string, pageParams *filter.Page) ([]*internalmodel.FuelEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.FuelEntry, error)
	CountTotalFuelCostByCarID(ctx context.Context, carID string) (float64, error)
	CountLatestAndAverageFuelConsumptionRateByCarID(ctx context.Context, carID string) (*fuelentrydto.ConsumptionRateStats, error)
	Delete(ctx context.Context, id string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.FuelEntry) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"fuel_type",
			"fuel_brand",
			"fuel_name",
			"fuel_price",
			"fuel_unit",
			"distance_traveled",
			"volume_filled",
			"filled_at",
			"total_price",
			"fuel_consumption_rate",
			"notes",
			"updated_at",
		}),
	})
	db = db.Where("car_id = ?", model.CarID)
	return db.Create(&model).Error
}

func (r repository) ListByIDs(ctx context.Context, ids []string) ([]*internalmodel.FuelEntry, error) {
	var models []*internalmodel.FuelEntry
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", ids)
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil

}

func (r repository) ListByCarID(ctx context.Context, carID string, pageParams *filter.Page) ([]*internalmodel.FuelEntry, error) {
	var models []*internalmodel.FuelEntry
	db := database.Get(ctx)
	db = database.GenerateJoinQuery(db, "INNER", "odometer_entries", "oe", "id", "fuel_entries", "odometer_entry_id")
	db = database.CountAffectedRecords(db, "fuel_entries")
	if pageParams != nil {
		db = database.GeneratePaginationQuery(db, pageParams.Limit, pageParams.Offset)
	}
	db = db.Order("oe.odometer_reading DESC")
	db = database.EqualsTo(db, "fuel_entries.car_id", carID)
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r repository) CountTotalFuelCostByCarID(ctx context.Context, carID string) (float64, error) {
	var totalCost float64
	db := database.Get(ctx)
	db = database.EqualsTo(db, "car_id", carID)
	err := db.Model(&internalmodel.FuelEntry{}).Select("SUM(total_price)").Scan(&totalCost).Error
	if err != nil {
		return 0, err
	}
	return totalCost, nil
}

func (r repository) CountLatestAndAverageFuelConsumptionRateByCarID(ctx context.Context, carID string) (*fuelentrydto.ConsumptionRateStats, error) {
	var stats fuelentrydto.ConsumptionRateStats
	db := database.Get(ctx)
	err := db.Raw(`
		WITH ranked AS (
			SELECT
				fuel_consumption_rate,
				total_price,
				ROW_NUMBER() OVER (ORDER BY filled_at DESC) AS rn,
				SUM(total_price) OVER () AS total_fuel_cost,
				AVG(fuel_consumption_rate) OVER () AS all_time_avg_rate,
				FIRST_VALUE(fuel_consumption_rate) OVER (ORDER BY filled_at DESC) AS current_rate,
				AVG(fuel_consumption_rate) OVER (
					ORDER BY filled_at DESC
					ROWS BETWEEN 1 FOLLOWING AND UNBOUNDED FOLLOWING
				) AS previous_avg_rate
			FROM fuel_entries
			WHERE car_id = ?
		)
		SELECT total_fuel_cost, all_time_avg_rate, current_rate, previous_avg_rate
		FROM ranked
		WHERE rn = 1
	`, carID).Scan(&stats).Error
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func (r repository) FindByID(ctx context.Context, id string) (*internalmodel.FuelEntry, error) {
	var model *internalmodel.FuelEntry
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", id)
	err := db.First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return model, nil
}

func (r repository) Delete(ctx context.Context, id string) error {
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", id)
	return db.Delete(&internalmodel.FuelEntry{}).Error
}
