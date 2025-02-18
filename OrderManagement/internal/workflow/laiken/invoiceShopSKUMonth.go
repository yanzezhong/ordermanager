package laiken

import (
	"OrderManagement/OrderManagement/internal/model"
	"fmt"
	"log"
	"sort"
	"time"
)

// sheet3 单店分品销量 主键： Year Month Customer SKU

type InvoiceShopSKUMonth struct {
	Year          int
	Month         int
	Customer      string `json:"customer"`       // 客户名称，记录与单据相关的客户
	CustomerLevel string `json:"customer_level"` // 客户级别，例如普通客户、VIP客户等
	Handler       string `json:"handler"`        // 经手人，记录处理单据的人员
	SKU           string
	SalesRevenue  float64 `json:"sales_revenue"` // 销售收入，记录销售的总收入
}

func ConverInvoiceShopSKUMonth(list []*model.InvoiceDetail) []*InvoiceShopSKUMonth {
	// 创建一个 map 来存储每个客户在每个月的每个 SKU 的销售数据
	dataMap := make(map[string]*InvoiceShopSKUMonth)

	// 遍历 InvoiceDetail 列表，填充 dataMap
	for _, detail := range list {
		date, err := time.Parse("2006/01/02", detail.DocumentDate)
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			continue
		}
		year, month := date.Year(), int(date.Month())

		// 生成长键
		key := fmt.Sprintf("%d-%02d-%s-%s-%s", year, month, detail.Customer, detail.CustomerLevel, detail.ProductName)

		if _, ok := dataMap[key]; !ok {
			dataMap[key] = &InvoiceShopSKUMonth{
				Year:          year,
				Month:         month,
				Customer:      detail.Customer,
				CustomerLevel: detail.CustomerLevel,
				Handler:       detail.Handler,
				SKU:           detail.ProductName,
				SalesRevenue:  0,
			}
		}

		dataMap[key].SalesRevenue += detail.SalesRevenue
	}

	// 将 map 转换为 InvoiceShopSKUMonth 切片
	var result []*InvoiceShopSKUMonth
	for _, value := range dataMap {
		result = append(result, value)
	}

	// 按年份和月份排序
	sort.Slice(result, func(i, j int) bool {
		if result[i].Year == result[j].Year {
			if result[i].Month == result[j].Month {
				return result[i].SKU < result[j].SKU
			}
			return result[i].Month < result[j].Month
		}
		return result[i].Year < result[j].Year
	})

	return result
}
