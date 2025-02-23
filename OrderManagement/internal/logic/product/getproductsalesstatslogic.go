package product

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSalesStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductSalesStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSalesStatsLogic {
	return &GetProductSalesStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductSalesStatsLogic) GetProductSalesStats(req *types.ProductSalesTrendReq) (resp *types.ProductSalesTrendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
