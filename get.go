package feishu

import "go.dtapp.net/golog"

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
