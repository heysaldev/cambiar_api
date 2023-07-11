package internal

import (
	corefeature "github.com/cambiar_api/internal/core/feature"
	handlerfeature "github.com/cambiar_api/internal/handler/feature"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func BuildContainer(database *mongo.Database) (*dig.Container, error) {
	container := dig.New()
	err := container.Provide(func() (*mongo.Database, error) {
		return database, nil
	})
	if err != nil {
		return nil, err
	}
	handlerfeature.BuildContainer(container)
	corefeature.BuildContainer(container)
	return container, nil
}
