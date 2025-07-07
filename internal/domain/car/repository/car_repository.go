// Package carrepository handles db queries for car domain
package carrepository

import (
	"context"
	"errors"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.Car) error
	List(ctx context.Context, userID string) ([]*internalmodel.Car, error)
	FindByID(ctx context.Context, id string, userID string) (*internalmodel.Car, error)
	Delete(ctx context.Context, id string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.Car) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"brand",
			"model",
			"manufacture_year",
			"cylinder_capacity",
			"vehicle_identity_number",
			"engine_number",
			"color",
			"fuel_type",
			"registration_year",
			"vehicle_ownership_document_number",
			"updated_at",
		}),
	})
	db = db.Where("id = ?", model.ID)
	return db.Create(&model).Error
}

func (r repository) List(ctx context.Context, userID string) ([]*internalmodel.Car, error) {
	var models []*internalmodel.Car
	db := database.Get(ctx)
	db = database.EqualsTo(db, "user_id", userID)
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r repository) FindByID(ctx context.Context, id string, userID string) (*internalmodel.Car, error) {
	var model *internalmodel.Car
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", id)
	db = database.EqualsTo(db, "user_id", userID)
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
	return db.Delete(&internalmodel.Car{}).Error
}
