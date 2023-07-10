package adapter

import (
	"context"
	"fmt"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/port"
	"github.com/cambiar_api/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	DB *mongo.Collection
}

func NewMongo(db *mongo.Database) port.IFeatureRepository {
	return &Mongo{
		DB: db.Collection("features"),
	}
}

func (m *Mongo) GetAllWithQuery(ctx context.Context, spec model.GetAllWithQuerySpec) []entity.Feature {
	filter := spec.ToBsonM()
	options := options.Find()
	cursor, err := m.DB.Find(context.TODO(), filter, options)
	if err != nil {
		fmt.Println("Failed to find documents:", err)
		return []entity.Feature{}
	}
	defer cursor.Close(context.TODO())
	var features []entity.Feature
	for cursor.Next(context.TODO()) {
		var feature entity.Feature
		err := cursor.Decode(&feature)
		if err != nil {
			fmt.Println("Failed to decode document:", err)
			return []entity.Feature{}
		}
		features = append(features, feature)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return []entity.Feature{}
	}
	return features
}
