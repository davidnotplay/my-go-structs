package mygostructs

// Avl is the classic binary search tree type data structure, with rebalancing. More info:
// https://en.wikipedia.org/wiki/AVL_tree
type Avl struct {
	Tree
}

// NewAvl creates an empty AVL tree.
func NewAvl() Avl {
	return Avl{Tree{rebalance: true}}
}
