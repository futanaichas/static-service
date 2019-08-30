package module

import (
	"encoding/json"

	"github.com/futanaichas/static-service/src/config"
)

// FailRes ...
func FailRes() []byte {
	result, _ := json.Marshal(&Response{
		Message: "异常错误",
		Code:    -1,
	})
	return result
}

// OkRes ...
func OkRes(hash string, url string) []byte {
	result, _ := json.Marshal(&Response{
		Code:    0,
		Message: "储存成功",
		Data: struct {
			Hash string `json:"hash"`
			URL  string `json:"url"`
		}{
			Hash: hash,
			URL:  "http://" + config.Host() + url,
		},
	})
	return result
}
