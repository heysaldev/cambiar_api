package handlerfeature

import (
	"github.com/cambiar_api/internal/core/feature/usecase"
	"github.com/cambiar_api/internal/handler/feature/controller"
	"go.uber.org/dig"
)

func ProvideFeatureController(uc usecase.IFeatureUsecase) controller.IFeatureController {
	return &controller.FeatureController{
		Usecase: uc,
	}
}

func BuildContainer(container *dig.Container) (*dig.Container, error) {
	err := container.Provide(ProvideFeatureController)
	if err != nil {
		return nil, err
	}

	return container, nil
}
