package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSalesStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesStatsLogic {
	return &GetShopSalesStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopSalesStatsLogic) GetShopSalesStats(req *types.ShopSalesStatsReq) (resp *types.ShopSalesStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
