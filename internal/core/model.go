package core

import "fmt"

// Node 表示流程图中的一个节点
type Node struct {
	ID      string
	Content string

	InEdges  []*Edge
	OutEdges []*Edge

	IsStart bool
}

// Edge 表示流程图中的一条边
type Edge struct {
	ID        string
	SourceID  string
	TargetID  string
	Condition string
}

// Pipeline 表示整个提示词流水线
type Pipeline struct {
	Nodes map[string]*Node
	Edges []*Edge
}

// NewPipeline 创建一个新的 Pipeline 实例
func NewPipeline() *Pipeline {
	return &Pipeline{
		Nodes: make(map[string]*Node),
		Edges: make([]*Edge, 0),
	}
}

// AddNode 添加一个节点到 Pipeline
func (p *Pipeline) AddNode(node *Node) {
	if _, exists := p.Nodes[node.ID]; exists {
		panic(fmt.Sprintf("node %s already exists", node.ID))
	}
	p.Nodes[node.ID] = node
}

// AddEdge 添加一条边到 Pipeline
func (p *Pipeline) AddEdge(edge *Edge) {
	// 同时更新节点的出边和入边
	source := p.Nodes[edge.SourceID]
	target := p.Nodes[edge.TargetID]
	source.OutEdges = append(source.OutEdges, edge)
	target.InEdges = append(target.InEdges, edge)

	p.Edges = append(p.Edges, edge)
}

// GetStartNode 获取起始节点
func (p *Pipeline) GetStartNode() *Node {
	for _, node := range p.Nodes {
		if node.IsStart {
			return node
		}
	}
	return nil
}

// GetOutEdges 获取指定节点的所有出边
func (p *Pipeline) GetOutEdges(nodeID string) []*Edge {
	return p.Nodes[nodeID].OutEdges
}

// GetInEdges 获取指定节点的所有入边
func (p *Pipeline) GetInEdges(nodeID string) []*Edge {
	return p.Nodes[nodeID].InEdges
}

// GetNode 通过 ID 获取节点
func (p *Pipeline) GetNode(id string) *Node {
	return p.Nodes[id]
}
