package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// utils function
// ==============

// i returns the address of the i param
func i(i int) *int {
	return &i
}

// vi returns the integer value inside of the value param.
func vi(value Value) int {
	v, _ := value.(IntValue)
	return v.value
}

// getAllValues returns a map with all values inside of `node` avl node.
func getAllValues(node *avlNode) (values map[int]*Value) {
	var getValue func(node, parent *avlNode)

	getValue = func(node, parent *avlNode) {
		if node == nil {
			return
		}
		getValue(node.ltree, node)
		value := node.value
		v, _ := node.value.(IntValue)
		values[v.value] = &value
		getValue(node.rtree, node)
	}

	values = map[int]*Value{}
	if node != nil {
		getValue(node, nil)
	}

	return
}
func createAvl(intValues ...int) (Avl, map[int]Value) {
	avl := NewAvl()

	values := map[int]Value{}

	for _, i := range intValues {
		values[i] = Iv(i)
		avl.Insert(values[i])
	}

	return avl, values
}

// checkVAndH checks the v value and the h height of the `node` avl node.
func checkVAndH(t *testing.T, node *avlNode, v, h int) {
	assert.Equal(t, vi(node.value), v)
	assert.Equal(t, node.height, h)
}

// checkTree check the int values of the children of `node` avl node.
func checkTree(t *testing.T, node *avlNode, l, r *int) {

	if l == nil {
		assert.Nil(t, node.ltree)
	} else {
		v, _ := node.ltree.value.(IntValue)
		assert.Equal(t, v.value, *l)
	}

	if r == nil {
		assert.Nil(t, node.rtree)
	} else {
		v, _ := node.rtree.value.(IntValue)
		assert.Equal(t, v.value, *r)
	}

}

//
// Start tests here
// ================
//

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
	tree1 := &avlNode{nil, nil, -1, Iv(1)}
	tree2 := &avlNode{nil, nil, -1, Iv(2)}
	tree3 := &avlNode{nil, nil, -1, Iv(3)}
	tree4 := &avlNode{nil, nil, -1, Iv(4)}
	tree5 := &avlNode{nil, nil, -1, Iv(5)}

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
	tree1 := &avlNode{nil, nil, -1, Iv(1)}
	tree2 := &avlNode{nil, nil, -1, Iv(2)}
	tree3 := &avlNode{nil, nil, -1, Iv(3)}
	tree4 := &avlNode{nil, nil, -1, Iv(4)}
	tree5 := &avlNode{nil, nil, -1, Iv(5)}

	tree4.ltree = tree3
	tree4.rtree = tree5
	tree2.ltree = tree1
	tree2.rtree = tree4

	ntree := tree2.rotateLeft()

	checkTree(t, ntree, i(2), i(5))
	checkTree(t, tree1, nil, nil)
	checkTree(t, tree2, i(1), i(3))
	checkTree(t, tree3, nil, nil)
	checkTree(t, tree4, i(2), i(5))
	checkTree(t, tree5, nil, nil)
}

func Test_avlNode_rotateRightLeft_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, Iv(1)}
	tree2 := &avlNode{nil, nil, -1, Iv(2)}
	tree3 := &avlNode{nil, nil, -1, Iv(3)}
	tree4 := &avlNode{nil, nil, -1, Iv(4)}
	tree5 := &avlNode{nil, nil, -1, Iv(5)}

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
	tree1 := &avlNode{nil, nil, -1, Iv(1)}
	tree2 := &avlNode{nil, nil, -1, Iv(2)}
	tree3 := &avlNode{nil, nil, -1, Iv(3)}
	tree4 := &avlNode{nil, nil, -1, Iv(4)}
	tree5 := &avlNode{nil, nil, -1, Iv(5)}

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

func Test_newAvl_func(t *testing.T) {
	as := assert.New(t)

	avl := NewAvl()
	as.Nil(avl.root, "root must be nil when the avl is made")
	as.Equal(avl.length, 0, "avl length must be 0 when the avl is made")
}

