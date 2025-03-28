package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuDetailLogic {
	return &GetMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuDetailLogic) GetMenuDetail(req *types.GetMenuDetailReq) (resp *types.GetMenuDetailResp, err error) {

	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// menu to map[string]interface{}

	newFunction(menu)
	resp = &types.GetMenuDetailResp{}
	return
}

func newFunction(menu *model.Menu) map[string]interface{} {
	menuMap := make(map[string]interface{})
	menuMap["id"] = menu.ID
	menuMap["name"] = menu.Name
	menuMap["parentId"] = menu.ParentId
	menuMap["routeName"] = menu.RouteName
	menuMap["routePath"] = menu.RoutePath
	menuMap["component"] = menu.Component
	menuMap["perm"] = menu.Perm
	menuMap["visible"] = menu.Visible
	menuMap["sort"] = menu.Sort
	menuMap["icon"] = menu.Icon
	menuMap["redirect"] = menu.Redirect
	menuMap["keepAlive"] = menu.KeepAlive
	menuMap["alwaysShow"] = menu.AlwaysShow
	menuMap["params"] = menu.Params
	var paramsList []map[string]interface{}
	for _, v := range menu.Params {
		paramsList = append(paramsList, map[string]interface{}{
			"key":   v.Key,
			"value": v.Value,
		})
	}
	menuMap["params"] = paramsList
}
