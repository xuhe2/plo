package drawio

import (
	"encoding/xml"
	"fmt"
	"html"
	"plo/internal/core"
	"strconv"
	"strings"
)

// Adapter 实现了 Draw.io XML 解析器
type Adapter struct{}

// NewAdapter 创建一个新的 DrawioAdapter
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Parse 解析 Draw.io XML 数据为 Pipeline
func (a *Adapter) Parse(data []byte) (*core.Pipeline, error) {
	var mxFile MxFile
	if err := xml.Unmarshal(data, &mxFile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	pipeline := core.NewPipeline()

	// 收集所有 cells 并建立索引
	cells := make(map[string]*MxCell)
	for i := range mxFile.Diagram.MxGraphModel.Root.MxCells {
		cell := &mxFile.Diagram.MxGraphModel.Root.MxCells[i]
		cells[cell.ID] = cell
	}

	// 第一遍：识别边并收集边的标签
	edgesByID := make(map[string]*MxCell)      // 边 ID -> 边 cell
	edgeLabels := make(map[string]string)       // 边 ID -> 标签文本

	for _, cell := range cells {
		if cell.IsEdge() {
			edgesByID[cell.ID] = cell
			// 边本身可能有 value
			if v := decodeValue(cell.Value); v != "" {
				edgeLabels[cell.ID] = v
			}
		}
	}

	// 查找边的标签（标签是顶点，parent 指向边）
	for _, cell := range cells {
		if cell.IsVertex() && cell.Parent != "" {
			if _, isEdge := edgesByID[cell.Parent]; isEdge {
				if v := decodeValue(cell.Value); v != "" {
					edgeLabels[cell.Parent] = v
				}
			}
		}
	}

	// 第二遍：收集实际的节点（顶点，但不是标签，也不是边）
	nodes := make(map[string]*nodeInfo)

	for _, cell := range cells {
		if cell.IsVertex() {
			// 跳过边的标签（parent 指向边的顶点）
			if cell.Parent != "" {
				if _, isEdge := edgesByID[cell.Parent]; isEdge {
					continue
				}
			}
			// 需要有 geometry 信息（有位置和大小）
			if cell.Geometry == nil {
				continue
			}
			// 跳过没有 value 的节点（可能只是容器）
			if cell.Value == "" {
				continue
			}
			// 跳过那些自身就是根节点的 cell (ID 为 0 或 1)
			if cell.ID == "0" || cell.ID == "1" {
				continue
			}

			info := &nodeInfo{
				id:   cell.ID,
				name: decodeValue(cell.Value),
				cell: cell,
			}
			if cell.Geometry != nil {
				info.x, _ = parseFloat(cell.Geometry.X)
				info.y, _ = parseFloat(cell.Geometry.Y)
			}
			nodes[cell.ID] = info
		}
	}

	// 找到起始节点（最左上的节点）
	var startNode *nodeInfo
	for _, n := range nodes {
		if startNode == nil {
			startNode = n
		} else {
			// 比较位置，最左上的为起始节点
			if n.y < startNode.y || (n.y == startNode.y && n.x < startNode.x) {
				startNode = n
			}
		}
	}

	// 添加节点到 Pipeline
	for _, info := range nodes {
		node := &core.Node{
			ID:      info.id,
			Name:    info.name,
			Prompt:  info.name,
			Content: info.name,
			IsStart: info == startNode,
		}
		pipeline.AddNode(node)
	}

	// 添加边到 Pipeline
	for edgeID, edgeCell := range edgesByID {
		if edgeCell.Source != "" && edgeCell.Target != "" {
			// 确保源和目标都是我们收集的节点
			if _, ok := nodes[edgeCell.Source]; !ok {
				continue
			}
			if _, ok := nodes[edgeCell.Target]; !ok {
				continue
			}

			condition := edgeLabels[edgeID]
			edge := &core.Edge{
				ID:        edgeID,
				SourceID:  edgeCell.Source,
				TargetID:  edgeCell.Target,
				Condition: condition,
			}
			pipeline.AddEdge(edge)
		}
	}

	return pipeline, nil
}

// findRootCells 找到根级别的 cell（parent 为 0 或 1）
func (a *Adapter) findRootCells(cells map[string]*MxCell) []*MxCell {
	var roots []*MxCell
	for _, cell := range cells {
		if cell.Parent == "0" || cell.Parent == "1" || cell.Parent == "" {
			if cell.ID != "0" && cell.ID != "1" {
				roots = append(roots, cell)
			}
		}
	}
	return roots
}

// nodeInfo 用于存储节点的临时信息
type nodeInfo struct {
	id   string
	name string
	x    float64
	y    float64
	cell *MxCell
}

// decodeValue 解码 HTML 实体并清理字符串
func decodeValue(s string) string {
	if s == "" {
		return ""
	}
	s = html.UnescapeString(s)
	s = stripHTMLTags(s)
	s = strings.TrimSpace(s)
	return s
}

// stripHTMLTags 简单地移除 HTML 标签
func stripHTMLTags(s string) string {
	var builder strings.Builder
	inTag := false
	for _, r := range s {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
		default:
			if !inTag {
				builder.WriteRune(r)
			}
		}
	}
	return builder.String()
}

// parseFloat 解析浮点数，失败返回 0
func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}
