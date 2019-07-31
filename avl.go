package structs

import (
	"fmt"
)

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




func (node *avlNode) stringifyNode(sep string, space string) string {
	if node == nil {
		return "NULL\n"
	}

	valueStr :=  (*node.value).String()
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

type AvlNode interface {
	Left()	 AvlNode
	Right()  AvlNode
	Value()  Value
	Height() int
}
