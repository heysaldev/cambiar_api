package adapter

import (
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
)

type Mongo struct{}

func NewMongo(dependencies Mongo) port.IFeatureRepository {
	return &dependencies
}

func (m *Mongo) GetAllWithQuery() []entity.Feature {
	return []entity.Feature{}
}
