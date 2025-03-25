package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRoutesLogic {
	return &GetMenuRoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuRoutesLogic) GetMenuRoutes(req *types.GetMenuRoutesReq) (resp *types.GetMenuRoutesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
