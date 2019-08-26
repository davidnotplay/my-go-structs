package structs


// Bst is the classic Binary search tree. more info:
// https://en.wikipedia.org/wiki/Binary_search_tree
type Bst struct {
	tree
}

// NewBst returns an empty Bst.
func NewBst() Bst {
	return Bst{tree{rebalance: false}}
}
