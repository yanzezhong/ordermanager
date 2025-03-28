package svc

import (
	"OrderManagement/OrderManagement/internal/config"
	"OrderManagement/OrderManagement/internal/model"
)

type ServiceContext struct {
	Config config.Config

	ShopModel    model.ShopModel
	OrderModel   model.OrderModel
	ProductModel model.ProductModel
	DeptModel    model.DeptModel
	MenuModel    model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ShopModel:    model.NewShopModel(c.Mongo.Url, c.Mongo.DB, model.CollectionShop),
		OrderModel:   model.NewOrderModel(c.Mongo.Url, c.Mongo.DB, model.CollectionOrder),
		ProductModel: model.NewProductModel(c.Mongo.Url, c.Mongo.DB, model.CollectionProduct),
		DeptModel:    model.NewDeptModel(c.Mongo.Url, c.Mongo.DB, model.CollectionDept),
		MenuModel:    model.NewMenuModel(c.Mongo.Url, c.Mongo.DB, model.CollectionMenu),
	}
}
