package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeptReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeptReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptReqLogic {
	return &DeleteDeptReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeptReqLogic) DeleteDeptReq(req *types.ListDeptOptionsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
