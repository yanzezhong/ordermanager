package workflow

import (
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncLaikenData struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncLaikenData(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLaikenData {
	return &SyncLaikenData{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}

}

// 将laiken数据同步至各个数据库
// 这个数据只是用来统计数据的,只是临时方案，当全部同步方案上线后，就要关闭
// todo 等待测试
func (m *SyncLaikenData) syncLaikenData() {
	cond := model.InvoiceDetailCond{}
	datas, _, err := m.svcCtx.InvoiceDetail.Search(m.ctx, &cond)
	if err != nil {
		m.Logger.Errorf("syncLaikenData search datas err: %+v", err)
		return
	}

	var shops []*model.Shop

	detailMap := map[string][]*model.InvoiceDetail{}

	for _, v := range datas {
		// 计算 ShopName 的 MD5 值
		shop := m.genShop(v)
		shops = append(shops, shop)

		if details, ok := detailMap[v.DocumentNumber]; ok {
			details = append(details, v)
			detailMap[v.DocumentNumber] = details

		} else {
			details := []*model.InvoiceDetail{}
			details = append(details, v)
			detailMap[v.DocumentNumber] = details
		}

	}

	_, err = m.svcCtx.ShopModel.InsertMany(m.ctx, shops)
	if err != nil {
		m.Logger.Errorf("syncLaikenData ShopModel InsertMany  err: %+v", err)
		return
	}

	// 生成订单记录

	shopMap := map[string]*model.Shop{}
	for _, v := range shops {
		shopMap[v.ID] = v
	}

	orders := []*model.Order{}
	productMap := map[string]*model.Product{}
	for documentNumber, details := range detailMap {
		if len(details) < 1 {
			continue
		}

		order := &model.Order{
			ID:          documentNumber,
			PurchaserId: details[0].Handler,
		}

		if shop, ok := shopMap[details[0].Customer]; ok {
			order.ShopId = shop.ID
			order.ShopName = shop.ShopName
		}

		result := []*model.Products{}
		for _, v := range details {
			hash := md5.Sum([]byte(v.ProductName))

			product := &model.Product{
				ID:   hex.EncodeToString(hash[:]),
				Name: v.ProductName,
			}

			if v.CustomerLevel == model.CustomerLevelTermianl.String() {
				product.Price = model.Price{
					WholeSale: v.UnitPrice,
				}
			} else if v.CustomerLevel == model.CustomerLevelWholeSale.String() {
				product.Price = model.Price{
					Terminal: v.UnitPrice,
				}
			}
			products := &model.Products{
				Product:      product,
				Count:        v.Amount,
				SalesRevenue: v.SalesQuantity,
			}
			result = append(result, products)

			if _, ok := productMap[product.ID]; !ok {
				productMap[product.ID] = product
			}
		}

		order.Products = result
		orders = append(orders, order)
	}

	_, err = m.svcCtx.OrderModel.InsertMany(m.ctx, orders)

	if err != nil {
		m.Logger.Errorf("syncLaikenData insert order datas err: %+v", err)
		return
	}

	productList := []*model.Product{}

	for _, v := range productMap {
		productList = append(productList, v)
	}

	_, err = m.svcCtx.ProductModel.InsertMany(m.ctx, productList)
	if err != nil {
		m.Logger.Errorf("syncLaikenData insert product datas err: %+v", err)
		return
	}

}

func (*SyncLaikenData) genShop(v *model.InvoiceDetail) *model.Shop {
	shop := &model.Shop{
		ShopName:      v.Customer,
		CustomerLevel: v.CustomerLevel,
		Address:       "",
		PhoneNumber:   "",
		UpdateAt:      v.UpdateAt,
		CreateAt:      v.CreateAt,
	}

	hash := md5.Sum([]byte(shop.ShopName))
	shop.ShopNameMD5 = hex.EncodeToString(hash[:])
	shop.ID = hex.EncodeToString(hash[:])
	return shop
}
