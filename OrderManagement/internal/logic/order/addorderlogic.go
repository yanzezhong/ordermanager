package order

import (
	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderLogic) AddOrder(req *types.PostOrder) error {
	// 将req中的同名字段添加到Order中
	order := &model.Order{
		Address:  req.Address,
		ShopId:   req.ShopId,
		ShopName: req.ShopName,
	}

	products := make([]*model.Products, 0)

	for _, product := range req.Products {
		p := &model.Products{
			Count:       product.Count,
			Price:       product.Price,
			ProductId:   product.ProductId,
			ProductName: product.ProductName,
		}
		products = append(products, p)
	}

	order.Products = products
	order.PurchaserId = req.PurchaserId

	if model.Payment(req.Payment).IsValid() {
		order.Payment = model.Payment(req.Payment)
	} else {
		return errorx.PaymentStatementInvalidError
	}
	order.ID = primitive.NewObjectID().Hex()
	return l.svcCtx.OrderModel.Insert(l.ctx, order)
}
