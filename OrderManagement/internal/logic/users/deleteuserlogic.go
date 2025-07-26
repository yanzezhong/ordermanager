package users

import (
	"context"
	"strconv"
	"strings"

	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleUserReq) (resp *types.NormalResp, err error) {
	// todo: add your logic here and delete this line

	ids := strings.Split(req.IDS, ",")
	if len(ids) > 5 {
		return nil, errorx.MaxDeleteUserError
	}
	for _, v := range ids {
		// string to int64
		userId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}

		user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
		if err != nil {
			return nil, err
		}

		user.Status = model.UserStateDelete

		_, err = l.svcCtx.UserModel.Update(l.ctx, user)
		if err != nil {
			return nil, err
		}

	}

	return
}
