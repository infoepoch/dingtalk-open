package utils

import "encoding/json"

// 转换 字符串 到 JSON
func TransStrToJSON(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), v)
	return err
}
