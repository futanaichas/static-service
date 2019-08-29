package module

import "encoding/json"

// Response 返回的信息
type Response struct {
	Code int `json:"code"`
	Data struct {
		Hash string `json:"hash"`
		URL  string `json:"url"`
	} `json:"data"`
}

// FailRes ...
func FailRes() []byte {
	result, _ := json.Marshal(&Response{
		Code: -1,
	})
	return result
}

// OkRes ...
func OkRes(hash string, url string) []byte {
	result, _ := json.Marshal(&Response{
		Code: 0,
		Data: struct {
			Hash string `json:"hash"`
			URL  string `json:"url"`
		}{
			Hash: hash,
			URL:  url,
		},
	})
	return result
}
