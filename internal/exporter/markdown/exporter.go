package markdown

import (
	"bytes"
	"fmt"
	"plo/internal/core"
	"sort"
	"text/template"
)

// 定义高度精简的模板
const skillTemplate = `
# 🛠 提示词流水线执行技能 (Skill: Plo-Pipeline)

## 第一部分：全局流程拓扑 (Workflow Topology)
> [PROTOCOL]: 此部分仅用于确定逻辑流转，严禁在此执行具体指令。

| 当前节点 ID | 节点名称 | 跳转条件 (Condition) | 下一个节点 (Target) |
| :--- | :--- | :--- | :--- |
{{- range .Flows}}
| **{{.Label}}** | {{.Name}} | {{if .HasEdges}}{{range .Edges}} [IF] "{{.Condition}}" <br> {{end}}{{else}} - {{end}} | {{if .HasEdges}}{{range .Edges}} ➡️ **{{.TargetLabel}}** ({{.TargetName}}) <br> {{end}}{{else}} 🏁 END_OF_FLOW {{end}} |
{{- end}}

---

## 第二部分：节点指令详情 (Node Specifications)
> [PROTOCOL]: 确定当前节点后，请严格执行以下 [INSTRUCTION] 内容。

{{range .Details}}
### 📍 {{.Label}} | {{.Name}}

**[INSTRUCTION]**
"""
{{if .Prompt}}{{.Prompt}}{{else}}(此节点无特定指令，请直接进行逻辑流转){{end}}
"""

**[CONTEXT/DATA]**
> {{if .Content}}{{.Content}}{{else}}无额外上下文{{end}}

---
{{end}}
`

type FlowNode struct {
	Label    string
	Name     string
	HasEdges bool
	Edges    []EdgeInfo
}

type EdgeInfo struct {
	Condition   string
	TargetLabel string
	TargetName  string
}

type NodeDetail struct {
	Label   string
	Name    string
	Prompt  string
	Content string
}

type Exporter struct{}

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e *Exporter) Export(p *core.Pipeline) ([]byte, error) {
	// 1. 对节点进行排序，保证 Node00X 编号的稳定性
	nodes := make([]*core.Node, 0, len(p.Nodes))
	for _, n := range p.Nodes {
		nodes = append(nodes, n)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].ID < nodes[j].ID
	})

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
			Label:    idToLabel[node.ID],
			Name:     node.Name,
			HasEdges: len(outEdges) > 0,
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
				TargetName:  target.Name,
			})
		}
		flowData = append(flowData, fn)

		// 构造第二部分：执行细节
		detailData = append(detailData, NodeDetail{
			Label:   idToLabel[node.ID],
			Name:    node.Name,
			Prompt:  node.Prompt,
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
