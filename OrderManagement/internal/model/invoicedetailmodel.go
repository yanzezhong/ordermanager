package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ InvoiceDetailModel = (*customInvoiceDetailModel)(nil)

type (
	// InvoiceDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvoiceDetailModel.
	InvoiceDetailModel interface {
		invoiceDetailModel
		InsertMany(ctx context.Context, data []*InvoiceDetail) error
		Search(ctx context.Context, cond *InvoiceDetailCond) ([]*InvoiceDetail, int64, error)
	}

	customInvoiceDetailModel struct {
		*defaultInvoiceDetailModel
	}
)

// NewInvoiceDetailModel returns a model for the mongo.
func NewInvoiceDetailModel(url, db, collection string) InvoiceDetailModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customInvoiceDetailModel{
		defaultInvoiceDetailModel: newDefaultInvoiceDetailModel(conn),
	}
}

func (m *defaultInvoiceDetailModel) InsertMany(ctx context.Context, data []*InvoiceDetail) error {

	results := []any{}
	for _, v := range data {
		if v.ID == "" {
			v.ID = primitive.NewObjectID().String()
		}

		v.CreateAt = time.Now()
		v.UpdateAt = time.Now()
		results = append(results, v)
	}

	_, err := m.conn.InsertMany(ctx, results)
	return err
}

// InvoiceDetailCond 定义查询条件结构体
type InvoiceDetailCond struct {
	DocumentDate       string
	DocumentNumber     string
	DocumentType       string
	Customer           string
	CustomerLevel      string
	SourceOrder        string
	Handler            string
	ProductName        string
	Specification      string
	SalesQuantity      *float64
	SalesSpecification string
	UnitPrice          *float64
	Amount             *float64
	SalesRevenue       *float64
	Weight             *float64
	LineItemAttribute  string
	DetailRemark       string
	DocumentRemark     string
	Quantity           *float64
	TotalAmount        *float64
	UpdateAt           *TimeRange
	CreateAt           *TimeRange
	PageParam
}

// genCond 根据查询条件结构体生成 MongoDB 的 bson.M 查询条件
func (m *InvoiceDetailCond) genCond() bson.M {
	cond := bson.M{}

	if m.DocumentDate != "" {
		cond["documentDate"] = m.DocumentDate
	}
	if m.DocumentNumber != "" {
		cond["documentNumber"] = m.DocumentNumber
	}
	if m.DocumentType != "" {
		cond["documentType"] = m.DocumentType
	}
	if m.Customer != "" {
		cond["customer"] = m.Customer
	}
	if m.CustomerLevel != "" {
		cond["customerLevel"] = m.CustomerLevel
	}
	if m.SourceOrder != "" {
		cond["sourceOrder"] = m.SourceOrder
	}
	if m.Handler != "" {
		cond["handler"] = m.Handler
	}
	if m.ProductName != "" {
		cond["productName"] = m.ProductName
	}
	if m.Specification != "" {
		cond["specification"] = m.Specification
	}
	if m.SalesQuantity != nil {
		cond["salesQuantity"] = *m.SalesQuantity
	}
	if m.SalesSpecification != "" {
		cond["salesSpecification"] = m.SalesSpecification
	}
	if m.UnitPrice != nil {
		cond["unitPrice"] = *m.UnitPrice
	}
	if m.Amount != nil {
		cond["amount"] = *m.Amount
	}
	if m.SalesRevenue != nil {
		cond["salesRevenue"] = *m.SalesRevenue
	}
	if m.Weight != nil {
		cond["weight"] = *m.Weight
	}
	if m.LineItemAttribute != "" {
		cond["lineItemAttribute"] = m.LineItemAttribute
	}
	if m.DetailRemark != "" {
		cond["detailRemark"] = m.DetailRemark
	}
	if m.DocumentRemark != "" {
		cond["documentRemark"] = m.DocumentRemark
	}
	if m.Quantity != nil {
		cond["quantity"] = *m.Quantity
	}
	if m.TotalAmount != nil {
		cond["totalAmount"] = *m.TotalAmount
	}
	// 处理 UpdateAt 时间范围查询
	if m.UpdateAt.Start != "" || m.UpdateAt.End != "" {
		timeCond := bson.M{}
		if m.UpdateAt.Start != "" {
			timeCond["$gte"] = m.UpdateAt.Start
		}
		if m.UpdateAt.End != "" {
			timeCond["$lte"] = m.UpdateAt.End
		}
		cond["updateAt"] = timeCond
	}

	// 处理 CreateAt 时间范围查询
	if m.CreateAt.Start != "" || m.CreateAt.End != "" {
		timeCond := bson.M{}
		if m.CreateAt.Start != "" {
			timeCond["$gte"] = m.CreateAt.Start
		}
		if m.CreateAt.End != "" {
			timeCond["$lte"] = m.CreateAt.End
		}
		cond["createAt"] = timeCond
	}

	return cond
}

func (m *defaultInvoiceDetailModel) Search(ctx context.Context, cond *InvoiceDetailCond) ([]*InvoiceDetail, int64, error) {

	query := cond.genCond()

	option := cond.GeneratePageOption()
	option.SetSort(bson.M{"_id": -1})

	var data []*InvoiceDetail

	_ = m.conn.Find(ctx, query, data, option)

	count, err := m.conn.CountDocuments(ctx, query)
	switch err {
	case nil:
		return data, count, nil
	case mon.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}
