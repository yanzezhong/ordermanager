package intergration

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 部门相关接口路径枚举
const (
	DeptBaseURL = "/v1/dept"
	DeptByIDURL = "/v1/dept/%d"
)

func TestDeptCRUD(t *testing.T) {
	deptName := fmt.Sprintf("测试部门_%d", time.Now().UnixNano())
	addBody := map[string]interface{}{
		"name":     deptName,
		"parentId": "",
		"status":   1,
		"remark":   "自动化测试",
	}
	addResp := doRequest(t, http.MethodPost, DeptBaseURL, "新增部门", addBody, 200)
	assert.Equal(t, "200", addResp["code"], "新增部门 code 应为0")
	assert.Equal(t, "success", addResp["message"], "新增部门 message 应为 success") // 如有 message 字段

	// 查询部门列表
	listResp := doRequest(t, http.MethodGet, DeptBaseURL, "查询部门列表", nil, 200)
	data, ok := listResp["data"].([]interface{})
	assert.True(t, ok, "[查询部门列表] data 不是数组: %v", listResp)
	var deptId float64
	for _, item := range data {
		m := item.(map[string]interface{})
		if m["name"] == deptName {
			deptId = m["id"].(float64)
			assert.Equal(t, deptName, m["name"])
			assert.Equal(t, float64(1), m["status"])
			break
		}
	}
	assert.NotEmpty(t, deptId, "[查询部门列表] 未找到新增的部门: %s", deptName)

	// 修改部门
	editBody := map[string]interface{}{
		"id":       deptId,
		"name":     deptName + "_修改",
		"status":   2,
		"remark":   "已修改",
		"code":     "123456",
		"sort":     1,
		"parentId": 0,
	}
	editURL := fmt.Sprintf(DeptByIDURL, int64(deptId))
	editResp := doRequest(t, http.MethodPut, editURL, "修改部门", editBody, 200)
	assert.Equal(t, float64(0), editResp["code"])

	// 再查
	listResp2 := doRequest(t, http.MethodGet, DeptBaseURL, "再次查询部门列表", nil, 200)
	found := false
	for _, item := range listResp2["data"].([]interface{}) {
		m := item.(map[string]interface{})
		if m["id"] == deptId {
			found = true
			assert.Equal(t, deptName+"_修改", m["name"])
			assert.Equal(t, float64(2), m["status"])
			break
		}
	}
	assert.True(t, found, "[再次查询部门列表] 部门修改未生效")

	// 删除
	delURL := fmt.Sprintf(DeptByIDURL, int64(deptId))
	delResp := doRequest(t, http.MethodDelete, delURL, "删除部门", nil, 200)
	assert.Equal(t, float64(0), delResp["code"])

	// 再查，断言已删除
	listResp3 := doRequest(t, http.MethodGet, DeptBaseURL, "删除后查询部门列表", nil, 200)
	for _, item := range listResp3["data"].([]interface{}) {
		m := item.(map[string]interface{})
		assert.NotEqual(t, deptId, m["id"], "[删除后查询部门列表] 部门未被删除")
	}
}
