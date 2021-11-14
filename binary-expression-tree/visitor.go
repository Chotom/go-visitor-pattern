package binary_expression_tree

// Visitor Defines the interface for Visitors that operate on tree nodes
type Visitor interface {
	VisitBinaryOperatorNode(BinaryOperatorNode)
	VisitNumericNode(NumericNode)
}
