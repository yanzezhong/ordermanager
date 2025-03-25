package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
