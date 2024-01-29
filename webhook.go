package feishu

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gotime"
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

// WebhookSend 发送消息
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
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

// WebhookSendSign 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSign(ctx context.Context, secret string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params["timestamp"] = gotime.Current().Timestamp()
	params["sign"], _ = c.genSign(secret, fmt.Sprintf("%v", params["timestamp"]))
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

func (c *Client) genSign(secret string, timestamp string) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
