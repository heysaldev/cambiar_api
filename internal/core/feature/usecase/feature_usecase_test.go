package usecase_test

import (
	"context"
	"testing"

	mock_port "github.com/cambiar_api/mocks/repository"

	"github.com/stretchr/testify/assert"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/usecase"
	"github.com/cambiar_api/internal/entity"
	"github.com/golang/mock/gomock"
)

func Test_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate object", func(t *testing.T) {
		repository := mock_port.NewMockIFeatureRepository(ctrl)
		usecase.NewFeatureUsecase(repository)
	})
}

func Test_GetAllWithQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mock_port.NewMockIFeatureRepository(ctrl)
	service := usecase.NewFeatureUsecase(repository)
	t.Run("Should return []entity.Feature when given GetAllWithQuerySpec with empty name and empty isEnabled", func(t *testing.T) {
		// Given
		spec := model.GetAllWithQuerySpec{
			Name:      nil,
			IsEnabled: nil,
		}
		// Mock
		resultMock := []entity.Feature{}
		resultMock = append(resultMock, entity.Feature{Name: "flag-1"})
		resultMock = append(resultMock, entity.Feature{Name: "flag-2"})
		repository.EXPECT().GetAllWithQuery(gomock.Any(), gomock.Any()).Return(resultMock)
		// Expected
		response := service.GetAllWithQuery(context.Background(), spec)
		assert.Equal(t, resultMock, response)
	})
}
