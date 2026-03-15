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
	inputFile := flag.String("input", "examples/example.drawio", "Input Draw.io file path")
	outputFile := flag.String("output", "", "Output Markdown file path (optional, defaults to stdout)")
	flag.Parse()

	if err := run(*inputFile, *outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(inputFile, outputFile string) error {
	// 1. Read file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// 2. Parse with DrawioAdapter
	parser := drawio.NewAdapter()
	pipeline, err := parser.Parse(data)
	if err != nil {
		return fmt.Errorf("failed to parse Draw.io file: %w", err)
	}

	// 3. Validate Pipeline
	if err := core.Validate(pipeline); err != nil {
		return fmt.Errorf("pipeline validation failed: %w", err)
	}

	// 4. Export with MarkdownExporter
	exporter := markdown.NewExporter()
	result, err := exporter.Export(pipeline)
	if err != nil {
		return fmt.Errorf("failed to export Markdown: %w", err)
	}

	// 5. Output result
	if outputFile != "" {
		if err := os.WriteFile(outputFile, result, 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Successfully exported to: %s\n", outputFile)
	} else {
		fmt.Println(string(result))
	}

	return nil
}
