package visitor

import (
	"fmt"
	"strconv"
)

// BinaryTreeVisitor Defines the interface for Visitors that operate on binary trees
type BinaryTreeVisitor interface {
	VisitNode(BinaryTreeNode)
	VisitLeaf(BinaryTreeLeaf)
	Report() string
}

//SummationVisitor implements a BinaryTreeVisitor that collects the sum of all BinaryTreeLeaf
type SummationVisitor struct {
	sum int
}

func NewSummationVisitor() *SummationVisitor {
	return &SummationVisitor{sum: 0}
}

func (s *SummationVisitor) VisitNode(node BinaryTreeNode) {
	node.left.Accept(s)
	node.right.Accept(s)
}

func (s *SummationVisitor) VisitLeaf(leaf BinaryTreeLeaf) {
	s.sum += leaf.Value()
}

func (s *SummationVisitor) Report() string {
	return fmt.Sprintf("SummationVisitor collected a sum of: %d", s.sum)
}

//TraversalVisitor implements a BinaryTreeVisitor that collects string representation of binary tree
type TraversalVisitor struct {
	result string
}

func NewTraversalVisitor() *TraversalVisitor {
	return &TraversalVisitor{result: ""}
}

func (t *TraversalVisitor) VisitNode(node BinaryTreeNode) {
	t.result += "{"
	node.Left().Accept(t)
	t.result += ","
	node.Right().Accept(t)
	t.result += "}"
}

func (t *TraversalVisitor) VisitLeaf(leaf BinaryTreeLeaf) {
	t.result += strconv.Itoa(leaf.Value())
}

func (t *TraversalVisitor) Report() string {
	return fmt.Sprintf("TraversalVisitor traversed the tree to: %s", t.result)
}
