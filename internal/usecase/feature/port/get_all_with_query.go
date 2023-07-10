package port

import "github.com/cambiar_api/internal/entity"

type GetAllWithQuerySpec struct {
	Name      *string
	IsEnabled *bool
}

type GetAllWithQuery interface {
	Execute(query GetAllWithQuerySpec) []entity.Feature
}
