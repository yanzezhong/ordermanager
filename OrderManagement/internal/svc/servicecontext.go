package svc

import (
	"OrderManagement/OrderManagement/internal/config"
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/utils"
)

type ServiceContext struct {
	Config config.Config

	ShopModel    model.ShopModel
	OrderModel   model.OrderModel
	ProductModel model.ProductModel
	UserModel    model.UserModel
	TokenCache   *utils.DataCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ShopModel:    model.NewShopModel(c.Mongo.Url, c.Mongo.DB, model.CollectionShop),
		OrderModel:   model.NewOrderModel(c.Mongo.Url, c.Mongo.DB, model.CollectionOrder),
		ProductModel: model.NewProductModel(c.Mongo.Url, c.Mongo.DB, model.CollectionProduct),
		UserModel:    model.NewUserModel(c.Mongo.Url, c.Mongo.DB, model.CollectionProduct),
		TokenCache:   utils.NewDataCache(),
	}
}
