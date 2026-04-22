// Package maintenancecategoryrepository handles db queries for maintenance category domain
package maintenancecategoryrepository

import (
	"context"
	"errors"

	maintenancecategoryconst "github.com/car-journal/api-backend/internal/domain/maintenancecategory/const"
	maintenancecategorypayload "github.com/car-journal/api-backend/internal/domain/maintenancecategory/payload"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.MaintenanceCategory) error
	List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error)
	FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error)
	Delete(ctx context.Context, id string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.MaintenanceCategory) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"description",
			"updated_at",
		}),
	})
	return db.Create(&model).Error
}

func (r repository) List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error) {
	var models []*internalmodel.MaintenanceCategory
	db := database.Get(ctx)
	db = database.CountAffectedRecords(db, internalmodel.MaintenanceCategoryTableName)

	if payload.Name != "" {
		db = database.ILike(db, "name", payload.Name)
	}

	if payload.Description != "" {
		db = database.ILike(db, "description", payload.Description)
	}

	if len(payload.Sorts) > 0 {
		db = database.GenerateOrderQuery(db, payload.Sorts, maintenancecategoryconst.ValidSorts)
	}

	db = database.GeneratePaginationQuery(db, payload.PageParams.Limit, payload.PageParams.Offset)

	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r repository) FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error) {
	var model *internalmodel.MaintenanceCategory
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
	return db.Delete(&internalmodel.MaintenanceCategory{}).Error
}
