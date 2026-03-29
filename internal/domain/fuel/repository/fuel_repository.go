// Package fuelrepository handles db queries for fuel domain
package fuelrepository

import (
	"context"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
)

type Interface interface {
	List(ctx context.Context, name string) ([]*internalmodel.Fuel, error)
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) List(ctx context.Context, name string) ([]*internalmodel.Fuel, error) {
	var models []*internalmodel.Fuel
	db := database.Get(ctx)
	db = database.ILike(db, "name", name)
	db = db.Order("name ASC")
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
