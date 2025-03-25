package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeptOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeptOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptOptionsLogic {
	return &ListDeptOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeptOptionsLogic) ListDeptOptions(req *types.ListDeptOptionsReq) (resp *types.ListDeptOptionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
