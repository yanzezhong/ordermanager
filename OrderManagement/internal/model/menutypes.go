package model

import (
	"time"
)

type KeyValue struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

type Menu struct {
	Authorization string     `bson:"authorization" header:"Authorization"`
	ID            int64      `bson:"id" json:"id"`
	ParentId      int64      `bson:"parentId" json:"parentId"`
	Name          string     `bson:"name" json:"name"`
	RouteName     string     `bson:"routeName" json:"routeName"`
	RoutePath     string     `bson:"routePath" json:"routePath"`
	Component     string     `bson:"component" json:"component"`
	Perm          string     `bson:"perm" json:"perm"`
	Visible       int32      `bson:"visible" json:"visible"`
	Sort          int32      `bson:"sort" json:"sort"`
	Icon          string     `bson:"icon" json:"icon"`
	Redirect      string     `bson:"redirect" json:"redirect"`
	KeepAlive     int32      `bson:"keepAlive" json:"keepAlive"`
	AlwaysShow    int32      `bson:"alwaysShow" json:"alwaysShow"`
	Params        []KeyValue `bson:"params" json:"params"`
	UpdateAt      time.Time  `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt      time.Time  `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
