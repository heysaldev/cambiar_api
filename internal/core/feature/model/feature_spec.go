package model

import "go.mongodb.org/mongo-driver/bson"

type GetAllWithQuerySpec struct {
	Name      *string `query:"name"`
	IsEnabled *bool   `query:"is_enabled"`
}

func (g *GetAllWithQuerySpec) ToBsonM() bson.M {
	filter := bson.M{}
	if g.Name != nil {
		filter["name"] = *g.Name
	}
	if g.IsEnabled != nil {
		filter["is_enabled"] = *g.IsEnabled
	}
	return filter
}
