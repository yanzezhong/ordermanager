package model

const (
	CollectionOrder         = "order"
	CollectionProduct       = "product"
	CollectionShop          = "shop"
	CollectionInvoiceDetail = "invoiceDetail"
)

// order
const (
	// StatePlaceOrder 下单未发货
	StatePlaceOrder State = 1
	// StateDelivering 发货中
	StateDelivering State = 2
	// StateDelivered 已送达
	StateDelivered State = 3
	// StateCancel 取消
	StateCancel State = 4
	// StateFinish 完成
	StateFinish State = 5

	/*   账期 不应该放在order 里 */

	// PaymentPaid 已支付
	PaymentPaid Payment = 1
	// PaymentUnpaid 未支付
	PaymentUnpaid Payment = 2

	// customerLevel

	CustomerLevelTermianl  CustomerType = "终端客户"
	CustomerLevelWholeSale CustomerType = "批发客户"
)

func (c CustomerType) String() string {
	return string(c)
}
func (s State) IsValid() bool {
	switch s {
	case StatePlaceOrder, StateDelivering, StateDelivered, StateCancel, StateFinish:
		return true
	}
	return false
}

func (s State) String() string {
	switch s {
	case StatePlaceOrder:
		return "下单未发货"
	case StateDelivering:
		return "发货中"
	case StateDelivered:
		return "已送达"
	case StateCancel:
		return "取消"
	case StateFinish:
		return "完成"
	}
	return "N/A"
}

func (s State) Value() int {
	return int(s)
}

func (p Payment) IsValid() bool {
	switch p {
	case PaymentPaid, PaymentUnpaid:
		return true
	}
	return false
}

func (p Payment) String() string {
	switch p {
	case PaymentPaid:
		return "已支付"
	case PaymentUnpaid:
		return "未支付"
	}
	return "N/A"
}

func (p Payment) Value() int {
	return int(p)
}
