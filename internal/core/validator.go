package core

import (
	"fmt"
)

// Validate 校验 Pipeline 的完整性和正确性
func Validate(p *Pipeline) error {
	if p == nil {
		return fmt.Errorf("pipeline is nil")
	}

	if err := validateSingleStartNode(p); err != nil {
		return err
	}

	if err := validateUnconditionalEdges(p); err != nil {
		return err
	}

	if err := validateEdgeReferences(p); err != nil {
		return err
	}

	return nil
}

// validateSingleStartNode 确保有且仅有一个起始节点
func validateSingleStartNode(p *Pipeline) error {
	startCount := 0
	for _, node := range p.Nodes {
		if node.IsStart {
			startCount++
		}
	}

	switch {
	case startCount == 0:
		return NewValidationError("IsStart", "no start node found")
	case startCount > 1:
		return NewValidationError("IsStart", fmt.Sprintf("multiple start nodes found: %d", startCount))
	default:
		return nil
	}
}

// validateUnconditionalEdges 确保每个节点最多有一个无条件出边
func validateUnconditionalEdges(p *Pipeline) error {
	// 为每个节点统计无条件出边
	unconditionalCount := make(map[string]int)

	for _, edge := range p.Edges {
		if edge.Condition == "" {
			unconditionalCount[edge.SourceID]++
		}
	}

	for nodeID, count := range unconditionalCount {
		if count > 1 {
			return NewValidationError(
				"Edges",
				fmt.Sprintf("node %s has %d unconditional edges, maximum 1 allowed", nodeID, count),
			)
		}
	}

	return nil
}

// validateEdgeReferences 确保边引用的节点都存在
func validateEdgeReferences(p *Pipeline) error {
	for _, edge := range p.Edges {
		if edge.SourceID == "" {
			return NewValidationError("Edge.SourceID", "edge has empty source ID")
		}
		if edge.TargetID == "" {
			return NewValidationError("Edge.TargetID", "edge has empty target ID")
		}
		if p.Nodes[edge.SourceID] == nil {
			return NewValidationError("Edge.SourceID", fmt.Sprintf("source node %s not found", edge.SourceID))
		}
		if p.Nodes[edge.TargetID] == nil {
			return NewValidationError("Edge.TargetID", fmt.Sprintf("target node %s not found", edge.TargetID))
		}
	}
	return nil
}
