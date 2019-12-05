package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewAvl_func(t *testing.T) {
	avl := NewAvl()

	assert.Nil(t, avl.root, "avl tree root isn't nil")
	assert.Equal(t, avl.length, 0, "avl tree length length is incorrect")
	assert.True(t, avl.rebalance, "avl tree must be rebalance")
	assert.False(t, avl.duplicated, "avl tree duplicated flag is incorrect")
}
