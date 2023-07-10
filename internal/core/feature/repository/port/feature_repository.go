package port

import "github.com/cambiar_api/internal/entity"

type IFeatureRepository interface {
	GetAllWithQuery() []entity.Feature
}
