package usecase_test

import (
	"testing"

	"github.com/cambiar_api/internal/core/feature/usecase"
	mock_port "github.com/cambiar_api/mocks/repository"
	"github.com/golang/mock/gomock"
)

func Test_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate object", func(t *testing.T) {
		repository := mock_port.NewMockIFeatureRepository(ctrl)
		usecase.NewFeatureUsecase(usecase.FeatureUsecase{Repo: repository})
	})
}
