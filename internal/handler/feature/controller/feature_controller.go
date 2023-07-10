package controller

import (
	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/usecase"
	"github.com/labstack/echo/v4"
)

type FeatureController struct {
	Usecase usecase.IFeatureUsecase
}

type IFeatureController interface {
	GetAll(ctx echo.Context) error
	GetByName()
	Create()
	Update()
	Delete()
}

func NewFeatureController(uc usecase.IFeatureUsecase) IFeatureController {
	return &FeatureController{
		Usecase: uc,
	}
}

func (f *FeatureController) GetAll(ctx echo.Context) error {
	params := new(model.GetAllWithQuerySpec)

	if err := ctx.Bind(params); err != nil {
		ctx.JSON(400, err)
		return err
	}

	result, err := f.Usecase.GetAllWithQuery(ctx.Request().Context(), *params)
	if err != nil {
		ctx.JSON(400, err)
		return err
	}
	ctx.JSON(200, result)
	return nil
}

func (f *FeatureController) GetByName() {

}

func (f *FeatureController) Create() {
}

func (f *FeatureController) Update() {
}

func (f *FeatureController) Delete() {
}
