package usecase

import (
	"context"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
)

type FeatureUsecase struct {
	Repo port.IFeatureRepository
}

type IFeatureUsecase interface {
	GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) []entity.Feature
}

func NewFeatureUsecase(repo port.IFeatureRepository) IFeatureUsecase {
	return &FeatureUsecase{
		Repo: repo,
	}
}

func (a *FeatureUsecase) GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) []entity.Feature {
	response := a.Repo.GetAllWithQuery(ctx, spec)
	return response
}
