package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// i returns the address of the i param
func i(i int) *int {
	return &i
}

// getAllItems returns a map with all items inside of `node` tree node.
func getAllItems(node *treeNode) (items map[int]*Item) {
	var getItem func(node *treeNode)

	getItem = func(node *treeNode) {
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

func createTreeAvl(intItems ...int) (Tree, map[int]Item) {
	tree := Tree{rebalance: true}

	items := map[int]Item{}

	for _, i := range intItems {
		items[i] = It(i)
		tree.Insert(items[i])
	}

	return tree, items
}

func createTreeBst(intItems ...int) (Tree, map[int]Item) {
	tree := Tree{rebalance: false}

	items := map[int]Item{}

	for _, i := range intItems {
		items[i] = It(i)
		tree.Insert(items[i])
	}

	return tree, items
}

// checkVAndH checks the v value and the h height of the `node` tree node.
func checkVAndH(t *testing.T, node *treeNode, v, h int) {
	assert.Equal(t, node.item.(IntItem).value, v)
	assert.Equal(t, node.height, h)
}

// checkTree check the int values of the children of `node` tree node.
func checkTree(t *testing.T, node *treeNode, l, r *int) {

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

func Test_treeNode_getHeight_func(t *testing.T) {
	as := assert.New(t)

	as.Equal((*treeNode)(nil).getHeight(), -1, "When node is nil, must returns -1")
	as.Equal((&treeNode{nil, nil, 33, nil}).getHeight(), 33, "Return invalid height")
}

func Test_treeNode_maxHeight_func(t *testing.T) {
	as := assert.New(t)

	node := treeNode{nil, nil, 3, nil}
	as.Equal(node.maxHeight(), -1)

	node = treeNode{&treeNode{nil, nil, 22, nil}, nil, 3, nil}
	as.Equal(node.maxHeight(), 22)

	node = treeNode{nil, &treeNode{nil, nil, 41, nil}, 3, nil}
	as.Equal(node.maxHeight(), 41)

	node = treeNode{&treeNode{nil, nil, 15, nil}, &treeNode{nil, nil, 30, nil}, 3, nil}
	as.Equal(node.maxHeight(), 30)
}

func Test_treeNode_rotateRight_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, -1, It(1)}
	tree2 := &treeNode{nil, nil, -1, It(2)}
	tree3 := &treeNode{nil, nil, -1, It(3)}
	tree4 := &treeNode{nil, nil, -1, It(4)}
	tree5 := &treeNode{nil, nil, -1, It(5)}

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

func Test_treeNode_rotateLeft_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, -1, It(1)}
	tree2 := &treeNode{nil, nil, -1, It(2)}
	tree3 := &treeNode{nil, nil, -1, It(3)}
	tree4 := &treeNode{nil, nil, -1, It(4)}
	tree5 := &treeNode{nil, nil, -1, It(5)}

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

func Test_treeNode_rotateRightLeft_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, -1, It(1)}
	tree2 := &treeNode{nil, nil, -1, It(2)}
	tree3 := &treeNode{nil, nil, -1, It(3)}
	tree4 := &treeNode{nil, nil, -1, It(4)}
	tree5 := &treeNode{nil, nil, -1, It(5)}

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

func Test_treeNode_rotateLeftRight_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, -1, It(1)}
	tree2 := &treeNode{nil, nil, -1, It(2)}
	tree3 := &treeNode{nil, nil, -1, It(3)}
	tree4 := &treeNode{nil, nil, -1, It(4)}
	tree5 := &treeNode{nil, nil, -1, It(5)}

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

