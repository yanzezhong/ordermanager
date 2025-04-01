package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRoutesLogic {
	return &GetMenuRoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuRoutesLogic) GetMenuRoutes(req *types.GetMenuRoutesReq) (resp *types.GetMenuRoutesResp, err error) {

	menus, _, err := l.svcCtx.MenuModel.Search(l.ctx, model.MenuCond{})
	if err != nil {
		return
	}
	// convert model.menu to RouteVO

	data := []types.RouteVO{}
	for _, menu := range menus {
		param := map[string]string{}
		for _, v := range menu.Params {
			param[v.Key] = v.Value
		}

		menuVO := types.RouteVO{
			Name:      menu.Name,
			Redirect:  menu.Redirect,
			Path:      menu.RoutePath,
			Component: menu.Component,
			Meta: types.Meta{
				Title:      menu.RouteName,
				Icon:       menu.Icon,
				Hidden:     menu.Visible == 1,
				KeepAlive:  menu.KeepAlive == 1,
				AlwaysShow: menu.AlwaysShow == 1,
				Params:     param,
			},
		}
		data = append(data, menuVO)

	}
	resp = &types.GetMenuRoutesResp{
		Code: "200",
		Data: data,
		Msg:  "success",
	}

	return
}
