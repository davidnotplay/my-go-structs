package structs

// max returns the param more large
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// avlNode is the internal AVL tree node
type avlNode struct {
	ltree, rtree *avlNode
	height       int
	item         Item
}

// Height returns the `node.height` property. If node is nil, then returns -1
func (node *avlNode) Height() int {
	if node != nil {
		return node.height
	}

	return -1
}

// maxHeights returns the max value of left tree height and right tree height
func (node avlNode) maxHeight() int {
	return max(node.ltree.Height(), node.rtree.Height())
}

// rotateRight execute an AVL tree right rotation.
func (node *avlNode) rotateRight() *avlNode {
	newNode := node.ltree
	node.ltree = newNode.rtree
	newNode.rtree = node

	newNode.height = newNode.maxHeight() + 1
	node.height = node.maxHeight() + 1

	return newNode
}

// rotateLeft execute an AVL tree left rotation.
func (node *avlNode) rotateLeft() *avlNode {
	newNode := node.rtree
	node.rtree = newNode.ltree
	newNode.ltree = node

	newNode.height = newNode.maxHeight() + 1
	node.height = node.maxHeight() + 1

	return newNode
}

// rotateRightLeft executes the AVL double rotation (right left)
func (node *avlNode) rotateRightLeft() *avlNode {
	node.rtree = node.rtree.rotateRight()
	return node.rotateLeft()
}

// rotateLeftRight executes the AVL double rotation (left right)
func (node *avlNode) rotateLeftRight() *avlNode {
	node.ltree = node.ltree.rotateLeft()
	return node.rotateRight()
}

// Avl is the struct where the AVL tree info will store.
type Avl struct {
	root   *avlNode // Tree root.
	length int      // Number of tree nodes.
}

// NewAvl creates a new empty AVL tree.
func NewAvl() Avl {
	return Avl{nil, 0}
}

// insertItem searchs the correct position inside of the param tree node, inserts the
// it item and rebalance the node. The item inserted must be unique. If it is repeated then
// item doesn't add to the tree. The function returns the node rebalanced and a flag indicating
// it item was added.
func insertItem(node *avlNode, it Item) (*avlNode, bool) {
	var inserted bool

	if node == nil {
		return &avlNode{nil, nil, 0, it}, true
	}

	if node.item.Eq(it) {
		// AVL tree doesn't allow repeated elements.
		return node, false
	}

	if it.Less(node.item) {
		node.ltree, inserted = insertItem(node.ltree, it)
	} else {
		node.rtree, inserted = insertItem(node.rtree, it)
	}

	if inserted {
		node = rebalance(node)
	}

	return node, inserted
}

// rebalance rebalances the node and return it.
func rebalance(node *avlNode) *avlNode {
	if node.ltree.Height()-node.rtree.Height() == 2 {
		ltree := node.ltree

		if ltree.ltree.Height() <= ltree.rtree.Height() {
			node = node.rotateLeftRight()
		} else {
			node = node.rotateRight()
		}

	} else if node.rtree.Height()-node.ltree.Height() == 2 {
		rtree := node.rtree

		if rtree.rtree.Height() <= rtree.ltree.Height() {
			node = node.rotateRightLeft()
		} else {
			node = node.rotateLeft()
		}
	}

	node.height = node.maxHeight() + 1
	return node
}

// Insert inserts the item to the avl tree. The item inserted must be unique in the tree;
// else, the item isn't inserted. The function returns a flag indicating if the item was
// inserted.
func (avl *Avl) Insert(it Item) bool {
	var inserted bool
	avl.root, inserted = insertItem(avl.root, it)

	if inserted {
		avl.length++
	}

	return inserted
}

// Length returns the number of elements in the avl tree.
func (avl *Avl) Length() int {
	return avl.length
}

// search searchs the item in the node tree. Returns the avlNode that contains the item or nil
// if the item isn't found. Also returns a flag indicating if the item exists.
func search(node *avlNode, it Item) (*avlNode, bool) {
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

// Search searchs the item in the avl tree. It returns the item found and a flag indicating if
// the item exists in the avl tree.
func (avl *Avl) Search(it Item) (Item, bool) {
	if node, found := search(avl.root, it); found {
		return node.item, true
	}

	return nil, false
}

// deleteAvl searchs the item in the node tree, it deletes it and rebalance the tree. The
// functions returns the node rebalanced, the item deleted and a flag indicating if the item
// existed in the tree.
func deleteAvl(node *avlNode, it Item) (*avlNode, Item, bool) {
	var (
		found	  bool
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

		var nodeTemp *avlNode
		nodeTemp = node.rtree

		for nodeTemp.ltree != nil {
			nodeTemp = nodeTemp.ltree
		}

		itDeleted = node.item
		node.item = nodeTemp.item
		node.rtree, _, _ = deleteAvl(node.rtree, nodeTemp.item)
		return node, itDeleted, true

	} else if node.item.Less(it) {
		node.rtree, itDeleted, found = deleteAvl(node.rtree, it)
	} else {
		node.ltree, itDeleted, found = deleteAvl(node.ltree, it)
	}

	if found {
		node = rebalance(node)
	}

	return node, itDeleted, found
}

// Delete deletes the item of the avl tree. Returns the item delete or nil and a flag indicating
// if the item existed in the tree.
func (avl *Avl) Delete(it Item) (itd Item, deleted bool) {
	avl.root, itd, deleted = deleteAvl(avl.root, it)
	if deleted {
		avl.length--
	}
	return
}
