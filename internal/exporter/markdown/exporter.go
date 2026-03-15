package markdown

import (
	"bytes"
	"fmt"
	"plo/internal/core"
	"text/template"
)

// 定义高度精简的模板
const skillTemplate = `
# 🛠 提示词流水线执行技能 (Skill: Plo-Pipeline)

## 第一部分：全局流程拓扑 (Workflow Topology)
> [PROTOCOL]: 此部分仅用于确定逻辑流转，严禁在此执行具体指令。

| 当前节点 ID | 节点摘要 | 跳转条件 (Condition) | 下一个节点 (Target) |
| :--- | :--- | :--- | :--- |
{{- range .Flows}}
| **{{.Label}}** | {{.ShortContent}} | {{if .HasEdges}}{{range .Edges}} [IF] "{{.Condition}}" <br> {{end}}{{else}} - {{end}} | {{if .HasEdges}}{{range .Edges}} ➡️ **{{.TargetLabel}}** <br> {{end}}{{else}} 🏁 END_OF_FLOW {{end}} |
{{- end}}

---

## 第二部分：节点指令详情 (Node Specifications)
> [PROTOCOL]: 确定当前节点后，请严格执行以下 [INSTRUCTION] 内容。

{{range .Details}}
### 📍 {{.Label}}

**[INSTRUCTION]**
"""
{{if .Content}}{{.Content}}{{else}}(此节点无特定指令，请直接进行逻辑流转){{end}}
"""

---
{{end}}
`

type FlowNode struct {
	Label        string
	ShortContent string
	HasEdges     bool
	Edges        []EdgeInfo
}

type EdgeInfo struct {
	Condition   string
	TargetLabel string
}

type NodeDetail struct {
	Label   string
	Content string
}

type Exporter struct{}

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e *Exporter) Export(p *core.Pipeline) ([]byte, error) {
	// 1. 从 StartNode 开始使用 BFS 遍历，保证拓扑顺序
	startNode := p.GetStartNode()
	if startNode == nil {
		return nil, fmt.Errorf("no start node found in pipeline")
	}

	visited := make(map[string]bool)
	queue := []*core.Node{startNode}
	visited[startNode.ID] = true
	var nodes []*core.Node

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodes = append(nodes, node)

		// 按顺序遍历出边，添加未访问的目标节点
		for _, edge := range node.OutEdges {
			target := p.GetNode(edge.TargetID)
			if !visited[target.ID] {
				visited[target.ID] = true
				queue = append(queue, target)
			}
		}
	}

	// 2. 建立 ID 到 NodeXXX 的映射
	idToLabel := make(map[string]string)
	for i, node := range nodes {
		idToLabel[node.ID] = fmt.Sprintf("Node%03d", i+1)
	}

	// 3. 构造模板数据
	flowData := make([]FlowNode, 0)
	detailData := make([]NodeDetail, 0)

	for _, node := range nodes {
		// 构造第一部分：纯净拓扑
		outEdges := p.GetOutEdges(node.ID)
		fn := FlowNode{
			Label:        idToLabel[node.ID],
			ShortContent: truncateContent(node.Content, 30),
			HasEdges:     len(outEdges) > 0,
		}
		for _, edge := range outEdges {
			target := p.GetNode(edge.TargetID)
			cond := edge.Condition
			if cond == "" {
				cond = "DEFAULT"
			}
			fn.Edges = append(fn.Edges, EdgeInfo{
				Condition:   cond,
				TargetLabel: idToLabel[target.ID],
			})
		}
		flowData = append(flowData, fn)

		// 构造第二部分：执行细节
		detailData = append(detailData, NodeDetail{
			Label:   idToLabel[node.ID],
			Content: node.Content,
		})
	}

	// 4. 渲染
	data := struct {
		Flows   []FlowNode
		Details []NodeDetail
	}{
		Flows:   flowData,
		Details: detailData,
	}

	tmpl, err := template.New("plo").Parse(skillTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.Bytes(), err
}

// truncateContent 截取内容，最多 maxChars 个字符，超出时添加 "..."
func truncateContent(content string, maxChars int) string {
	if content == "" {
		return "(无内容)"
	}
	// 转换为 []rune 以正确处理中文字符
	runes := []rune(content)
	if len(runes) <= maxChars {
		return content
	}
	return string(runes[:maxChars]) + "..."
}
