package mygostructs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// i returns the address of the i param
func i(i int) *int {
	return &i
}

// checkVAndH checks the v value and the h height of the `node` tree node.
func checkVAndH(t *testing.T, node *treeNode, v, h int) {
	assert.Equal(t, node.item.(IntItem).value, v)
	assert.Equal(t, node.height, h)
}

// checkTreeNode check the int values of the children of `node` tree node.
func checkTreeNode(t *testing.T, node *treeNode, l, r *int) {

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

// struct for test if a tree is well formed. This struct is used in the function check tree
type treeTest struct {
	value  int  // node value
	height int  // node height
	lvalue *int // value left child. If it is nil then the left child doesn't exist
	rvalue *int // value of the right child. If it is nil then the right child doesn't exist
}

// checkTree checks if the root matchs with the results of the array.
func checkTree(t *testing.T, root *treeNode, results []treeTest) {
	var checkNode func(*treeNode)

	i := 0
	as := assert.New(t)

	checkNode = func(node *treeNode) {
		if node == nil {
			return
		}

		checkNode(node.ltree)

		// Check number of results
		as.GreaterOrEqual(len(results), i, "number of nodes is great than results slice")

		msg := fmt.Sprintf("in node %s", node.item.String())

		// check value and height
		result := results[i]
		value := node.item.(IntItem).value
		as.Equal(value, result.value, "%s: the value doesn't match", msg)
		as.Equal(node.height, result.height, "%s: the height doesn't match", msg)

		// check left node
		if result.lvalue == nil {
			if node.ltree != nil {
				valueStr := node.ltree.item.String()
				as.Fail(
					"left node isn't nil",
					"%s: left node isn't nil, value: %s",
					msg,
					valueStr,
				)
			}
		} else {
			if node.ltree == nil {
				as.Fail(
					"left node is nil",
					"%s: left node is nil, value expected: %d",
					msg,
					*result.lvalue,
				)
			} else {
				value := node.ltree.item.(IntItem).value
				as.Equal(
					value,
					*result.lvalue,
					"%s: value of the left node doesn't match",
					msg,
				)
			}
		}

		// check right node
		if result.rvalue == nil {
			if node.rtree != nil {
				valueStr := node.rtree.item.String()
				as.Fail(
					"right node isn't nil",
					"%s: right node isn't nil, value: %s",
					msg,
					valueStr,
				)
			}
		} else {
			if node.rtree == nil {
				as.Fail(
					"right node is nil",
					"%s: right node is nil, value expected: %d",
					msg,
					*result.rvalue,
				)
			} else {
				value := node.rtree.item.(IntItem).value
				as.Equal(
					value,
					*result.rvalue,
					"%s: value of the right node doesn't match",
					msg,
				)
			}
		}
		i++

		checkNode(node.rtree)
	}

	checkNode(root)

	// check the length
	as.Equal(i, len(results), "number of nodes visited doesn't match with result length")
}

func treeChangeProp(tree *Tree, size int, done chan bool) {
	for i := 0; i < size; i++ {
		tree.mutex.Lock()
		root, length := tree.root, tree.length
		tree.root, tree.length = nil, -1
		time.Sleep(time.Nanosecond)
		tree.root, tree.length = root, length
		tree.mutex.Unlock()
	}

	done <- true
}

//
// Start tests
// ===========
//
func Test_max_func(t *testing.T) {
	as := assert.New(t)
	as.Equal(max(2, 3), 3, "no returns the highest value")
	as.Equal(max(3, 3), 3, "no returns the highest value")
	as.Equal(max(3, 2), 3, "no returns the highest value")
}

func Test_treeNode_getHeight_func(t *testing.T) {
	as := assert.New(t)

	as.Equal((*treeNode)(nil).getHeight(), -1, "when node is nil, must returns -1")
	as.Equal((&treeNode{nil, nil, 1, nil}).getHeight(), 1, "node height doesn't match")
}

func Test_treeNode_maxHeight_func(t *testing.T) {
	as := assert.New(t)

	node := treeNode{nil, nil, 3, nil}
	as.Equal(node.maxHeight(), -1, "children are nil, must returns -1")

	node = treeNode{&treeNode{nil, nil, 10, nil}, nil, 3, nil}
	as.Equal(node.maxHeight(), 10, "value doesn't match with the height of left child")

	node = treeNode{nil, &treeNode{nil, nil, 10, nil}, 3, nil}
	as.Equal(node.maxHeight(), 10, "value doesn't match with the height of right child")

	node = treeNode{&treeNode{nil, nil, 10, nil}, &treeNode{nil, nil, 11, nil}, 3, nil}
	as.Equal(
		node.maxHeight(),
		11,
		"must be 11. Height of left child is 10, height of right child is 11",
	)
}

func Test_treeNode_rotateRight_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, 0, It(1)}
	tree2 := &treeNode{nil, nil, 0, It(2)}
	tree3 := &treeNode{nil, nil, 0, It(3)}
	tree4 := &treeNode{nil, nil, 0, It(4)}
	tree5 := &treeNode{nil, nil, 0, It(5)}

	tree2.ltree = tree1
	tree2.rtree = tree3
	tree4.ltree = tree2
	tree4.rtree = tree5
	/*
			Tree created:
			        4
			      2   5
		             1 3

	*/

	ntree := tree4.rotateRight()
	/*
		Tree after right rotation:
			2
		      1   4
			 3 5
	*/

	checkTree(t, ntree, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 1, i(3), i(5)},
		{5, 0, nil, nil},
	})
}

