package logger

import (
	"encoding/json"
	"fmt"
)

type IJson interface {
	JSON(obj any) string
}

type DefaultJson struct{}

func (d *DefaultJson) JSON(obj any) string {
	if obj == nil {
		return "null"
	}

	// 快路径：最常见成功情况
	b, err := json.Marshal(obj)
	if err == nil {
		return string(b)
	}

	// 失败兜底：Go 原生打印
	return fmt.Sprintf("%+v", obj)
}
