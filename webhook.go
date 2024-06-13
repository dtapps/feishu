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
	"go.opentelemetry.io/otel/codes"
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
func (c *Client) WebhookSend(ctx context.Context, key string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "open-apis/bot/v2/hook")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	request, err := c.request(ctx, fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return newWebhookSendResult(WebhookSendResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response WebhookSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

// WebhookSendSign 发送消息签名版
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
func (c *Client) WebhookSendSign(ctx context.Context, key string, secret string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "open-apis/bot/v2/hook")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params["timestamp"] = gotime.Current().Timestamp()
	params["sign"], _ = c.webhookSendSignGenSign(secret, fmt.Sprintf("%v", params["timestamp"]))

	// 请求
	request, err := c.request(ctx, fmt.Sprintf("open-apis/bot/v2/hook/%s", key), params)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return newWebhookSendResult(WebhookSendResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response WebhookSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

func (c *Client) webhookSendSignGenSign(secret string, timestamp string) (string, error) {
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