func Test_treeNode_rotateLeft_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, 0, It(1)}
	tree2 := &treeNode{nil, nil, 0, It(2)}
	tree3 := &treeNode{nil, nil, 0, It(3)}
	tree4 := &treeNode{nil, nil, 0, It(4)}
	tree5 := &treeNode{nil, nil, 0, It(5)}

	tree4.ltree = tree3
	tree4.rtree = tree5
	tree2.ltree = tree1
	tree2.rtree = tree4
	/*
		Test created:
			2
		      1   4
			 3 5
	*/

	ntree := tree2.rotateLeft()
	/*
			Test after left rotation:
			        4
			      2   5
		             1 3
	*/
	checkTree(t, ntree, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), i(3)},
		{3, 0, nil, nil},
		{4, 1, i(2), i(5)},
		{5, 0, nil, nil},
	})
}

func Test_treeNode_rotateRightLeft_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, 0, It(1)}
	tree2 := &treeNode{nil, nil, 0, It(2)}
	tree3 := &treeNode{nil, nil, 0, It(3)}
	tree4 := &treeNode{nil, nil, 0, It(4)}
	tree5 := &treeNode{nil, nil, 0, It(5)}

	tree4.ltree = tree3
	tree4.rtree = tree5
	tree2.ltree = tree1
	tree2.rtree = tree4
	/*
		Tree created:
			2
		      1   4
		         3 5
	*/

	ntree := tree2.rotateRightLeft()
	/*
		Tree after right left rotation:
		        3
		      2   4
		     1     5
	*/
	checkTree(t, ntree, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), nil},
		{3, 2, i(2), i(4)},
		{4, 1, nil, i(5)},
		{5, 0, nil, nil},
	})
}

func Test_treeNode_rotateLeftRight_func(t *testing.T) {
	tree1 := &treeNode{nil, nil, 0, It(1)}
	tree2 := &treeNode{nil, nil, 0, It(2)}
	tree3 := &treeNode{nil, nil, 0, It(3)}
	tree4 := &treeNode{nil, nil, 0, It(4)}
	tree5 := &treeNode{nil, nil, 0, It(5)}

	tree2.ltree = tree1
	tree2.rtree = tree3
	tree4.ltree = tree2
	tree4.rtree = tree5
	/*
		Tree created
		        4
		      2	  5
		     1 3
	*/

	ntree := tree4.rotateLeftRight()
	/*
		Tree after left right rotation
		        3
		      2   4
		     1     5
	*/
	checkTree(t, ntree, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), nil},
		{3, 2, i(2), i(4)},
		{4, 1, nil, i(5)},
		{5, 0, nil, nil},
	})
}

