package mygostructs

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

// Height returns the `node.height` property. If node is nil, then returns -1
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

// rotateRightLeft executes the AVL double rotation (right left)
func (node *treeNode) rotateRightLeft() *treeNode {
	node.rtree = node.rtree.rotateRight()
	return node.rotateLeft()
}

// rotateLeftRight executes the AVL double rotation (left right)
func (node *treeNode) rotateLeftRight() *treeNode {
	node.ltree = node.ltree.rotateLeft()
	return node.rotateRight()
}

// tree is the struct where the tree info will store.
type tree struct {
	root       *treeNode // Tree root.
	length     int       // Number of tree nodes.
	rebalance  bool      // Rebalance the tree after modify it.
	duplicated bool      // Flag indicating if allows duplicated items.
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

// Insert inserts the item in the tree. The function resturns true, if the item was inserted or
// false if the item already in the tree (duplicated item).
func (tr *tree) Insert(it Item) bool {
	var inserted bool
	tr.root, inserted = insertItem(tr.root, it, tr.rebalance, tr.duplicated)

	if inserted {
		tr.length++
	}

	return inserted
}

// Length returns the number of items in the tree.
func (tr *tree) Length() int {
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
func (tr *tree) Search(it Item) (Item, bool) {
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

		var nodeTemp *treeNode
		nodeTemp = node.rtree

		for nodeTemp.ltree != nil {
			nodeTemp = nodeTemp.ltree
		}

		itDeleted = node.item
		node.item = nodeTemp.item
		node.rtree, _, _ = deleteNode(node.rtree, nodeTemp.item, rebalanceIt)
		return node, itDeleted, true

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
func (tr *tree) Delete(it Item) (itd Item, deleted bool) {
	tr.root, itd, deleted = deleteNode(tr.root, it, tr.rebalance)
	if deleted {
		tr.length--
	}
	return
}