func Test_Tree_Insert_func(t *testing.T) {
	tr := Tree{rebalance: true}

	l := func(l int) {
		assert.Equal(t, tr.Length(), l)
	}

	in := func(e int) {
		assert.True(t, tr.Insert(It(e)))
	}

	// Insert first item. Item 5
	in(10)
	checkVAndH(t, tr.root, 10, 0)
	l(1)

	// Insert item 5
	in(5)
	checkVAndH(t, tr.root, 10, 1)
	checkVAndH(t, tr.root.ltree, 5, 0)
	l(2)

	// Insert item 3. Re-balance tree. Right rotate
	in(3)
	checkVAndH(t, tr.root, 5, 1)
	checkVAndH(t, tr.root.ltree, 3, 0)
	checkVAndH(t, tr.root.rtree, 10, 0)
	l(3)

	// Insert item 15
	in(15)
	checkVAndH(t, tr.root, 5, 2)
	checkVAndH(t, tr.root.ltree, 3, 0)
	checkVAndH(t, tr.root.rtree, 10, 1)
	checkVAndH(t, tr.root.rtree.rtree, 15, 0)
	l(4)

	// Insert item 7
	in(7)
	checkVAndH(t, tr.root, 5, 2)
	checkVAndH(t, tr.root.ltree, 3, 0)
	checkVAndH(t, tr.root.rtree, 10, 1)
	checkVAndH(t, tr.root.rtree.rtree, 15, 0)
	checkVAndH(t, tr.root.rtree.ltree, 7, 0)
	l(5)

	// Insert item 20. Re-balance tree. Left rotate
	in(20)
	checkVAndH(t, tr.root, 10, 2)
	checkVAndH(t, tr.root.ltree, 5, 1)
	checkVAndH(t, tr.root.ltree.ltree, 3, 0)
	checkVAndH(t, tr.root.ltree.rtree, 7, 0)

	checkVAndH(t, tr.root.rtree, 15, 1)
	checkVAndH(t, tr.root.rtree.rtree, 20, 0)
	l(6)

	// New tree to check the double rotations. Left right
	tr = Tree{rebalance: true}
	in(10)
	in(5)
	in(7)
	checkVAndH(t, tr.root, 7, 1)
	checkVAndH(t, tr.root.ltree, 5, 0)
	checkVAndH(t, tr.root.rtree, 10, 0)
	l(3)

	// Insert 15 and 12. Generate double rotation right and left
	in(15)
	in(12)
	checkVAndH(t, tr.root, 7, 2)
	checkVAndH(t, tr.root.ltree, 5, 0)
	checkVAndH(t, tr.root.rtree, 12, 1)
	checkVAndH(t, tr.root.rtree.ltree, 10, 0)
	checkVAndH(t, tr.root.rtree.rtree, 15, 0)
	l(5)

	// Insert duplicate item
	assert.False(t, tr.Insert(It(5)))
	checkVAndH(t, tr.root, 7, 2)
	checkVAndH(t, tr.root.ltree, 5, 0)
	checkVAndH(t, tr.root.rtree, 12, 1)
	checkVAndH(t, tr.root.rtree.ltree, 10, 0)
	checkVAndH(t, tr.root.rtree.rtree, 15, 0)
	l(5)

	//insert without rebalancer
	tr, _ = createTreeBst(1, 2, 3, 4, 5, 6)
	for trNode, i := tr.root, 1; i <= tr.length; trNode, i = trNode.rtree, i+1 {
		assert.Equal(t, trNode.item.(IntItem).value, i)
		assert.Nil(t, trNode.ltree)

		if i == 6 {
			assert.Nil(t, trNode.rtree)
		}
	}

	// allows duplicate items.
	tr = Tree{rebalance: true, duplicated: true}

	for _, i := range []int{1, 2, 3, 1} {
		inserted := tr.Insert(It(i))
		assert.True(t, inserted)
	}

	checkVAndH(t, tr.root, 2, 2)
	checkTree(t, tr.root, i(1), i(3))

	checkVAndH(t, tr.root.ltree, 1, 1)
	checkTree(t, tr.root.ltree, nil, i(1))

	checkVAndH(t, tr.root.ltree.rtree, 1, 0)
	checkTree(t, tr.root.ltree.rtree, nil, nil)

	checkVAndH(t, tr.root.rtree, 3, 0)
	checkTree(t, tr.root.rtree, nil, nil)
}

func Test_search_func(t *testing.T) {
	as := assert.New(t)
	tr := Tree{rebalance: true}

	v1 := It(1)
	v2 := It(2)
	v3 := It(3)
	v4 := It(4)

	tr.Insert(v1)
	tr.Insert(v2)
	tr.Insert(v3)
	tr.Insert(v4)

	node, found := search(tr.root, It(2))
	as.True(found)
	as.Equal(tr.root, node)

	node, found = search(tr.root, It(4))
	as.True(found)
	as.Equal(node.item.(IntItem).value, 4)

	node, found = search(tr.root, It(7))
	as.False(found)
	as.Nil(node)
}

func Test_Tree_Search_func(t *testing.T) {
	as := assert.New(t)
	tr := Tree{rebalance: true}

	items := map[int]Item{
		1:  It(1),
		2:  It(2),
		5:  It(5),
		12: It(12),
		8:  It(8),
		33: It(33),
	}

	for _, v := range items {
		tr.Insert(v)
	}

	for number, v := range items {
		result, found := tr.Search(It(number))
		as.Truef(found, "item %s not found", v.String())
		as.Truef(result.Eq(v), "items doesn't match, Item %s", result.String())
	}

	// Values hasn't in the tree
	invalidVals := []int{-1, 3, 4, 6, 9, 11, 13, 20, 30, 32, 34, 19322}
	for _, number := range invalidVals {
		result, found := tr.Search(It(number))
		as.False(found, "number  %d found in the tree", number)
		as.Nil(result, "result %d isn't nil", number)
	}
}

