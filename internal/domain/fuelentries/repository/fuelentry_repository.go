package fuelentriesrepository

import (
	"context"
	"errors"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Interface interface {
	Save(ctx context.Context, model internalmodel.UserProfile) error
	FindByUserID(ctx context.Context, userID string) (*internalmodel.UserProfile, error)
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) Save(ctx context.Context, model internalmodel.UserProfile) error {
	db := database.Get(ctx)
	db = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"gender", "first_name", "last_name", "date_of_birth", "picture_url", "updated_at",
		}),
	})
	db = db.Where("user_id = ?", model.UserID)
	return db.Create(&model).Error
}

func (r repository) FindByUserID(ctx context.Context, userID string) (*internalmodel.UserProfile, error) {
	var model *internalmodel.UserProfile
	db := database.Get(ctx)
	db = database.EqualsTo(db, "user_id", userID)
	err := db.First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return model, nil
}
