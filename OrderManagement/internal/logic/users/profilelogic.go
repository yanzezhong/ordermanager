package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile(req *types.UserProfileReq) (resp *types.UserFormResp, err error) {

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	resp = &types.UserFormResp{
		Code: "200",
		Data: &types.UserForm{
			Avatar:   user.Avatar,
			Email:    user.Email,
			Gender:   user.Gender,
			Mobile:   user.Mobile,
			Nickname: user.Nickname,
			OpenID:   user.OpenID,
			Status:   int64(user.Status),
			ID:       user.ID,
		},
		Msg: "success",
	}
	return
}
