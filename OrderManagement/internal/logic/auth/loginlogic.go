package auth

import (
	"context"

	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	// search user by username
	cond := model.UserCond{
		Username: req.Username,
	}

	users, _, err := l.svcCtx.UserModel.Search(l.ctx, cond)
	if err != nil {
		return nil, err
	}
	//todo 整理代码
	if len(users) < 1 {
		return nil, errorx.OrderNotFoundError

	}

	if users[0].Password == req.Password {
		accessToken, refreshToken, err := utils.GenerateToken(l.svcCtx.Config, users[0].UserName, users[0].Password)
		if err != nil {
			return nil, err
		}
		//todo 整理代码
		resp = &types.LoginResp{
			Code: "200",
			Data: &types.AuthenticationToken{
				AccessToken:  accessToken,
				ExpiresIn:    l.svcCtx.Config.Auth.AccessExpire,
				RefreshToken: refreshToken,
				TokenType:    "Bearer",
			},
			Msg: "success",
		}
	} else {
		return nil, errorx.OrderNotFoundError
	}

	return
}
