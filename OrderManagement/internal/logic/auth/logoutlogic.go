package auth

import (
	"context"
	"strings"

	"OrderManagement/OrderManagement/internal/config"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.NormalResp, err error) {
	//todo 替换为redis
	claims, err := getTokenClaims(req.Authorization, l.svcCtx.Config)

	l.svcCtx.TokenCache.Delete(claims.Username)

	resp = &types.NormalResp{
		Code: "200",
		Msg:  "success",
	}

	return
}

func getTokenClaims(authorization string, c config.Config) (*utils.JwtClaims, error) {
	token := strings.Replace(authorization, "Bearer ", "", 1)

	claims, err := utils.ParseToken(c, token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
