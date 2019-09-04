package mygostructs

// Bst is the classic binary search tree type data structure. More info:
// https://en.wikipedia.org/wiki/Binary_search_tree
type Bst struct {
	Tree
}

// NewBst returns an empty Bst.
func NewBst() Bst {
	return Bst{Tree{rebalance: false}}
}
