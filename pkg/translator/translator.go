package translator

import (
	"github.com/STX5/llm-translate/util"
)

// 翻译器接口
type Translator interface {
	// 翻译文本片段
	TranslateSection(text []byte) ([]byte, error)
	// 翻译整个文档
	TranslateDocument(content interface{}) ([]byte, error)
	// 使用流式API翻译文本片段
	TranslateSectionStream(text []byte) (*util.StreamReader, error)
	// 使用流式API翻译整个文档
	TranslateDocumentStream(content interface{}) (*util.StreamReader, error)
}
