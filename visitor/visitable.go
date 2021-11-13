package visitor

// Visitable defines the interface for nodes that can accept Visitor.
type Visitable interface {
	Accept(BinaryTreeVisitor)
}

//BinaryTreeNode represents a tree node with two children
type BinaryTreeNode struct {
	left  Visitable
	right Visitable
}

func NewBinaryTreeNode(left, right Visitable) *BinaryTreeNode {
	return &BinaryTreeNode{left: left, right: right}
}

func (b *BinaryTreeNode) Left() Visitable {
	return b.left
}

func (b *BinaryTreeNode) Right() Visitable {
	return b.right
}

func (b *BinaryTreeNode) Accept(visitor BinaryTreeVisitor) {
	visitor.VisitNode(*b)
}

//BinaryTreeLeaf represents a leaf node that stores value
type BinaryTreeLeaf struct {
	value int
}

func NewBinaryTreeLeaf(value int) *BinaryTreeLeaf {
	return &BinaryTreeLeaf{value: value}
}

func (b *BinaryTreeLeaf) Value() int {
	return b.value
}

func (b *BinaryTreeLeaf) Accept(visitor BinaryTreeVisitor) {
	visitor.VisitLeaf(*b)
}
