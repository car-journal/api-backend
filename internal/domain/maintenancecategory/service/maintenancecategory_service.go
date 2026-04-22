// Package maintenancecategoryservice handles services for maintenance category domain
package maintenancecategoryservice

import (
	"context"
	"fmt"

	maintenancecategorypayload "github.com/car-journal/api-backend/internal/domain/maintenancecategory/payload"
	maintenancecategoryrepository "github.com/car-journal/api-backend/internal/domain/maintenancecategory/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, payload maintenancecategorypayload.CreatePayload) error
	List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error)
	FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error)
	Update(ctx context.Context, payload maintenancecategorypayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	maintenanceCategoryRepository maintenancecategoryrepository.Interface
	uuidLib                       uuid.UUIDInterface
}

func Service(
	maintenanceCategoryRepository maintenancecategoryrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		maintenanceCategoryRepository: maintenanceCategoryRepository,
		uuidLib:                       uuidLib,
	}
}

func (s service) Create(ctx context.Context, payload maintenancecategorypayload.CreatePayload) error {
	id := s.uuidLib.GenerateNewUUID()
	return s.maintenanceCategoryRepository.Save(ctx, internalmodel.MaintenanceCategory{
		BaseModel: internalmodel.BaseModel{
			ID: id,
		},
		Name:        payload.Name,
		Description: payload.Description,
	})
}

func (s service) List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error) {
	return s.maintenanceCategoryRepository.List(ctx, payload)
}

func (s service) FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error) {
	category, err := s.maintenanceCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance category id %s doesn't exists", id))
	}

	return category, nil
}

func (s service) Update(ctx context.Context, payload maintenancecategorypayload.UpdatePayload) error {
	category, err := s.maintenanceCategoryRepository.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if category == nil {
		return httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance category id %s doesn't exists", payload.ID))
	}

	if payload.Name != nil {
		category.Name = *payload.Name
	}

	if payload.Description.Valid {
		category.Description = payload.Description.Value
	}

	return s.maintenanceCategoryRepository.Save(ctx, *category)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.maintenanceCategoryRepository.Delete(ctx, id)
}
