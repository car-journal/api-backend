package container

import (
	authrepository "github.com/car-journal/api-backend/internal/domain/auth/repository"
	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	fuelrepository "github.com/car-journal/api-backend/internal/domain/fuel/repository"
	fuelentryrepository "github.com/car-journal/api-backend/internal/domain/fuelentry/repository"
	odometerentryrepository "github.com/car-journal/api-backend/internal/domain/odometerentry/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	userprofilerepository "github.com/car-journal/api-backend/internal/domain/userprofile/repository"
)

type RepositoryContainer struct {
	Auth          authrepository.Interface
	Car           carrepository.Interface
	Fuel          fuelrepository.Interface
	FuelEntry     fuelentryrepository.Interface
	OdometerEntry odometerentryrepository.Interface
	User          userrepository.Interface
	UserProfile   userprofilerepository.Interface
}

func CreateRepositoryContainer(clientContainer ClientContainer) RepositoryContainer {
	return RepositoryContainer{
		Auth:          authrepository.Repository(),
		Car:           carrepository.Repository(),
		Fuel:          fuelrepository.Repository(),
		FuelEntry:     fuelentryrepository.Repository(),
		OdometerEntry: odometerentryrepository.Repository(),
		User:          userrepository.Repository(),
		UserProfile:   userprofilerepository.Repository(),
	}
}
