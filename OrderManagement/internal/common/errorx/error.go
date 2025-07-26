package errorx

import "OrderManagement/OrderManagement/internal/common/errorx/errorcode"

var (
	PaymentStatementInvalidError = NewCodeError(errorcode.ErrPaymentStateInvalid, "订单支付状态错误")
	StateInvalidError            = NewCodeError(errorcode.ErrStateInvalid, "订单状态错误")

	OrderNotFoundError = NewCodeError(errorcode.ErrOrderNotFound, "未找到当前订单")
	// 删除数量过多
	DeleteCountTooManyError = NewCodeError(errorcode.ErrDeleteCountTooMany, "删除数量过多")
	// no id
	NoIdError = NewCodeError(errorcode.ErrNoId, "未传入ID")

	UserOrPasswordNotFound = NewCodeError(errorcode.ErrUserOrPasswordNotFound, "用户名或密码错误")

	// "最多只能删除5个用户"
	MaxDeleteUserError = NewCodeError(errorcode.ErrMaxDeleteUser, "最多只能删除5个用户")
)
