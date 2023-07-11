package internal

import (
	handlerfeature "github.com/cambiar_api/internal/handler/feature"
	"github.com/cambiar_api/internal/handler/feature/controller"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func RegisterRoutes(e *echo.Echo, container *dig.Container) error {
	err := container.Invoke(func(controller controller.IFeatureController) {
		handlerfeature.NewRouter(controller).RegisterRoutes(e)
	})
	if err != nil {
		return err
	}

	return nil
}
