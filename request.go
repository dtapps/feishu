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

	// 日志
	if c.log.gorm {
		go c.log.logGormClient.GormMiddleware(ctx, request, Version)
	}
	if c.log.mongo {
		go c.log.logMongoClient.MongoMiddleware(ctx, request, Version)
	}

	return request, err
}
