// Package carservice handles services for car domain
package carservice

import (
	"context"
	"fmt"

	cardto "github.com/car-journal/api-backend/internal/domain/car/dto"
	carpayload "github.com/car-journal/api-backend/internal/domain/car/payload"
	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	fuelentryrepository "github.com/car-journal/api-backend/internal/domain/fuelentry/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, payload carpayload.CreatePayload) error
	ListCarsWithAverageFuelConsumptionRate(ctx context.Context, userID string, pageParams *filter.Page) ([]*cardto.CarWithAverageFuelConsumptionRate, error)
	FindByIDWithFuelSummary(ctx context.Context, id string, userID string) (*cardto.CarWithFuelSummary, error)
	FindByID(ctx context.Context, ID string, userID string) (*internalmodel.Car, error)
	Update(ctx context.Context, payload carpayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	carRepository       carrepository.Interface
	fuelEntryRepository fuelentryrepository.Interface
	userRepository      userrepository.Interface
	uuidLib             uuid.UUIDInterface
}

func Service(
	carRepository carrepository.Interface,
	fuelEntryRepository fuelentryrepository.Interface,
	userRepository userrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		carRepository:       carRepository,
		fuelEntryRepository: fuelEntryRepository,
		userRepository:      userRepository,
		uuidLib:             uuidLib,
	}
}

func (s service) Create(ctx context.Context, payload carpayload.CreatePayload) error {
	if isExist := s.userRepository.IsIDExists(ctx, payload.UserID); !isExist {
		return httperror.New(errortype.RecordNotFound, fmt.Errorf("id %s doesn't exists", payload.UserID))
	}

	id := s.uuidLib.GenerateNewUUID()
	userID, err := uuid.StringToUUID(payload.UserID)
	if err != nil {
		return err
	}
	return s.carRepository.Save(ctx, internalmodel.Car{
		BaseModel: internalmodel.BaseModel{
			ID: id,
		},
		UserID:                         userID,
		Brand:                          payload.Brand,
		Model:                          payload.Model,
		ManufactureYear:                payload.ManufactureYear,
		CylinderCapacity:               payload.CylinderCapacity,
		VehicleIdentityNumber:          payload.VehicleIdentityNumber,
		EngineNumber:                   payload.EngineNumber,
		Color:                          payload.Color,
		FuelType:                       payload.FuelType,
		RegistrationYear:               payload.RegistrationYear,
		VehicleOwnershipDocumentNumber: payload.VehicleOwnershipDocumentNumber,
	})
}

func (s service) ListCarsWithAverageFuelConsumptionRate(ctx context.Context, userID string, pageParams *filter.Page) ([]*cardto.CarWithAverageFuelConsumptionRate, error) {
	return s.carRepository.ListCarsWithAverageFuelConsumptionRate(ctx, userID, pageParams)
}

func (s service) FindByIDWithFuelSummary(ctx context.Context, id string, userID string) (*cardto.CarWithFuelSummary, error) {
	car, err := s.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	fuelEntries, err := s.fuelEntryRepository.ListByCarID(ctx, car.ID.String(), &filter.Page{
		Limit:  5,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	var currentFuelConsumptionRate float64
	var previousFuelConsumptionRate float64
	for i, val := range fuelEntries {
		currentFuelConsumptionRate += val.FuelConsumptionRate

		if i > 0 {
			previousFuelConsumptionRate += val.FuelConsumptionRate
		}
	}

	var currentAverageFuelConsumptionRate float64
	var previousAverageFuelConsumptionRate float64
	if len(fuelEntries) > 0 {
		currentAverageFuelConsumptionRate = currentFuelConsumptionRate / float64(len(fuelEntries))
	}

	if len(fuelEntries) > 1 {
		previousAverageFuelConsumptionRate = previousFuelConsumptionRate / float64(len(fuelEntries)-1)
	}

	delta := currentAverageFuelConsumptionRate - previousAverageFuelConsumptionRate
	trend := delta * previousAverageFuelConsumptionRate / 100

	return &cardto.CarWithFuelSummary{
		Car: car,
		FuelSummary: cardto.FuelSummary{
			AverageFuelConsumptionRate: currentAverageFuelConsumptionRate,
			FuelConsumptionRateTrend:   trend,
			RecentFuelEntries:          fuelEntries,
		},
	}, nil
}

func (s service) FindByID(ctx context.Context, id string, userID string) (*internalmodel.Car, error) {
	car, err := s.carRepository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	if car == nil {
		return nil, httperror.New(errortype.RecordNotFound, fmt.Errorf("car id %s doesn't exists", id))
	}

	return car, nil
}

func (s service) Update(ctx context.Context, payload carpayload.UpdatePayload) error {
	car, err := s.carRepository.FindByID(ctx, payload.ID, payload.UserID)
	if err != nil {
		return err
	}

	if car == nil {
		return httperror.New(errortype.RecordNotFound, fmt.Errorf("car id %s doesn't exists", payload.ID))
	}

	if payload.Brand != nil {
		car.Brand = *payload.Brand
	}

	if payload.Model != nil {
		car.Model = *payload.Model
	}

	if payload.ManufactureYear.Valid {
		car.ManufactureYear = payload.ManufactureYear.Value
	}

	if payload.CylinderCapacity.Valid {
		car.CylinderCapacity = payload.CylinderCapacity.Value
	}

	if payload.VehicleIdentityNumber.Valid {
		car.VehicleIdentityNumber = payload.VehicleIdentityNumber.Value
	}

	if payload.EngineNumber.Valid {
		car.EngineNumber = payload.EngineNumber.Value
	}

	if payload.Color.Valid {
		car.Color = payload.Color.Value
	}

	if payload.FuelType != nil {
		car.FuelType = *payload.FuelType
	}

	if payload.RegistrationYear.Valid {
		car.RegistrationYear = payload.RegistrationYear.Value
	}

	if payload.VehicleOwnershipDocumentNumber.Valid {
		car.VehicleOwnershipDocumentNumber = payload.VehicleOwnershipDocumentNumber.Value
	}

	return s.carRepository.Save(ctx, *car)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.carRepository.Delete(ctx, id)
}
