package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuOptionsLogic {
	return &GetMenuOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuOptionsLogic) GetMenuOptions(req *types.GetMenuOptionsReq) (resp *types.GetMenuOptionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
