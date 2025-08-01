package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"OrderManagement/OrderManagement/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddShopLogic {
	return &AddShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddShopLogic) AddShop(req *types.PostShop) error {
	// todo: add your logic here and delete this line
	shop := &model.Shop{
		Address:       req.Address,
		CustomerLevel: req.CustomerLevel,
		PhoneNumber:   req.PhoneNumber,
		ShopName:      req.ShopName,
	}

	// shopName MD5
	shop.ShopNameMD5 = utils.GetMD5(shop.ShopName)
	shop.ID = shop.ShopNameMD5

	_, err := l.svcCtx.ShopModel.InsertMany(l.ctx, []*model.Shop{shop})

	if err != nil {
		return errorx.NewDBError(err.Error())
	}
	return nil
}
