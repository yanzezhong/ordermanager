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

	// find one
	dept, err := l.svcCtx.DeptModel.FindOne(l.ctx, req.DeptID)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	// gen resp for all depts
	resp = &types.ListDeptFromResp{
		Code: "200",
		Data: &types.DeptForm{
			Code:     dept.Code,
			ID:       dept.ID,
			Name:     dept.Name,
			ParentID: dept.ParentID,
			Sort:     dept.Sort,
			Status:   dept.Status,
		},
		Msg: "success",
	}
	return
}
