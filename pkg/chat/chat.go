package chat

import (
	"github.com/STX5/llm-translate/util"
)

type ChatCli interface {
	Chat(chatRequest *util.ChatRequest) (*util.ChatResponse, error)
	ChatStream(chatRequest *util.ChatRequest) (*util.StreamReader, error)
}
