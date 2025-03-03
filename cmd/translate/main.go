package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/STX5/llm-translate/pkg/chat/qwen"
	"github.com/STX5/llm-translate/pkg/document/txt"
	"github.com/STX5/llm-translate/pkg/translator"
	"github.com/STX5/llm-translate/util"
)

func main() {
	// 解析命令行参数
	inputFile := flag.String("input", "", "输入文件路径")
	outputFile := flag.String("output", "", "输出文件路径")
	apiKey := flag.String("apikey", util.ApiKey, "通义千问API密钥")
	maxLength := flag.Int("max-length", 2000, "每个片段的最大长度")
	debug := flag.Bool("debug", false, "是否开启调试模式")
	flag.Parse()

	// 检查必要参数
	if *inputFile == "" {
		fmt.Println("请指定输入文件路径，使用 --input 参数")
		os.Exit(1)
	}

	if *apiKey == "" {
		fmt.Println("请指定通义千问API密钥，使用 --apikey 参数")
		os.Exit(1)
	}

	// 创建通义千问客户端
	chatCli := qwen.NewClient(*apiKey, *debug)

	// 创建翻译器
	qwenTranslator := translator.NewQwenTranslator(chatCli)

	// 创建TXT翻译器
	txtTranslator, err := txt.NewTxtTranslator(*inputFile)
	if err != nil {
		fmt.Printf("创建TXT翻译器失败: %v\n", err)
		os.Exit(1)
	}

	// 设置翻译器和参数
	txtTranslator.SetTranslator(qwenTranslator)
	txtTranslator.SetMaxLength(*maxLength)
	if *outputFile != "" {
		txtTranslator.SetOutputPath(*outputFile)
	}

	// 执行翻译
	fmt.Println("开始翻译文档...")
	err = txtTranslator.TranslateDocument()
	if err != nil {
		fmt.Printf("翻译文档失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("文档翻译完成！")
	if *outputFile != "" {
		fmt.Printf("翻译结果已保存到: %s\n", *outputFile)
	} else {
		ext := ".txt"
		if len(*inputFile) > 4 {
			ext = (*inputFile)[len(*inputFile)-4:]
		}
		fmt.Printf("翻译结果已保存到: %s.translated%s\n", (*inputFile)[:len(*inputFile)-len(ext)], ext)
	}
}
