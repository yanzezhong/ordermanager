package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuVisibilityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuVisibilityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuVisibilityLogic {
	return &UpdateMenuVisibilityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuVisibilityLogic) UpdateMenuVisibility(req *types.UpdateMenuVisibilityReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
