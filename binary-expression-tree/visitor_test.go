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
