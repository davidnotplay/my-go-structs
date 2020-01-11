package mygostructs

import "sync"

// listNode is the node for the List struct.
type listNode struct {
	prev *listNode
	next *listNode
	item Item
}

// Less checks if item in the listNode is less than the the parameter item.
func (ln listNode) Less(it Item) bool {
	lnn, valid := it.(*listNode)
	return valid && ln.item.Less(lnn.item)
}

// Eq checks if item in the listNode is equal to the parameter item.
func (ln listNode) Eq(it Item) bool {
	lnn, valid := it.(*listNode)
	return valid && ln.item.Eq(lnn.item)
}

//String transforms and returns the item in the listNode as an String.
func (ln listNode) String() string {
	return ln.item.String()
}

// List is a struct it implements a doubly linked list type data structure. The items are inserted
// linearly. It can access and manipulate any item of the list. Also it allows to search quickly
// items.
//
// The struct is adapted to run in multithread code.
type List struct {
	fnode *listNode  // pointer to the first node of the list.
	lnode *listNode  // ponter to the last node of the list
	pnode *listNode  // Internal pointer. It is moved using the struct functions.
	avl   Tree       // avl tree
	mutex sync.Mutex // Lock for avoid the concurrence when manipulate the struct.
}

// NewList returns an empty List. The parameter is flag indicating if the list allows items
// duplicated.
func NewList(duplicated bool) List {
	return List{avl: Tree{rebalance: true, duplicated: duplicated}}
}

// AddAfter adds the item after the item pointed by internal pointer and moves the internal
// pointer to the new item inserted. Returns a flag indicating if the item was added successfully.
func (l *List) AddAfter(it Item) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node := listNode{}
	node.item = it

	// Insert in tree
	if !l.avl.Insert(&node) {
		return false
	}

	if l.fnode == nil {
		// List is empty. Add the first node
		l.fnode = &node
		l.lnode = &node
		l.pnode = &node
		return true
	}

	node.next = l.pnode.next
	node.prev = l.pnode
	l.pnode.next = &node

	if node.next != nil {
		node.next.prev = &node
	}

	if node.next == nil {
		// The node has been inserted in the last position.
		l.lnode = &node
	}

	l.pnode = &node
	return true
}

// AddBefore adds the item before the item pointed by internal pointer and moves the internal
// pointer to the new item inserted. Returns a flag indicating if the item was added successfully
func (l *List) AddBefore(it Item) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node := listNode{}
	node.item = it

	// Insert in tree
	if !l.avl.Insert(&node) {
		return false
	}

	if l.fnode == nil {
		// List is empty. Add the first node
		l.fnode = &node
		l.lnode = &node
		l.pnode = &node
		return true
	}

	node.next = l.pnode
	node.prev = l.pnode.prev
	l.pnode.prev = &node

	if node.prev != nil {
		node.prev.next = &node
	}

	if node.prev == nil {
		// the value inserted is the first.
		l.fnode = &node
	}

	l.pnode = &node
	return true
}

// Next moves the internal pointer to the next item. Returns a flag indicating if the operation
// was possible.
func (l *List) Next() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pnode != nil && l.pnode.next != nil {
		l.pnode = l.pnode.next
		return true
	}

	return false
}

// Prev moves the internal pointer to the previous item. Returns a flag indicating if the operation
// was possible.
func (l *List) Prev() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pnode != nil && l.pnode.prev != nil {
		l.pnode = l.pnode.prev
		return true
	}

	return false
}

// First moves the internal pointer to the first item of the list.
func (l *List) First() {
	l.mutex.Lock()
	l.pnode = l.fnode
	l.mutex.Unlock()
}

// Last moves the internal pointer to the last item of the list.
func (l *List) Last() {
	l.mutex.Lock()
	l.pnode = l.lnode
	l.mutex.Unlock()
}

// Advance advances the internal pointer one position and returns the item pointed. The second
// value returned is a flag indicating if the operation was successfully.
func (l *List) Advance() (Item, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pnode != nil && l.pnode.next != nil {
		it := l.pnode.next.item
		l.pnode = l.pnode.next
		return it, true
	}

	return nil, false
}

