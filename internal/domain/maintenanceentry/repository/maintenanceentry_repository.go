// Package maintenanceentryrepository handles db queries for maintenance entry domain
package maintenanceentryrepository

import (
	"context"
	"errors"

	maintenanceentryconst "github.com/car-journal/api-backend/internal/domain/maintenanceentry/const"
	maintenanceentrypayload "github.com/car-journal/api-backend/internal/domain/maintenanceentry/payload"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.MaintenanceEntry) error
	List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error)
	Delete(ctx context.Context, id string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.MaintenanceEntry) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"car_id",
			"odometer_entry_id",
			"category_id",
			"brand",
			"name",
			"price",
			"performed_at",
			"notes",
			"updated_at",
		}),
	})
	return db.Create(&model).Error
}

func (r repository) List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error) {
	var models []*internalmodel.MaintenanceEntry
	db := database.Get(ctx)
	db = database.CountAffectedRecords(db, internalmodel.MaintenanceEntryTableName)

	if payload.CarID != "" {
		db = database.EqualsTo(db, "car_id", payload.CarID)
	}

	if payload.Name != "" {
		db = database.ILike(db, "name", payload.Name)
	}

	if len(payload.CategoryIDs) > 0 {
		db = database.In(db, "category_id", payload.CategoryIDs)
	}

	if len(payload.Sorts) > 0 {
		db = database.GenerateOrderQuery(db, payload.Sorts, maintenanceentryconst.ValidSorts)
	}

	db = database.GeneratePaginationQuery(db, payload.PageParams.Limit, payload.PageParams.Offset)

	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r repository) FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error) {
	var model *internalmodel.MaintenanceEntry
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
	return db.Delete(&internalmodel.MaintenanceEntry{}).Error
}
