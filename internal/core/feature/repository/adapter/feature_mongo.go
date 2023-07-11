package adapter

import (
	"context"

	"github.com/cambiar_api/internal/common/database"
	"github.com/cambiar_api/internal/common/helper"
	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeatureMongo struct {
	Col database.DB
}

var EMPTY_DATA []entity.Feature = []entity.Feature{}

func NewFeatureMongo(db *mongo.Database) port.IFeatureRepository {

	return &FeatureMongo{
		Col: database.NewMongo(db.Collection("features")),
	}
}

func (m *FeatureMongo) GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) ([]entity.Feature, error) {
	filter := spec.ToBsonM()
	results, err := m.Col.Find(context.TODO(), filter)
	if err != nil {
		return EMPTY_DATA, err
	}
	var features []entity.Feature
	if err := helper.JSONToStruct(results, &features); err != nil {
		return EMPTY_DATA, err
	}
	return features, nil
}
