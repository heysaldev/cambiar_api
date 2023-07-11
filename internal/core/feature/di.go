package corefeature

import (
	"github.com/cambiar_api/internal/common/database"
	"github.com/cambiar_api/internal/core/feature/repository/adapter"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/core/feature/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func ProvideFeatureMongo(db *mongo.Database) port.IFeatureRepository {
	return &adapter.FeatureMongo{
		Col: database.NewMongo(db.Collection("features")),
	}
}

func ProvideFeatureUsecase(featureRepo port.IFeatureRepository) usecase.IFeatureUsecase {
	return &usecase.FeatureUsecase{
		Repo: featureRepo,
	}
}

func BuildContainer(container *dig.Container) (*dig.Container, error) {
	err := container.Provide(ProvideFeatureUsecase)
	if err != nil {
		return nil, err
	}
	err = container.Provide(ProvideFeatureMongo)
	if err != nil {
		return nil, err
	}
	return container, nil
}
