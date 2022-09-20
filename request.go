package feishu

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}) (gorequest.Response, error) {

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.log.status {
		go c.log.client.Middleware(ctx, request, Version)
	}

	return request, err
}
