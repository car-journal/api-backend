package containers

import (
	authrepository "github.com/car-journal/api-backend/internal/domain/auth/repository"
	carrepository "github.com/car-journal/api-backend/internal/domain/car/repository"
	userrepository "github.com/car-journal/api-backend/internal/domain/user/repository"
	userprofilerepository "github.com/car-journal/api-backend/internal/domain/userprofile/repository"
)

type RepositoryContainer struct {
	Auth        authrepository.Interface
	Car         carrepository.Interface
	User        userrepository.Interface
	UserProfile userprofilerepository.Interface
}

func CreateRepositoryContainer(clientContainer ClientContainer) RepositoryContainer {
	return RepositoryContainer{
		Auth:        authrepository.Repository(),
		Car:         carrepository.Repository(),
		User:        userrepository.Repository(),
		UserProfile: userprofilerepository.Repository(),
	}
}
