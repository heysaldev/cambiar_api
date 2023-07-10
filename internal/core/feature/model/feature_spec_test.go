package model_test

import (
	"testing"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_ToBsonM(t *testing.T) {
	t.Run("Should return bsonM contain name and is_enabled when given GetAllWithQuerySpec contain name and is_enabled", func(t *testing.T) {
		// Given
		name := "flag-test-1"
		isEnabled := true
		spec := model.GetAllWithQuerySpec{
			Name:      &name,
			IsEnabled: &isEnabled,
		}
		// Expected
		expected := bson.M{}
		expected["name"] = name
		expected["is_enabled"] = isEnabled
		// Result
		result := spec.ToBsonM()
		assert.Equal(t, expected, result)
	})
	t.Run("Should return bsonM contain name only when given GetAllWithQuerySpec contain name only", func(t *testing.T) {
		// Given
		name := "flag-test-1"
		spec := model.GetAllWithQuerySpec{
			Name:      &name,
			IsEnabled: nil,
		}
		// Expected
		expected := bson.M{}
		expected["name"] = name
		// Result
		result := spec.ToBsonM()
		assert.Equal(t, expected, result)
	})
	t.Run("Should return bsonM contain is_enabled only when given GetAllWithQuerySpec contain is_enabled only", func(t *testing.T) {
		// Given
		isEnabled := true
		spec := model.GetAllWithQuerySpec{
			Name:      nil,
			IsEnabled: &isEnabled,
		}
		// Expected
		expected := bson.M{}
		expected["is_enabled"] = isEnabled
		// Result
		result := spec.ToBsonM()
		assert.Equal(t, expected, result)
	})
}
