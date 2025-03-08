package model

import (
	"time"
)

type UserState int
type User struct {
	ID int64 `bson:"_id" json:"id,omitempty"`
	// TODO: Fill your own fields
	UserName string    `bson:"userName" json:"userName,omitempty"`
	Password string    `bson:"password" json:"password,omitempty"`
	Phone    string    `bson:"phone" json:"phone,omitempty"`
	State    UserState `bson:"state" json:"state,omitempty"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
