// Package fuelservice handles services for fuel domain
package fuelservice

import (
	"context"

	fuelrepository "github.com/car-journal/api-backend/internal/domain/fuel/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
)

type Interface interface {
	List(ctx context.Context, name string) ([]*internalmodel.Fuel, error)
}

type service struct {
	fuelRepository fuelrepository.Interface
}

func Service(fuelRepository fuelrepository.Interface) Interface {
	return &service{
		fuelRepository: fuelRepository,
	}
}

func (s *service) List(ctx context.Context, name string) ([]*internalmodel.Fuel, error) {
	return s.fuelRepository.List(ctx, name)
}
