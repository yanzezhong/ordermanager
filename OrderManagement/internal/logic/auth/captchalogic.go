package auth

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req *types.CaptchaReq) (resp *types.CaptchaResp, err error) {
	captcha := utils.Captcha()

	resp = &types.CaptchaResp{
		Code: "200",
		Data: &types.CAPTCHAInfo{
			CAPTCHABase64: &captcha.Encode,
			CAPTCHAKey:    &captcha.Id,
		},
		Msg: "success",
	}
	return
}
