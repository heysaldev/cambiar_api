package port

import (
	"context"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/entity"
)

type IFeatureRepository interface {
	GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) []entity.Feature
}