func Test_Tree_Delete_func(t *testing.T) {
	as := assert.New(t)

	// Delete leaf
	tr, _ := createTreeAvl(3, 2, 1, 4)
	item, found := tr.Delete(It(1))
	as.Equal(item.(IntItem).value, 1)
	as.True(found)
	as.Equal(tr.Length(), 3)

	checkVAndH(t, tr.root, 3, 1)
	checkTree(t, tr.root, i(2), i(4))

	checkVAndH(t, tr.root.ltree, 2, 0)
	checkTree(t, tr.root.ltree, nil, nil)

	checkVAndH(t, tr.root.rtree, 4, 0)
	checkTree(t, tr.root.rtree, nil, nil)

	// Delete node with one child. Left child
	tr, _ = createTreeAvl(4, 2, 1, 3)
	item, found = tr.Delete(It(4))
	as.Equal(item.(IntItem).value, 4)
	as.True(found)
	as.Equal(tr.Length(), 3)

	checkVAndH(t, tr.root, 2, 1)
	checkTree(t, tr.root, i(1), i(3))

	checkVAndH(t, tr.root.ltree, 1, 0)
	checkTree(t, tr.root.ltree, nil, nil)

	checkVAndH(t, tr.root.rtree, 3, 0)
	checkTree(t, tr.root.rtree, nil, nil)

	// Delete node with one child. Right child
	tr, _ = createTreeAvl(3, 2, 1, 4)
	item, found = tr.Delete(It(3))
	as.Equal(item.(IntItem).value, 3)
	as.True(found)
	as.Equal(tr.Length(), 3)

	checkVAndH(t, tr.root, 2, 1)
	checkTree(t, tr.root, i(1), i(4))

	checkVAndH(t, tr.root.ltree, 1, 0)
	checkTree(t, tr.root.ltree, nil, nil)

	checkVAndH(t, tr.root.rtree, 4, 0)
	checkTree(t, tr.root.rtree, nil, nil)

	// Delete node with 2 children.
	tr, _ = createTreeAvl(4, 2, 1, 3, 5, 6, 7)
	item, found = tr.Delete(It(4))
	as.Equal(item.(IntItem).value, 4)
	as.True(found)
	as.Equal(tr.Length(), 6)

	checkVAndH(t, tr.root, 5, 2)
	checkTree(t, tr.root, i(2), i(6))

	checkVAndH(t, tr.root.ltree, 2, 1)
	checkTree(t, tr.root.ltree, i(1), i(3))

	checkVAndH(t, tr.root.ltree.ltree, 1, 0)
	checkTree(t, tr.root.ltree.ltree, nil, nil)

	checkVAndH(t, tr.root.ltree.rtree, 3, 0)
	checkTree(t, tr.root.ltree.rtree, nil, nil)

	checkVAndH(t, tr.root.rtree, 6, 1)
	checkTree(t, tr.root.rtree, nil, i(7))

	checkVAndH(t, tr.root.rtree.rtree, 7, 0)
	checkTree(t, tr.root.rtree.rtree, nil, nil)

	// Delete another node with 2 children.
	item, found = tr.Delete(It(2))
	as.Equal(item.(IntItem).value, 2)
	as.True(found)
	as.Equal(tr.Length(), 5)

	checkVAndH(t, tr.root, 5, 2)
	checkTree(t, tr.root, i(3), i(6))

	checkVAndH(t, tr.root.ltree, 3, 1)
	checkTree(t, tr.root.ltree, i(1), nil)

	checkVAndH(t, tr.root.ltree.ltree, 1, 0)
	checkTree(t, tr.root.ltree.ltree, nil, nil)

	checkVAndH(t, tr.root.rtree, 6, 1)
	checkTree(t, tr.root.rtree, nil, i(7))

	checkVAndH(t, tr.root.rtree.rtree, 7, 0)
	checkTree(t, tr.root.rtree.rtree, nil, nil)

	// try delete an node it isn't in the tree
	item, found = tr.Delete(It(-1))
	as.Nil(item)
	as.False(found)

	// delete node and without rebalance.
	tr, _ = createTreeBst(1, 2, 3, 4, 5, 6)
	for i := 1; i <= 6; i++ {
		tr.Delete(It(i))

		for trNode, j := tr.root, i+1; j <= 6; trNode, j = trNode.rtree, j+1 {
			assert.Equal(t, trNode.item.(IntItem).value, j)
			assert.Nil(t, trNode.ltree)

			if i == 6 {
				assert.Nil(t, trNode.rtree)
			}
		}
	}
}

func Test_Tree_Clear_func(t *testing.T) {
	params := []struct {
		rebalance, duplicated bool
	}{
		{false, false},
		{false, true},
		{true, false},
		{true, true},
	}

	for _, param := range params {
		tr := Tree{rebalance: param.rebalance, duplicated: param.duplicated}

		for i := 1; i <= 10; i++ {
			tr.Insert(It(i))
		}

		assert.NotNil(t, tr.root)
		assert.Equal(t, tr.length, 10)
		assert.Equal(t, tr.rebalance, param.rebalance)
		assert.Equal(t, tr.duplicated, param.duplicated)

		// Clear
		tr.Clear()
		assert.Nil(t, tr.root)
		assert.Equal(t, tr.length, 0)
		assert.Equal(t, tr.rebalance, param.rebalance)
		assert.Equal(t, tr.duplicated, param.duplicated)
	}
}
