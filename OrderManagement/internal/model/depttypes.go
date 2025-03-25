package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dept struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name,omitempty"`
	ParentID int64              `json:"parentId"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
