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

// Left returns the `node` left child. If node is nil then returns nil.
func (node *avlNode) Left() *avlNode {
	if node != nil {
		return node.ltree
	}

	return nil
}

// Right returns the `node` right child. If node is nil then returns nil.
func (node *avlNode) Right() *avlNode {
	if node != nil {
		return node.rtree
	}

	return nil
}

// Value returns the `node` value. If node is nil then returns nil.
func (node *avlNode) Value() Value {
	if node != nil {
		return *node.value
	}

	return nil
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
	var duplicated bool

	if node == nil {
		return &avlNode{nil, nil, 0, v}, true
	}

	if node.Value().Eq(*v) {
		// AVL tree doesn't allow repeated elements.
		return node, false
	}

	if (*v).Less(node.Value()) {
		node.ltree, duplicated = insertValue(node.ltree, v)

		if node.ltree.Height() - node.rtree.Height() == 2 {
			// Balance tree
			if (*v).Less(node.ltree.Value()) {
				node = node.rotateRight()
			} else {
				node = node.rotateLeftRight()
			}
		}
	} else {
		node.rtree, duplicated = insertValue(node.rtree, v)
		if node.rtree.Height() - node.ltree.Height() == 2 {
			// Balance tree
			if node.rtree.Value().Less(*v) {
				node = node.rotateLeft()
			} else {
				node = node.rotateRightLeft()
			}
		}
	}

	node.height = node.maxHeight() + 1
	return node, duplicated
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


// func (node *avlNode) stringifyNode(sep string, space string) string {
// 	if node == nil {
// 		return "NULL\n"
// 	}

// 	valueStr :=  fmt.Sprintf("%s (%d)", (*node.value).String(), node.Height())
// 	lChildStr := node.ltree.stringifyNode(fmt.Sprintf("%s│%s", sep, space), space)
// 	rChildStr := node.rtree.stringifyNode(fmt.Sprintf("%s%s%s", sep, space, space), space)
// 	return fmt.Sprintf("%s\n%s├─%s%s└─%s", valueStr, sep, lChildStr, sep, rChildStr)
// }

// func (node avlNode) Stringify() string {
// 	return node.StringifyWithIndent(" ")
// }

// func (node avlNode) StringifyWithIndent(indent string) string {
// 	return node.stringifyNode("", indent)
// }

type AvlNode interface {
	Left()	 AvlNode
	Right()  AvlNode
	Value()  Value
	Height() int
}
