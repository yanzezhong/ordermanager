package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFormLogic {
	return &UserFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFormLogic) UserForm(req *types.UserFormReq) (resp *types.UserFormResp, err error) {
	// todo: add your logic here and delete this line

	return
}