func Test_insertGetAdy_func(t *testing.T) {
	var (
		root      *treeNode
		prev      *Item
		inserted  bool
		items     []int
		prevItems []*int
	)

	as := assert.New(t)

	/*
		Test function:
			- rebalance
			- no duplicated items
	*/
	items = []int{3, 4, 2, 9, 7, 6, 5, 1, 0, 8}
	prevItems = []*int{nil, i(3), nil, i(4), i(4), i(4), i(4), nil, nil, i(7)}
	for indx, num := range items {
		item := It(num)
		root, prev, inserted = insertGetAdy(root, item, true, false)
		as.True(inserted, "item %s wasn't inserted", item)

		if prevItems[indx] == nil {
			as.Nil(prev, "item %s is min, but previous item isn't nil", item)
			continue
		}

		if prev == nil {
			as.Fail(
				"prev item is nil",
				"item %s isn't min but previous item is nil",
				item,
			)
			continue
		}

		as.Equal(
			(*prev).(IntItem).value,
			*prevItems[indx],
			"previous item doesn't match, item inserted is %s",
			item,
		)
	}

	// check the tree created
	checkTree(t, root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), nil},
		{2, 2, i(1), i(3)},
		{3, 0, nil, nil},
		{4, 3, i(2), i(7)},
		{5, 0, nil, nil},
		{6, 1, i(5), nil},
		{7, 2, i(6), i(9)},
		{8, 0, nil, nil},
		{9, 1, i(8), nil},
	})

	// insert item duplicated.
	root, prev, inserted = insertGetAdy(root, It(1), true, false)
	as.False(inserted, "duplicated item was inserted")
	as.Nil(prev, "previous item isn't nil when item wasn't inserted")

	// check the root didn't change.
	checkTree(t, root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), nil},
		{2, 2, i(1), i(3)},
		{3, 0, nil, nil},
		{4, 3, i(2), i(7)},
		{5, 0, nil, nil},
		{6, 1, i(5), nil},
		{7, 2, i(6), i(9)},
		{8, 0, nil, nil},
		{9, 1, i(8), nil},
	})
	/*
		Test function:
			- no rebalance
			- no duplicated items
	*/
	root = nil
	items = []int{4, 8, 3, 2, 7, 6, 9, 1, 0, 5}
	prevItems = []*int{nil, i(4), nil, nil, i(4), i(4), i(8), nil, nil, i(4)}
	for indx, num := range items {
		item := It(num)
		root, prev, inserted = insertGetAdy(root, item, false, false)
		as.True(inserted, "the item %s wasn't inserted", item)

		if prevItems[indx] == nil {
			as.Nil(prev, "item %s is min, but previous item isn't nil", item)
			continue
		}

		if prev == nil {
			as.Fail(
				"prev item is nil",
				"item %s isn't min but previous item is nil",
				item,
			)
			continue
		}

		as.Equal(
			(*prev).(IntItem).value,
			*prevItems[indx],
			"previous item doesn't match, item inserted is %s",
			item,
		)
	}

	checkTree(t, root, []treeTest{
		{0, 0, nil, nil},
		{1, 0, i(0), nil},
		{2, 0, i(1), nil},
		{3, 0, i(2), nil},
		{4, 0, i(3), i(8)},
		{5, 0, nil, nil},
		{6, 0, i(5), nil},
		{7, 0, i(6), nil},
		{8, 0, i(7), i(9)},
		{9, 0, nil, nil},
	})

	// insert item duplicated.
	root, prev, inserted = insertGetAdy(root, It(1), false, false)
	as.False(inserted, "duplicated item was inserted")
	as.Nil(prev, "previous item isn't nil when item wasn't inserted")

	/*
		Test function:
			- rebalance
			- duplicated items
	*/
	root = nil
	items = []int{4, 3, 3, 2, 1, 0, 4, 5, 1, 2}
	prevItems = []*int{nil, nil, i(3), nil, nil, nil, i(4), i(4), i(1), i(2)}
	for indx, num := range items {
		item := It(num)
		root, prev, inserted = insertGetAdy(root, item, true, true)
		as.True(inserted, "the item %s wasn't inserted", item)

		if prevItems[indx] == nil {
			as.Nil(prev, "item %s is min, but previous item isn't nil", item)
			continue
		}

		if prev == nil {
			as.Fail(
				"prev item is nil",
				"item %s isn't min but previous item is nil",
				item,
			)
			continue
		}

		as.Equal(
			(*prev).(IntItem).value,
			*prevItems[indx],
			"previous item doesn't match, item inserted is %s",
			item,
		)
	}

	/*
		       2
		   1       3
		 0   1   3   4
		- - - - 3 - 4 5
	*/
	checkTree(t, root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(1)},
		{1, 0, nil, nil},
		{2, 3, i(1), i(3)},
		{2, 0, nil, nil},
		{3, 1, i(2), nil},
		{3, 2, i(3), i(4)},
		{4, 0, nil, nil},
		{4, 1, i(4), i(5)},
		{5, 0, nil, nil},
	})

	/*
		Test function:
			- mo rebalance
			- duplicated items
	*/
	root = nil
	items = []int{4, 3, 3, 2, 1, 0, 4, 5, 1, 2}
	prevItems = []*int{nil, nil, i(3), nil, nil, nil, i(4), i(4), i(1), i(2)}
	for indx, num := range items {
		item := It(num)
		root, prev, inserted = insertGetAdy(root, item, false, true)
		as.True(inserted, "the item %s wasn't inserted", item)

		if prevItems[indx] == nil {
			as.Nil(prev, "item %s is min, but previous item isn't nil", item)
			continue
		}

		if prev == nil {
			as.Fail(
				"prev item is nil",
				"item %s isn't min but previous item is nil",
				item,
			)
			continue
		}

		as.Equal(
			(*prev).(IntItem).value,
			*prevItems[indx],
			"previous item doesn't match, item inserted is %s",
			item,
		)
	}

	/*
		               4
		       3               4
		   2       3       -       5
		 1   2   -   -   -   -   -   -
		0 1 - - - - - - - - - - - - - -
	*/
	checkTree(t, root, []treeTest{
		{0, 0, nil, nil},
		{1, 0, i(0), i(1)},
		{1, 0, nil, nil},
		{2, 0, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 0, i(2), i(3)},
		{3, 0, nil, nil},
		{4, 0, i(3), i(4)},
		{4, 0, nil, i(5)},
		{5, 0, nil, nil},
	})
}

