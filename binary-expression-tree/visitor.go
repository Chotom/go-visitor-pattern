package binary_expression_tree

import "fmt"

// Visitor defines the interface for Visitors that operate on tree nodes
type Visitor interface {
	VisitBinaryOperatorNode(BinaryOperatorNode)
	VisitNumericNode(NumericNode)
	Result() string
}

//InFixVisitor represents visitor that operate on tree node and save expression in InFix form
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

//PostFixVisitor represents visitor that operate on tree node and save expression in PostFix form
type PostFixVisitor struct {
	result string
}

func (i *PostFixVisitor) VisitBinaryOperatorNode(node BinaryOperatorNode) {
	node.Left().Accept(i)
	i.result += " "
	node.Right().Accept(i)
	i.result += " " + node.GetLabel()
}

func (i *PostFixVisitor) VisitNumericNode(node NumericNode) {
	i.result += node.GetLabel()
}

func (i *PostFixVisitor) Result() string {
	return i.result
}

//EvaluateVisitor represents visitor that operate on tree node and calculate result of expression
type EvaluateVisitor struct {
	result int
}

func (e *EvaluateVisitor) VisitBinaryOperatorNode(node BinaryOperatorNode) {
	node.Left().Accept(e)
	leftValue := e.result

	node.Right().Accept(e)
	rightValue := e.result

	e.result = node.Compute(leftValue, rightValue)
}

func (e *EvaluateVisitor) VisitNumericNode(node NumericNode) {
	e.result = node.Value()
}

func (e *EvaluateVisitor) Result() string {
	return fmt.Sprintf("%d", e.result)
}
