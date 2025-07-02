// Package container handles middleware, client, repository and service containers
package container

import "github.com/car-journal/api-backend/transport/http/container"

// AppContainer handle all requirement for app to run properly
type AppContainer struct {
	Middleware   container.MiddlewareContainer
	Clients      ClientContainer
	Repositories RepositoryContainer
	Services     ServiceContainer
}

// CreateAppContainer construct all requirement for app
func CreateAppContainer() AppContainer {
	clientContainer := CreateClientContainer()
	repoContainer := CreateRepositoryContainer(clientContainer)
	return AppContainer{
		Middleware:   container.CreateMiddlewareContainer(),
		Services:     CreateServiceContainer(repoContainer, clientContainer),
		Clients:      clientContainer,
		Repositories: repoContainer,
	}
}
