package intergration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 订单相关接口路径枚举
const (
	OrderBaseURL = "/v1/order"
	OrderByIDURL = "/v1/order/%s"
)

func TestOrderCRUD(t *testing.T) {
	orderName := fmt.Sprintf("测试订单_%d", time.Now().UnixNano())
	// 1. 新增订单
	addBody := map[string]interface{}{
		"shopId":   "shop001",
		"shopName": "测试商店",
		"address":  "测试地址",
		"products": []map[string]interface{}{
			{"productId": "prod001", "productName": "测试商品", "price": 10.5, "count": 2},
		},
		"state":       1,
		"payment":     1,
		"purchaserId": "user001",
		"driverId":    "driver001",
		"picture":     "",
	}
	addResp := doRequest(t, "POST", OrderBaseURL, "新增订单", addBody, 200)
	assert.Equal(t, "200", addResp["code"], "新增订单 code 应为0")

	// 2. 查询订单列表

	listResp := doRequest(t, "GET", OrderBaseURL, "查询订单列表", nil, 200)
	data, ok := listResp["items"].([]interface{})
	assert.True(t, ok, "[查询订单列表] data 不是数组: %v", listResp)
	var orderId string
	for _, item := range data {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		if m["shopName"] == "测试商店" && m["address"] == "测试地址" {
			if _, ok := m["orderId"]; ok {
				orderId = m["orderId"].(string)
			}
			assert.Equal(t, "测试商店", m["shopName"])
			assert.Equal(t, "测试地址", m["address"])
			break
		}
	}
	assert.NotEmpty(t, orderId, "[查询订单列表] 未找到新增的订单")

	// 3. 修改订单
	editBody := map[string]interface{}{
		"id":          orderId,
		"shopId":      "shop001",
		"shopName":    orderName + "_修改",
		"address":     "新地址",
		"products":    []map[string]interface{}{{"productId": "prod001", "productName": "测试商品", "price": 10.5, "count": 3}},
		"state":       2,
		"payment":     2,
		"purchaserId": "user001",
		"driverId":    "driver001",
		"picture":     "",
	}
	editURL := fmt.Sprintf(OrderByIDURL, orderId)
	_ = doRequest(t, "PUT", editURL, "修改订单", editBody, 200)

	// 4. 再查
	listResp2 := doRequest(t, "GET", OrderBaseURL, "再次查询订单列表", nil, 200)
	found := false
	for _, item := range listResp2["items"].([]interface{}) {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		if m["orderId"] == orderId {
			found = true
			assert.Equal(t, "新地址", m["address"])
			assert.Equal(t, float64(2), m["state"])
			break
		}
	}
	assert.True(t, found, "[再次查询订单列表] 订单修改未生效")

}
