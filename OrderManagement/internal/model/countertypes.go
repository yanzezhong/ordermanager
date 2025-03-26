package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Counter struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ModelName string             `bson:"modelName,omitempty" json:"modelName,omitempty"`
	Counter   int64              `bson:"counter,omitempty" json:"counter,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
