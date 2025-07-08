// Package userrepository handles db queries for user domain
package userrepository

import (
	"context"
	"errors"

	userdto "github.com/car-journal/api-backend/internal/domain/user/dto"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
	"gorm.io/gorm"
)

type Interface interface {
	Create(ctx context.Context, payload *internalmodel.User) error
	ListByIDs(ctx context.Context, ids []string) ([]*internalmodel.User, error)
	IsIDExists(ctx context.Context, id string) bool
	IsEmailExists(ctx context.Context, email string) bool
	FindByID(ctx context.Context, id string, preloads ...string) (*userdto.UserWithUserProfile, error)
	FindByEmail(ctx context.Context, email string) (*internalmodel.User, error)
	UpdatePassword(ctx context.Context, model *internalmodel.User) error
}

type repository struct {
}

func Repository() Interface {
	return &repository{}
}

func (r repository) Create(ctx context.Context, payload *internalmodel.User) error {
	db := database.Get(ctx)
	return db.Create(payload).Error
}

func (r repository) ListByIDs(ctx context.Context, ids []string) ([]*internalmodel.User, error) {
	var models []*internalmodel.User
	db := database.Get(ctx)
	database.In(db, "id", ids)
	if err := db.Find(&models).Error; nil != err {
		return nil, err
	}

	return models, nil
}

func (r repository) IsIDExists(ctx context.Context, id string) bool {
	model, _ := r.FindByID(ctx, id)
	return model != nil
}

func (r repository) IsEmailExists(ctx context.Context, email string) bool {
	model, _ := r.FindByEmail(ctx, email)
	return model != nil
}

func (r repository) FindByID(ctx context.Context, id string, preloads ...string) (*userdto.UserWithUserProfile, error) {
	var model *userdto.UserWithUserProfile
	db := database.Get(ctx)
	db = db.Table(internalmodel.UserTableName)
	db = database.EqualsTo(db, "id", id)
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err := db.First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return model, nil
}

func (r repository) FindByEmail(ctx context.Context, email string) (*internalmodel.User, error) {
	var model *internalmodel.User
	db := database.Get(ctx)
	db = database.EqualsTo(db, "email", email)
	err := db.First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return model, nil
}

func (r repository) UpdatePassword(ctx context.Context, model *internalmodel.User) error {
	db := database.Get(ctx)
	return db.Select("password").Updates(model).Error
}
