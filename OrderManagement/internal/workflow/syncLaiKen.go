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
func (m *SyncLaikenData) syncLaikenData() {
	cond := model.InvoiceDetailCond{}
	datas, _, err := m.svcCtx.InvoiceDetail.Search(m.ctx, &cond)
	if err != nil {
		m.Logger.Errorf("syncLaikenData search datas err: %+v", err)
		return
	}

	var shops []*model.Shop

	for _, v := range datas {
		shop := &model.Shop{
			ID:           v.ID,
			ShopName:     v.Customer,
			CustomerType: v.CustomerLevel,
			Address:      "",
			PhoneNumber:  "",
			UpdateAt:     v.UpdateAt,
			CreateAt:     v.CreateAt,
		}
		// 计算 ShopName 的 MD5 值
		hash := md5.Sum([]byte(shop.ShopName))
		shop.ShopNameMD5 = hex.EncodeToString(hash[:])
		shops = append(shops, shop)
	}

	_, err = m.svcCtx.ShopModel.InsertMany(m.ctx, shops)
	if err != nil {
		m.Logger.Errorf("syncLaikenData ShopModel InsertMany  err: %+v", err)
		return
	}

}
