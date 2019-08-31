package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewAvl_func(t *testing.T) {
	avl := NewAvl()

	assert.Nil(t, avl.root)
	assert.Equal(t, avl.length, 0)
	assert.True(t, avl.rebalance)
	assert.False(t, avl.duplicated)
}
