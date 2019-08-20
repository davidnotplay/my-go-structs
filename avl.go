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
	value        Value
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

// NewAvl creates a new empty value
func NewAvl() Avl {
	return Avl{nil, 0}
}

// InsertValues searchs the correct position inside of the param tree node, inserts the
// v value and rebalance the node. The value inserted must be unique. If it is repeated then
// value doesn't add to the tree. The function returns the node rebalanced and a flag indicating
// v value was added.
func insertValue(node *avlNode, v Value) (*avlNode, bool) {
	var inserted bool

	if node == nil {
		return &avlNode{nil, nil, 0, v}, true
	}

	if node.value.Eq(v) {
		// AVL tree doesn't allow repeated elements.
		return node, false
	}

	if v.Less(node.value) {
		node.ltree, inserted = insertValue(node.ltree, v)
	} else {
		node.rtree, inserted = insertValue(node.rtree, v)
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

// Insert inserts the value to the avl tree. The value inserted must be unique in the tree;
// else, the value isn't inserted. The function returns a flag indicating if v value was
// inserted.
func (avl *Avl) Insert(v Value) bool {
	var inserted bool
	avl.root, inserted = insertValue(avl.root, v)

	if inserted {
		avl.length++
	}

	return inserted
}

// Length returns the number of elements in the avl tree.
func (avl *Avl) Length() int {
	return avl.length
}

// search searchs the value in the node tree. Returns the avlNode that contains the value or nil
// if v isn't found. Also returns a flag indicating if the value exists.
func search(node *avlNode, v Value) (*avlNode, bool) {
	if node == nil {
		// v value not found
		return nil, false
	}

	if node.value.Eq(v) {
		return node, true
	}

	if node.value.Less(v) {
		return search(node.rtree, v)
	}

	return search(node.ltree, v)
}

// Search searchs the value in the avl tree. It returns the value found and a flag indicating if
// the value exists in the avl tree.
func (avl *Avl) Search(v Value) (Value, bool) {
	if node, found := search(avl.root, v); found {
		return node.value, true
	}

	return nil, false
}

// deleteAvl searchs the value in the node tree, it deletes it and rebalance the tree.  The
// functions returns the node rebalanced, the node deleted and a flag indicating if the value
// existed in the tree.
func deleteAvl(node *avlNode, v Value) (*avlNode, Value, bool) {
	var (
		found    bool
		vDeleted Value
	)

	if node == nil {
		return node, nil, false
	}

	if node.value.Eq(v) {
		if node.ltree == nil {
			return node.rtree, node.value, true
		} else if node.rtree == nil {
			return node.ltree, node.value, true
		}

		var nodeTemp *avlNode
		nodeTemp = node.rtree

		for nodeTemp.ltree != nil {
			nodeTemp = nodeTemp.ltree
		}

		vDeleted = node.value
		node.value = nodeTemp.value
		node.rtree, _, _ = deleteAvl(node.rtree, nodeTemp.value)
		return node, vDeleted, true

	} else if node.value.Less(v) {
		node.rtree, vDeleted, found = deleteAvl(node.rtree, v)
	} else {
		node.ltree, vDeleted, found = deleteAvl(node.ltree, v)
	}

	if found {
		node = rebalance(node)
	}

	return node, vDeleted, found
}

// Delete deletes value v of the avl tree. Returns the value delete or nil and a flag indicating
// if the value existed in the tree.
func (avl *Avl) Delete(v Value) (vd Value, deleted bool) {
	avl.root, vd, deleted = deleteAvl(avl.root, v)
	if deleted {
		avl.length--
	}
	return
}
