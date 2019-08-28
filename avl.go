package mygostructs

// Avl is the clasic data-struct AVL tree. More info:

// Avl is the classic binary search tree type data structure, with rebalancing. More info:
// https://en.wikipedia.org/wiki/AVL_tree
type Avl struct {
	tree
}

// NewAvl creates an empty AVL tree.
func NewAvl() Avl{
	return Avl{tree{rebalance: true}}
}
