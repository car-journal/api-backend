// Package odometerentryservice handles user odometer entries.
package odometerentryservice

import (
	"context"
	"fmt"

	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	fuelentrydto "github.com/car-journal/api-backend/internal/domain/fuelentry/dto"
	fuelentryrepository "github.com/car-journal/api-backend/internal/domain/fuelentry/repository"
	odometerentrydto "github.com/car-journal/api-backend/internal/domain/odometerentry/dto"
	odometerentryrepository "github.com/car-journal/api-backend/internal/domain/odometerentry/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	ListByCarID(ctx context.Context, carID string, userID string) ([]*odometerentrydto.OdometerEntrySummary, error)
	FindByID(ctx context.Context, id string) (*internalmodel.OdometerEntry, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	carRepository           carrepository.Interface
	fuelEntryRepository     fuelentryrepository.Interface
	odometerEntryRepository odometerentryrepository.Interface
	uuidLib                 uuid.UUIDInterface
}

func Service(
	carRepository carrepository.Interface,
	fuelEntryRepository fuelentryrepository.Interface,
	odometerEntryRepository odometerentryrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		carRepository:           carRepository,
		fuelEntryRepository:     fuelEntryRepository,
		odometerEntryRepository: odometerEntryRepository,
		uuidLib:                 uuidLib,
	}
}

func (s service) ListByCarID(ctx context.Context, carID string, userID string) ([]*odometerentrydto.OdometerEntrySummary, error) {
	car, err := s.findCarByID(ctx, carID, userID)
	if err != nil {
		return nil, err
	}

	fuelEntries, err := s.fuelEntryRepository.ListByCarID(ctx, car.ID.String(), nil)
	if err != nil {
		return nil, err
	}

	mappedFuelEntries := make(map[uuid.UUID]*fuelentrydto.FuelEntry, len(fuelEntries))
	for _, fuelEntry := range fuelEntries {
		mappedFuelEntries[fuelEntry.OdometerEntryID] = &fuelentrydto.FuelEntry{
			ID:                  fuelEntry.ID,
			FuelBrand:           fuelEntry.FuelBrand,
			FuelName:            fuelEntry.FuelName,
			FuelUnit:            fuelEntry.FuelUnit,
			VolumeFilled:        fuelEntry.VolumeFilled,
			TotalPrice:          fuelEntry.TotalPrice,
			FuelConsumptionRate: fuelEntry.FuelConsumptionRate,
		}
	}

	odometerEntries, err := s.odometerEntryRepository.ListByCarID(ctx, car.ID.String())
	if err != nil {
		return nil, err
	}

	odometerEntriesSummary := make([]*odometerentrydto.OdometerEntrySummary, len(odometerEntries))

	for idx, odometerEntry := range odometerEntries {
		odometerEntriesSummary[idx] = &odometerentrydto.OdometerEntrySummary{
			OdometerEntry: odometerEntry,
			FuelEntries:   mappedFuelEntries[odometerEntry.ID],
		}
	}

	return odometerEntriesSummary, nil
}

func (s service) FindByID(ctx context.Context, id string) (*internalmodel.OdometerEntry, error) {
	odometerEntry, err := s.odometerEntryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if odometerEntry == nil {
		return nil, httperror.New(errortype.RecordNotFound, fmt.Errorf("odometer entry id %s doesn't exists", id))
	}

	return odometerEntry, nil
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.odometerEntryRepository.Delete(ctx, id)
}

func (s service) findCarByID(ctx context.Context, carID string, userID string) (*internalmodel.Car, error) {
	car, err := s.carRepository.FindByID(ctx, carID, userID)
	if err != nil {
		return nil, err
	}

	if car == nil {
		return nil, httperror.New(errortype.RecordNotFound, fmt.Errorf("car id %s doesn't exists", carID))
	}

	return car, nil
}
