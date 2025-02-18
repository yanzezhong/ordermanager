package laiken

import (
	"OrderManagement/OrderManagement/internal/model"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

// sheet 1 每个人按月的销量数据

type InvoiceMonth struct {
	Year          int
	Month         int
	Handler       string   `json:"handler"`       // 经手人，记录处理单据的人员
	SalesRevenue  float64  `json:"sales_revenue"` // 销售收入，记录销售的总收入
	ActiveShopNum int      // 活跃店数
	ActiveShop    []string // 活跃店名
	DisActiveShop []string // 失去活跃商店
	NewActiveShop []string // 新增活跃商店
}

// ConvertInvoiceMonth 将 InvoiceDetail 列表转换为按月统计的 InvoiceMonth 列表
func ConvertInvoiceMonth(list []*model.InvoiceDetail) []InvoiceMonth {
	// 使用 map 来按年份和月份聚合数据
	monthlyData := make(map[string]map[string]float64)

	for _, detail := range list {
		// 解析 DocumentDate
		date, err := time.Parse("2006/01/02", detail.DocumentDate)
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			continue
		}

		// 获取年份和月份
		year, month := date.Year(), int(date.Month())

		// 创建或获取年份的 map
		if _, ok := monthlyData[detail.Handler]; !ok {
			monthlyData[detail.Handler] = make(map[string]float64)
		}

		// 创建或获取月份的聚合数据
		key := fmt.Sprintf("%d-%02d", year, month)
		if _, ok := monthlyData[detail.Handler][key]; !ok {
			monthlyData[detail.Handler][key] = 0
		}

		// 累加 SalesRevenue
		monthlyData[detail.Handler][key] += detail.SalesRevenue
	}

	// 将 map 转换为 InvoiceMonth 列表
	var result []InvoiceMonth
	for handler, months := range monthlyData {
		for key, revenue := range months {
			year, month := splitYearMonth(key)
			result = append(result, InvoiceMonth{
				Year:          year,
				Month:         month,
				Handler:       handler,
				SalesRevenue:  revenue,
				ActiveShopNum: 0,          // 这里需要根据实际需求填充活跃店数
				ActiveShop:    []string{}, // 这里需要根据实际需求填充活跃店名
				DisActiveShop: []string{}, // 这里需要根据实际需求填充失去活跃商店
				NewActiveShop: []string{}, // 这里需要根据实际需求填充新增活跃商店
			})
		}
	}

	// 按年份和月份排序
	sort.Slice(result, func(i, j int) bool {
		if result[i].Year == result[j].Year {
			return result[i].Month < result[j].Month
		}
		return result[i].Year < result[j].Year
	})

	return generateActive(list, result)
}

func generateActive(list []*model.InvoiceDetail, result []InvoiceMonth) []InvoiceMonth {
	// 创建一个 map 来存储每个经手人在每个月的活跃店铺
	activeShops := make(map[string]map[int]map[string]bool)

	// 遍历 InvoiceDetail 列表，填充 activeShops
	for _, detail := range list {
		date, err := time.Parse("2006/01/02", detail.DocumentDate)
		if err != nil {
			log.Printf("Failed to parse date: %v", err)
			continue
		}
		year, month := date.Year(), int(date.Month())
		handlerKey := fmt.Sprintf("%s-%02d", detail.Handler, month)

		if _, ok := activeShops[handlerKey]; !ok {
			activeShops[handlerKey] = make(map[int]map[string]bool)
		}
		if _, ok := activeShops[handlerKey][year]; !ok {
			activeShops[handlerKey][year] = make(map[string]bool)
		}
		activeShops[handlerKey][year][detail.Customer] = true
	}

	// 遍历 InvoiceMonth 列表，填充 ActiveShopNum, ActiveShop, DisActiveShop, NewActiveShop
	for i := range result {
		handlerKey := fmt.Sprintf("%s-%02d", result[i].Handler, result[i].Month)
		currentYear := result[i].Year
		previousYear := currentYear - 1

		if shops, ok := activeShops[handlerKey]; ok {
			result[i].ActiveShop = make([]string, 0, len(shops[currentYear]))
			for shop := range shops[currentYear] {
				result[i].ActiveShop = append(result[i].ActiveShop, shop)
			}
			result[i].ActiveShopNum = len(result[i].ActiveShop)
		}

		// 如果存在上一年的同月数据，计算 DisActiveShop 和 NewActiveShop
		if _, ok := activeShops[handlerKey][previousYear]; ok {
			for shop := range activeShops[handlerKey][previousYear] {
				if !activeShops[handlerKey][currentYear][shop] {
					result[i].DisActiveShop = append(result[i].DisActiveShop, shop)
				}
			}

			for shop := range activeShops[handlerKey][currentYear] {
				if !activeShops[handlerKey][previousYear][shop] {
					result[i].NewActiveShop = append(result[i].NewActiveShop, shop)
				}
			}
		}
	}

	return result
}

// splitYearMonth 将 "年-月" 格式的字符串拆分为年份和月份
func splitYearMonth(key string) (int, int) {
	parts := strings.Split(key, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	return year, month
}
