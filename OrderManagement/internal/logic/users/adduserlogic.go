package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdduserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdduserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdduserLogic {
	return &AdduserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdduserLogic) Adduser(req *types.AddUserReq) (resp *types.NormalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
