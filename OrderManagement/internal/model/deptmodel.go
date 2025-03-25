package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ DeptModel = (*customDeptModel)(nil)

type (
	// DeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeptModel.
	DeptModel interface {
		deptModel
	}

	customDeptModel struct {
		*defaultDeptModel
	}
)

// NewDeptModel returns a model for the mongo.
func NewDeptModel(url, db, collection string) DeptModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customDeptModel{
		defaultDeptModel: newDefaultDeptModel(conn),
	}
}
