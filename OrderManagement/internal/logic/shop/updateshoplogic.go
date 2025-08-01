package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShopLogic {
	return &UpdateShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShopLogic) UpdateShop(req *types.PostShop) error {
	// 根据店铺ID获取现有店铺信息
	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return err
	}

	// 更新店铺字段
	shop.ShopName = req.ShopName
	shop.CustomerLevel = model.CustomerType(req.CustomerLevel)
	shop.Address = req.Address
	shop.PhoneNumber = req.PhoneNumber
	shop.CustomerID = req.CustomerID
	shop.CustomerSource = req.CustomerSource
	shop.Category = req.Category
	shop.SettlementMethod = req.SettlementMethod
	shop.Remarks = req.Remarks
	shop.MnemonicCode = req.MnemonicCode
	shop.CollectionPeriod = req.CollectionPeriod
	shop.CreditLimit = req.CreditLimit
	shop.ArrearsBalance = req.ArrearsBalance
	shop.PrepaymentBalance = req.PrepaymentBalance
	shop.Longitude = req.Longitude
	shop.Latitude = req.Latitude
	shop.AdCode = req.AdCode

	// 更新店铺信息
	_, err = l.svcCtx.ShopModel.Update(l.ctx, shop)
	return err
}
