package shop

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShopLogic {
	return &ListShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListShopLogic) ListShop(req *types.ListShopReq) (resp *types.ListShopResp, err error) {

	cond := &model.ShopCond{
		Address:      req.Address,
		CustomerType: req.CustomerLevel,
		PageParam: model.PageParam{
			Page: req.Page,
			Size: req.Size,
		},
		PhoneNumber:         req.PhoneNumber,
		ShopName:            req.ShopName,
		CustomerID:          req.CustomerID,
		CustomerSource:      req.CustomerSource,
		Category:            req.Category,
		SettlementMethod:    req.SettlementMethod,
		Remarks:             req.Remarks,
		MnemonicCode:        req.MnemonicCode,
		CollectionPeriod:    req.CollectionPeriod,
		CreditLimit:         req.CreditLimit,
		ArrearsBalance:      req.ArrearsBalance,
		PrepaymentBalance:   req.PrepaymentBalance,
		LastTransactionTime: req.LastTransactionTime,
		Longitude:           req.Longitude,
		Latitude:            req.Latitude,
		AdCode:              req.AdCode,
		LocationTime:        req.LocationTime,
	}

	shops, count, err := l.svcCtx.ShopModel.Search(l.ctx, cond)

	if err != nil {
		return nil, err
	}

	resp = &types.ListShopResp{
		Count: count,
		Items: convertShop(shops),
	}
	return
}

func convertShop(shops []*model.Shop) []*types.Shop {
	result := []*types.Shop{}
	for _, shop := range shops {
		result = append(result, &types.Shop{
			Address:             shop.Address,
			CustomerLevel:       string(shop.CustomerLevel),
			ID:                  shop.ID,
			PhoneNumber:         shop.PhoneNumber,
			ShopName:            shop.ShopName,
			ShopNameMD5:         shop.ShopNameMD5,
			CustomerID:          shop.CustomerID,
			CustomerSource:      shop.CustomerSource,
			Category:            shop.Category,
			SettlementMethod:    shop.SettlementMethod,
			Remarks:             shop.Remarks,
			MnemonicCode:        shop.MnemonicCode,
			CollectionPeriod:    shop.CollectionPeriod,
			CreditLimit:         shop.CreditLimit,
			ArrearsBalance:      shop.ArrearsBalance,
			PrepaymentBalance:   shop.PrepaymentBalance,
			LastTransactionTime: shop.LastTransactionTime.Unix(),
			UpdateAt:            shop.UpdateAt.Unix(),
			CreateAt:            shop.CreateAt.Unix(),
			Longitude:           shop.Longitude,
			Latitude:            shop.Latitude,
			AdCode:              shop.AdCode,
			LocationTime:        shop.LocationTime.Unix(),
		})
	}

	return result
}
