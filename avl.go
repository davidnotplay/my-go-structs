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

// treeNode is the interanal AVL tree node
type treeNode struct {
	ltree, rtree *treeNode
	height       int
	value        Value
}

// Height returns the `node.height` property. If node is nil, then returns -1
func (node *treeNode) Height() int {
	if node != nil {
		return node.height
	}

	return -1
}

// maxHeights returns the max value of left tree height and right tree height
func (node treeNode) maxHeight() int {
	return max(node.ltree.Height(), node.rtree.Height())
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




func (node *treeNode) stringifyNode(sep string, space string) string {
	if node == nil {
		return "NULL\n"
	}

	valueStr :=  node.value.String()
	lChildStr := node.ltree.stringifyNode(fmt.Sprintf("%s│%s", sep, space), space)
	rChildStr := node.rtree.stringifyNode(fmt.Sprintf("%s%s%s", sep, space, space), space)
	return fmt.Sprintf("%s\n%s├─%s%s└─%s", valueStr, sep, lChildStr, sep, rChildStr)
}

func (node treeNode) Stringify() string {
	return node.StringifyWithIndent(" ")
}

func (node treeNode) StringifyWithIndent(indent string) string {
	return node.stringifyNode("", indent)
}
