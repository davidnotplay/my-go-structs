package mygostructs

import "sync"

// max returns the param more large
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// treeNode is the internal tree node
type treeNode struct {
	ltree, rtree *treeNode
	height       int
	item         Item
}

// getHeight returns the `node.height` property. If node is nil, then returns -1
func (node *treeNode) getHeight() int {
	if node != nil {
		return node.height
	}

	return -1
}

// maxHeights returns the max value of left tree height and right tree height
func (node treeNode) maxHeight() int {
	return max(node.ltree.getHeight(), node.rtree.getHeight())
}

// rotateRight execute an AVL tree right rotation.
func (node *treeNode) rotateRight() *treeNode {
	newNode := node.ltree
	node.ltree = newNode.rtree
	newNode.rtree = node

	newNode.height = newNode.maxHeight() + 1
	node.height = node.maxHeight() + 1

	return newNode
}

// rotateLeft execute an AVL tree left rotation.
func (node *treeNode) rotateLeft() *treeNode {
	newNode := node.rtree
	node.rtree = newNode.ltree
	newNode.ltree = node

	newNode.height = newNode.maxHeight() + 1
	node.height = node.maxHeight() + 1

	return newNode
}

// rotateRightLeft executes the AVL double rotation (right, left)
func (node *treeNode) rotateRightLeft() *treeNode {
	node.rtree = node.rtree.rotateRight()
	return node.rotateLeft()
}

// rotateLeftRight executes the AVL double rotation (left, right)
func (node *treeNode) rotateLeftRight() *treeNode {
	node.ltree = node.ltree.rotateLeft()
	return node.rotateRight()
}

// Tree struct is the base for the Bst struct and AVL struct.
type Tree struct {
	root       *treeNode  // Tree root.
	length     int        // Number of tree nodes.
	rebalance  bool       // Rebalance the tree after modify it.
	duplicated bool       // Flag indicating if allows duplicated items.
	mutex      sync.Mutex // Lock for avoid the concurrence when manipulate the struct.
}

// insertItem searchs the correct position inside of the param tree node, inserts the
// it item and rebalance the node, if rebalance flag is true. If duplicated paramater is false, the
// item inserted must be unique. The function returns the node rebalanced and a flag indicating it
// item was added. The item isn't added if the item isn't unique and duplicated flag is false.
func insertItem(node *treeNode, it Item, rebalanceIt bool, duplicated bool) (*treeNode, bool) {
	var inserted bool

	if node == nil {
		return &treeNode{nil, nil, 0, it}, true
	}

	if node.item.Eq(it) && !duplicated {
		return node, false
	}

	if it.Less(node.item) {
		node.ltree, inserted = insertItem(node.ltree, it, rebalanceIt, duplicated)
	} else {
		node.rtree, inserted = insertItem(node.rtree, it, rebalanceIt, duplicated)
	}

	if inserted && rebalanceIt {
		node = rebalance(node)
	}

	return node, inserted
}

// insertGetAdy searchs the position in the node, inserts the item and rebalance the node if reb
// flag is true. If duplicated paramater is false, the item inserted must be unique. The function
// returns the node rebalanced, the item previous already inserted in the tree, and a flag
// indicating it item was added. The item isn't added if the item isn't unique and duplicated flag
// is false.
func insertGetAdy(node *treeNode, item Item, reb, duplicated bool) (*treeNode, *Item, bool) {
	var (
		inserted bool
		prev     *Item
	)

	if node == nil {
		return &treeNode{nil, nil, 0, item}, nil, true
	}

	if node.item.Eq(item) && !duplicated {
		return node, nil, false
	}

	if item.Less(node.item) {
		node.ltree, prev, inserted = insertGetAdy(node.ltree, item, reb, duplicated)
	} else {
		node.rtree, prev, inserted = insertGetAdy(node.rtree, item, reb, duplicated)
	}

	if prev == nil && (node.item.Less(item) || node.item.Eq(item)) {
		prev = &node.item
	}

	if inserted && reb {
		node = rebalance(node)
	}

	return node, prev, inserted
}

// rebalance rebalances the node and return it.
func rebalance(node *treeNode) *treeNode {
	if node.ltree.getHeight()-node.rtree.getHeight() == 2 {
		ltree := node.ltree

		if ltree.ltree.getHeight() <= ltree.rtree.getHeight() {
			node = node.rotateLeftRight()
		} else {
			node = node.rotateRight()
		}

	} else if node.rtree.getHeight()-node.ltree.getHeight() == 2 {
		rtree := node.rtree

		if rtree.rtree.getHeight() <= rtree.ltree.getHeight() {
			node = node.rotateRightLeft()
		} else {
			node = node.rotateLeft()
		}
	}

	node.height = node.maxHeight() + 1
	return node
}

// Insert inserts the item in the tree. The function returns a flag indicating if the operation
// was success or the item cannot be inserted because it was duplicated.
func (tr *Tree) Insert(it Item) bool {
	var inserted bool

	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	tr.root, inserted = insertItem(tr.root, it, tr.rebalance, tr.duplicated)

	if inserted {
		tr.length++
	}

	return inserted
}

// Length returns the number of items in the tree.
func (tr *Tree) Length() int {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	return tr.length
}

// search searchs the item in the node tree. Returns the treeNode that contains the item or nil
// if the item isn't found. Also returns a flag indicating if the item exists.
func search(node *treeNode, it Item) (*treeNode, bool) {
	if node == nil {
		// item  not found
		return nil, false
	}

	if node.item.Eq(it) {
		return node, true
	}

	if node.item.Less(it) {
		return search(node.rtree, it)
	}

	return search(node.ltree, it)
}

// Search searchs the item in the tree. It returns the item found and a flag indicating if
// the item exists in the tree tree.
func (tr *Tree) Search(it Item) (Item, bool) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	if node, found := search(tr.root, it); found {
		return node.item, true
	}

	return nil, false
}

// deleteNode searchs the item in the node tree, it deletes it and rebalance the tree, if the flag
// is true. The functions returns the node rebalanced, the item deleted and a flag indicating if
// the item existed in the tree.
func deleteNode(node *treeNode, it Item, rebalanceIt bool) (*treeNode, Item, bool) {
	var (
		found     bool
		itDeleted Item
	)

	if node == nil {
		return node, nil, false
	}

	if node.item.Eq(it) {
		if node.ltree == nil {
			return node.rtree, node.item, true
		} else if node.rtree == nil {
			return node.ltree, node.item, true
		}

		nodeTemp := node.rtree
		for nodeTemp.ltree != nil {
			nodeTemp = nodeTemp.ltree
		}

		itDeleted = node.item
		node.item = nodeTemp.item
		node.rtree, _, _ = deleteNode(node.rtree, nodeTemp.item, rebalanceIt)
		found = true

	} else if node.item.Less(it) {
		node.rtree, itDeleted, found = deleteNode(node.rtree, it, rebalanceIt)
	} else {
		node.ltree, itDeleted, found = deleteNode(node.ltree, it, rebalanceIt)
	}

	if found && rebalanceIt {
		node = rebalance(node)
	}

	return node, itDeleted, found
}

// Delete deletes the item of the tree. Returns the item deleted and a flag indicating if the item
// existed in the tree.
func (tr *Tree) Delete(it Item) (itd Item, deleted bool) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	tr.root, itd, deleted = deleteNode(tr.root, it, tr.rebalance)
	if deleted {
		tr.length--
	}
	return
}

// Clear clears the tree.
func (tr *Tree) Clear() {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	tr.root = nil
	tr.length = 0
}
