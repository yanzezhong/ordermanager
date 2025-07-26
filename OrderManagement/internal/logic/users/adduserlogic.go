package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdduserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdduserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdduserLogic {
	return &AdduserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdduserLogic) Adduser(req *types.AddUserReq) (resp *types.NormalResp, err error) {

	// gengerate user
	user := model.User{
		Avatar:   req.Avatar,
		DeptID:   req.DeptID,
		Email:    req.Email,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		OpenID:   req.OpenID,
		RoleIDS:  req.RoleIDS,
		UserName: req.Username,
		Password: req.Password,
		Status:   model.UserStateNormal,
	}
	//

	err = l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return &types.NormalResp{
			Code: "500",
			Msg:  "insert user error",
			Data: nil,
		}, err
	}
	return &types.NormalResp{
		Code: "200",
		Msg:  "success",
		Data: nil,
	}, nil
}
