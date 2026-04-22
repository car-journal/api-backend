package container

import (
	authservice "github.com/car-journal/api-backend/internal/domain/auth/service"
	carservice "github.com/car-journal/api-backend/internal/domain/car/service"
	fuelservice "github.com/car-journal/api-backend/internal/domain/fuel/service"
	fuelentryservice "github.com/car-journal/api-backend/internal/domain/fuelentry/service"
	maintenancecategoryservice "github.com/car-journal/api-backend/internal/domain/maintenancecategory/service"
	maintenanceentryservice "github.com/car-journal/api-backend/internal/domain/maintenanceentry/service"
	odometerentryservice "github.com/car-journal/api-backend/internal/domain/odometerentry/service"
	userservice "github.com/car-journal/api-backend/internal/domain/user/service"
	userprofileservice "github.com/car-journal/api-backend/internal/domain/userprofile/service"
)

// ServiceContainer handle all service used in project
type ServiceContainer struct {
	Auth                authservice.Interface
	Car                 carservice.Interface
	Fuel                fuelservice.Interface
	FuelEntry           fuelentryservice.Interface
	MaintenanceCategory maintenancecategoryservice.Interface
	MaintenanceEntry    maintenanceentryservice.Interface
	OdometerEntry       odometerentryservice.Interface
	User                userservice.Interface
	UserProfile         userprofileservice.Interface
}

// CreateServiceContainer construct all services available in the app
func CreateServiceContainer(repoContainer RepositoryContainer, clientContainer ClientContainer) ServiceContainer {
	// application service
	return ServiceContainer{
		Auth:                authservice.Service(repoContainer.Auth, clientContainer.Hash, repoContainer.User, repoContainer.UserProfile, clientContainer.UUID),
		Car:                 carservice.Service(repoContainer.Car, repoContainer.FuelEntry, repoContainer.MaintenanceEntry, repoContainer.User, clientContainer.UUID),
		Fuel:                fuelservice.Service(repoContainer.Fuel),
		FuelEntry:           fuelentryservice.Service(repoContainer.Car, repoContainer.FuelEntry, repoContainer.OdometerEntry, repoContainer.User, clientContainer.UUID),
		MaintenanceCategory: maintenancecategoryservice.Service(repoContainer.MaintenanceCategory, clientContainer.UUID),
		MaintenanceEntry:    maintenanceentryservice.Service(repoContainer.MaintenanceCategory, repoContainer.MaintenanceEntry, repoContainer.OdometerEntry, clientContainer.UUID),
		OdometerEntry:       odometerentryservice.Service(repoContainer.Car, repoContainer.FuelEntry, repoContainer.OdometerEntry, clientContainer.UUID),
		User:                userservice.Service(clientContainer.Hash, repoContainer.User, repoContainer.UserProfile, clientContainer.UUID),
		UserProfile:         userprofileservice.Service(repoContainer.UserProfile, clientContainer.UUID),
	}
}
