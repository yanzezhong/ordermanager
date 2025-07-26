package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Search(ctx context.Context, cond UserCond) ([]*User, int64, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string) UserModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}

// search user func

type UserCond struct {
	CreateAt *TimeRange
	PageParam
	Username string
	UpdateAt *TimeRange
}

// gencond

func (c *UserCond) genCond() bson.M {
	filter := make(bson.M)
	if c.Username != "" {
		filter["username"] = c.Username
	}
	if c.CreateAt != nil {
		filter["createAt"] = bson.M{
			"$gte": c.CreateAt.Start,
			"$lte": c.CreateAt.End,
		}
	}

	if c.UpdateAt != nil {
		filter["updateAt"] = bson.M{
			"$gte": c.CreateAt.Start,
			"$lte": c.CreateAt.End,
		}
	}

	return filter
}

func (m *customUserModel) Search(ctx context.Context, cond UserCond) ([]*User, int64, error) {
	var users []*User

	option := cond.GeneratePageOption()
	option.SetSort(bson.M{"_id": -1})

	filter := cond.genCond()
	err := m.conn.Find(ctx, &users, filter, option)
	if err != nil {
		return nil, 0, err
	}

	count, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}
