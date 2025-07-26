package users

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line

	// 生成 model search cond
	cond := model.UserCond{}

	users, count, err := l.svcCtx.UserModel.Search(l.ctx, cond)
	if err != nil {
		return nil, err
	}

	userPages := make([]types.UserPageVO, 0)
	for _, v := range users {
		u := types.UserPageVO{
			Avatar:     v.Avatar,
			CreateTime: v.CreateAt.Unix(),
			Email:      v.Email,
			Gender:     v.Gender,
			ID:         v.ID,
			Mobile:     v.Mobile,
			Nickname:   v.Nickname,
			RoleNames:  v.RoleIDS,
		}
		userPages = append(userPages, u)
	}

	return &types.UserListResp{
		Code: "200",
		Data: &types.DataUserPageVO{
			List:  userPages,
			Total: count,
		},
		Msg: "success",
	}, nil
}
