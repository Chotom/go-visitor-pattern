package binary_expression_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	treeStack        []VisitableTreeNode
	operatorStack    []string
	operatorPriority map[string]int
}

func NewParser() *Parser {
	return &Parser{
		operatorPriority: map[string]int{
			"*": 2,
			"/": 2,
			"+": 1,
			"-": 1,
		},
	}
}

// Parse return pointer to root of final tree
func (p *Parser) Parse(inFixExpression string) (VisitableTreeNode, error) {
	tokens := strings.Split(inFixExpression, " ")
	for _, token := range tokens {
		if err := p.handleToken(token); err != nil {
			return nil, err
		}
	}

	for len(p.operatorStack) != 0 {
		p.popConnectPush()
	}

	return p.treeStack[len(p.treeStack)-1], nil
}

func (p *Parser) handleToken(token string) error {
	if token == "(" {
		p.operatorStack = append(p.operatorStack, token)
	} else if numericValue, err := strconv.Atoi(token); err == nil {
		p.treeStack = append(p.treeStack, NewNumericNode(numericValue))
	} else if priority := p.getOperatorPriority(token); priority >= 0 {
		p.handleOperator(token)
	} else if token == ")" {
		p.handleEndOfParentheses()
	} else {
		return fmt.Errorf("error in expression")
	}

	return nil
}

func (p *Parser) handleOperator(token string) {
	for {
		if opStackLen := len(p.operatorStack); opStackLen == 0 {
			break
		} else if top := p.operatorStack[opStackLen-1]; top == "(" {
			break
		} else if p.getOperatorPriority(top) < p.getOperatorPriority(token) {
			break
		} else {
			p.popConnectPush()
		}
	}
	p.operatorStack = append(p.operatorStack, token)
}

func (p *Parser) handleEndOfParentheses() {
	for p.operatorStack[len(p.operatorStack)-1] != "(" {
		p.popConnectPush()
	}
	p.operatorStack = p.operatorStack[:len(p.operatorStack)-1]
}

func (p *Parser) popConnectPush() {
	right := p.treeStackPop()
	left := p.treeStackPop()

	opTop := p.operatorStack[len(p.operatorStack)-1]
	p.operatorStack = p.operatorStack[:len(p.operatorStack)-1]

	tree := p.binaryOperatorNodeFactory(opTop, left, right)
	p.treeStack = append(p.treeStack, &tree)
}

func (p *Parser) treeStackPop() VisitableTreeNode {
	treeStackLen := len(p.treeStack)

	top := p.treeStack[treeStackLen-1]
	p.treeStack = p.treeStack[:treeStackLen-1]
	return top
}

func (p *Parser) getOperatorPriority(token string) int {
	if priority, ok := p.operatorPriority[token]; ok {
		return priority
	} else {
		return -1
	}
}

func (p *Parser) binaryOperatorNodeFactory(operator string, left, right VisitableTreeNode) BinaryOperatorNode {
	var node OperatorNode
	switch operator {
	case "*":
		node = NewMultiplicationNode(left, right)
	case "/":
		node = NewDivisionNode(left, right)
	case "-":
		node = NewSubtractionNode(left, right)
	case "+":
		node = NewAdditionNode(left, right)
	}
	return *NewBinaryOperatorNode(node)
}
