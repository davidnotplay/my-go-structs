package structs

// Avl is the clasic data-struct AVL tree. More info:
// https://en.wikipedia.org/wiki/AVL_tree
type Avl struct {
	tree
}

// NewAvl creates an empty AVL tree.
func NewAvl() Avl{
	return Avl{tree{rebalance: true}}
}
