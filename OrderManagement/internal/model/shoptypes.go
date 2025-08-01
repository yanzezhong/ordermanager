package model

import (
	"time"
)

type CustomerType string

type Shop struct {
	ID string `bson:"_id,omitempty" json:"id,omitempty"`

	// 店铺基本信息
	ShopName      string       `bson:"shopName,omitempty" json:"shopName,omitempty"`
	CustomerLevel CustomerType `bson:"customerLevel,omitempty" json:"customerLevel,omitempty"`
	Address       string       `bson:"address,omitempty" json:"address,omitempty"`
	ShopNameMD5   string       `bson:"shopNameMD5,omitempty" json:"shopNameMD5,omitempty"`
	PhoneNumber   string       `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`

	// 时间信息
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	// 客户资料补充字段
	CustomerID       string `bson:"customerId,omitempty" json:"customer_id,omitempty"`             // 客户编号
	CustomerSource   string `bson:"customerSource,omitempty" json:"customer_source,omitempty"`     // 客户来源
	Category         string `bson:"category,omitempty" json:"category,omitempty"`                  // 所属分类
	SettlementMethod string `bson:"settlementMethod,omitempty" json:"settlement_method,omitempty"` // 结款方式
	Remarks          string `bson:"remarks,omitempty" json:"remarks,omitempty"`                    // 备注
	MnemonicCode     string `bson:"mnemonicCode,omitempty" json:"mnemonic_code,omitempty"`         // 助记码

	// 账期与财务信息
	CollectionPeriod    int       `bson:"collectionPeriod,omitempty" json:"collection_period,omitempty"`        // 收款期限
	CreditLimit         int       `bson:"creditLimit,omitempty" json:"credit_limit,omitempty"`                  // 信用额度
	ArrearsBalance      float64   `bson:"arrearsBalance,omitempty" json:"arrears_balance,omitempty"`            // 欠款余额
	PrepaymentBalance   float64   `bson:"prepaymentBalance,omitempty" json:"prepayment_balance,omitempty"`      // 预收款余额
	LastTransactionTime time.Time `bson:"lastTransactionTime,omitempty" json:"last_transaction_time,omitempty"` // 最近交易时间

	// 联系人
	Contacts []Contact `bson:"contacts,omitempty" json:"contacts,omitempty"`

	// 预留定位字段
	Longitude    float64   `bson:"longitude,omitempty" json:"longitude,omitempty"`        // 经度
	Latitude     float64   `bson:"latitude,omitempty" json:"latitude,omitempty"`          // 纬度
	AdCode       string    `bson:"adCode,omitempty" json:"ad_code,omitempty"`             // 行政区划编码
	LocationTime time.Time `bson:"locationTime,omitempty" json:"location_time,omitempty"` // 定位时间
}

type Contact struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty" json:"phone_number,omitempty"`
}
