package controller

import (
	"strings"
)

var errMap = map[string]string{
	"duplicate key value violates":      "主键字段值已存在!",
	"Table not set, please set it like": "数据表设置错误，请联系管理员!",
	//""
}

//错误处理
func parseError(err error) string {
	if err == nil {
		return "未知错误"
	}

	errStr := err.Error()
	for k, v := range errMap {
		if strings.Contains(errStr, k) {
			return v
		}
	}
	return errStr
}
