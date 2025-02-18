// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productModel interface {
	Insert(ctx context.Context, data *Product) error
	FindOne(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, data *Product) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultProductModel struct {
	conn *mon.Model
}

func newDefaultProductModel(conn *mon.Model) *defaultProductModel {
	return &defaultProductModel{conn: conn}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) error {
	if data.ID == "" {
		data.ID = primitive.NewObjectID().String()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id string) (*Product, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Product

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultProductModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
