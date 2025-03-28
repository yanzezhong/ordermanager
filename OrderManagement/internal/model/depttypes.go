package model

import (
	"time"
)

type Dept struct {
	ID       int64     `bson:"_id" json:"id"`
	Code     string    `bson:"code" json:"code"`
	Name     string    `bson:"name" json:"name"`
	ParentID int64     `bson:"parentId" json:"parentId"`
	Status   int64     `bson:"status" json:"status"`
	Sort     int64     `bson:"sort" json:"sort" `
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
