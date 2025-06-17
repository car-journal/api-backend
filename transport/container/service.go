package container

import (
	authservice "github.com/car-journal/internal/domain/auth/service"
	userservice "github.com/car-journal/internal/domain/user/service"
	userprofileservice "github.com/car-journal/internal/domain/userprofile/service"
)

// ServiceContainer handle all service used in project
type ServiceContainer struct {
	Auth        authservice.Interface
	User        userservice.Interface
	UserProfile userprofileservice.Interface
}

// CreateServiceContainer construct all services available in the app
func CreateServiceContainer(repoContainer RepositoryContainer, clientContainer ClientContainer) ServiceContainer {
	// application service
	return ServiceContainer{
		Auth:        authservice.Service(repoContainer.Auth, clientContainer.Hash, repoContainer.User, repoContainer.UserProfile, clientContainer.UUID),
		User:        userservice.Service(clientContainer.Hash, repoContainer.User, repoContainer.UserProfile, clientContainer.UUID),
		UserProfile: userprofileservice.Service(repoContainer.UserProfile, clientContainer.UUID),
	}
}
