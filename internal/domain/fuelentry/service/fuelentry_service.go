// Package fuelentryservice handles user authentication and profile management.
package fuelentryservice

import (
	"context"
	"fmt"

	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	fuelentrypayload "github.com/car-journal/api-backend/internal/domain/fuelentry/payload"
	fuelentryrepository "github.com/car-journal/api-backend/internal/domain/fuelentry/repository"
	odometerentryrepository "github.com/car-journal/api-backend/internal/domain/odometerentry/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/fuel"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, payload fuelentrypayload.CreatePayload) error
	ListByCarID(ctx context.Context, carID string, userID string) ([]*internalmodel.FuelEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.FuelEntry, error)
	Update(ctx context.Context, payload fuelentrypayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	carRepository           carrepository.Interface
	fuelEntryRepository     fuelentryrepository.Interface
	odometerEntryRepository odometerentryrepository.Interface
	userRepository          userrepository.Interface
	uuidLib                 uuid.UUIDInterface
}

func Service(
	carRepository carrepository.Interface,
	fuelEntryRepository fuelentryrepository.Interface,
	odometerEntryRepository odometerentryrepository.Interface,
	userRepository userrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		carRepository:           carRepository,
		fuelEntryRepository:     fuelEntryRepository,
		odometerEntryRepository: odometerEntryRepository,
		userRepository:          userRepository,
		uuidLib:                 uuidLib,
	}
}

func (s service) Create(ctx context.Context, payload fuelentrypayload.CreatePayload) error {
	if isExist := s.userRepository.IsIDExists(ctx, payload.UserID); !isExist {
		return httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("id %s doesn't exists", payload.UserID))
	}

	fuelEntryID := s.uuidLib.GenerateNewUUID()
	odometerEntryID := s.uuidLib.GenerateNewUUID()

	carID, err := uuid.StringToUUID(payload.CarID)
	if err != nil {
		return err
	}

	fuelConsumptionRate := fuel.CalculateFuelConsumptionRate(payload.DistanceTraveled, payload.VolumeFilled)
	totalPrice := fuel.CalculateTotalPrice(payload.FuelPrice, payload.VolumeFilled)

	if err := s.odometerEntryRepository.Save(ctx, internalmodel.OdometerEntry{
		BaseModel: internalmodel.BaseModel{
			ID: odometerEntryID,
		},
		CarID:           carID,
		OdometerReading: payload.OdometerReading,
		ReadingUnit:     payload.ReadingUnit,
	}); err != nil {
		return err
	}

	return s.fuelEntryRepository.Save(ctx, internalmodel.FuelEntry{
		BaseModel: internalmodel.BaseModel{
			ID: fuelEntryID,
		},
		CarID:               carID,
		OdometerEntryID:     odometerEntryID,
		FuelType:            payload.FuelType,
		FuelBrand:           payload.FuelBrand,
		FuelName:            payload.FuelName,
		FuelPrice:           payload.FuelPrice,
		FuelUnit:            payload.FuelUnit,
		DistanceTraveled:    payload.DistanceTraveled,
		VolumeFilled:        payload.VolumeFilled,
		TotalPrice:          totalPrice,
		FuelConsumptionRate: fuelConsumptionRate,
		Notes:               payload.Notes,
	})
}

func (s service) ListByCarID(ctx context.Context, carID string, userID string) ([]*internalmodel.FuelEntry, error) {
	car, err := s.findCarByID(ctx, carID, userID)
	if err != nil {
		return nil, err
	}

	return s.fuelEntryRepository.ListByCarID(ctx, car.ID.String())
}

func (s service) FindByID(ctx context.Context, id string) (*internalmodel.FuelEntry, error) {
	fuelEntry, err := s.fuelEntryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if fuelEntry == nil {
		return nil, httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("fuel entry id %s doesn't exists", id))
	}

	return fuelEntry, nil
}

func (s service) Update(ctx context.Context, payload fuelentrypayload.UpdatePayload) error {
	_, err := s.findCarByID(ctx, payload.CarID, payload.UserID)
	if err != nil {
		return err
	}

	fuelEntry, err := s.fuelEntryRepository.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if fuelEntry == nil {
		return httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("fuel entry id %s doesn't exists", payload.ID))
	}

	odometerEntry, err := s.odometerEntryRepository.FindByID(ctx, fuelEntry.OdometerEntryID.String())
	if err != nil {
		return err
	}

	if odometerEntry == nil {
		return httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("odometer entry id %s doesn't exists", fuelEntry.OdometerEntryID))
	}

	if payload.OdometerReading != nil {
		odometerEntry.OdometerReading = *payload.OdometerReading
		odometerEntry.ReadingUnit = *payload.ReadingUnit
	}

	if err := s.odometerEntryRepository.Save(ctx, *odometerEntry); err != nil {
		return err
	}

	if payload.FuelType != nil {
		fuelEntry.FuelType = *payload.FuelType
	}

	if payload.FuelBrand != nil {
		fuelEntry.FuelBrand = *payload.FuelBrand
	}

	if payload.FuelName != nil {
		fuelEntry.FuelName = *payload.FuelName
	}

	var pricePerUnit *float64
	if payload.FuelPrice != nil {
		fuelEntry.FuelPrice = *payload.FuelPrice
		pricePerUnit = payload.FuelPrice
	}

	if payload.FuelUnit != nil {
		fuelEntry.FuelUnit = *payload.FuelUnit
	}

	var distanceTraveled *float64
	if payload.DistanceTraveled != nil {
		fuelEntry.DistanceTraveled = *payload.DistanceTraveled
		distanceTraveled = payload.DistanceTraveled
	}

	var volumeFilled *float64
	if payload.VolumeFilled != nil {
		fuelEntry.VolumeFilled = *payload.VolumeFilled
		volumeFilled = payload.VolumeFilled
	}

	if distanceTraveled != nil && volumeFilled != nil {
		fuelEntry.FuelConsumptionRate = fuel.CalculateFuelConsumptionRate(*distanceTraveled, *volumeFilled)
	}

	if pricePerUnit != nil && volumeFilled != nil {
		fuelEntry.TotalPrice = fuel.CalculateTotalPrice(*pricePerUnit, *volumeFilled)
	}

	if payload.Notes.Valid {
		fuelEntry.Notes = payload.Notes.Value
	}

	return s.fuelEntryRepository.Save(ctx, *fuelEntry)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.fuelEntryRepository.Delete(ctx, id)
}

func (s service) findCarByID(ctx context.Context, carID string, userID string) (*internalmodel.Car, error) {
	car, err := s.carRepository.FindByID(ctx, carID, userID)
	if err != nil {
		return nil, err
	}

	if car == nil {
		return nil, httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("car id %s doesn't exists", carID))
	}

	return car, nil
}