func Test_Tree_Insert_func(t *testing.T) {
	var (
		tree     Tree
		inserted bool
		items    []int
	)

	as := assert.New(t)

	/*
		Test function:
			- rebalance
			- no duplicated items
	*/
	tree = Tree{rebalance: true, duplicated: false}
	items = []int{3, 4, 2, 9, 7, 6, 5, 1, 0, 8}
	for _, num := range items {
		item := It(num)
		inserted = tree.Insert(item)
		as.True(inserted, "item %s wasn't inserted", item)
	}

	// check the tree created
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), nil},
		{2, 2, i(1), i(3)},
		{3, 0, nil, nil},
		{4, 3, i(2), i(7)},
		{5, 0, nil, nil},
		{6, 1, i(5), nil},
		{7, 2, i(6), i(9)},
		{8, 0, nil, nil},
		{9, 1, i(8), nil},
	})

	// insert item duplicated.
	inserted = tree.Insert(It(0))
	as.False(inserted, "duplicated item was inserted")

	// check the root didn't change.
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), nil},
		{2, 2, i(1), i(3)},
		{3, 0, nil, nil},
		{4, 3, i(2), i(7)},
		{5, 0, nil, nil},
		{6, 1, i(5), nil},
		{7, 2, i(6), i(9)},
		{8, 0, nil, nil},
		{9, 1, i(8), nil},
	})
	/*
		Test function:
			- no rebalance
			- no duplicated items
	*/
	tree = Tree{rebalance: false, duplicated: false}
	items = []int{4, 8, 3, 2, 7, 6, 9, 1, 0, 5}
	for _, num := range items {
		item := It(num)
		inserted = tree.Insert(item)
		as.True(inserted, "item %s wasn't inserted", item)
	}

	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 0, i(0), nil},
		{2, 0, i(1), nil},
		{3, 0, i(2), nil},
		{4, 0, i(3), i(8)},
		{5, 0, nil, nil},
		{6, 0, i(5), nil},
		{7, 0, i(6), nil},
		{8, 0, i(7), i(9)},
		{9, 0, nil, nil},
	})

	// insert item duplicated.
	inserted = tree.Insert(It(0))
	as.False(inserted, "duplicated item was inserted")

	/*
		Test function:
			- rebalance
			- duplicated items
	*/
	tree = Tree{rebalance: true, duplicated: true}
	items = []int{4, 3, 3, 2, 1, 0, 4, 5, 1, 2}
	for _, num := range items {
		item := It(num)
		inserted = tree.Insert(item)
		as.True(inserted, "item %s wasn't inserted", item)
	}

	/*
		       2
		   1       3
		 0   1   3   4
		- - - - 3 - 4 5
	*/
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(1)},
		{1, 0, nil, nil},
		{2, 3, i(1), i(3)},
		{2, 0, nil, nil},
		{3, 1, i(2), nil},
		{3, 2, i(3), i(4)},
		{4, 0, nil, nil},
		{4, 1, i(4), i(5)},
		{5, 0, nil, nil},
	})

	/*
		Test function:
			- mo rebalance
			- duplicated items
	*/
	tree = Tree{rebalance: false, duplicated: true}
	items = []int{4, 3, 3, 2, 1, 0, 4, 5, 1, 2}
	for _, num := range items {
		item := It(num)
		inserted = tree.Insert(item)
		as.True(inserted, "item %s wasn't inserted", item)
	}

	/*
		               4
		       3               4
		   2       3       -       5
		 1   2   -   -   -   -   -   -
		0 1 - - - - - - - - - - - - - -
	*/
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 0, i(0), i(1)},
		{1, 0, nil, nil},
		{2, 0, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 0, i(2), i(3)},
		{3, 0, nil, nil},
		{4, 0, i(3), i(4)},
		{4, 0, nil, i(5)},
		{5, 0, nil, nil},
	})

}

