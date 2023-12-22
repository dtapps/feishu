package feishu

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type WebhookSendResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type WebhookSendResult struct {
	Result WebhookSendResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest.Response) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http}
}

// WebhookSend https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSend(ctx context.Context, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/open-apis/bot/v2/hook/%s", c.GetKey()), params)
	if err != nil {
		return newWebhookSendResult(WebhookSendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WebhookSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}
