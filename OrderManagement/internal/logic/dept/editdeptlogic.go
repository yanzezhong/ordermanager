package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
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
	dept := model.Dept{
		ID:       req.ID,
		Name:     req.Name,
		ParentID: req.ParentID,
		Status:   req.Status,
		Sort:     req.Sort,
		Code:     req.Code,
	}

	_, err = l.svcCtx.DeptModel.Update(l.ctx, &dept)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp = &types.CommonResponse{
		Code: "200",
		Msg:  "success",
	}
	return
}
