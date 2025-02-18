package order

import (
	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.PutOrderReq) error {

	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, req.ProductId)
	if err != nil {
		return errorx.OrderNotFoundError
	}

	if req.Address != "" {
		order.Address = req.Address
	}

	if req.Payment != 0 {
		if model.Payment(req.Payment).IsValid() {
			order.Payment = model.Payment(req.Payment)
		} else {
			return errorx.PaymentStatementInvalidError
		}
	}

	if req.State !=0{
		if model.State(req.State).IsValid() {
			order.State = model.State(req.State)
		} else {
			return errorx.StateInvalidError
		}
	}

	_, err = l.svcCtx.OrderModel.Update(l.ctx, order)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update order error: %v", err)
		return err
	}
	return nil
}
