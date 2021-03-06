package binary_expression_tree

import "fmt"

//VisitableTreeNode defines the interface for tree nodes that can accept Visitor.
type VisitableTreeNode interface {
	TreeNode
	Visitable
}

//Visitable defines the interface that can accept Visitor.
type Visitable interface {
	Accept(Visitor)
}

//TreeNode defines the interface for tree nodes.
type TreeNode interface {
	GetLabel() string
}

//BinaryParent defines the interface for node with two children.
type BinaryParent interface {
	Left() VisitableTreeNode
	Right() VisitableTreeNode
}

//OperatorNode defines the interface for parent nodes with binary operator expression.
type OperatorNode interface {
	TreeNode
	BinaryParent
	Compute(int, int) int
}

//NumericNode represents a node with numeric expression.
type NumericNode struct {
	value int
}

func NewNumericNode(value int) *NumericNode {
	return &NumericNode{value: value}
}

func (n *NumericNode) Accept(visitor Visitor) {
	visitor.VisitNumericNode(*n)
}

func (n *NumericNode) GetLabel() string {
	return fmt.Sprintf("%d", n.value)
}

func (n *NumericNode) Value() int {
	return n.value
}

//ParentNode represents a node with two children.
type ParentNode struct {
	left  VisitableTreeNode
	right VisitableTreeNode
}

func (b *ParentNode) Left() VisitableTreeNode {
	return b.left
}

func (b *ParentNode) Right() VisitableTreeNode {
	return b.right
}

//BinaryOperatorNode represents parent nodes with binary operator expression.
type BinaryOperatorNode struct {
	OperatorNode
}

func NewBinaryOperatorNode(node OperatorNode) *BinaryOperatorNode {
	return &BinaryOperatorNode{OperatorNode: node}
}

func (b *BinaryOperatorNode) Accept(visitor Visitor) {
	visitor.VisitBinaryOperatorNode(*b)
}

//AdditionNode represents a node with '+' operator expression.
type AdditionNode struct {
	ParentNode
}

func NewAdditionNode(left, right VisitableTreeNode) *AdditionNode {
	return &AdditionNode{
		ParentNode: ParentNode{left: left, right: right},
	}
}

func (a *AdditionNode) Compute(left int, right int) int {
	return left + right
}

func (a *AdditionNode) GetLabel() string {
	return "+"
}

//SubtractionNode represents a node with '-' operator expression.
type SubtractionNode struct {
	ParentNode
}

func NewSubtractionNode(left, right VisitableTreeNode) *SubtractionNode {
	return &SubtractionNode{
		ParentNode: ParentNode{left: left, right: right},
	}
}

func (a *SubtractionNode) Compute(left int, right int) int {
	return left - right
}

func (a *SubtractionNode) GetLabel() string {
	return "-"
}

//MultiplicationNode represents a node with '*' operator expression.
type MultiplicationNode struct {
	ParentNode
}

func NewMultiplicationNode(left, right VisitableTreeNode) *MultiplicationNode {
	return &MultiplicationNode{
		ParentNode: ParentNode{left: left, right: right},
	}
}

func (a *MultiplicationNode) Compute(left int, right int) int {
	return left * right
}

func (a *MultiplicationNode) GetLabel() string {
	return "*"
}

//DivisionNode represents a node with '/' operator expression.
type DivisionNode struct {
	ParentNode
}

func NewDivisionNode(left, right VisitableTreeNode) *DivisionNode {
	return &DivisionNode{
		ParentNode: ParentNode{left: left, right: right},
	}
}

func (a *DivisionNode) Compute(left int, right int) int {
	return left / right
}

func (a *DivisionNode) GetLabel() string {
	return "/"
}

//ModuloNode represents a node with '/' operator expression.
type ModuloNode struct {
	ParentNode
}

func NewModuloNode(left, right VisitableTreeNode) *ModuloNode {
	return &ModuloNode{
		ParentNode: ParentNode{left: left, right: right},
	}
}

func (a *ModuloNode) Compute(left int, right int) int {
	return left % right
}

func (a *ModuloNode) GetLabel() string {
	return "%"
}
