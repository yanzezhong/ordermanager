package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuReqLogic {
	return &UpdateMenuReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuReqLogic) UpdateMenuReq(req *types.UpdateMenuReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
