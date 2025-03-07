package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShopLogic {
	return &UpdateShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShopLogic) UpdateShop(req *types.PostShop) error {
	// 暂时不写
	return nil
}
