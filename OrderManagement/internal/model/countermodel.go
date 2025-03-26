package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ CounterModel = (*customCounterModel)(nil)

type (
	// CounterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCounterModel.
	CounterModel interface {
		counterModel
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
