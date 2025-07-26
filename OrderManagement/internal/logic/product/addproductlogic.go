package product

import (
	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"context"

	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductLogic) AddProduct(req *types.PostProduct) error {

	product := &model.Product{
		BarCode:       req.BarCode,
		Image:         req.Image,
		Name:          req.Name,
		NickName:      req.NickName,
		Specification: req.Specification,
	}

	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	product.Price = model.Price{
		Cost:      req.Price.Cost,
		SRP:       req.Price.SRP,
		Terminal:  req.Price.Terminal,
		Warning:   req.Price.Warning,
		WholeSale: req.Price.WholeSale,
	}

	if err := l.svcCtx.ProductModel.Insert(l.ctx, product); err != nil {
		return errorx.NewDBError(err.Error())
	}
	return nil
}
