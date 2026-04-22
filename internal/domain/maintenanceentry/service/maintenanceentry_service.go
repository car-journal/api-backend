// Package maintenanceentryservice handles services for maintenance entry domain
package maintenanceentryservice

import (
	"context"
	"fmt"

	maintenancecategoryrepository "github.com/car-journal/api-backend/internal/domain/maintenancecategory/repository"
	maintenanceentrypayload "github.com/car-journal/api-backend/internal/domain/maintenanceentry/payload"
	maintenanceentryrepository "github.com/car-journal/api-backend/internal/domain/maintenanceentry/repository"
	odometerentryrepository "github.com/car-journal/api-backend/internal/domain/odometerentry/repository"
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
)

type Interface interface {
	Create(ctx context.Context, payload maintenanceentrypayload.CreatePayload) error
	List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error)
	FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error)
	Update(ctx context.Context, payload maintenanceentrypayload.UpdatePayload) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	maintenanceCategoryRepository maintenancecategoryrepository.Interface
	maintenanceEntryRepository    maintenanceentryrepository.Interface
	odometerEntryRepository       odometerentryrepository.Interface
	uuidLib                       uuid.UUIDInterface
}

func Service(
	maintenanceCategoryRepository maintenancecategoryrepository.Interface,
	maintenanceEntryRepository maintenanceentryrepository.Interface,
	odometerEntryRepository odometerentryrepository.Interface,
	uuidLib uuid.UUIDInterface,
) Interface {
	return &service{
		maintenanceCategoryRepository: maintenanceCategoryRepository,
		maintenanceEntryRepository:    maintenanceEntryRepository,
		odometerEntryRepository:       odometerEntryRepository,
		uuidLib:                       uuidLib,
	}
}

func (s service) Create(ctx context.Context, payload maintenanceentrypayload.CreatePayload) error {
	// Verify category exists
	category, err := s.maintenanceCategoryRepository.FindByID(ctx, payload.CategoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance category id %s doesn't exists", payload.CategoryID))
	}

	carID, err := uuid.StringToUUID(payload.CarID)
	if err != nil {
		return err
	}

	categoryID, err := uuid.StringToUUID(payload.CategoryID)
	if err != nil {
		return err
	}

	// Resolve odometer entry
	odometerEntryID, err := s.resolveOdometerEntry(ctx, carID, payload.OdometerReading, payload.ReadingUnit)
	if err != nil {
		return err
	}

	entryID := s.uuidLib.GenerateNewUUID()
	return s.maintenanceEntryRepository.Save(ctx, internalmodel.MaintenanceEntry{
		BaseModel: internalmodel.BaseModel{
			ID: entryID,
		},
		CarID:           carID,
		OdometerEntryID: odometerEntryID,
		CategoryID:      categoryID,
		Brand:           payload.Brand,
		Name:            payload.Name,
		Price:           payload.Price,
		PerformedAt:     payload.PerformedAt,
		Notes:           payload.Notes,
	})
}

func (s service) List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error) {
	return s.maintenanceEntryRepository.List(ctx, payload)
}

func (s service) FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error) {
	entry, err := s.maintenanceEntryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance entry id %s doesn't exists", id))
	}

	return entry, nil
}

func (s service) Update(ctx context.Context, payload maintenanceentrypayload.UpdatePayload) error {
	entry, err := s.maintenanceEntryRepository.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if entry == nil {
		return httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance entry id %s doesn't exists", payload.ID))
	}

	// Handle odometer upsert if reading is provided
	if payload.OdometerReading != nil {
		readingUnit := ""
		if payload.ReadingUnit != nil {
			readingUnit = *payload.ReadingUnit
		}
		odometerEntryID, err := s.resolveOdometerEntry(ctx, entry.CarID, *payload.OdometerReading, readingUnit)
		if err != nil {
			return err
		}
		entry.OdometerEntryID = odometerEntryID
	}

	// Verify and update category if provided
	if payload.CategoryID != nil {
		category, err := s.maintenanceCategoryRepository.FindByID(ctx, *payload.CategoryID)
		if err != nil {
			return err
		}
		if category == nil {
			return httperror.New(errortype.RecordNotFound, fmt.Errorf("maintenance category id %s doesn't exists", *payload.CategoryID))
		}
		categoryID, err := uuid.StringToUUID(*payload.CategoryID)
		if err != nil {
			return err
		}
		entry.CategoryID = categoryID
	}

	if payload.Brand != nil {
		entry.Brand = *payload.Brand
	}

	if payload.Name != nil {
		entry.Name = *payload.Name
	}

	if payload.Price != nil {
		entry.Price = *payload.Price
	}

	if payload.PerformedAt.Valid {
		entry.PerformedAt = *payload.PerformedAt.Value
	}

	if payload.Notes.Valid {
		entry.Notes = payload.Notes.Value
	}

	return s.maintenanceEntryRepository.Save(ctx, *entry)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.maintenanceEntryRepository.Delete(ctx, id)
}

// resolveOdometerEntry checks if an odometer entry with the given reading exists.
// If it does, it returns the existing ID. Otherwise, it creates a new one and returns its ID.
func (s service) resolveOdometerEntry(ctx context.Context, carID uuid.UUID, reading float64, readingUnit string) (uuid.UUID, error) {
	existing, err := s.odometerEntryRepository.FindByReading(ctx, reading)
	if err != nil {
		return uuid.UUID{}, err
	}

	if existing != nil {
		return existing.ID, nil
	}

	newID := s.uuidLib.GenerateNewUUID()
	if err := s.odometerEntryRepository.Save(ctx, internalmodel.OdometerEntry{
		BaseModel: internalmodel.BaseModel{
			ID: newID,
		},
		CarID:           carID,
		OdometerReading: reading,
		ReadingUnit:     readingUnit,
	}); err != nil {
		return uuid.UUID{}, err
	}

	return newID, nil
}
