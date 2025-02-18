package model

import (
	"time"
)

type Product struct {
	ID string `bson:"_id" json:"id"`
	// 商品名称
	Name string `bson:"name" json:"name"`
	// 商品价格
	Price Price `bson:"price" json:"price"`
	// 规格
	Specification int `bson:"specification" json:"specification"`
	// 是否生效
	IsActive bool `bson:"isActive,omitempty" json:"isActive,omitempty"`
	// 图片
	Image string `bson:"image,omitempty" json:"image,omitempty"`
	// 标签
	Tag string `bson:"tag,omitempty" json:"tag,omitempty"`
	// 品牌
	BrandId string `bson:"brandId" json:"brandId" `
	// 别名
	NickName string `bson:"nickName,omitempty" json:"nickName,omitempty"`
	//条码
	BarCode  string    `bson:"barCode" json:"barCode"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

// Price 价格应该写在订单上,该Price只用作展示
type Price struct {
	Terminal  float64 `bson:"terminal,omitempty" json:"terminal,omitempty"`   // 终端零售价
	WholeSale float64 `bson:"wholesale,omitempty" json:"wholesale,omitempty"` // 批发零售价
	Cost      float64 `bson:"cost,omitempty" json:"cost,omitempty"`           // 进价
	SRP       float64 `bson:"srp,omitempty" json:"srp,omitempty"`             // 建议零售价 Suggested Retail Price
	Warning   float64 `bson:"warning,omitempty" json:"warning,omitempty"`     // 报警价格
}
