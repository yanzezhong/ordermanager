package logic

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderManagementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderManagementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderManagementLogic {
	return &OrderManagementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderManagementLogic) OrderManagement(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