func Test_Tree_Insert_func_sync(t *testing.T) {
	as := assert.New(t)
	tree := Tree{rebalance: true}
	done := make(chan bool)
	concurrence := 8
	size := 2000

	insert := func(min, max int) {
		for i := min; i < max; i++ {
			item := It(i)
			as.Truef(tree.Insert(item), "item %s is duplicated", item)
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go insert(i*size, (i+1)*size)
		go treeChangeProp(&tree, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}

	for i := 0; i < size*concurrence; i++ {
		item, found := tree.Search(It(i))
		as.Equal(item.(IntItem).value, i, "item found is incorrect")
		as.True(found, "item not found")
	}

	as.Equal(tree.Length(), concurrence*size, "tree length doesn't match")
}

func Test_Tree_Length_func(t *testing.T) {
	as := assert.New(t)
	tree := Tree{rebalance: true}

	// inserting items
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
		as.Equalf(tree.Length(), i, "tree length doesn't match")
	}

	// inserting duplicated items
	tree = Tree{rebalance: true}
	for i := 1; i <= 10; i++ {
		tree.Insert(It(1))
		assert.Equalf(t, tree.Length(), 1, "tree length doesn't match")
	}

	// tree that allow duplicated items
	tree = Tree{rebalance: true, duplicated: true}
	for i := 1; i <= 10; i++ {
		tree.Insert(It(1))
		assert.Equalf(t, tree.Length(), i, "tree length doesn't match")
	}
}

func Test_Tree_Length_func_sync(t *testing.T) {
	done := make(chan bool)
	concurrence := 10
	size := 1000
	tree := Tree{}

	lengthFunc := func() {
		for i := 0; i < size; i++ {
			assert.Equalf(t, tree.Length(), size, "tree length doesn't match")
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		tree.Insert(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go lengthFunc()
		go treeChangeProp(&tree, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}
}

func Test_Tree_Search_func(t *testing.T) {
	as := assert.New(t)
	tr := Tree{rebalance: true}

	for i := 0; i < 50; i++ {
		tr.Insert(It(i * 2))
	}

	// items that exists in the tree
	for i := 0; i < 50; i++ {
		item := It(i * 2)
		result, found := tr.Search(item)
		as.True(found, "item %s not found", item)
		as.Truef(result.Eq(item), "item found doesn't match")
	}

	// items that doesn't exist in the tree
	for i := 0; i < 50; i++ {
		item := It(i*2 + 1)
		result, found := tr.Search(item)
		as.False(found, "item %s found", item)
		if result != nil {
			as.Fail("not nil", "result isn't nil: result: %s", result)
		}
	}
}

func Test_Tree_Search_func_sync(t *testing.T) {
	as := assert.New(t)
	tree := Tree{rebalance: true, duplicated: false}
	size := 1000
	concurrence := 8
	done := make(chan bool)

	search := func(min, max int) {
		for i := min; i < max; i++ {
			item := It(i)
			_, found := tree.Search(item)
			as.True(found, "item %s not found", item)
		}

		done <- true
	}

	// create the tree
	for i := 0; i < size*concurrence; i++ {
		tree.Insert(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go search(i*size, (i+1)*size)
		go treeChangeProp(&tree, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}
}

func Test_Tree_Delete_func(t *testing.T) {
	var (
		tree    Tree
		item    Item
		deleted bool
	)
	as := assert.New(t)

	/*
		Tree:
		  - rebalance
		  - no duplicated
	*/
	tree = Tree{rebalance: true, duplicated: false}

	for i := 0; i < 10; i++ {
		tree.Insert(It(i))
	}

	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(7)},
		{4, 0, nil, nil},
		{5, 1, i(4), i(6)},
		{6, 0, nil, nil},
		{7, 2, i(5), i(8)},
		{8, 1, nil, i(9)},
		{9, 0, nil, nil},
	})

	// remove one leaf
	item, deleted = tree.Delete(It(9))
	as.Equal(item.(IntItem).value, 9, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(7)},
		{4, 0, nil, nil},
		{5, 1, i(4), i(6)},
		{6, 0, nil, nil},
		{7, 2, i(5), i(8)},
		{8, 0, nil, nil},
	})

	// remove another leaf, the tree must rebalance it
	item, deleted = tree.Delete(It(8))
	as.Equal(item.(IntItem).value, 8, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(6)},
		{4, 0, nil, nil},
		{5, 1, i(4), nil},
		{6, 2, i(5), i(7)},
		{7, 0, nil, nil},
	})

	// delete one node it has only the left child.
	item, deleted = tree.Delete(It(5))
	as.Equal(item.(IntItem).value, 5, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 1, i(0), i(2)},
		{2, 0, nil, nil},
		{3, 2, i(1), i(6)},
		{4, 0, nil, nil},
		{6, 1, i(4), i(7)},
		{7, 0, nil, nil},
	})

	// delete leaf
	item, deleted = tree.Delete(It(0))
	as.Equal(item.(IntItem).value, 0, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")

	// delete one node it has only the right child.
	item, deleted = tree.Delete(It(1))
	as.Equal(item.(IntItem).value, 1, "item deleted is incorrect")
	as.True(deleted, "rw")
	checkTree(t, tree.root, []treeTest{
		{2, 0, nil, nil},
		{3, 2, i(2), i(6)},
		{4, 0, nil, nil},
		{6, 1, i(4), i(7)},
		{7, 0, nil, nil},
	})

	// delete node with 2 children
	item, deleted = tree.Delete(It(3))
	as.Equal(item.(IntItem).value, 3, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{2, 0, nil, nil},
		{4, 2, i(2), i(6)},
		{6, 1, nil, i(7)},
		{7, 0, nil, nil},
	})

	// Delete another leaf and rebalance
	item, deleted = tree.Delete(It(2))
	as.Equal(item.(IntItem).value, 2, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{4, 0, nil, nil},
		{6, 1, i(4), i(7)},
		{7, 0, nil, nil},
	})

	// delete a item it doesn't exist in the tree
	item, deleted = tree.Delete(It(9))
	if item != nil {
		as.Fail("item isn't nil", "item %s was deleted", item)
	}
	as.False(deleted, "flag indicates item was deleted")

	/*
		Tree:
		  - no rebalance
		  - no duplicated
	*/
	tree = Tree{rebalance: false, duplicated: false}

	for _, i := range []int{5, 3, 4, 2, 1, 0, 9, 7, 6, 8} {
		tree.Insert(It(i))
	}

	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 0, i(0), nil},
		{2, 0, i(1), nil},
		{3, 0, i(2), i(4)},
		{4, 0, nil, nil},
		{5, 0, i(3), i(9)},
		{6, 0, nil, nil},
		{7, 0, i(6), i(8)},
		{8, 0, nil, nil},
		{9, 0, i(7), nil},
	})

	// Delete a leaf
	item, deleted = tree.Delete(It(0))
	as.Equal(item.(IntItem).value, 0, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{2, 0, i(1), nil},
		{3, 0, i(2), i(4)},
		{4, 0, nil, nil},
		{5, 0, i(3), i(9)},
		{6, 0, nil, nil},
		{7, 0, i(6), i(8)},
		{8, 0, nil, nil},
		{9, 0, i(7), nil},
	})

	// delete one node it has only the left child.
	item, deleted = tree.Delete(It(2))
	as.Equal(item.(IntItem).value, 2, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{3, 0, i(1), i(4)},
		{4, 0, nil, nil},
		{5, 0, i(3), i(9)},
		{6, 0, nil, nil},
		{7, 0, i(6), i(8)},
		{8, 0, nil, nil},
		{9, 0, i(7), nil},
	})

	// delete leaf
	item, deleted = tree.Delete(It(6))
	as.Equal(item.(IntItem).value, 6, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")

	// delete node with 2 children
	item, deleted = tree.Delete(It(5))
	as.Equal(item.(IntItem).value, 5, "item deleted is incorrect")
	as.True(deleted, "item deleted is incorrect")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{3, 0, i(1), i(4)},
		{4, 0, nil, nil},
		{7, 0, i(3), i(9)},
		{8, 0, nil, nil},
		{9, 0, i(8), nil},
	})

	// delete a item it doesn't exist in the tree
	item, deleted = tree.Delete(It(10))
	if item != nil {
		as.Fail("item isn't nil", "item %s was deleted", item)
	}
	as.False(deleted, "flag indicates item was deleted")

	/*
		Tree:
		  - rebalance
		  - duplicated
	*/
	tree = Tree{rebalance: true, duplicated: true}

	for _, i := range []int{3, 3, 2, 1, 0, 4, 5, 4, 1, 2} {
		tree.Insert(It(i))
	}

	checkTree(t, tree.root, []treeTest{
		{0, 0, nil, nil},
		{1, 2, i(0), i(2)},
		{1, 0, nil, nil},
		{2, 1, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 2, i(3), i(5)},
		{4, 0, nil, nil},
		{5, 1, i(4), nil},
	})

	// delete leaf
	item, deleted = tree.Delete(It(0))
	as.Equal(item.(IntItem).value, 0, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{1, 2, i(1), i(2)},
		{2, 1, nil, i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 2, i(3), i(5)},
		{4, 0, nil, nil},
		{5, 1, i(4), nil},
	})

	// delete node with right child.
	item, deleted = tree.Delete(It(2))
	as.Equal(item.(IntItem).value, 2, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{1, 1, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 3, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 2, i(3), i(5)},
		{4, 0, nil, nil},
		{5, 1, i(4), nil},
	})

	// delete node with 2 children.
	item, deleted = tree.Delete(It(4))
	as.Equal(item.(IntItem).value, 4, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{1, 1, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 2, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 1, i(3), i(5)},
		{5, 0, nil, nil},
	})

	// Insert and delete the 3
	tree.Insert(It(3))
	item, deleted = tree.Delete(It(3))
	as.Equal(item.(IntItem).value, 3, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{1, 1, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 2, i(1), i(4)},
		{3, 0, nil, nil},
		{4, 1, i(3), i(5)},
		{5, 0, nil, nil},
	})

	// delete item number 3
	item, deleted = tree.Delete(It(3))
	as.Equal(item.(IntItem).value, 3, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{1, 1, i(1), i(2)},
		{2, 0, nil, nil},
		{3, 2, i(1), i(4)},
		{4, 1, nil, i(5)},
		{5, 0, nil, nil},
	})

	// delete item number 1
	item, deleted = tree.Delete(It(1))
	as.Equal(item.(IntItem).value, 1, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), nil},
		{3, 2, i(2), i(4)},
		{4, 1, nil, i(5)},
		{5, 0, nil, nil},
	})

	// delete item number 4, 5
	item, deleted = tree.Delete(It(4))
	as.Equal(item.(IntItem).value, 4, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	item, deleted = tree.Delete(It(5))
	as.Equal(item.(IntItem).value, 5, "item deleted is incorrect")
	as.True(deleted, "item wasn't deleted")
	checkTree(t, tree.root, []treeTest{
		{1, 0, nil, nil},
		{2, 1, i(1), i(3)},
		{3, 0, nil, nil},
	})
}

