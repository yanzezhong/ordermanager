package product

import (
	"context"

	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductLogic {
	return &ListProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// IsActive  bool   `json:"isActive,optional"` // 是否生效
// 	NickName  string `json:"nickName,optional"` // 别名
// 	BarCode   string `json:"barCode,optional"`  //条码

func (l *ListProductLogic) ListProduct(req *types.ListProductReq) (resp *types.ListProductResp, err error) {
	cond := model.ProductCond{
		Name:     req.Name,
		Id:       req.ProductId,
		IsActive: req.IsActive,
		BarCode:  req.BarCode,
		NickName: req.NickName,
	}
	products, count, err := l.svcCtx.ProductModel.Search(l.ctx, cond)
	if err != nil {
		return nil, err
	}

	items := []*types.Product{}
	for _, v := range products {
		item := ConvertModelToProduct(v)
		items = append(items, item)
	}

	resp = &types.ListProductResp{
		Items: items,
		Count: count,
	}

	return
}

// ConvertModelToProduct 函数用于将 ModelProduct 转换为 Product
func ConvertModelToProduct(modelProduct *model.Product) *types.Product {
	return &types.Product{
		ProductId:     modelProduct.ID,
		Name:          modelProduct.Name,
		Specification: modelProduct.Specification,
		IsActive:      modelProduct.IsActive,
		Image:         modelProduct.Image,
		NickName:      modelProduct.NickName,
		BarCode:       modelProduct.BarCode,
		Price:         ConvertModelPriceToPrice(modelProduct.Price),
	}
}

func ConvertModelPriceToPrice(modelPrice model.Price) types.Price {
	return types.Price{
		Terminal:  modelPrice.Terminal,
		WholeSale: modelPrice.WholeSale,
		Cost:      modelPrice.Cost,
		SRP:       modelPrice.SRP,
		Warning:   modelPrice.Warning,
	}
}
