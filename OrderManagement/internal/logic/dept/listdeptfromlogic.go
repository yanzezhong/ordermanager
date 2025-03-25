package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeptFromLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeptFromLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptFromLogic {
	return &ListDeptFromLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeptFromLogic) ListDeptFrom(req *types.ListDeptFromReq) (resp *types.ListDeptFromResp, err error) {
	// todo: add your logic here and delete this line

	return
}
