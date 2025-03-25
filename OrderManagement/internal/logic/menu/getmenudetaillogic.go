package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuDetailLogic {
	return &GetMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuDetailLogic) GetMenuDetail(req *types.GetMenuDetailReq) (resp *types.GetMenuDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
