package usecase

import (
	"context"
	"fmt"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
)

type FeatureUsecase struct {
	Repo port.IFeatureRepository
}

type IFeatureUsecase interface {
	GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) ([]entity.Feature, error)
}

func NewFeatureUsecase(repo port.IFeatureRepository) IFeatureUsecase {
	return &FeatureUsecase{
		Repo: repo,
	}
}

func (a *FeatureUsecase) GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) ([]entity.Feature, error) {
	response, err := a.Repo.GetAllWithQuery(ctx, spec)
	if err != nil {
		fmt.Println("Failed GetAllWithQuery:", err)
		return []entity.Feature{}, err
	}
	return response, nil
}