// Rewind rewinds the internal pointer one position and returns the item pointed. The second value
// returned is a flag indicating if the operation was successfully.
func (l *List) Rewind() (Item, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pnode != nil && l.pnode.prev != nil {
		it := l.pnode.prev.item
		l.pnode = l.pnode.prev
		return it, true
	}

	return nil, false
}

// Get gets the item pointed by the internal pointer. Returns the item and a flag indicating if
// it was possible get the item.
func (l *List) Get() (Item, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pnode != nil {
		return l.pnode.item, true
	}

	return nil, false

}

// Replace replaces the item pointed by the internal pointer by the item of parameter.
// Returns a flag indicating if the operatio was successfully.
func (l *List) Replace(it Item) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.avl.length == 0 {
		return false //list empty
	}

	nodeReplaced := &listNode{l.pnode.prev, l.pnode.next, it}
	l.avl.Delete(l.pnode)

	if l.pnode == l.fnode {
		l.fnode = nodeReplaced
	} else {
		nodeReplaced.prev.next = nodeReplaced
	}

	if l.pnode == l.lnode {
		l.lnode = nodeReplaced
	} else {
		nodeReplaced.next.prev = nodeReplaced
	}

	l.pnode = nodeReplaced
	l.avl.Insert(nodeReplaced)

	return true
}

// Search searchs the item in the list. Returns the item searched and a flag indicating if the
// item was found.
func (l *List) Search(it Item) (Item, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node, found := l.avl.Search(&listNode{item: it})
	if found {
		l.pnode = node.(*listNode)
		return l.pnode.item, true
	}

	return nil, false
}

// Delete deletes the item pointed by the internal pointer and it moves the internal pointer to
// the begining of the list. The second value indicates if the item was deleted.
func (l *List) Delete() (Item, bool) {
	var item Item

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.avl.length == 0 {
		return item, false
	}

	// Delete from tree
	l.avl.Delete(l.pnode)

	if l.avl.length == 0 {
		item = l.pnode.item

		// create empty list.
		l.fnode = nil
		l.pnode = nil
		l.lnode = nil

		return item, true
	}

	tmpPnode := l.pnode
	item = tmpPnode.item

	if tmpPnode.prev != nil {
		tmpPnode.prev.next = tmpPnode.next
	} else {
		l.fnode = tmpPnode.next
	}

	if tmpPnode.next != nil {
		tmpPnode.next.prev = tmpPnode.prev
	} else {
		l.lnode = tmpPnode.prev
	}

	l.pnode = l.fnode
	return item, true
}

// Length returns the number of items in the list.
func (l *List) Length() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.avl.length
}

// Clear clears the list.
func (l *List) Clear() {
	l.mutex.Lock()
	l.avl = Tree{rebalance: true, duplicated: l.avl.duplicated}
	l.fnode = nil
	l.pnode = nil
	l.lnode = nil
	l.mutex.Unlock()
}

// ForEach excutes the function of the parameter in all items of the list, consecutively and from
// the begining. The behaviour of this function isn't defined if you modify the list inside of the
// function or in another thread while this method is executing
func (l *List) ForEach(f func(Item)) {

	if l.Length() == 0 {
		// empty list
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	for node := l.fnode; node != nil; node = node.next {
		l.mutex.Unlock()
		f(node.item)
		l.mutex.Lock()
	}
}

// Map creates a new list using the results of parser function execution in all items of the list.
// The behaviour of this function isn't defined if you modify the list inside of the function or
// in another thread while this method is executing
func (l *List) Map(parser func(Item) Item) *List {
	var (
		newList List
		forFunc func(Item)
	)

	newList = NewList(l.avl.duplicated)
	forFunc = func(it Item) {
		newList.AddAfter(parser(it))
	}

	l.ForEach(forFunc)
	return &newList
}

// Filter creates a new list with all items that pass the test implemented in the filter function.
// The behaviour of this function isn't defined if you modify the list inside of the function or
// in another thread while this method is executing
func (l *List) Filter(filter func(Item) bool) *List {
	var (
		newList List
		forFunc func(Item)
	)

	newList = NewList(l.avl.duplicated)
	forFunc = func(it Item) {
		if filter(it) {
			newList.AddAfter(it)
		}
	}

	l.ForEach(forFunc)
	return &newList
}