func Test_Avl_Insert_func(t *testing.T) {
	avl := NewAvl()

	l := func(l int) {
		assert.Equal(t, avl.Length(), l, "inavlid length")
	}

	in := func(e int) {
		assert.True(t, avl.Insert(Iv(e)), "element duplicated")
	}

	// Insert first element. Element 5
	in(10)
	checkVAndH(t, avl.root, 10, 0)
	l(1)

	// Insert element 5
	in(5)
	checkVAndH(t, avl.root, 10, 1)
	checkVAndH(t, avl.root.ltree, 5, 0)
	l(2)

	// Insert element 3. Re-balance tree. Right rotate
	in(3)
	checkVAndH(t, avl.root, 5, 1)
	checkVAndH(t, avl.root.ltree, 3, 0)
	checkVAndH(t, avl.root.rtree, 10, 0)
	l(3)

	// Insert element 15
	in(15)
	checkVAndH(t, avl.root, 5, 2)
	checkVAndH(t, avl.root.ltree, 3, 0)
	checkVAndH(t, avl.root.rtree, 10, 1)
	checkVAndH(t, avl.root.rtree.rtree, 15, 0)
	l(4)

	// Insert element 7
	in(7)
	checkVAndH(t, avl.root, 5, 2)
	checkVAndH(t, avl.root.ltree, 3, 0)
	checkVAndH(t, avl.root.rtree, 10, 1)
	checkVAndH(t, avl.root.rtree.rtree, 15, 0)
	checkVAndH(t, avl.root.rtree.ltree, 7, 0)
	l(5)

	// Insert element 20. Re-balance tree. Left rotate
	in(20)
	checkVAndH(t, avl.root, 10, 2)
	checkVAndH(t, avl.root.ltree, 5, 1)
	checkVAndH(t, avl.root.ltree.ltree, 3, 0)
	checkVAndH(t, avl.root.ltree.rtree, 7, 0)

	checkVAndH(t, avl.root.rtree, 15, 1)
	checkVAndH(t, avl.root.rtree.rtree, 20, 0)
	l(6)

	// New tree to check the double rotations. Left right
	avl = NewAvl()
	in(10)
	in(5)
	in(7)
	checkVAndH(t, avl.root, 7, 1)
	checkVAndH(t, avl.root.ltree, 5, 0)
	checkVAndH(t, avl.root.rtree, 10, 0)
	l(3)

	// Insert 15 and 12. Generate double rotation right and left
	in(15)
	in(12)
	checkVAndH(t, avl.root, 7, 2)
	checkVAndH(t, avl.root.ltree, 5, 0)
	checkVAndH(t, avl.root.rtree, 12, 1)
	checkVAndH(t, avl.root.rtree.ltree, 10, 0)
	checkVAndH(t, avl.root.rtree.rtree, 15, 0)
	l(5)

	// Insert duplicate element
	assert.False(t, avl.Insert(Iv(5)))
	checkVAndH(t, avl.root, 7, 2)
	checkVAndH(t, avl.root.ltree, 5, 0)
	checkVAndH(t, avl.root.rtree, 12, 1)
	checkVAndH(t, avl.root.rtree.ltree, 10, 0)
	checkVAndH(t, avl.root.rtree.rtree, 15, 0)
	l(5)
}

func Test_search_func(t *testing.T) {
	as := assert.New(t)
	avl := NewAvl()

	v1 := Iv(1)
	v2 := Iv(2)
	v3 := Iv(3)
	v4 := Iv(4)

	avl.Insert(v1)
	avl.Insert(v2)
	avl.Insert(v3)
	avl.Insert(v4)

	node, found := search(2, avl.root)
	as.True(found)
	as.Equal(avl.root, node)

	node, found = search(4, avl.root)
	as.True(found)
	as.Equal(vi(node.value), 4)

	node, found = search(7, avl.root)
	as.False(found)
	as.Nil(node)
}

