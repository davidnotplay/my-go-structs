package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

type intValue struct {
	value int
}

func (i intValue) String() string {
	return fmt.Sprintf("%d", i.value)
}

func (i intValue) Less(v Value) bool {
	vi, _ := v.(*intValue)
	return i.value < vi.value
}

func (i intValue) Eq(v Value) bool {
	vi, _ := v.(*intValue)
	return i.value == vi.value
}

func i(i int) *int {
	return &i
}

func iv(i int) *Value{
	var v Value = intValue{i}
	return &v
}

func checkTree(t *testing.T, node *avlNode, l, r *int) {

	if l == nil {
		assert.Nil(t, node.ltree)
	} else {
		v, _ := (*node.ltree.value).(intValue)
		assert.Equal(t, v.value, *l)
	}

	if r == nil {
		assert.Nil(t, node.rtree)
	} else {
		v, _ := (*node.rtree.value).(intValue)
		assert.Equal(t, v.value, *r)
	}

}

// Tests
func Test_max_func(t *testing.T) {
	assert.Equal(t, max(2, 3), 3)
	assert.Equal(t, max(4, 1), 4)
	assert.Equal(t, max(2, 2), 2)

	assert.Equal(t, max(-2, -3), -2)
	assert.Equal(t, max(-4, -1), -1)
	assert.Equal(t, max(-2, -2), -2)

	assert.Equal(t, max(2, -2), 2)
	assert.Equal(t, max(1, -2), 1)
	assert.Equal(t, max(-1, 2), 2)
}

func Test_avlNode_Left_func(t *testing.T) {
	as := assert.New(t)

	// node is nil
	as.Nil((*avlNode)(nil).Left(), "When node is nil, must returns nil")

	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree1.ltree = &avlNode{nil, nil, -1, iv(2)}
	v, _ := (*tree1.Left().value).(intValue)
	as.Equal(v.value, 2, "left child must has the value 2")
}

func Test_avlNode_Right_func(t *testing.T) {
	as := assert.New(t)

	// node is nil
	as.Nil((*avlNode)(nil).Right(), "When node is nil, must returns nil")

	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree1.rtree = &avlNode{nil, nil, -1, iv(2)}
	v, _ := (*tree1.Right().value).(intValue)
	as.Equal(v.value, 2, "left child must has the value 2")
}

func Test_avlNode_Value_func(t *testing.T) {
	as := assert.New(t)

	// node is nil
	as.Nil((*avlNode)(nil).Value(), "When node is nil, must returns nil")

	tree1 := &avlNode{nil, nil, -1, iv(11)}
	v, _ := tree1.Value().(intValue)
	as.Equal(v.value, 11, "left child must has the value 11")
}

func Test_avlNode_Height_func(t *testing.T) {
	as := assert.New(t)

	as.Equal((*avlNode)(nil).Height(), -1, "When node is nil, must returns -1")
	as.Equal((&avlNode{nil, nil, 33, nil}).Height(), 33, "Return invalid height")
}

func Test_avlNode_maxHeight_func(t *testing.T) {
	as := assert.New(t)

	node := avlNode{nil, nil, 3, nil}
	as.Equal(node.maxHeight(), -1, "children are nil, so the height of the node must be -1")

	node = avlNode{&avlNode{nil, nil, 22, nil}, nil, 3, nil}
	as.Equal(node.maxHeight(), 22, "height value is the height value of the left child")

	node = avlNode{nil, &avlNode{nil, nil, 41, nil}, 3, nil}
	as.Equal(node.maxHeight(), 41, "height value is the height value of the left child")

	node = avlNode{&avlNode{nil, nil, 15, nil}, &avlNode{nil, nil, 30, nil}, 3, nil}
	as.Equal(node.maxHeight(), 30, "height must be 30")
}

func Test_avlNode_rotateRight_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree2 := &avlNode{nil, nil, -1, iv(2)}
	tree3 := &avlNode{nil, nil, -1, iv(3)}
	tree4 := &avlNode{nil, nil, -1, iv(4)}
	tree5 := &avlNode{nil, nil, -1, iv(5)}

	tree2.ltree = tree1
	tree2.rtree = tree3
	tree4.ltree = tree2
	tree4.rtree = tree5

	ntree := tree4.rotateRight()

	checkTree(t, ntree, i(1), i(4))
	checkTree(t, tree1, nil, nil)
	checkTree(t, tree2, i(1), i(4))
	checkTree(t, tree3, nil, nil)
	checkTree(t, tree4, i(3), i(5))
	checkTree(t, tree5, nil, nil)
}

func Test_avlNode_rotateLeft_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree2 := &avlNode{nil, nil, -1, iv(2)}
	tree3 := &avlNode{nil, nil, -1, iv(3)}
	tree4 := &avlNode{nil, nil, -1, iv(4)}
	tree5 := &avlNode{nil, nil, -1, iv(5)}

	tree4.ltree = tree3
	tree4.rtree = tree5
	tree2.ltree = tree1
	tree2.rtree = tree4

	ntree := tree2.rotateLeft();

	checkTree(t, ntree, i(2), i(5))
	checkTree(t, tree1, nil, nil)
	checkTree(t, tree2, i(1), i(3))
	checkTree(t, tree3, nil, nil)
	checkTree(t, tree4, i(2), i(5))
	checkTree(t, tree5, nil, nil)
}

func Test_avlNode_rotateRightLeft_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree2 := &avlNode{nil, nil, -1, iv(2)}
	tree3 := &avlNode{nil, nil, -1, iv(3)}
	tree4 := &avlNode{nil, nil, -1, iv(4)}
	tree5 := &avlNode{nil, nil, -1, iv(5)}

	tree4.ltree = tree3
	tree4.rtree = tree5
	tree2.ltree = tree1
	tree2.rtree = tree4

	ntree := tree2.rotateRightLeft()

	checkTree(t, ntree, i(2), i(4))
	checkTree(t, tree1, nil, nil)
	checkTree(t, tree2, i(1), nil)
	checkTree(t, tree3, i(2), i(4))
	checkTree(t, tree4, nil, i(5))
	checkTree(t, tree5, nil, nil)
}

func Test_avlNode_rotateLeftRight_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, iv(1)}
	tree2 := &avlNode{nil, nil, -1, iv(2)}
	tree3 := &avlNode{nil, nil, -1, iv(3)}
	tree4 := &avlNode{nil, nil, -1, iv(4)}
	tree5 := &avlNode{nil, nil, -1, iv(5)}

	tree2.ltree = tree1
	tree2.rtree = tree3
	tree4.ltree = tree2
	tree4.rtree = tree5

	ntree := tree4.rotateLeftRight()

	checkTree(t, ntree, i(2), i(4))
	checkTree(t, tree1, nil, nil)
	checkTree(t, tree2, i(1), nil)
	checkTree(t, tree3, i(2), i(4))
	checkTree(t, tree4, nil, i(5))
	checkTree(t, tree1, nil, nil)
}
