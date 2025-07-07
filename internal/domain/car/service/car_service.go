// Package carservice handles services for car domain
package carservice

import (
	"context"
	"fmt"

	carpayload "github.com/car-journal/api-backend/internal/domain/car/payload"
	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, payload carpayload.CreatePayload) error
	List(ctx context.Context, userID string) ([]*internalmodel.Car, error)
	FindByID(ctx context.Context, ID string, userID string) (*internalmodel.Car, error)
	Update(ctx context.Context, payload carpayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	carRepository  carrepository.Interface
	userRepository userrepository.Interface
	uuidLib        uuid.UUIDInterface
}

func Service(
	carRepository carrepository.Interface,
	userRepository userrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		carRepository:  carRepository,
		userRepository: userRepository,
		uuidLib:        uuidLib,
	}
}

func (s service) Create(ctx context.Context, payload carpayload.CreatePayload) error {
	if isExist := s.userRepository.IsIDExists(ctx, payload.UserID); !isExist {
		return httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("id %s doesn't exists", payload.UserID))
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

func (s service) List(ctx context.Context, userID string) ([]*internalmodel.Car, error) {
	return s.carRepository.List(ctx, userID)
}

func (s service) FindByID(ctx context.Context, id string, userID string) (*internalmodel.Car, error) {
	car, err := s.carRepository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	if car == nil {
		return nil, httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("car id %s doesn't exists", id))
	}

	return car, nil
}

func (s service) Update(ctx context.Context, payload carpayload.UpdatePayload) error {
	car, err := s.carRepository.FindByID(ctx, payload.ID, payload.UserID)
	if err != nil {
		return err
	}

	if car == nil {
		return httperror.New(errortype.RECORD_NOT_FOUND, fmt.Errorf("car id %s doesn't exists", payload.ID))
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
