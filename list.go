package mygostructs

// listNode is the node for the List struct.
type listNode struct {
	prev  *listNode
	next  *listNode
	item  Item
}

// List is a doubly linked list type data structure. More info:
// https://en.wikipedia.org/wiki/Doubly_linked_list
type List struct {
	fnode *listNode // pointer to the first node of the list.
	lnode *listNode // ponter to the last node of the list
	pnode *listNode // Internal pointer. It is moved using the struct functions.
	length int      // List size
}


// NewList returns an empty List.
func NewList() List{
	return List{}
}

// AddAfter adds the item after the item pointed by internal pointer and moves the internal
// pointer to the new item inserted.
func (l *List) AddAfter(it Item) {
	node := listNode{}
	node.item = it

	if l.fnode == nil {
		// List is empty. Add the first node
		l.fnode = &node
		l.lnode = &node
		l.pnode = &node
		l.length++
		return
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

	l.pnode =  &node
}

// AddBefore adds the item before the item pointed by internal pointer and moves the internal
// pointer to the new item inserted.
func (l *List) AddBefore(it Item) {
	node := listNode{}
	node.item = it

	if l.fnode == nil {
		// List is empty. Add the first node
		l.fnode = &node
		l.lnode = &node
		l.pnode = &node
		l.length++
		return
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

// Get gets the item pointed by the internal pointer. Returns the item and a boolean flag with
// the true value if the item was getted or false if the list is empty.
func (l *List) Get() (Item, bool) {
	if l.pnode != nil {
		return l.pnode.item, true
	}

	return nil, false

}

// Delete deletes the item pointed by the internal pointer and move the internal pointer to the
// first of the list. It returns 2 values: The item deleted and a flag indicating if the value
// was deleted.
func (l *List) Delete() (Item, bool) {
	var item Item

	if l.length == 0 {
		return item, false
	}

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

// Search searchs the item in the list. It returns the item found and a flag indicating if the item
// exists in the list. This function also move the internal pointer to the item found.
func (l *List) Search(it Item) (Item, bool) {
	if (l.fnode == nil) {
		return nil, false // list is empty.
	}

	var tmpn *listNode

	for tmpn = l.fnode; tmpn != nil && !tmpn.item.Eq(it); tmpn = tmpn.next {
	}

	if tmpn != nil {
		l.pnode = tmpn
		return tmpn.item, true
	}

	return nil, false
}

// Length returns the number of items in the list.
func (l *List)Length() int {
	return l.length
}
