package model

import (
	"time"
)

type InvoiceDetail struct {
	ID                 string    `bson:"_id,omitempty" json:"id,omitempty"`
	DocumentDate       string    `bson:"documentDate" json:"documentDate"`             // 单据日期，记录单据的创建日期
	DocumentNumber     string    `bson:"documentNumber" json:"documentNumber"`         // 单据编号，唯一标识单据的编号
	DocumentType       string    `bson:"documentType" json:"documentType"`             // 单据类型，例如销售发票、采购订单等
	Customer           string    `bson:"customer" json:"customer"`                     // 客户名称，记录与单据相关的客户
	CustomerLevel      string    `bson:"customerLevel" json:"customerLevel"`           // 客户级别，例如普通客户、VIP客户等
	SourceOrder        string    `bson:"sourceOrder" json:"sourceOrder"`               // 来源订单，记录单据的来源订单编号
	Handler            string    `bson:"handler" json:"handler"`                       // 经手人，记录处理单据的人员
	ProductName        string    `bson:"productName" json:"productName"`               // 商品名称，记录销售的商品名称
	Specification      string    `bson:"specification" json:"specification"`           // 规格，记录商品的规格信息
	SalesQuantity      float64   `bson:"salesQuantity" json:"salesQuantity"`           // 销售数量，记录销售的商品数量
	SalesSpecification string    `bson:"salesSpecification" json:"salesSpecification"` // 销售规格，记录商品的销售规格
	UnitPrice          float64   `bson:"unitPrice" json:"unitPrice"`                   // 单价，记录商品的单价
	Amount             float64   `bson:"amount" json:"amount"`                         // 金额，记录商品的总金额
	SalesRevenue       float64   `bson:"salesRevenue" json:"salesRevenue"`             // 销售收入，记录销售的总收入
	Weight             float64   `bson:"weight" json:"weight"`                         // 重量（kg），记录商品的重量
	LineItemAttribute  string    `bson:"lineItemAttribute" json:"lineItemAttribute"`   // 商品行属性，记录商品行的特殊属性
	DetailRemark       string    `bson:"detailRemark" json:"detailRemark"`             // 明细备注，记录单据明细的备注信息
	DocumentRemark     string    `bson:"documentRemark" json:"documentRemark"`         // 单据备注，记录单据整体的备注信息
	Quantity           float64   `bson:"quantity" json:"quantity"`                     // 数量，记录商品的总数量
	TotalAmount        float64   `bson:"totalAmount" json:"totalAmount"`               // 金额，记录商品的总金额
	UpdateAt           time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt           time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
