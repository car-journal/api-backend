// Package carservice handles services for car domain
package carservice

import (
	"context"
	"fmt"

	"github.com/car-journal/api-backend/config"
	cardto "github.com/car-journal/api-backend/internal/domain/car/dto"
	carpayload "github.com/car-journal/api-backend/internal/domain/car/payload"
	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	fuelentryrepository "github.com/car-journal/api-backend/internal/domain/fuelentry/repository"
	maintenanceentryconst "github.com/car-journal/api-backend/internal/domain/maintenanceentry/const"
	maintenanceentrypayload "github.com/car-journal/api-backend/internal/domain/maintenanceentry/payload"
	maintenanceentryrepository "github.com/car-journal/api-backend/internal/domain/maintenanceentry/repository"
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
	FindByIDWithFuelAndMaintenanceSummary(ctx context.Context, id string, userID string) (*cardto.CarWithFuelAndMaintenanceSummary, error)
	FindByID(ctx context.Context, ID string, userID string) (*internalmodel.Car, error)
	Update(ctx context.Context, payload carpayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	carRepository              carrepository.Interface
	fuelEntryRepository        fuelentryrepository.Interface
	maintenanceEntryRepository maintenanceentryrepository.Interface
	userRepository             userrepository.Interface
	uuidLib                    uuid.UUIDInterface
}

func Service(
	carRepository carrepository.Interface,
	fuelEntryRepository fuelentryrepository.Interface,
	maintenanceEntryRepository maintenanceentryrepository.Interface,
	userRepository userrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		carRepository:              carRepository,
		fuelEntryRepository:        fuelEntryRepository,
		maintenanceEntryRepository: maintenanceEntryRepository,
		userRepository:             userRepository,
		uuidLib:                    uuidLib,
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

func (s service) FindByIDWithFuelAndMaintenanceSummary(ctx context.Context, id string, userID string) (*cardto.CarWithFuelAndMaintenanceSummary, error) {
	car, err := s.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	pageParams := &filter.Page{
		Limit:  5,
		Offset: 0,
	}

	fuelEntries, err := s.fuelEntryRepository.ListByCarID(ctx, car.ID.String(), pageParams)
	if err != nil {
		return nil, err
	}

	stats, err := s.fuelEntryRepository.CountLatestAndAverageFuelConsumptionRateByCarID(ctx, car.ID.String())
	if err != nil {
		return nil, err
	}

	var trend float64
	if stats != nil && stats.PreviousAvgRate > 0 {
		delta := stats.CurrentRate - stats.PreviousAvgRate
		trend = (delta / stats.PreviousAvgRate) * 100
	}
	maintenanceEntries, err := s.maintenanceEntryRepository.List(ctx, maintenanceentrypayload.ListPayload{
		CarID: id,
		Sorts: []string{
			fmt.Sprintf("%s:%s", maintenanceentryconst.PerformedAt, config.Descending),
		},
		PageParams: pageParams,
	})
	if err != nil {
		return nil, err
	}

	totalMaintenanceCost, err := s.maintenanceEntryRepository.CountTotalMaintenanceCostByCarID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &cardto.CarWithFuelAndMaintenanceSummary{
		Car: car,
		FuelSummary: cardto.FuelSummary{
			AverageFuelConsumptionRate:  stats.CurrentRate,
			FuelConsumptionRateTrend:    trend,
			FuelConsumptionRateIncrease: stats.AllTimeAvgRate - stats.PreviousAvgRate,
			TotalFuelCost:               stats.TotalFuelCost,
			RecentFuelEntries:           fuelEntries,
		},
		MaintenanceSummary: cardto.MaintenanceSummary{
			TotalMaintenanceCost:     totalMaintenanceCost,
			RecentMaintenanceEntries: maintenanceEntries,
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
