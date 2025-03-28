package menu

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.MenuListReq) (resp *types.MenuListResp, err error) {

	resp = &types.MenuListResp{}

	// 构建查询条件

	// Status 1->正常；0->禁用

	cond := model.MenuCond{
		Keywords: req.Keywords,
		Status:   convertStatus(req.Status),
	}

	// 执行搜索
	menuList, _, err := l.svcCtx.MenuModel.Search(l.ctx, cond)
	if err != nil {
		return nil, err
	}

	// 转换数据
	for _, menu := range menuList {

		menuVO := types.MenuVO{
			Id:        menu.ID,
			ParentId:  menu.ParentId,
			Name:      menu.Name,
			RouteName: menu.RouteName,
			RoutePath: menu.RoutePath,
			Component: menu.Component,
			Perm:      menu.Perm,
			Visible:   menu.Visible,
			Sort:      menu.Sort,
			Icon:      menu.Icon,
			Redirect:  menu.Redirect,
		}
		resp.Data = append(resp.Data)
		// 递归获取子菜单
		children, err := l.getChildrenMenus(menu.ID)
		if err != nil {
			return nil, err
		}
		menuVO.Children = children
		resp.Data = append(resp.Data, menuVO)

	}

	resp.Code = "200"
	resp.Msg = "success"

	return
}

func convertStatus(status string) int32 {
	if status == "正常" {
		return 1
	} else if status == "禁用" {
		return 0
	}
	return 0
}

// getChildrenMenus 递归获取子菜单
func (l *GetMenuListLogic) getChildrenMenus(parentId int64) ([]types.MenuVO, error) {
	cond := model.MenuCond{
		ParentId: parentId,
	}

	childMenus, _, err := l.svcCtx.MenuModel.Search(l.ctx, cond)
	if err != nil {
		return nil, err
	}

	var children []types.MenuVO
	for _, child := range childMenus {
		childVO := types.MenuVO{
			Id:        child.ID,
			ParentId:  child.ParentId,
			Name:      child.Name,
			RouteName: child.RouteName,
			RoutePath: child.RoutePath,
			Component: child.Component,
			Perm:      child.Perm,
			Visible:   child.Visible,
			Sort:      child.Sort,
			Icon:      child.Icon,
			Redirect:  child.Redirect,
		}

		// 递归获取子菜单的子菜单
		subChildren, err := l.getChildrenMenus(child.ID)
		if err != nil {
			return nil, err
		}
		childVO.Children = subChildren
		children = append(children, childVO)
	}

	return children, nil
}
