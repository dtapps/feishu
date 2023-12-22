package feishu

import (
	"go.dtapp.net/golog"
)

func (c *Client) Config(key string) *Client {
	c.config.key = key
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}
