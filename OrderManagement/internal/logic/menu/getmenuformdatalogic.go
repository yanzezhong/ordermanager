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

	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	data := types.MenuForm{
		Id:         menu.ID,
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		RouteName:  menu.RouteName,
		RoutePath:  menu.RoutePath,
		Component:  menu.Component,
		Perm:       menu.Perm,
		Visible:    menu.Visible,
		Sort:       menu.Sort,
		Icon:       menu.Icon,
		Redirect:   menu.Redirect,
		KeepAlive:  menu.KeepAlive,
		AlwaysShow: menu.AlwaysShow,
	}
	// gen param
	for _, v := range menu.Params {
		data.Params = append(data.Params, types.KeyValue{
			Key:   v.Key,
			Value: v.Value,
		})
	}
	// gen menu form
	resp = &types.GetMenuFormDataResp{
		Data: data,
		Code: "200",
		Msg:  "success",
	}

	return
}
