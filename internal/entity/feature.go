package entity

type Feature struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	IsEnabled bool   `json:"is_enabled" bson:"is_enabled"`
}
