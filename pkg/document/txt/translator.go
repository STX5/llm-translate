package txt

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/STX5/llm-translate/pkg/translator"
)

type TxtTranslator struct {
	BaseDocument
	Translator translator.Translator
	MaxLength  int
	OutputPath string
}

func NewTxtTranslator(path string) (*TxtTranslator, error) {
	return &TxtTranslator{
		BaseDocument: BaseDocument{Path: path},
		MaxLength:    4000, // 默认最大长度
	}, nil
}

// SetTranslator 设置翻译器
func (t *TxtTranslator) SetTranslator(trans translator.Translator) {
	t.Translator = trans
}

// SetMaxLength 设置最大长度
func (t *TxtTranslator) SetMaxLength(maxLength int) {
	t.MaxLength = maxLength
}

// SetOutputPath 设置输出路径
func (t *TxtTranslator) SetOutputPath(outputPath string) {
	t.OutputPath = outputPath
}

// TranslateDocument 翻译文档
func (t *TxtTranslator) TranslateDocument() error {
	// 检查翻译器是否已设置
	if t.Translator == nil {
		return errors.New("translator not set")
	}

	// 读取文件内容
	if len(t.Content) == 0 {
		_, err := t.Read(t.Path)
		if err != nil {
			return err
		}
	}

	// 获取要翻译的文本片段
	textParts, err := t.ReadForTranslation(t.MaxLength)
	if err != nil {
		return err
	}

	// 翻译每个片段
	translatedParts := make([][]byte, 0, len(textParts))
	for _, part := range textParts {
		translatedContent, err := t.Translator.TranslateSection(part)
		if err != nil {
			return err
		}

		translatedParts = append(translatedParts, translatedContent)
	}

	// 合并翻译后的内容
	mergedContent, err := t.MergeTranslatedContent(translatedParts)
	if err != nil {
		return err
	}

	// 确定输出路径
	outputPath := t.OutputPath
	if outputPath == "" {
		// 如果未指定输出路径，则使用原文件名加上.translated后缀
		ext := filepath.Ext(t.Path)
		baseName := strings.TrimSuffix(t.Path, ext)
		outputPath = baseName + ".translated" + ext
	}

	// 保存翻译后的内容
	return ioutil.WriteFile(outputPath, mergedContent, 0644)
}
