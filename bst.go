package mygostructs

// Bst is a struct it implements a binary search tree type data structure.
type Bst struct {
	Tree
}

// NewBst returns an empty Bst.
func NewBst() Bst {
	return Bst{Tree{rebalance: false}}
}
