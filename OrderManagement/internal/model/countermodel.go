package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ CounterModel = (*customCounterModel)(nil)

type (
	// CounterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCounterModel.
	CounterModel interface {
		counterModel
		FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (int64, error)
	}

	customCounterModel struct {
		*defaultCounterModel
	}
)

// NewCounterModel returns a model for the mongo.
func NewCounterModel(url, db, collection string) CounterModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCounterModel{
		defaultCounterModel: newDefaultCounterModel(conn),
	}
}

// FindOneAndUpdate 根据条件查找并更新文档，并返回更新后的结果
func (m *customCounterModel) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (counter int64, err error) {
	var result Counter
	opts := options.FindOneAndUpdate().SetUpsert(true)
	if err := m.conn.FindOneAndUpdate(ctx, &result, filter, update, opts); err != nil {
		return 0, err
	}
	// 日志

	return result.Counter, nil
}
