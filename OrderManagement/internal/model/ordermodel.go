package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		Search(ctx context.Context, cond OrderCond) ([]*Order, int64, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a Order for the mongo.
func NewOrderModel(url, db, collection string) OrderModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customOrderModel{
		defaultOrderModel: newDefaultOrderModel(conn),
	}
}

type OrderCond struct {
	State     int
	ShopId    string
	ProductID string
	UpdateAt  *TimeRange
	CreateAt  *TimeRange
	PageParam
}

func (c *OrderCond) genCond() bson.M {
	filter := bson.M{}

	if c.State > 0 {
		filter["state"] = c.State
	}
	if c.ShopId != "" {
		filter["shopId"] = c.ShopId
	}
	if c.ProductID != "" {
		filter["products.productId"] = c.ProductID
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
func (m *customOrderModel) Search(ctx context.Context, cond OrderCond) ([]*Order, int64, error) {
	option := cond.GeneratePageOption()
	option.SetSort(bson.M{"_id": -1})

	var r []*Order
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
