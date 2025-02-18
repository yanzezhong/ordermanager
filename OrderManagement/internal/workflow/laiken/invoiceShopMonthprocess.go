package laiken

import (
	"OrderManagement/OrderManagement/internal/model"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
)

// sheet2 每个店的月度销量 Year Month Customer

type InvoiceShopMonth struct {
	Year          int
	Month         int
	Customer      string   `json:"customer"`       // 客户名称，记录与单据相关的客户
	CustomerLevel string   `json:"customer_level"` // 客户级别，例如普通客户、VIP客户等
	Handler       []string `json:"handler"`        // 经手人，记录处理单据的人员
	SalesRevenue  float64  `json:"sales_revenue"`  // 销售收入，记录销售的总收入
	SKUNum        int
	SKU           []string
	DisSKU        []string
	NewSKU        []string
}

// ConverInvoiceShopMonth 将 InvoiceDetail 列表转换为按年份和月份聚合的 InvoiceShopMonth 列表
func ConverInvoiceShopMonth(list []*model.InvoiceDetail) []*InvoiceShopMonth {
	// 创建一个 map 来存储每个客户在每个月的销售数据
	monthlyData := make(map[string]map[string]*InvoiceShopMonth)

	// 遍历 InvoiceDetail 列表，填充 monthlyData
	for _, detail := range list {
		date, err := time.Parse("2006/01/02", detail.DocumentDate)
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			continue
		}
		year, month := date.Year(), int(date.Month())
		key := fmt.Sprintf("%d-%02d", year, month)
		customerKey := fmt.Sprintf("%s-%s", detail.Customer, detail.CustomerLevel)

		if _, ok := monthlyData[key]; !ok {
			monthlyData[key] = make(map[string]*InvoiceShopMonth)
		}
		// if _, ok := monthlyData[key][customerKey]; !ok {
		// 	monthlyData[key][customerKey] = &InvoiceShopMonth{}
		// }
		if _, ok := monthlyData[key][customerKey]; !ok {

			dataMap := monthlyData[key]
			dataMap[customerKey] = &InvoiceShopMonth{
				Year:          year,
				Month:         month,
				Customer:      detail.Customer,
				CustomerLevel: detail.CustomerLevel,
				Handler:       []string{detail.Handler},
				SalesRevenue:  0,
				SKUNum:        0,
				SKU:           []string{},
				DisSKU:        []string{},
				NewSKU:        []string{},
			}
			monthlyData[key] = dataMap
		}

		data := monthlyData[key][customerKey]
		data.SalesRevenue += detail.SalesRevenue
		set := collection.NewSet()
		for _, v := range data.Handler {
			set.Add(v)
		}
		if !set.Contains(detail.Handler) {
			data.Handler = append(data.Handler, detail.Handler)
		}

	}

	// 将 map 转换为 InvoiceShopMonth 切片
	var result []*InvoiceShopMonth
	for _, customers := range monthlyData {
		for _, value := range customers {
			result = append(result, value)

		}
	}

	// 按年份和月份排序
	sort.Slice(result, func(i, j int) bool {
		if result[i].Year == result[j].Year {
			return result[i].Month < result[j].Month
		}
		return result[i].Year < result[j].Year
	})

	return genSKU(list, result)
}

// 不要变更排序
func genSKU(list []*model.InvoiceDetail, result []*InvoiceShopMonth) []*InvoiceShopMonth {
	// 创建一个 map 来存储每个经手人在每个月的活跃店铺
	SKUMap := make(map[string]map[int]*collection.Set)

	// 遍历 InvoiceDetail 列表，填充 activeShops
	for _, detail := range list {
		date, err := time.Parse("2006/01/02", detail.DocumentDate)
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			continue
		}
		year, month := date.Year(), int(date.Month())
		shopKey := fmt.Sprintf("%s-%02d", detail.Customer, month)

		if _, ok := SKUMap[shopKey]; !ok {
			SKUMap[shopKey] = make(map[int]*collection.Set)
		}
		if _, ok := SKUMap[shopKey][year]; !ok {
			skuSet := collection.NewSet()
			skuSet.Add(detail.ProductName)
			SKUMap[shopKey][year] = skuSet
		} else {
			set := SKUMap[shopKey][year]
			set.Add(detail.ProductName)
		}
	}

	// 遍历 InvoiceMonth 列表，填充 ActiveShopNum, ActiveShop, DisActiveShop, NewActiveShop
	for i := range result {
		shopKey := fmt.Sprintf("%s-%02d", result[i].Customer, result[i].Month)
		currentYear := result[i].Year
		previousYear := currentYear - 1

		if shops, ok := SKUMap[shopKey]; ok {

			set := shops[currentYear]
			result[i].SKU = set.KeysStr()
			result[i].NewSKU = set.KeysStr()
			result[i].SKUNum = len(result[i].SKU)
		}

		// 如果存在上一年的同月数据，计算 DisActiveShop 和 NewActiveShop
		if _, ok := SKUMap[shopKey][previousYear]; ok {
			for _, sku := range SKUMap[shopKey][previousYear].KeysStr() {
				if !SKUMap[shopKey][currentYear].Contains(sku) {
					result[i].DisSKU = append(result[i].DisSKU, sku)
				}
			}
			skuNew := []string{}

			for _, sku := range SKUMap[shopKey][currentYear].KeysStr() {
				if !SKUMap[shopKey][previousYear].Contains(sku) {
					skuNew = append(skuNew, sku)
				}
			}

			if len(skuNew) > 0 {
				result[i].NewSKU = skuNew
			}
		}
	}
	return result
}