func Test_Avl_Search_func(t *testing.T) {
	as := assert.New(t)
	avl := NewAvl()

	values := map[int]Value{
		1:  Iv(1),
		2:  Iv(2),
		5:  Iv(5),
		12: Iv(12),
		8:  Iv(8),
		33: Iv(33),
	}

	for _, v := range values {
		avl.Insert(v)
	}

	for number, v := range values {
		result, found := avl.Search(Iv(number))
		as.Truef(found, "value %s not found", v.String())
		as.Truef(result.Eq(v), "Values doesn't match, key %d", number)
	}

	// Values hasn't in the tree
	invalidVals := []int{-1, 3, 4, 6, 9, 11, 13, 20, 30, 32, 34, 19322}
	for _, number := range invalidVals {
		result, found := avl.Search(Iv(number))
		as.False(found, "number  %d found in the tree", number)
		as.Nil(result, "result %d isn't nil", number)
	}
}

func Test_Avl_SearchKey_func(t *testing.T) {
	as := assert.New(t)
	avl := NewAvl()

	values := map[int]Value{
		1:  Iv(1),
		2:  Iv(2),
		5:  Iv(5),
		12: Iv(12),
		8:  Iv(8),
		33: Iv(33),
	}

	for _, v := range values {
		avl.Insert(v)
	}

	for number, v := range values {
		result, found := avl.SearchKey(number)
		as.Truef(found, "value %s not found", v.String())
		as.Truef(result.Eq(v), "Values doesn't match, key %d", number)
	}

	// Values hasn't in the tree
	invalidVals := []int{-1, 3, 4, 6, 9, 11, 13, 20, 30, 32, 34, 19322}
	for _, number := range invalidVals {
		result, found := avl.SearchKey(number)
		as.False(found, "number  %d found in the tree", number)
		as.Nil(result, "result %d isn't nil", number)
	}
}

func Test_Avl_Delete_func(t *testing.T) {
	as := assert.New(t)

	// Delete leaf
	avl, _ := createAvl(3, 2, 1, 4)
	value, found := avl.Delete(Iv(1))
	as.Equal(vi(value), 1)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 3, 1)
	checkTree(t, avl.root, i(2), i(4))

	checkVAndH(t, avl.root.ltree, 2, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 4, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with one child. Left child
	avl, _ = createAvl(4, 2, 1, 3)
	value, found = avl.Delete(Iv(4))
	as.Equal(vi(value), 4)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 2, 1)
	checkTree(t, avl.root, i(1), i(3))

	checkVAndH(t, avl.root.ltree, 1, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 3, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with one child. Right child
	avl, _ = createAvl(3, 2, 1, 4)
	value, found = avl.Delete(Iv(3))
	as.Equal(vi(value), 3)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 2, 1)
	checkTree(t, avl.root, i(1), i(4))

	checkVAndH(t, avl.root.ltree, 1, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 4, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with 2 children.
	avl, _ = createAvl(4, 2, 1, 3, 5, 6, 7)
	value, found = avl.Delete(Iv(4))
	as.Equal(vi(value), 4)
	as.True(found)
	as.Equal(avl.Length(), 6)

	checkVAndH(t, avl.root, 5, 2)
	checkTree(t, avl.root, i(2), i(6))

	checkVAndH(t, avl.root.ltree, 2, 1)
	checkTree(t, avl.root.ltree, i(1), i(3))

	checkVAndH(t, avl.root.ltree.ltree, 1, 0)
	checkTree(t, avl.root.ltree.ltree, nil, nil)

	checkVAndH(t, avl.root.ltree.rtree, 3, 0)
	checkTree(t, avl.root.ltree.rtree, nil, nil)

	checkVAndH(t, avl.root.rtree, 6, 1)
	checkTree(t, avl.root.rtree, nil, i(7))

	checkVAndH(t, avl.root.rtree.rtree, 7, 0)
	checkTree(t, avl.root.rtree.rtree, nil, nil)

	// Delete another node with 2 children.
	value, found = avl.Delete(Iv(2))
	as.Equal(vi(value), 2)
	as.True(found)
	as.Equal(avl.Length(), 5)

	checkVAndH(t, avl.root, 5, 2)
	checkTree(t, avl.root, i(3), i(6))

	checkVAndH(t, avl.root.ltree, 3, 1)
	checkTree(t, avl.root.ltree, i(1), nil)

	checkVAndH(t, avl.root.ltree.ltree, 1, 0)
	checkTree(t, avl.root.ltree.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 6, 1)
	checkTree(t, avl.root.rtree, nil, i(7))

	checkVAndH(t, avl.root.rtree.rtree, 7, 0)
	checkTree(t, avl.root.rtree.rtree, nil, nil)
}

