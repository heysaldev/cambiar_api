package adapter

import (
	"github.com/cambiar_api/internal/entity"
	"github.com/cambiar_api/internal/usecase/feature/port"
)

type getAllWithQuery struct {
}

func NewGetAllWithQuery(dependencies getAllWithQuery) port.GetAllWithQuery {
	return &dependencies
}

func (a *getAllWithQuery) Execute(spec port.GetAllWithQuerySpec) []entity.Feature {
	return []entity.Feature{}
}
