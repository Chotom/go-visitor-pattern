package binary_expression_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_ParseWithUnknownOperator(t *testing.T) {
	// given
	parser := NewParser()
	_, err := parser.Parse("1 + ( 2 - 3 * 4 + 5 ) / 5")

	// then
	assert.Nil(t, err)
}

func TestInFixVisitor_Result(t *testing.T) {
	// given
	parser := NewParser()
	visitor := InFixVisitor{}
	rootNode, err := parser.Parse("1 + ( 2 - 3 * 4 ) / 5")

	//when
	rootNode.Accept(&visitor)

	// then
	assert.Nil(t, err)
	assert.Equal(t, "(1 + ((2 - (3 * 4)) / 5))", visitor.Result())
}
