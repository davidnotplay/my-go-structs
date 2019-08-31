package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBst_func(t *testing.T) {
	bst := NewBst()

	assert.Nil(t, bst.root)
	assert.Equal(t, bst.length, 0)
	assert.False(t, bst.rebalance)
	assert.False(t, bst.duplicated)
}
