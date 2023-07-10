package database

import (
	"context"
)

type DB interface {
	Find(ctx context.Context, filter map[string]interface{}) ([]map[string]interface{}, error)
	FindOne()
	Create()
	Update()
	Delete()
}
