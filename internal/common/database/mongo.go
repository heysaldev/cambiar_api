package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongod struct {
	Col *mongo.Collection
}

func NewMongo(col *mongo.Collection) DB {
	return &mongod{
		Col: col,
	}
}

func (m *mongod) Find(ctx context.Context, filter map[string]interface{}) ([]map[string]interface{}, error) {
	cursor, err := m.Col.Find(context.TODO(), filter)
	if err != nil {
		return []map[string]interface{}{}, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(context.TODO())
	var results []map[string]interface{}
	for cursor.Next(context.TODO()) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, fmt.Errorf("failed to decode document: %w", err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return []map[string]interface{}{}, fmt.Errorf("cursor error: %w", err)
	}
	fmt.Printf("Mongo query result: %+v", results)
	return results, nil
}
func (m *mongod) FindOne() {}
func (m *mongod) Create()  {}
func (m *mongod) Update()  {}
func (m *mongod) Delete()  {}
