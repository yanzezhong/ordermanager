package dept

import (
	"context"
	"strings"

	"OrderManagement/OrderManagement/internal/common/errorx"
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

func (l *DeleteDeptReqLogic) DeleteDeptReq(req *types.DeleteDeptReq) (resp *types.CommonResponse, err error) {

	ids := strings.Split(req.IDS, ",")

	if len(ids) > 10 {
		//
		return nil, errorx.DeleteCountTooManyError
	}
	for _, v := range ids {
		_, err = l.svcCtx.DeptModel.Delete(l.ctx, v)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
	}

	return
}
