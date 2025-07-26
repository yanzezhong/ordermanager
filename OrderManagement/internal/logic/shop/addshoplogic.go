package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddShopLogic {
	return &AddShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddShopLogic) AddShop(req *types.PostShop) error {
	// todo: add your logic here and delete this line

	return nil
}
