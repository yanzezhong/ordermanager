package order

import (
	"OrderManagement/OrderManagement/internal/model"
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrderLogic) ListOrder(req *types.ListOrderRequest) (resp *types.ListOrderResp, err error) {
	orders, count, err := l.svcCtx.OrderModel.Search(l.ctx, model.OrderCond{
		State:     req.State,
		ShopId:    req.ShopId,
		ProductID: req.ProductID,
		PageParam: model.PageParam{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return nil, err
	}
	// orders to types.order
	resp = &types.ListOrderResp{
		Count: count,
		Items: convertOrder(orders),
	}
	return
}

func convertOrder(orders []*model.Order) []*types.Order {
	results := make([]*types.Order, len(orders))
	for _, order := range orders {
		o := &types.Order{
			OrderId:     order.ID,
			ShopId:      order.ShopId,
			ShopName:    order.ShopName,
			Address:     order.Address,
			DriverId:    order.DriverId,
			PurchaserId: order.PurchaserId,
			State:       order.State.Value(),
			Payment:     order.Payment.Value(),
			CreateTime:  order.CreateAt.UnixMilli(),
			UpdateTime:  order.UpdateAt.UnixMilli(),
		}
		//todo 重写

		// for _, product := range order.Products {
		// 	// o.Products = append(o.Products, &types.Products{
		// 	// 	ProductId:   product.ProductId,
		// 	// 	ProductName: product.ProductName,
		// 	// 	Price:       product.Price,
		// 	// 	Count:       product.Count,
		// 	// })
		// }

		results = append(results, o)
	}

	return results
}
