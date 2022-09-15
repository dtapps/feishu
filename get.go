package feishu

import "go.dtapp.net/golog"

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
