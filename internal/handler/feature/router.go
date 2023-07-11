package handlerfeature

import (
	"github.com/cambiar_api/internal/handler/feature/controller"
	"github.com/labstack/echo/v4"
)

type Router struct {
	app controller.IFeatureController
}

func NewRouter(app controller.IFeatureController) *Router {
	return &Router{
		app: app,
	}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/features", r.app.GetAll)
}
