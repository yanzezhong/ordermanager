package errorcode

/**
errorcode 占用情况：
*/

const (
	ErrPaymentStateInvalid = 1000001 // 支付状态无效
	ErrOrderNotFound       = 1000002 // 订单不存在
	ErrStateInvalid        = 1000003 // 订单状态无效

	ErrUserOrPasswordNotFound = 1010005

	ErrDBError = 500001
)
