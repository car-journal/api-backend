package container

import (
	authrepository "github.com/car-journal/internal/domain/auth/repository"
	userrepository "github.com/car-journal/internal/domain/user/repository"
	userprofilerepository "github.com/car-journal/internal/domain/userprofile/repository"
)

type RepositoryContainer struct {
	Auth        authrepository.Interface
	User        userrepository.Interface
	UserProfile userprofilerepository.Interface
}

func CreateRepositoryContainer(clientContainer ClientContainer) RepositoryContainer {
	return RepositoryContainer{
		Auth:        authrepository.Repository(),
		User:        userrepository.Repository(),
		UserProfile: userprofilerepository.Repository(),
	}
}
