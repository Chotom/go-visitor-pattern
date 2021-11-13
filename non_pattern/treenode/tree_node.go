package treenode

type TreeNode interface {
	ToInFix() string
	ToPostFix() string
	ToPreFix() string
	Evaluate() (float64, error)
}

// treeNodeValue holds representation of tree node
type treeNodeValue string

func newTreeNodeValue(value string) treeNodeValue {
	return treeNodeValue(value)
}
