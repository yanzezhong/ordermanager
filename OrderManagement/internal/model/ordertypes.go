package model

import (
	"time"
)

type State int
type Payment int

type Order struct {
	ID string `bson:"_id" json:"id"`
	// 商品列表
	Products []*Products `bson:"products" json:"products"`
	// 商店Id
	ShopId string `bson:"shopId" json:"shopId"`
	// 商店名称
	ShopName string `bson:"shopName" json:"shopName"`
	// 收货地址
	Address string `bson:"address" json:"address"`
	// 配送状态
	State State `bson:"state" json:"state"`
	// 支付状态
	Payment Payment `bson:"payment" json:"payment"`
	// 下单人
	PurchaserId string `bson:"purchaserId" json:"purchaserId"`
	// 配送司机
	DriverId string `bson:"driverId" json:"driverId"`
	// picture 回单照片
	Picture string `bson:"picture" json:"picture"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type Products struct {
	// 商品ID
	ProductId string `bson:"productId" json:"productId"`
	// 商品名称
	ProductName string `bson:"productName" json:"productName"`
	// 商品价格
	Price  float64  `bson:"price" json:"price"`
	// 商品数量
	Count  int    `bson:"count" json:"count"`
}