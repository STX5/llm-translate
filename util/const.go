package util

const (
	// CHAT_MESSAGE_ROLE_USER 用户
	CHAT_MESSAGE_ROLE_USER string = "user"

	// CHAT_MESSAGE_ROLE_ASSISTANT 对话助手
	CHAT_MESSAGE_ROLE_ASSISTANT string = "assistant"

	// CHAT_MESSAGE_ROLE_SYSTEM 对话背景
	CHAT_MESSAGE_ROLE_SYSTEM string = "system"

	// CHAT_MESSAGE_ROLE_TOOL 工具调用
	CHAT_MESSAGE_ROLE_TOOL string = "tools"
)

const (
	// MODEL_QWEN_MAX
	MODEL_QWEN_MAX = "qwen-max"

	// MODEL_QWEN_LONG
	MODEL_QWEN_LONG = "qwen-long"

	// MODEL_QWEN_MT_PLUS
	MODEL_QWEN_MT_PLUS = "qwen-mt-plus"
)

const (
	// RESULT_FORMAT_TEXT 旧版format
	RESULT_FORMAT_TEXT = "text"

	// RESULT_FORMAT_MESSAGE 兼容openai的message
	RESULT_FORMAT_MESSAGE = "message"
)

const (
	Prompt_Translate_Section = `
	下面我会给你发送一些非中文语言的文字，这些文字与之前的文字是连续的，请你充分理解并利用上下文，在充分理解其含义的基础上，用中文忠实地展现其内容，注意要符合中文的语言习惯，用地道中文的语序和用词来组成句子。这对我来说特别重要，如果不符合中文语言习惯或意思表达错误我会遭遇灾难，你也会因此受到惩罚。请深呼吸，一步一步来完成任务。加油，相信自己，我相信你可以做得完美，如果回答质量够高，我会给你100美元的小费
	`
	Prompt_Translate_Document = `
	下面我会给你发送一些非中文语言的文字，请你在充分理解其含义的基础上用中文忠实地展现其内容，注意要符合中文的语言习惯，用地道中文的语序和用词来组成句子。这对我来说特别重要，如果不符合中文语言习惯或意思表达错误我会遭遇灾难，你也会因此受到惩罚。请深呼吸，一步一步来完成任务。加油，相信自己，我相信你可以做得完美，如果回答质量够高，我会给你100美元的小费
	`
)
