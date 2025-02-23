package product

import (
	"context"

	"OrderManagement/OrderManagement/internal/common/errorx"
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.PostProduct) error {
	// todo: add your logic here and delete this line
	//

	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, req.ProductID)
	if err != nil {
		return err
	}

	if product == nil {
		return errorx.NewDefaultError("product not found")
	}

	product = SyncProductWithReq(req, product)

	if _, err = l.svcCtx.ProductModel.Update(l.ctx, product); err != nil {
		return errorx.NewDBError(err.Error())
	}

	return nil
}

func SyncProductWithReq(req *types.PostProduct, product *model.Product) *model.Product {
	if req.BarCode != "" {
		product.BarCode = req.BarCode
	}
	if req.Image != "" {
		product.Image = req.Image
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.NickName != "" {
		product.NickName = req.NickName
	}
	if req.Specification != 0 {
		product.Specification = req.Specification
	}

	if req.Price != nil {
		product.Price = model.Price{
			Cost:      req.Price.Cost,
			SRP:       req.Price.SRP,
			Terminal:  req.Price.Terminal,
			Warning:   req.Price.Warning,
			WholeSale: req.Price.WholeSale,
		}
	}

	return product
}
