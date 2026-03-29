// Package carrepository handles db queries for car domain
package carrepository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	cardto "github.com/car-journal/api-backend/internal/domain/car/dto"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.Car) error
	List(ctx context.Context, userID string) ([]*internalmodel.Car, error)
	ListCarsWithAverageFuelConsumptionRate(ctx context.Context, userID string, pageParams *filter.Page) ([]*cardto.CarWithAverageFuelConsumptionRate, error)
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

func (r repository) ListCarsWithAverageFuelConsumptionRate(ctx context.Context, userID string, pageParams *filter.Page) ([]*cardto.CarWithAverageFuelConsumptionRate, error) {
	var models []*cardto.CarWithAverageFuelConsumptionRate
	db := database.Get(ctx)
	sql := fmt.Sprintf(`
		SELECT
			c.id,
			c.brand,
			c.model,
			c.manufacture_year,
			c.cylinder_capacity,
			c.color,
			c.fuel_type,
			AVG(f.fuel_consumption_rate) AS average_fuel_consumption_rate,
			COUNT (*) OVER() AS affected_records,
			c.created_at,
			c.updated_at,
			c.deleted_at
		FROM
			cars c
		LEFT JOIN fuel_entries f ON f.car_id = c.id
		WHERE c.user_id = '%s'
		GROUP BY c.id
		ORDER BY c.brand ASC
		LIMIT %d
		OFFSET %d
	`, userID, pageParams.Limit, pageParams.Offset)
	logger.LoggerInterface.Log("ListCarsWithAverageFuelConsumptionRate query:" + sql)
	if err := db.Raw(sql).Scan(&models).Error; err != nil {
		logger.LoggerInterface.Log(err.Error())
		return nil, err
	}
	modeslBt, _ := json.Marshal(models)
	logger.LoggerInterface.Log("ListCarsWithAverageFuelConsumptionRate result" + string(modeslBt))

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
	modelBt, _ := json.Marshal(model)
	logger.LoggerInterface.Log("FindByID Car " + string(modelBt))

	return model, nil
}

func (r repository) Delete(ctx context.Context, id string) error {
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", id)
	return db.Delete(&internalmodel.Car{}).Error
}