func Test_Tree_Delete_func_sync(t *testing.T) {
	as := assert.New(t)
	tree := Tree{rebalance: true}
	size := 1000
	concurrence := 8
	done := make(chan bool)
	deleted := func(min, max int) {
		for i := min; i < max; i++ {
			value, deleted := tree.Delete(It(i))
			as.True(deleted, "item %d wasn't deleted", i)
			as.Equal(value.(IntItem).value, i, "the value deleted is incorrect")
		}

		done <- true
	}

	for i := 0; i < size*concurrence; i++ {
		tree.Insert(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go deleted(i*size, (i+1)*size)
		go treeChangeProp(&tree, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}

	as.Equal(tree.Length(), 0, "the tree isn't empty")
}

func Test_Tree_Clear_func(t *testing.T) {
	as := assert.New(t)
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

		tr.Clear()
		as.Nil(tr.root, "root isn't nil when the tree is empty")
		as.Nil(tr.root, "tree is empty, but root isn't nil")
		as.Equal(tr.length, 0, "tree is empty, but length isn't 0")
		as.Equal(tr.rebalance, param.rebalance, "rebalance property is incorrect")
		as.Equal(tr.duplicated, param.duplicated, "duplicated property is incorrect")
	}
}

func Test_Tree_Clear_func_sync(t *testing.T) {
	// dummy test
	concurrence := 8
	tree := Tree{rebalance: true}
	done := make(chan bool)
	size := 1000
	clear := func() {
		for i := 0; i < size; i++ {
			tree.Insert(It(i))
			tree.Clear()
		}
		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go clear()
		go treeChangeProp(&tree, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}

	assert.Nil(t, tree.root, "tree root isn't nil")
	assert.Equal(t, tree.length, 0, "tree length isn't 0")
}
