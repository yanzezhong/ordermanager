package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditDeptLogic {
	return &EditDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditDeptLogic) EditDept(req *types.EditDeptReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
