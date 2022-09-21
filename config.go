package feishu

func (c *Client) Config(key string) *Client {
	c.config.key = key
	return c
}
