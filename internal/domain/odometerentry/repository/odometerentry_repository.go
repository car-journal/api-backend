// Package odometerentryrepository handles db queries for odometer entry domain
package odometerentryrepository

import (
	"context"
	"errors"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.OdometerEntry) error
	ListByCarID(ctx context.Context, carID string) ([]*internalmodel.OdometerEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.OdometerEntry, error)
	FindByReading(ctx context.Context, reading float64) (*internalmodel.OdometerEntry, error)
	Delete(ctx context.Context, id string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.OdometerEntry) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"car_id",
			"odometer_reading",
			"reading_unit",
			"updated_at",
		}),
	})
	db = db.Where("car_id = ?", model.CarID)
	return db.Create(&model).Error
}

func (r repository) ListByCarID(ctx context.Context, carID string) ([]*internalmodel.OdometerEntry, error) {
	var models []*internalmodel.OdometerEntry
	db := database.Get(ctx)
	db = database.EqualsTo(db, "car_id", carID)
	db = db.Order("odometer_reading DESC")
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r repository) FindByID(ctx context.Context, id string) (*internalmodel.OdometerEntry, error) {
	var model *internalmodel.OdometerEntry
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

func (r repository) FindByReading(ctx context.Context, reading float64) (*internalmodel.OdometerEntry, error) {
	var model *internalmodel.OdometerEntry
	db := database.Get(ctx)
	db = database.EqualsTo(db, "odometer_reading", reading)
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
	return db.Delete(&internalmodel.OdometerEntry{}).Error
}
