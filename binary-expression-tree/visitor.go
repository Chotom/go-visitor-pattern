package binary_expression_tree

// Visitor defines the interface for Visitors that operate on tree nodes
type Visitor interface {
	VisitBinaryOperatorNode(BinaryOperatorNode)
	VisitNumericNode(NumericNode)
	Result() string
}

type InFixVisitor struct {
	result string
}

func (i *InFixVisitor) VisitBinaryOperatorNode(node BinaryOperatorNode) {
	i.result += "("
	node.Left().Accept(i)
	i.result += " " + node.GetLabel() + " "
	node.Right().Accept(i)
	i.result += ")"
}

func (i *InFixVisitor) VisitNumericNode(node NumericNode) {
	i.result += node.GetLabel()
}

func (i *InFixVisitor) Result() string {
	return i.result
}
