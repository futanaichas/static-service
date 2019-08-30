package module

// Response 返回的信息
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Hash string `json:"hash"`
		URL  string `json:"url"`
	} `json:"data"`
}
