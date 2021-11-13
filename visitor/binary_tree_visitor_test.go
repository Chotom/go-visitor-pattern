package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSummationVisitor_Report(t *testing.T) {
	// given
	sumVisitor := NewSummationVisitor()
	root := getTree()

	// when
	root.Accept(sumVisitor)

	// then
	assert.Equal(t, "SummationVisitor collected a sum of: 6", sumVisitor.Report())
}

func TestTraversalVisitor_Report(t *testing.T) {
	// given
	traversalVisitor := NewTraversalVisitor()
	root := getTree()

	// when
	root.Accept(traversalVisitor)

	// then
	assert.Equal(t, "TraversalVisitor traversed the tree to: {{1,2},3}", traversalVisitor.Report())
}

// getTree Returns root of tree:
//          root
//         /    \
//     regN      3
//    /    \
//   1      2
func getTree() *BinaryTreeNode {
	one := NewBinaryTreeLeaf(1)
	two := NewBinaryTreeLeaf(2)
	three := NewBinaryTreeLeaf(3)

	regN := NewBinaryTreeNode(one, two)
	root := NewBinaryTreeNode(regN, three)

	return root
}
