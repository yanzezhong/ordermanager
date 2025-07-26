package model

import (
	"time"
)

type UserState int
type User struct {
	ID       int64     `bson:"_id" json:"id"`
	Avatar   string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
	DeptID   int64     `bson:"deptID,omitempty" json:"deptID,omitempty"`
	Email    string    `bson:"email,omitempty" json:"email,omitempty"`
	Gender   int64     `bson:"gender,omitempty" json:"gender,omitempty"`
	Mobile   string    `bson:"mobile,omitempty" json:"mobile,omitempty"`
	Nickname string    `bson:"nickname,omitempty" json:"nickname,omitempty"`
	OpenID   string    `bson:"openID,omitempty" json:"openID,omitempty"`
	RoleIDS  []string  `bson:"roleIDS,omitempty" json:"roleIDS,omitempty"`
	Status   UserState `bson:"status,omitempty" json:"status,omitempty"`
	UserName string    `bson:"username,omitempty" json:"username,omitempty"`
	Password string    `bson:"password,omitempty" json:"password,omitempty"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
