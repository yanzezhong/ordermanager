package intergration

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8888" // 替换为你的实际端口

// 通用请求函数，带断言
func doRequest(t *testing.T, method, path string, caseName string, body interface{}, expectCode int) map[string]interface{} {
	t.Logf("[%s] %s %s start", caseName, method, path)
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			t.Fatalf("json marshal error: %v", err)
		}
	}
	req, err := http.NewRequest(method, baseURL+path, bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("new request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http do error: %v", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body error: %v", err)
	}

	var respMap map[string]interface{}

	// base64 解码
	// 先把 bodyBytes 解析成字符串（去掉引号）
	// 确定是否是 base64 编码

	var targetBytes []byte
	targetBytes = bodyBytes
	var base64Str string
	err = json.Unmarshal(bodyBytes, &base64Str)
	if err != nil {
		t.Errorf("decode base64 string error: %v, body: %s", err, string(bodyBytes))
	} else {
		// base64 解码
		targetBytes, err = base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			t.Errorf("base64 decode error: %v, body: %s", err, base64Str)
		}
	}

	if len(bodyBytes) == 0 {
		// 空响应体，返回空map或自定义内容
		respMap = map[string]interface{}{}
	} else if err := json.Unmarshal(targetBytes, &respMap); err != nil {
		t.Fatalf("decode response error: %v, body: %s", err, string(targetBytes))
	}
	if resp.StatusCode != expectCode {
		t.Fatalf("expect status %d, got %d, body: %v", expectCode, resp.StatusCode, respMap)
	}

	t.Logf("[%s] %s %s success", caseName, method, path)

	return respMap
}
