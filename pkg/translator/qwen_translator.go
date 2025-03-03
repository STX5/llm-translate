package translator

import (
	"errors"

	"github.com/STX5/llm-translate/pkg/chat"
	"github.com/STX5/llm-translate/util"
)

// QwenTranslator 通义千问翻译器实现
type QwenTranslator struct {
	ChatCli chat.ChatCli
}

// NewQwenTranslator 创建新的通义千问翻译器
func NewQwenTranslator(chatCli chat.ChatCli) *QwenTranslator {
	return &QwenTranslator{
		ChatCli: chatCli,
	}
}

// TranslateSection 翻译文本片段
func (t *QwenTranslator) TranslateSection(text []byte) ([]byte, error) {
	if len(text) == 0 {
		return nil, errors.New("empty text to translate")
	}

	// 创建翻译请求
	chatRequest := &util.ChatRequest{
		Model: util.MODEL_QWEN_MT_PLUS,
		Input: util.Input{
			Messages: []util.MessageInfo{
				{
					Role:    util.CHAT_MESSAGE_ROLE_USER,
					Content: string(text),
				},
			},
		},
		Parameters: util.Parameters{
			ResultFormat: util.RESULT_FORMAT_TEXT,
			TranslationOptions: util.TranslationOptions{
				SourceLang: "English",
				TargetLang: "Chinese",
			},
		},
	}

	// 发送翻译请求
	response, err := t.ChatCli.Chat(chatRequest)
	if err != nil {
		return nil, err
	}

	// 从choices[0].message.content获取翻译结果
	if len(response.Output.Choices) > 0 && response.Output.Choices[0].Message.Content != "" {
		return []byte(response.Output.Choices[0].Message.Content), nil
	}

	// 如果choices为空，尝试从Text字段获取
	if response.Output.Text != "" {
		return []byte(response.Output.Text), nil
	}

	return nil, errors.New("no translation result found in API response")
}

// TranslateDocument 翻译整个文档
func (t *QwenTranslator) TranslateDocument(content interface{}) ([]byte, error) {
	contentBytes, ok := content.([]byte)
	if !ok {
		return nil, errors.New("content is not a []byte")
	}

	return t.TranslateSection(contentBytes)
}

// TranslateSectionStream 使用流式API翻译文本片段
func (t *QwenTranslator) TranslateSectionStream(text []byte) (*util.StreamReader, error) {
	if len(text) == 0 {
		return nil, errors.New("empty text to translate")
	}

	// 创建翻译请求
	chatRequest := &util.ChatRequest{
		Model: util.MODEL_QWEN_MT_PLUS,
		Input: util.Input{
			Messages: []util.MessageInfo{
				{
					Role:    util.CHAT_MESSAGE_ROLE_USER,
					Content: string(text),
				},
			},
		},
		Parameters: util.Parameters{
			ResultFormat:      util.RESULT_FORMAT_TEXT,
			IncrementalOutput: true, // 启用流式输出
			TranslationOptions: util.TranslationOptions{
				SourceLang: "English",
				TargetLang: "Chinese",
			},
		},
	}

	// 发送流式翻译请求
	return t.ChatCli.ChatStream(chatRequest)
}

// TranslateDocumentStream 使用流式API翻译整个文档
func (t *QwenTranslator) TranslateDocumentStream(content interface{}) (*util.StreamReader, error) {
	contentBytes, ok := content.([]byte)
	if !ok {
		return nil, errors.New("content is not a []byte")
	}

	return t.TranslateSectionStream(contentBytes)
}
