package qwen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/STX5/llm-translate/util"

	"github.com/swxctx/ghttp"
	"github.com/swxctx/xlog"
)

// Chat 对话方法
func (c *QwenCli) chat(chatRequest *util.ChatRequest) (*util.ChatResponse, error) {
	// new request
	req := ghttp.Request{
		Url:       c.baseUri,
		Method:    "POST",
		ShowDebug: c.debug,
		Body:      chatRequest,
	}
	req.AddHeader("Authorization", "Bearer "+c.apiKey)
	req.AddHeader("Content-Type", "application/json")

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("qwen: Chat err, err is-> %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("qwen: Chat http response code not 200, code is -> %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// read body
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("qwen: Chat read resp body err-> %v", err)
	}
	if c.debug {
		xlog.Debugf("qwen: chat resp-> %s", string(respBs))
	}

	// unmarshal data
	var (
		chatResp *util.ChatResponse
	)

	err = json.Unmarshal(respBs, &chatResp)
	if err != nil {
		return nil, fmt.Errorf("qwen: Chat data unmarshal err-> %v", err)
	}
	return chatResp, nil
}

// ChatStream 流式对话方法
func (c *QwenCli) chatStream(chatRequest *util.ChatRequest) (*util.StreamReader, error) {
	// new request
	req := ghttp.Request{
		Url:       c.baseUri,
		Method:    "POST",
		ShowDebug: c.debug,
		Body:      chatRequest,
	}
	req.AddHeader("Authorization", "Bearer "+c.apiKey)
	req.AddHeader("Content-Type", "application/json")
	req.AddHeader("Accept", "text/event-stream")

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("qwen: Chat err, err is-> %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("qwen: Chat http response code not 200, code is -> %d", resp.StatusCode)
	}

	// 交给外部调用逻辑处理
	return util.NewStreamReader(resp, c.maxEmptyMessageCount), nil
}
