package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShopLogic {
	return &ListShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListShopLogic) ListShop(req *types.ListShopReq) (resp *types.ListShopResp, err error) {
	// todo: add your logic here and delete this line

	return
}
