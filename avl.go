package structs

import "fmt"

// max returns the param more large
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// avlNode is the interanal AVL tree node
type avlNode struct {
	ltree, rtree *avlNode
	height       int
	value        *Value
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
	root	*avlNode // Tree root.
	length  int      // Number of tree nodes.
}


// newAvl creates a new empty value
func NewAvl() Avl {
	return Avl{nil, 0}
}

// insertValues insert the `v` value in the `node` tree or in one the his sub-trees, and
// rotate the nodes for balance the new tree. No allow elements duplicated.
// Returns the new tree and and a flag indicating if the element has been inserted.
func insertValue(node *avlNode, v *Value) (*avlNode, bool){
	var inserted bool

	if node == nil {
		return &avlNode{nil, nil, 0, v}, true
	}

	if (*node.value).Eq(*v) {
		// AVL tree doesn't allow repeated elements.
		return node, false
	}

	if (*v).Less(*node.value) {
		node.ltree, inserted = insertValue(node.ltree, v)
	} else {
		node.rtree, inserted = insertValue(node.rtree, v)
	}

	if inserted {
		node = rebalance(node)
	}

	return node, inserted
}

// rebalance Reblance the `node` avlNode and return it.
func rebalance(node *avlNode) *avlNode {
	if node.ltree.Height() - node.rtree.Height() == 2 {
		ltree := node.ltree

		if ltree.ltree.Height() <= ltree.rtree.Height() {
			node = node.rotateLeftRight()
		} else {
			node = node.rotateRight()
		}

	} else if node.rtree.Height() - node.ltree.Height() == 2 {
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

// Insert inserts the `v` value in the avl tree.
// The function returns true if the element has been inserted.
// False if `v` element is duplicated inside the tree
func (avl *Avl) Insert(v *Value) bool {
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

// search searchs the `v` value in the avlNode or in one of their trees
// The function returns:
//	- node where is the value. If value doesn't exist then returns nil.
//	- Flag indicating if the value has been found.
func search(v *Value, node *avlNode) (*avlNode, bool) {
	if node == nil {
		// v value not found
		return nil, false
	}


	if (*node.value).Eq(*v) {
		return node, true
	}

	if (*node.value).Less(*v) {
		return search(v, node.rtree)
	}

	return search(v, node.ltree)
}

// Search searchs the `v` value in the avl tree.
// The function returns: The value found in the avl tree. If `v` doesn't exists, then returns nil.
// And a flag indicating if the value exists or not.
func (avl *Avl) Search(v *Value) (Value, bool) {
	if node, found := search(v, avl.root); found {
		return *node.value, true
	}

	return nil, false
}

// deleteAvl searchs the `v` value in the `node` tree,  delete it if it is found, and re-balance
// `node` tree. The function returns tne node deleted, the new node re-balanced and a flag
// indicating if the value was found.
func deleteAvl(node *avlNode, v *Value) (*avlNode, *Value, bool) {
	var (
		found bool
		vDeleted *Value
	)

	if node == nil {
		return node, nil, false
	}

	if (*node.value).Eq(*v) {
		if node.ltree == nil {
			return node.rtree, node.value, true
		} else if node.rtree == nil {
			return node.ltree, node.value, true
		}

		var nodeTemp *avlNode
		nodeTemp = node.rtree;

		for nodeTemp.ltree != nil {
			nodeTemp = nodeTemp.ltree
		}

		vDeleted = node.value
		node.value = nodeTemp.value
		node.rtree, _, _ = deleteAvl(node.rtree, nodeTemp.value)
		return node, vDeleted, true

	} else if (*node.value).Less(*v) {
		node.rtree, vDeleted, found = deleteAvl(node.rtree, v)
	} else {
		node.ltree, vDeleted, found = deleteAvl(node.ltree, v)
	}

	if found {
		node = rebalance(node)
	}


	return node, vDeleted, found
}

// Delete delete the `v` value of the avl tree.
// Returns the value deleted and a flag indicating if `v` was found in the tree.
func (avl *Avl)Delete(v *Value) (Value, bool){
	var (
		vDeleted *Value = nil
		found bool = false
	)

	if avl.root, vDeleted, found = deleteAvl(avl.root, v); found {
		avl.length--
	}

	return *vDeleted, found
}

func (node *avlNode) stringifyNode(sep string, space string) string {
	if node == nil {
		return "NULL\n"
	}

	valueStr :=  fmt.Sprintf("%s (%d)", (*node.value).String(), node.Height())
	lChildStr := node.ltree.stringifyNode(fmt.Sprintf("%s│%s", sep, space), space)
	rChildStr := node.rtree.stringifyNode(fmt.Sprintf("%s%s%s", sep, space, space), space)
	return fmt.Sprintf("%s\n%s├─%s%s└─%s", valueStr, sep, lChildStr, sep, rChildStr)
}

func (node avlNode) Stringify() string {
	return node.StringifyWithIndent(" ")
}

func (node avlNode) StringifyWithIndent(indent string) string {
	return node.stringifyNode("", indent)
}
