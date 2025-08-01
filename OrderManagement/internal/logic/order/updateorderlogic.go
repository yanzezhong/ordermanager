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

	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, req.OrderId)
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

	if req.State != 0 {
		if model.State(req.State).IsValid() {
			order.State = model.State(req.State)
		} else {
			return errorx.StateInvalidError
		}
	}

	if req.Items != nil {
		data := make([]*model.Products, 0)
		for _, product := range req.Items {
			// 检查产品是否存在
			_, err := l.svcCtx.ProductModel.FindOne(l.ctx, product.ProductId)
			if err != nil {
				logx.WithContext(l.ctx).Errorf("product not found: %s, error: %v", product.ProductId, err)
				return errorx.ProductNotFoundError
			}

			products := &model.Products{
				Product: &model.Product{
					ID:            product.ProductId,
					Name:          product.ProductName,
					Price:         model.Price{Terminal: product.Price}, // 使用终端价格
					Specification: 0,                                    // 默认值
					IsActive:      true,                                 // 默认值
					Image:         "",
					NickName:      "",
					BarCode:       "",
				},
				Count:        float64(product.Count),
				SalesRevenue: product.Price * float64(product.Count),
			}
			data = append(data, products)
		}
		order.Products = data
	}

	if req.Address != "" {
		order.Address = req.Address
	}

	_, err = l.svcCtx.OrderModel.Update(l.ctx, order)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update order error: %v", err)
		return err
	}
	return nil
}
