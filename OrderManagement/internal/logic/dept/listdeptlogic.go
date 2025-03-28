package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptLogic {
	return &ListDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeptLogic) ListDept(req *types.ListDeptReq) (resp *types.ListDeptVOResp, err error) {

	// change stauts to int64 1->正常；0->禁用

	cond := model.DeptCond{
		Keywords: req.Keywords,
		Status:   convertStatus(req.Status),
	}
	depts, _, err := l.svcCtx.DeptModel.Search(l.ctx, cond)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp = &types.ListDeptVOResp{
		Code: "200",
		Data: convertDeptsToDeptVOs(depts),
		Msg:  "success",
	}
	return
}

// function convert status to int64 1->正常；0->禁用
func convertStatus(status string) int64 {
	if status == "正常" {
		return 1
	} else if status == "禁用" {
		return 0
	}
	return 0
}

// convertDeptsToDeptVOs
func convertDeptsToDeptVOs(depts []*model.Dept) []types.DeptVO {
	deptVOs := make([]types.DeptVO, 0, len(depts))
	for _, dept := range depts {
		deptVO := types.DeptVO{
			ID:   &dept.ID,
			Name: &dept.Name,
			Code: &dept.Code,
		}
		deptVOs = append(deptVOs, deptVO)

	}
	return deptVOs
}
