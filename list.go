package mygostructs

// listNode is the node for the List struct.
type listNode struct {
	prev *listNode
	next *listNode
	item Item
}

func (ln listNode) Less(it Item) bool {
	lnn, valid := it.(*listNode)
	return valid && ln.item.Less(lnn.item)
}

func (ln listNode) Eq(it Item) bool {
	lnn, valid := it.(*listNode)
	return valid && ln.item.Eq(lnn.item)
}

func (ln listNode) String() string {
	return ln.item.String()
}

// List is a doubly linked list type data structure. More info:
// https://en.wikipedia.org/wiki/Doubly_linked_list
type List struct {
	fnode  *listNode // pointer to the first node of the list.
	lnode  *listNode // ponter to the last node of the list
	pnode  *listNode // Internal pointer. It is moved using the struct functions.
	length int       // List size
	avl    tree      // avl tree
}

// NewList returns an empty List. The duplicated parameter is flag indicating if the list allows
// items duplicated.
func NewList(duplicated bool) List {
	return List{avl: tree{rebalance: true, duplicated: duplicated}}
}

// AddAfter adds the item after the item pointed by internal pointer and moves the internal
// pointer to the new item inserted. The function returns true if the item was inserted or false
// if the item is duplicated, and duplicated property is false.
func (l *List) AddAfter(it Item) bool {
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
		l.length++
		return true
	}

	node.next = l.pnode.next
	node.prev = l.pnode
	l.pnode.next = &node

	if node.next != nil {
		node.next.prev = &node
	}
	l.length++

	if node.next == nil {
		// the node has been inserted in the last position.
		l.lnode = &node
	}

	l.pnode = &node
	return true
}

// AddBefore adds the item before the item pointed by internal pointer and moves the internal
// pointer to the new item inserted. The function returns true if the item was inserted or false
// if the item is duplicated, and duplicated property is false.
func (l *List) AddBefore(it Item) bool {
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
		l.length++
		return true
	}

	node.next = l.pnode
	node.prev = l.pnode.prev
	l.pnode.prev = &node

	if node.prev != nil {
		node.prev.next = &node
	}
	l.length++

	if node.prev == nil {
		// the value inserted is the first.
		l.fnode = &node
	}

	l.pnode = &node
	return true
}

// Next moves the internal pointer to the next item, if it possible. Returns true if the pointer
// was moved correctly or false if it is impossible (the list is empty or the internal pointer is
// pointed to the last item)
func (l *List) Next() bool {
	if l.pnode.next != nil {
		l.pnode = l.pnode.next
		return true
	}

	return false
}

// Prev moves the internal pointer to the previous item, if it possible. Returns true if the
// pointer was moved correctly or false if it is impossible (the list is empty or the internal
// pointer is pointed to the first item)
func (l *List) Prev() bool {
	if l.pnode.prev != nil {
		l.pnode = l.pnode.prev
		return true
	}

	return false
}

// First moves the internal pointer to the first item of the list.
func (l *List) First() {
	l.pnode = l.fnode
}

// Last moves the internal pointer to the last item of the list.
func (l *List) Last() {
	l.pnode = l.lnode
}

// Advance advances the internal pointer one position and returns the item pointed and the true
// value. If it isn't possible advance more, because the internal pointer at end, the function
// returns nil and false then.
func (l *List) Advance() (Item, bool) {
	if l.pnode.next != nil {
		defer l.Next()
		return l.pnode.next.item, true
	}

	return nil, false
}

// Rewind rewinds the internal pointer one position and returns the item pointed and the true
// value. If it can not rewind, because the internal pointer is in the begin, the function returns
// nil and false then.
func (l *List) Rewind() (Item, bool) {
	if l.pnode.prev != nil {
		defer l.Prev()
		return l.pnode.prev.item, true
	}

	return nil, false
}

// Get gets the item pointed by the internal pointer. Returns the item and a boolean flag with
// the true value if the item was getted or false if the list is empty.
func (l *List) Get() (Item, bool) {
	if l.pnode != nil {
		return l.pnode.item, true
	}

	return nil, false

}

// Replace replaces the item pointed by the internal pointer by the item of parameter.
// Returns true if the function has replaced the item. False if it is impossible: The list is
// empty.
func (l *List)Replace(it Item) bool {
	if l.Length() == 0 {
		return false //list empty
	}

	l.pnode.item = it
	return true
}

// Delete deletes the item pointed by the internal pointer and move the internal pointer to the
// first of the list. It returns 2 values: The item deleted and a flag indicating if the value
// was deleted.
func (l *List) Delete() (Item, bool) {
	var item Item

	if l.length == 0 {
		return item, false
	}

	// Delete from tree
	l.avl.Delete(l.pnode)

	if l.length == 1 {
		item = l.pnode.item

		// create empty list.
		l.fnode = nil
		l.pnode = nil
		l.lnode = nil
		l.length = 0

		return item, true
	}

	tmpPnode := l.pnode
	item = tmpPnode.item
	l.length--

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

	l.First()
	return item, true
}

// Clear clears the list.
func (l *List) Clear() {
	duplicated := l.avl.duplicated
	*l = NewList(duplicated)
}

// Search searchs the item in the list. It returns the item found and a flag indicating if the item
// exists in the list. This function also move the internal pointer to the item found.
func (l *List) Search(it Item) (Item, bool) {
	node, found := l.avl.Search(&listNode{item: it})
	if found {
		l.pnode = node.(*listNode)
		return l.pnode.item, true
	}

	return nil, false
}

// Length returns the number of items in the list.
func (l *List) Length() int {
	return l.length
}

// ForEach excutes the function of the parameter in all elements of the list, consecutively and
// from the begining.
func (l *List) ForEach(f func(Item)) {
	var item Item

	if l.Length() == 0 {
		// empty list
		return
	}

	oldpnode := l.pnode
	l.First()
	for cont := true; cont; cont = l.Next() {
		item, _ = l.Get()
		f(item)
	}

	l.pnode = oldpnode
}

// Map creates a new list using the results of execute parser function in all elements of the list.
func (l *List)Map(parser func(Item)Item) *List {
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

// Filter create a new list with all elements that pass the test implemented in the filter
// function.
func (l *List)Filter(filter func(Item)bool) *List {
	var (
		newList List
		forFunc func(Item)
	)

	newList = NewList(l.avl.duplicated)
	forFunc = func(it Item) {
		if (filter(it)) {
			newList.AddAfter(it)
		}
	}

	l.ForEach(forFunc)
	return &newList
}
