package config

import (
	"encoding/json"
	"io/ioutil"
)

var _host string

// Host 使用函数返回防止被修改
func Host() string {
	return _host
}

// 将配置内容设置为全局变量
func setConfig(etc *Config) {
	_host = etc.Host
}

func init() {
	var result = &Config{}
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(configFile, result)
	if err != nil {
		panic(err.Error())
	}

	setConfig(result)
}
