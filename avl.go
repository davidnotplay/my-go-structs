package mygostructs

// Avl is a struct it implements a AVL tree type data structure.
//
// The struct is adapted to run in multithread code.
type Avl struct {
	Tree
}

// NewAvl creates an empty AVL tree.
func NewAvl() Avl {
	return Avl{Tree{rebalance: true}}
}
