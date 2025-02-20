package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		InsertMany(ctx context.Context, data []*Product) (*mongo.InsertManyResult, error)
		Search(ctx context.Context, cond ProductCond) ([]*Order, int64, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the mongo.
func NewProductModel(url, db, collection string) ProductModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProductModel{
		defaultProductModel: newDefaultProductModel(conn),
	}
}

type ProductCond struct {
	Id        string
	Name      string
	Terminal  float64
	WholeSale float64
	Cost      float64
	SRP       float64
	Warning   float64
	IsActive  *bool
	BarCode   string
	NickName  string
	UpdateAt  *TimeRange
	CreateAt  *TimeRange
	PageParam
}

func (c *ProductCond) genCond() bson.M {
	filter := bson.M{}
	if c.Id != "" {
		filter["_id"] = c.Id
	}
	if c.Name != "" {
		filter["name"] = primitive.Regex{Pattern: c.Name, Options: "i"}
	}
	if c.Terminal != 0 {
		filter["price.terminal"] = c.Terminal
	}
	if c.WholeSale != 0 {
		filter["price.wholesale"] = c.WholeSale
	}
	if c.Cost != 0 {
		filter["price.cost"] = c.Cost
	}
	if c.SRP != 0 {
		filter["price.srp"] = c.SRP
	}
	if c.Warning != 0 {
		filter["price.warning"] = c.Warning
	}
	if c.IsActive != nil {
		filter["isActive"] = c.IsActive
	}
	if c.BarCode != "" {
		filter["barCode"] = c.BarCode
	}
	if c.NickName != "" {
		filter["nickName"] = primitive.Regex{Pattern: c.NickName, Options: "i"}
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

func (m *customProductModel) Search(ctx context.Context, cond ProductCond) ([]*Order, int64, error) {
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

// InsertMany 批量插入多个 order 记录
func (m *customProductModel) InsertMany(ctx context.Context, data []*Product) (*mongo.InsertManyResult, error) {
	now := time.Now()
	documents := make([]interface{}, len(data))

	for _, product := range data {
		if product.ID == "" {
			product.ID = primitive.NewObjectID().String()
		}
		product.CreateAt = now
		product.UpdateAt = now
		documents = append(documents, product)
	}

	return m.conn.InsertMany(ctx, documents, options.InsertMany())
}
