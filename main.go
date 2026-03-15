package main

import (
	"flag"
	"fmt"
	"os"
	"plo/internal/adapter/drawio"
	"plo/internal/core"
	"plo/internal/exporter/markdown"
)

func main() {
	inputFile := flag.String("input", "example.drawio", "输入 Draw.io 文件路径")
	outputFile := flag.String("output", "", "输出 Markdown 文件路径 (可选，不指定则输出到 stdout)")
	flag.Parse()

	if err := run(*inputFile, *outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

func run(inputFile, outputFile string) error {
	// 1. 读取文件
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	// 2. 使用 DrawioAdapter 解析
	parser := drawio.NewAdapter()
	pipeline, err := parser.Parse(data)
	if err != nil {
		return fmt.Errorf("解析 Draw.io 文件失败: %w", err)
	}

	// 3. 校验 Pipeline
	if err := core.Validate(pipeline); err != nil {
		return fmt.Errorf("Pipeline 校验失败: %w", err)
	}

	// 4. 使用 MarkdownExporter 导出
	exporter := markdown.NewExporter()
	result, err := exporter.Export(pipeline)
	if err != nil {
		return fmt.Errorf("导出 Markdown 失败: %w", err)
	}

	// 5. 输出结果
	if outputFile != "" {
		if err := os.WriteFile(outputFile, result, 0644); err != nil {
			return fmt.Errorf("写入输出文件失败: %w", err)
		}
		fmt.Printf("成功导出到: %s\n", outputFile)
	} else {
		fmt.Println(string(result))
	}

	return nil
}
