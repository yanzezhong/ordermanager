package model

import (
	"time"
)

type CustomerType string

type Shop struct {
	ID string `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	ShopName      string    `bson:"shopName,omitempty" json:"shopName,omitempty"`
	CustomerLevel string    `bson:"customerLevel,omitempty" json:"customerLevel,omitempty"`
	Address       string    `bson:"address,omitempty" json:"address,omitempty"`
	ShopNameMD5   string    `bson:"shopNameMD5,omitempty" json:"shopNameMD5,omitempty"`
	PhoneNumber   string    `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	UpdateAt      time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt      time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
