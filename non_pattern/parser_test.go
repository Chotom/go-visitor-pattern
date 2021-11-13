package non_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_ParseWithUnknownOperator(t *testing.T) {
	// given
	parser := NewParser()
	_, err := parser.Parse("1 + ( 2 @ 3 * 4 % 5 ) / 5")

	// then
	assert.NotNil(t, err)
}

func TestTreeNode_toInFix(t *testing.T) {
	// given
	parser := NewParser()
	rootNode, err := parser.Parse("1 + ( 2 - 3 * 4 ) / 5")

	// then
	assert.Nil(t, err)
	assert.Equal(t, "(1 + ((2 - (3 * 4)) / 5))", rootNode.ToInFix())
}

func TestTreeNode_toPostFix(t *testing.T) {
	// given
	parser := NewParser()
	rootNode, err := parser.Parse("1 + ( 2 - 3 * 4 ) / 5")

	// then
	assert.Nil(t, err)
	assert.Equal(t, "1 2 3 4 * - 5 / +", rootNode.ToPostFix())
}

func TestTreeNode_toPreFix(t *testing.T) {
	// given
	parser := NewParser()
	rootNode, err := parser.Parse("1 + ( 2 - 3 * 4 ) / 5")

	// then
	assert.Nil(t, err)
	assert.Equal(t, "+ 1 / - 2 * 3 4 5", rootNode.ToPreFix())
}

func TestTreeNode_Evaluate(t *testing.T) {
	// given
	parser := NewParser()
	rootNode, _ := parser.Parse("1 + ( 2 - 3 * 4 % 5 ) / 5")

	// when
	result, err := rootNode.Evaluate()

	// then
	assert.Nil(t, err)
	assert.Equal(t, float64(1), result)
}
