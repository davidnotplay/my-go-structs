package mygostructs

// SortedList is a struct it implements a ordered doubly linked list type data structure. The items
// are inserted ordered. It can access and manipulate any item of the list. Also it allows search
// quickly, if an item exists in the list.
type SortedList struct {
	list List
}

// NewSortedList creates and returns a new empty sorted list.
func NewSortedList(duplicated bool) SortedList{
	return SortedList{NewList(duplicated)}
}

// Add adds the item to sorted list. Returns a flag indicating if the item wass added
// successfully.
func (so *SortedList) Add(item Item) bool {
	var (
		prev	 *Item
		inserted bool
		node     *listNode
	)

	node = &listNode{item: item}

	so.list.avl.root, prev, inserted = insertGetAdy(
		so.list.avl.root,
		node,
		so.list.avl.rebalance,
		so.list.avl.duplicated)

	if !inserted {
		return false
	}

	so.list.pnode = node
	so.list.avl.length++

	switch true {
	case so.list.avl.length == 1:
		so.list.fnode = node
		so.list.lnode = node

	case prev == nil:
		// the node inserted is min
		node.next = so.list.fnode
		so.list.fnode.prev = node
		so.list.fnode = node

	default:
		prevItem := (*prev).(*listNode)
		node.prev = prevItem
		node.next = prevItem.next
		prevItem.next = node

		if node.next != nil {
			node.next.prev = node
		} else {
			so.list.lnode = node
		}
	}

	return true
}

// Next moves the internal pointer to the next item. Returns a flag indicating if the operation
// was possible.
func (so *SortedList) Next() bool {
	return so.list.Next()
}

// Prev moves the internal pointer to the previous item. Returns a flag indicating if the operation
// was possible.
func (so *SortedList) Prev() bool {
	return so.list.Prev()
}

// First moves the internal pointer to the first item of the list.
func (so *SortedList) First() {
	so.list.First()
}

// Last moves the internal pointer to the last item of the list.
func (so *SortedList) Last() {
	so.list.Last()
}

// Advance advances the internal pointer one position and returns the item pointed. The second
// value returned is a flag indicating if the operation was successfully.
func (so *SortedList) Advance() (Item, bool) {
	return so.list.Advance()
}

// Rewind rewinds the internal pointer one position and returns the item pointed. The second value
// returned is a flag indicating if the operation was successfully.
func (so *SortedList) Rewind() (Item, bool) {
	return so.list.Rewind()
}

// Search searchs the item in the list. Returns the item searched and a flag indicating if the
// item was found.
func (so *SortedList) Get() (Item, bool) {
	return so.list.Get()
}

// Search searchs the item in the list. It returns the item found and a flag indicating if the item
// exists in the list. This function also move the internal pointer to the item found.
func (so *SortedList) Search(it Item) (Item, bool) {
	return so.list.Search(it)
}

// Delete deletes the item pointed by the internal pointer and it moves the internal pointer to
// the begining of the list. The second value indicates if the item was deleted.
func (so *SortedList) Delete() (Item, bool) {
	return so.list.Delete()
}

// Clear clears the list.
func (so *SortedList) Clear() {
	so.list.Clear()
}

// Length returns the number of items in the list.
func (so *SortedList) Length() int {
	return so.list.Length()
}

// ForEach excutes the function of the parameter in all items of the list, consecutively and
// from the begining.
func (so *SortedList) ForEach(f func(Item)) {
	so.list.ForEach(f)
}

// Map creates a new list using the results of parser function execution in all items of
// the list.
func (so *SortedList) Map(parser func(Item) Item) *SortedList {
	return &SortedList{*(so.list.Map(parser))}
}

// Filter create a new list with all items that pass the test implemented in the filter
// function.
func (so *SortedList) Filter(filter func(Item) bool) *SortedList {
	return &SortedList{*(so.list.Filter(filter))}
}
