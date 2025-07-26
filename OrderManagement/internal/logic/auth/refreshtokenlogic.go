package auth

import (
	"context"
	"time"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.NormalResp, err error) {

	claims, err := getTokenClaims(req.Authorization, l.svcCtx.Config)
	_, ok := l.svcCtx.TokenCache.Get(claims.Username)
	if ok {
		return &types.NormalResp{
			Code: "401",
			Msg:  "user logout",
		}, nil
	}
	l.svcCtx.TokenCache.Set(req.RefreshToken, req.RefreshToken, 24*time.Hour)

	accessToken, refreshToken, err := utils.GenerateToken(l.svcCtx.Config, claims.UserID, claims.Username, claims.Password)
	if err != nil {
		return nil, err
	}
	resp = &types.NormalResp{
		Code: "200",
		Data: &types.AuthenticationToken{
			AccessToken:  accessToken,
			ExpiresIn:    l.svcCtx.Config.Auth.AccessExpire,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
		},
		Msg: "success",
	}
	return
}
