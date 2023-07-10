package entity

import "time"

type Feature struct {
	Name      string    `json:"name" bson:"_id"`
	IsEnabled bool      `json:"is_enabled" bson:"is_enabled"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
