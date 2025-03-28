package dept

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeptLogic) AddDept(req *types.AddDeptReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	// 生成插入逻辑
	// https://www.mongodb.com/resources/products/platform/mongodb-auto-increment 自增怎么写

	dept := model.Dept{
		ID:       req.ID,
		Code:     req.Code,
		Name:     req.Name,
		ParentID: req.ParentID,
		Status:   req.Status,
		Sort:     req.Sort,
	}

	err = l.svcCtx.DeptModel.Insert(l.ctx, &dept)
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
