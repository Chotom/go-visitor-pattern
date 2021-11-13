package treenode

import "strconv"

type simpleTreeNode struct {
	label treeNodeValue
}

func NewSimpleTreeNode(label string) TreeNode {
	return &simpleTreeNode{
		label: newTreeNodeValue(label),
	}
}

func (s *simpleTreeNode) ToInFix() string {
	return string(s.label)
}

func (s *simpleTreeNode) ToPostFix() string {
	return string(s.label)
}

func (s *simpleTreeNode) ToPreFix() string {
	return string(s.label)
}

func (s *simpleTreeNode) Evaluate() (float64, error) {
	if result, err := strconv.ParseFloat(string(s.label), 64); err != nil {
		return 0, err
	} else {
		return result, nil
	}
}
