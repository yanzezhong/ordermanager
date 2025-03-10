package errorx

import "OrderManagement/OrderManagement/internal/common/errorx/errorcode"

var (
	PaymentStatementInvalidError = NewCodeError(errorcode.ErrPaymentStateInvalid, "订单支付状态错误")
	StateInvalidError            = NewCodeError(errorcode.ErrStateInvalid, "订单状态错误")

	OrderNotFoundError = NewCodeError(errorcode.ErrOrderNotFound, "未找到当前订单")

	UserOrPasswordNotFound = NewCodeError(errorcode.ErrUserOrPasswordNotFound, "用户名或密码错误")
)
