package usecase

import (
	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
)

type FeatureUsecase struct {
	Repo port.IFeatureRepository
}

type IFeatureUsecase interface {
	GetAllWithQuery(spec model.GetAllWithQuerySpec) []entity.Feature
}

func NewFeatureUsecase(dependencies FeatureUsecase) IFeatureUsecase {
	return &dependencies
}

func (a *FeatureUsecase) GetAllWithQuery(spec model.GetAllWithQuerySpec) []entity.Feature {
	return []entity.Feature{}
}
