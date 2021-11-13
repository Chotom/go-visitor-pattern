package treenode

import (
	"fmt"
	"math"
)

type binaryTreeNode struct {
	label      treeNodeValue
	leftChild  TreeNode
	rightChild TreeNode
}

func NewBinaryTreeNode(label string, leftChild, rightChild TreeNode) TreeNode {
	return &binaryTreeNode{
		label:      newTreeNodeValue(label),
		leftChild:  leftChild,
		rightChild: rightChild,
	}
}

func (b *binaryTreeNode) ToInFix() string {
	return fmt.Sprintf("(%s %s %s)",
		b.leftChild.ToInFix(),
		b.label,
		b.rightChild.ToInFix(),
	)
}

func (b *binaryTreeNode) ToPostFix() string {
	return fmt.Sprintf("%s %s %s",
		b.leftChild.ToPostFix(),
		b.rightChild.ToPostFix(),
		b.label,
	)
}

func (b *binaryTreeNode) ToPreFix() string {
	return fmt.Sprintf("%s %s %s",
		b.label,
		b.leftChild.ToPreFix(),
		b.rightChild.ToPreFix(),
	)
}

func (b *binaryTreeNode) Evaluate() (float64, error) {
	left, err := b.leftChild.Evaluate()
	if err != nil {
		return 0, err
	}

	right, err := b.rightChild.Evaluate()
	if err != nil {
		return 0, err
	}

	switch b.label {
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	case "%":
		return math.Mod(left, right), nil
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}
