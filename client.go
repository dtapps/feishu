package feishu

import (
	"go.dtapp.net/golog"
)

// Client 实例
type Client struct {
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient() (*Client, error) {
	c := &Client{}
	return c, nil
}
