package exporter

import "plo/internal/core"

// Exporter 定义了导出器接口，将 Pipeline 转换为特定格式
type Exporter interface {
	Export(p *core.Pipeline) ([]byte, error)
}
