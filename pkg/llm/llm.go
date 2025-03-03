package llm

import (
	"github.com/STX5/llm-translate/util"
)

type LLMCli interface {
	Chat(chatRequest *util.ChatRequest) (*util.ChatResponse, error)
	ChatStream(chatRequest *util.ChatRequest) (*util.StreamReader, error)
}
