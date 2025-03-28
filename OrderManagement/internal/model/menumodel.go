package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		Search(ctx context.Context, cond MenuCond) ([]*Menu, int64, error)
	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the mongo.
func NewMenuModel(url, db, collection string) MenuModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMenuModel{
		defaultMenuModel: newDefaultMenuModel(conn),
	}
}

type MenuCond struct {
	ID        int64      `json:"id"`
	ParentId  int64      `json:"parentId"`
	Name      string     `json:"name"`
	RouteName string     `json:"routeName"`
	RoutePath string     `json:"routePath"`
	Component string     `json:"component"`
	Perm      string     `json:"perm"`
	Visible   int32      `json:"visible"`
	Keywords  string     `json:"keywords"`
	Status    int32      `json:"status"`
	CreateAt  *TimeRange `json:"createAt"`
	UpdateAt  *TimeRange `json:"updateAt"`
	Page      int64      `json:"page"`
	PageSize  int64      `json:"pageSize"`
	PageParam
}

// 实现 Search 方法
func (m *customMenuModel) Search(ctx context.Context, cond MenuCond) ([]*Menu, int64, error) {
	option := cond.GeneratePageOption()
	option.SetSort(bson.M{"_id": -1})

	var r []*Menu
	filter := cond.genCond()
	err := m.conn.Find(ctx, &r, filter, option)
	if err != nil {
		return nil, 0, err
	}

	count, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return r, count, nil
}

func (c *MenuCond) genCond() bson.M {
	filter := bson.M{}

	if c.ID > 0 {
		filter["_id"] = c.ID
	}
	if c.ParentId > 0 {
		filter["parentId"] = c.ParentId
	}
	if c.Name != "" {
		filter["name"] = c.Name
	}
	if c.RouteName != "" {
		filter["routeName"] = c.RouteName
	}

	if c.Status > 0 {
		filter["status"] = c.Status
	}
	if c.RoutePath != "" {
		filter["routePath"] = c.RoutePath
	}
	if c.Component != "" {
		filter["component"] = c.Component
	}
	if c.Perm != "" {
		filter["perm"] = c.Perm
	}
	if c.Visible > 0 {
		filter["visible"] = c.Visible
	}

	if c.Keywords != "" {
		filter["$or"] = []bson.M{
			{"name": bson.M{"$regex": c.Keywords}},
			{"routeName": bson.M{"$regex": c.Keywords}},
			{"routePath": bson.M{"$regex": c.Keywords}},
			{"component": bson.M{"$regex": c.Keywords}},
		}
	}

	if c.CreateAt != nil {
		filter["createAt"] = bson.M{
			"$gte": c.CreateAt.Start,
			"$lte": c.CreateAt.End,
		}
	}
	if c.UpdateAt != nil {
		filter["updateAt"] = bson.M{
			"$gte": c.UpdateAt.Start,
			"$lte": c.UpdateAt.End,
		}
	}
	return filter
}
