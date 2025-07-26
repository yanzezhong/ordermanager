package model

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ShopModel = (*customShopModel)(nil)

type (
	// ShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShopModel.
	ShopModel interface {
		shopModel
		InsertMany(ctx context.Context, data []*Shop) (*mongo.InsertManyResult, error)
		Search(ctx context.Context, cond *ShopCond) ([]*Shop, int64, error)
	}

	customShopModel struct {
		*defaultShopModel
	}
)

// NewShopModel returns a model for the mongo.
func NewShopModel(url, db, collection string) ShopModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customShopModel{
		defaultShopModel: newDefaultShopModel(conn),
	}
}

// ShopCond 店铺查询条件结构体
type ShopCond struct {
	ShopName     string    `json:"shopName"`
	CustomerType string    `json:"customerType"`
	Address      string    `json:"address"`
	PhoneNumber  string    `json:"phoneNumber"`
	UpdataAt     TimeRange `json:"UpdataAt"`
	CreateAt     TimeRange `json:"CreateAt"`
	PageParam
}

// genCond 根据查询条件结构体生成 MongoDB 的 bson.M 查询条件
func (m *ShopCond) genCond() bson.M {
	cond := bson.M{}

	if m.ShopName != "" {
		// 生成店名的 MD5 值
		hash := md5.Sum([]byte(m.ShopName))
		shopNameMD5 := hex.EncodeToString(hash[:])
		cond["shopNameMD5"] = shopNameMD5
	}
	if m.CustomerType != "" {
		cond["customerType"] = m.CustomerType
	}
	if m.Address != "" {
		cond["address"] = m.Address
	}
	if m.PhoneNumber != "" {
		cond["phoneNumber"] = m.PhoneNumber
	}

	// 处理 UpdateAt 时间范围查询
	if m.UpdataAt.Start != "" || m.UpdataAt.End != "" {
		timeCond := bson.M{}
		if m.UpdataAt.Start != "" {
			startTime, _ := time.Parse(time.RFC3339, m.UpdataAt.Start)
			timeCond["$gte"] = startTime
		}
		if m.UpdataAt.End != "" {
			endTime, _ := time.Parse(time.RFC3339, m.UpdataAt.End)
			timeCond["$lte"] = endTime
		}
		cond["updateAt"] = timeCond
	}

	// 处理 CreateAt 时间范围查询
	if m.CreateAt.Start != "" || m.CreateAt.End != "" {
		timeCond := bson.M{}
		if m.CreateAt.Start != "" {
			startTime, _ := time.Parse(time.RFC3339, m.CreateAt.Start)
			timeCond["$gte"] = startTime
		}
		if m.CreateAt.End != "" {
			endTime, _ := time.Parse(time.RFC3339, m.CreateAt.End)
			timeCond["$lte"] = endTime
		}
		cond["createAt"] = timeCond
	}

	return cond
}

// defaultShopModel 默认店铺模型结构体

// Search 根据条件搜索店铺
func (m *defaultShopModel) Search(ctx context.Context, cond *ShopCond) ([]*Shop, int64, error) {

	options := cond.GeneratePageOption()
	options.SetSort(bson.M{"_id": -1})
	// 生成查询条件
	query := cond.genCond()

	var data []*Shop
	_ = m.conn.Find(ctx, query, data)

	count, err := m.conn.CountDocuments(ctx, query)

	switch err {
	case nil:
		return data, count, nil
	case mongo.ErrNoDocuments:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

// InsertMany 批量插入多个 Shop 记录
func (m *defaultShopModel) InsertMany(ctx context.Context, data []*Shop) (*mongo.InsertManyResult, error) {
	now := time.Now()
	documents := make([]interface{}, len(data))

	for _, shop := range data {
		if shop.ID == "" {
			shop.ID = primitive.NewObjectID().String()
		}
		shop.CreateAt = now
		shop.UpdateAt = now
		documents = append(documents, shop)
	}

	return m.conn.InsertMany(ctx, documents, options.InsertMany())
}
