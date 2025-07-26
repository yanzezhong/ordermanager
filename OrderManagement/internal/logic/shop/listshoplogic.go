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
		PhoneNumber: req.PhoneNumber,
		ShopName:    req.ShopName,
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
			Address:       shop.Address,
			CustomerLevel: shop.CustomerLevel,
			ID:            shop.ID,
			PhoneNumber:   shop.PhoneNumber,
			ShopName:      shop.ShopName,
			ShopNameMD5:   shop.ShopNameMD5,
			UpdateAt:      shop.UpdateAt.Unix(),
			CreateAt:      shop.CreateAt.Unix(),
		})
	}

	return result
}
