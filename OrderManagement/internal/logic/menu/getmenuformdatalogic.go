package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuFormDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuFormDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuFormDataLogic {
	return &GetMenuFormDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuFormDataLogic) GetMenuFormData(req *types.GetMenuFormDataReq) (resp *types.GetMenuFormDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
