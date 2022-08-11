package feishu

import "go.dtapp.net/gorequest"

func (c *Client) request(url string, params map[string]interface{}) (gorequest.Response, error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddleware(request, Version)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddleware(request)
	}

	return request, err
}
