package feishu

import "go.dtapp.net/golog"

func (c *Client) Config(key string) *Client {
	c.config.key = key
	return c
}

// ConfigApiClientFun 日志配置
func (c *Client) ConfigApiClientFun(apiClientFun golog.ApiClientFun) {
	apiClient := apiClientFun()
	if apiClient != nil {
		c.log.client = apiClient
		c.log.status = true
	}
}
