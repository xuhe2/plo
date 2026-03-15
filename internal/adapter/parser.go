package adapter

import "plo/internal/core"

// Parser 定义了解析器接口，将原始数据转换为 Pipeline
type Parser interface {
	Parse(data []byte) (*core.Pipeline, error)
}
