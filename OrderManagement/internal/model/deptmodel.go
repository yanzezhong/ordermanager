package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ DeptModel = (*customDeptModel)(nil)

type (
	// DeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeptModel.
	DeptModel interface {
		deptModel
		Search(ctx context.Context, cond DeptCond) ([]*Dept, int64, error)
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

type DeptCond struct {
	// ID 定义了部门ID的条件。
	ID int64
	// Name 定义了部门名称的条件。
	Name string
	//
	Keywords string
	// ParentID 定义了父部门ID的条件。
	ParentID int64
	// Status 定义了部门状态的条件。
	Status int64
	// Sort 定义了部门排序的条件。
	Sort     int64
	CreateAt *TimeRange
	UpdateAt *TimeRange
	PageParam
}

// search

func (c *DeptCond) genCond() bson.M {
	filter := bson.M{}

	if c.ID > 0 {
		filter["_id"] = c.ID
	}
	if c.Name != "" {
		filter["name"] = c.Name
	}
	if c.ParentID > 0 {
		filter["parentId"] = c.ParentID
	}
	if c.Status > 0 {
		filter["status"] = c.Status
	}
	if c.Sort > 0 {
		filter["sort"] = c.Sort
	}

	if c.Keywords != "" {
		filter["$or"] = []bson.M{
			{"name": bson.M{"$regex": c.Keywords}},
			{"parentId": bson.M{"$regex": c.Keywords}},
			{"_id": bson.M{"$regex": c.Keywords}},
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
func (m *customDeptModel) Search(ctx context.Context, cond DeptCond) ([]*Dept, int64, error) {
	option := cond.GeneratePageOption()
	option.SetSort(bson.M{"_id": -1})

	var r []*Dept
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
