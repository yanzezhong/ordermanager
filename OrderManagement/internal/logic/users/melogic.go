package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	_, userId, err := utils.GetUserProfile(req.Authorization, l.svcCtx.Config)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	data := &types.UserInfoVO{
		Avatar:   user.Avatar,
		UserID:   userId,
		Nickname: user.Nickname,
		Roles:    user.RoleIDS,
		Username: user.UserName,
	}

	resp = &types.UserResp{
		Code: "200",
		Data: data,
		Msg:  "success",
	}
	return
}
