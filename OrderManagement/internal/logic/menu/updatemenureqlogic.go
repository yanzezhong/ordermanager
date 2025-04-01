package menu

import (
	"context"
	"errors"
	"time"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuReqLogic {
	return &UpdateMenuReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuReqLogic) UpdateMenuReq(req *types.UpdateMenuReq) (resp *types.CommonResponse, err error) {

	if req.Id == 0 {
		return nil, errors.ErrUnsupported
	}

	menu := &model.Menu{
		ID:         req.Id,
		ParentId:   req.ParentId,
		Name:       req.Name,
		RouteName:  req.RouteName,
		RoutePath:  req.RoutePath,
		Component:  req.Component,
		Perm:       req.Perm,
		Visible:    req.Visible,
		Sort:       req.Sort,
		Icon:       req.Icon,
		Redirect:   req.Redirect,
		KeepAlive:  req.KeepAlive,
		AlwaysShow: req.AlwaysShow,
		UpdateAt:   time.Now(), // Set the update time to the current time
		// CreateAt:      time.Now(), // Uncomment if you want to update the creation time as well
	}

	for _, v := range req.Params {
		menu.Params = append(menu.Params, model.KeyValue{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	_, err = l.svcCtx.MenuModel.Update(l.ctx, menu)
	if err != nil {
		l.Logger.Errorf("failed to update menu: %v", err)
		return nil, err
	}

	// Return a success response
	return &types.CommonResponse{
		Code: "200",
		Msg:  "menu updated successfully",
	}, nil
}
