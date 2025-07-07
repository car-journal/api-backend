// Package fuelentryrepository handles db queries for fuel entry domain
package fuelentryrepository

import (
	"context"
	"errors"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.FuelEntry) error
	ListByIDs(ctx context.Context, ids []string) ([]*internalmodel.FuelEntry, error)
	ListByCarID(ctx context.Context, carID string) ([]*internalmodel.FuelEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.FuelEntry, error)
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

func (r repository) ListByCarID(ctx context.Context, carID string) ([]*internalmodel.FuelEntry, error) {
	var models []*internalmodel.FuelEntry
	db := database.Get(ctx)
	db = database.EqualsTo(db, "car_id", carID)
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
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
