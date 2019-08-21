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

// getAllItems returns a map with all items inside of `node` avl node.
func getAllItems(node *avlNode) (items map[int]*Item) {
	var getItem func(node *avlNode)

	getItem = func(node *avlNode) {
		if node == nil {
			return
		}
		getItem(node.ltree)
		item := node.item
		it, _ := node.item.(IntItem)
		items[it.value] = &item
		getItem(node.rtree)
	}

	items = map[int]*Item{}
	if node != nil {
		getItem(node)
	}

	return
}
func createAvl(intItems ...int) (Avl, map[int]Item) {
	avl := NewAvl()

	items := map[int]Item{}

	for _, i := range intItems {
		items[i] = It(i)
		avl.Insert(items[i])
	}

	return avl, items
}

// checkVAndH checks the v value and the h height of the `node` avl node.
func checkVAndH(t *testing.T, node *avlNode, v, h int) {
	assert.Equal(t, node.item.(IntItem).value, v)
	assert.Equal(t, node.height, h)
}

// checkTree check the int values of the children of `node` avl node.
func checkTree(t *testing.T, node *avlNode, l, r *int) {

	if l == nil {
		assert.Nil(t, node.ltree)
	} else {
		assert.Equal(t, node.ltree.item.(IntItem).value, *l)
	}

	if r == nil {
		assert.Nil(t, node.rtree)
	} else {
		assert.Equal(t, node.rtree.item.(IntItem).value, *r)
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
	as.Equal(node.maxHeight(), -1)

	node = avlNode{&avlNode{nil, nil, 22, nil}, nil, 3, nil}
	as.Equal(node.maxHeight(), 22)

	node = avlNode{nil, &avlNode{nil, nil, 41, nil}, 3, nil}
	as.Equal(node.maxHeight(), 41)

	node = avlNode{&avlNode{nil, nil, 15, nil}, &avlNode{nil, nil, 30, nil}, 3, nil}
	as.Equal(node.maxHeight(), 30)
}

func Test_avlNode_rotateRight_func(t *testing.T) {
	tree1 := &avlNode{nil, nil, -1, It(1)}
	tree2 := &avlNode{nil, nil, -1, It(2)}
	tree3 := &avlNode{nil, nil, -1, It(3)}
	tree4 := &avlNode{nil, nil, -1, It(4)}
	tree5 := &avlNode{nil, nil, -1, It(5)}

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
	tree1 := &avlNode{nil, nil, -1, It(1)}
	tree2 := &avlNode{nil, nil, -1, It(2)}
	tree3 := &avlNode{nil, nil, -1, It(3)}
	tree4 := &avlNode{nil, nil, -1, It(4)}
	tree5 := &avlNode{nil, nil, -1, It(5)}

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
	tree1 := &avlNode{nil, nil, -1, It(1)}
	tree2 := &avlNode{nil, nil, -1, It(2)}
	tree3 := &avlNode{nil, nil, -1, It(3)}
	tree4 := &avlNode{nil, nil, -1, It(4)}
	tree5 := &avlNode{nil, nil, -1, It(5)}

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
	tree1 := &avlNode{nil, nil, -1, It(1)}
	tree2 := &avlNode{nil, nil, -1, It(2)}
	tree3 := &avlNode{nil, nil, -1, It(3)}
	tree4 := &avlNode{nil, nil, -1, It(4)}
	tree5 := &avlNode{nil, nil, -1, It(5)}

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
		assert.True(t, avl.Insert(It(e)), "element duplicated")
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
	assert.False(t, avl.Insert(It(5)))
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

	v1 := It(1)
	v2 := It(2)
	v3 := It(3)
	v4 := It(4)

	avl.Insert(v1)
	avl.Insert(v2)
	avl.Insert(v3)
	avl.Insert(v4)

	node, found := search(avl.root, It(2))
	as.True(found)
	as.Equal(avl.root, node)

	node, found = search(avl.root, It(4))
	as.True(found)
	as.Equal(node.item.(IntItem).value, 4)

	node, found = search(avl.root, It(7))
	as.False(found)
	as.Nil(node)
}

func Test_Avl_Search_func(t *testing.T) {
	as := assert.New(t)
	avl := NewAvl()

	items := map[int]Item{
		1:  It(1),
		2:  It(2),
		5:  It(5),
		12: It(12),
		8:  It(8),
		33: It(33),
	}

	for _, v := range items {
		avl.Insert(v)
	}

	for number, v := range items {
		result, found := avl.Search(It(number))
		as.Truef(found, "item %s not found", v.String())
		as.Truef(result.Eq(v), "items doesn't match, Item %s", result.String())
	}

	// Values hasn't in the tree
	invalidVals := []int{-1, 3, 4, 6, 9, 11, 13, 20, 30, 32, 34, 19322}
	for _, number := range invalidVals {
		result, found := avl.Search(It(number))
		as.False(found, "number  %d found in the tree", number)
		as.Nil(result, "result %d isn't nil", number)
	}
}

func Test_Avl_Delete_func(t *testing.T) {
	as := assert.New(t)

	// Delete leaf
	avl, _ := createAvl(3, 2, 1, 4)
	item, found := avl.Delete(It(1))
	as.Equal(item.(IntItem).value, 1)
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
	item, found = avl.Delete(It(4))
	as.Equal(item.(IntItem).value, 4)
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
	item, found = avl.Delete(It(3))
	as.Equal(item.(IntItem).value, 3)
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
	item, found = avl.Delete(It(4))
	as.Equal(item.(IntItem).value, 4)
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
	item, found = avl.Delete(It(2))
	as.Equal(item.(IntItem).value, 2)
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
