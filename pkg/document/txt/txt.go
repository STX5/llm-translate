package txt

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math"
	"unicode"
)

type TxtDocument struct {
	BaseDocument
}

// 基本文档实现
type BaseDocument struct {
	Path     string
	Content  []byte
	Encoding string
}

// 创建新的文档处理器
func NewDocument(path string) (*BaseDocument, error) {
	doc := &BaseDocument{
		Path: path,
	}

	// 立即读取文件内容
	content, err := doc.Read(path)
	if err != nil {
		return nil, err
	}

	// 确保内容已保存到doc.Content
	doc.Content = content.([]byte)

	return doc, nil
}

func (d *BaseDocument) Read(path string) (interface{}, error) {
	// 简单地读取文本文件
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 保存内容到d.Content
	d.Content = content
	return content, nil
}

func (d *BaseDocument) ProcessDocument() (interface{}, error) {
	return d.Content, nil
}

func (d *BaseDocument) ExceedsMaxLength(maxLength int) bool {
	return len(d.Content) > maxLength
}

func (d *BaseDocument) SplitContent(content interface{}, maxLength int) (interface{}, error) {
	contentBytes, ok := content.([]byte)
	if !ok {
		return nil, errors.New("content is not a []byte")
	}
	if len(contentBytes) <= maxLength {
		return [][]byte{contentBytes}, nil
	}

	// 计算需要分割的片段数量
	numParts := int(math.Ceil(float64(len(contentBytes)) / float64(maxLength)))
	parts := make([][]byte, 0, numParts)

	// 尝试在句子边界分割
	text := string(contentBytes)
	currentPos := 0

	for currentPos < len(text) {
		endPos := currentPos + maxLength
		if endPos > len(text) {
			endPos = len(text)
		} else {
			// 尝试在句子边界分割
			for i := endPos - 1; i > currentPos; i-- {
				r := []rune(text[i : i+1])[0]
				if r == '.' || r == '。' || r == '!' || r == '！' || r == '?' || r == '？' {
					// 找到句子结束标记
					endPos = i + 1
					break
				}
			}
		}

		// 如果没有找到合适的句子边界，尝试在单词边界分割
		if endPos == currentPos+maxLength && currentPos+maxLength < len(text) {
			for i := endPos - 1; i > currentPos; i-- {
				if unicode.IsSpace(rune(text[i])) {
					endPos = i + 1
					break
				}
			}
		}

		parts = append(parts, []byte(text[currentPos:endPos]))
		currentPos = endPos
	}

	return parts, nil
}

// 实现TranslatableDocument接口
func (d *BaseDocument) ReadForTranslation(maxLength int) ([][]byte, error) {
	// 如果内容为空，先处理文档
	if len(d.Content) == 0 {
		_, err := d.ProcessDocument()
		if err != nil {
			return nil, err
		}
	}

	// 检查是否需要分割
	if d.ExceedsMaxLength(maxLength) {
		parts, err := d.SplitContent(d.Content, maxLength)
		if err != nil {
			return nil, err
		}
		return parts.([][]byte), nil
	}

	return [][]byte{d.Content}, nil
}

// 合并翻译后的内容
func (d *BaseDocument) MergeTranslatedContent(translatedParts [][]byte) ([]byte, error) {
	var result bytes.Buffer

	for _, part := range translatedParts {
		result.Write(part)
		// 可以在这里添加分隔符，如果需要的话
		// result.WriteString("\n\n")
	}

	return result.Bytes(), nil
}

// 保存翻译后的内容到文件
func (d *BaseDocument) SaveTranslatedContent(translatedContent []byte, outputPath string) error {
	return ioutil.WriteFile(outputPath, translatedContent, 0644)
}
