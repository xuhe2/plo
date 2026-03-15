package core

// Node 表示流程图中的一个节点
type Node struct {
	ID      string
	Name    string
	Prompt  string
	Content string
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
	p.Nodes[node.ID] = node
}

// AddEdge 添加一条边到 Pipeline
func (p *Pipeline) AddEdge(edge *Edge) {
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
	var edges []*Edge
	for _, edge := range p.Edges {
		if edge.SourceID == nodeID {
			edges = append(edges, edge)
		}
	}
	return edges
}

// GetNode 通过 ID 获取节点
func (p *Pipeline) GetNode(id string) *Node {
	return p.Nodes[id]
}
