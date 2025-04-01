package menu

import (
	"context"
	"time"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	// gen model

	// 生成插入逻辑
	menu := model.Menu{
		AlwaysShow: req.AlwaysShow,
		Component:  req.Component,
		Icon:       req.Icon,
		ID:         req.Id,
		ParentId:   req.ParentId,
		Name:       req.Name,
		RouteName:  req.RouteName,
		RoutePath:  req.RoutePath,
		Perm:       req.Perm,
		Visible:    req.Visible,
		Sort:       req.Sort,
		Redirect:   req.Redirect,
		KeepAlive:  req.KeepAlive,
		UpdateAt:   time.Now(),
		CreateAt:   time.Now(),
	}

	for _, v := range req.Params {
		param := model.KeyValue{
			Key:   v.Key,
			Value: v.Value,
		}
		menu.Params = append(menu.Params, param)
	}

	err = l.svcCtx.MenuModel.Insert(l.ctx, &menu)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	resp = &types.CommonResponse{
		Code: "200",
		Msg:  "success",
	}
	return
}
