package qwen

import (
	"fmt"

	"github.com/STX5/llm-translate/pkg/chat"
	"github.com/STX5/llm-translate/util"

	"github.com/swxctx/xlog"
)

/*
	接口文档：https://help.aliyun.com/zh/dashscope/developer-reference/api-details?spm=a2c4g.11186623.0.0.e66d23edk4jpy6#b8ebf6b25eul6
*/

// QwenCli API请求客户端
type QwenCli struct {
	// 基础请求api
	baseUri string

	// API Key
	apiKey string

	// 是否调试模式[调试模式可以输出详细的信息]
	debug bool

	// 最大空消息数量
	maxEmptyMessageCount int
}

// NewClient 初始化请求客户端
func NewClient(apiKey string, debug ...bool) chat.ChatCli {
	client := &QwenCli{
		apiKey:               apiKey,
		baseUri:              "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
		maxEmptyMessageCount: 900,
	}
	if len(debug) > 0 {
		client.debug = debug[0]
	}

	if client.debug {
		xlog.SetLevel("debug")
	}

	return client
}

func (c *QwenCli) String() string {
	return fmt.Sprintf("apiKey: %s, baseUri: %s, debug: %t, maxEmptyMessageCount: %d}", c.apiKey, c.baseUri, c.debug, c.maxEmptyMessageCount)
}

// SetMaxEmptyMessageCount 最大空消息数量
func (c *QwenCli) SetMaxEmptyMessageCount(count int) {
	c.maxEmptyMessageCount = count
}

// SetDebug debug开关
func (c *QwenCli) SetDebug(debug bool) {
	c.debug = debug
}

// Chat 对话接口
func (c *QwenCli) Chat(chatRequest *util.ChatRequest) (*util.ChatResponse, error) {
	return c.chat(chatRequest)
}

// ChatStream 流式对话接口
func (c *QwenCli) ChatStream(chatRequest *util.ChatRequest) (*util.StreamReader, error) {
	return c.chatStream(chatRequest)
}
