package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBst_func(t *testing.T) {
	bst := NewBst()

	assert.Nil(t, bst.root, "bst root isn't nil")
	assert.Equal(t, bst.length, 0, "bst length is incorrect")
	assert.False(t, bst.rebalance, "bst mustn't be rebalanced")
	assert.False(t, bst.duplicated, "bst duplicated flag is incorrect")
}
