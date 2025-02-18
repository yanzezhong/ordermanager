package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ InvoiceDetailModel = (*customInvoiceDetailModel)(nil)

type (
	// InvoiceDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvoiceDetailModel.
	InvoiceDetailModel interface {
		invoiceDetailModel
		InsertMany(ctx context.Context, data []*InvoiceDetail) error
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
