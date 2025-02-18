package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ShopModel = (*customShopModel)(nil)

type (
	// ShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShopModel.
	ShopModel interface {
		shopModel
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