func Test_Avl_DeleteKey_func(t *testing.T) {
	as := assert.New(t)

	// Delete leaf
	avl, _ := createAvl(3, 2, 1, 4)
	value, found := avl.DeleteKey(1)
	as.Equal(vi(value), 1)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 3, 1)
	checkTree(t, avl.root, i(2), i(4))

	checkVAndH(t, avl.root.ltree, 2, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 4, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with one child. Left child
	avl, _ = createAvl(4, 2, 1, 3)
	value, found = avl.DeleteKey(4)
	as.Equal(vi(value), 4)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 2, 1)
	checkTree(t, avl.root, i(1), i(3))

	checkVAndH(t, avl.root.ltree, 1, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 3, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with one child. Right child
	avl, _ = createAvl(3, 2, 1, 4)
	value, found = avl.DeleteKey(3)
	as.Equal(vi(value), 3)
	as.True(found)
	as.Equal(avl.Length(), 3)

	checkVAndH(t, avl.root, 2, 1)
	checkTree(t, avl.root, i(1), i(4))

	checkVAndH(t, avl.root.ltree, 1, 0)
	checkTree(t, avl.root.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 4, 0)
	checkTree(t, avl.root.rtree, nil, nil)

	// Delete node with 2 children.
	avl, _ = createAvl(4, 2, 1, 3, 5, 6, 7)
	value, found = avl.DeleteKey(4)
	as.Equal(vi(value), 4)
	as.True(found)
	as.Equal(avl.Length(), 6)

	checkVAndH(t, avl.root, 5, 2)
	checkTree(t, avl.root, i(2), i(6))

	checkVAndH(t, avl.root.ltree, 2, 1)
	checkTree(t, avl.root.ltree, i(1), i(3))

	checkVAndH(t, avl.root.ltree.ltree, 1, 0)
	checkTree(t, avl.root.ltree.ltree, nil, nil)

	checkVAndH(t, avl.root.ltree.rtree, 3, 0)
	checkTree(t, avl.root.ltree.rtree, nil, nil)

	checkVAndH(t, avl.root.rtree, 6, 1)
	checkTree(t, avl.root.rtree, nil, i(7))

	checkVAndH(t, avl.root.rtree.rtree, 7, 0)
	checkTree(t, avl.root.rtree.rtree, nil, nil)

	// Delete another node with 2 children.
	value, found = avl.DeleteKey(2)
	as.Equal(vi(value), 2)
	as.True(found)
	as.Equal(avl.Length(), 5)

	checkVAndH(t, avl.root, 5, 2)
	checkTree(t, avl.root, i(3), i(6))

	checkVAndH(t, avl.root.ltree, 3, 1)
	checkTree(t, avl.root.ltree, i(1), nil)

	checkVAndH(t, avl.root.ltree.ltree, 1, 0)
	checkTree(t, avl.root.ltree.ltree, nil, nil)

	checkVAndH(t, avl.root.rtree, 6, 1)
	checkTree(t, avl.root.rtree, nil, i(7))

	checkVAndH(t, avl.root.rtree.rtree, 7, 0)
	checkTree(t, avl.root.rtree.rtree, nil, nil)

}
