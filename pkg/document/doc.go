package document

type Document interface {
	FileReader
	FileProcessor
}

type FileReader interface {
	// default read, with context length limit
	Read(path string) (interface{}, error)
	// // read page, with context length limit
	// ReadPage(start int, end int) (interface{}, error)
	// // read line, with context length limit
	// ReadLine(start int, end int) (interface{}, error)
}

type FileProcessor interface {
	// 处理文档并返回内容，支持多种格式（pdf、txt等）
	ProcessDocument() (interface{}, error)
	// 检查内容是否超过指定的最大长度
	ExceedsMaxLength(maxLength int) bool
	// 将内容分割成多个片段，每个片段不超过指定的最大长度
	SplitContent(content interface{}, maxLength int) (interface{}, error)
}
