package structs

// listNode is the node for the internal list.
type listNode struct {
	prev  *listNode
	next  *listNode
	item  Item
}

// internalList is the internal list struct.
type internalList struct {
	fnode *listNode // pointer to the first node of the list.
	lnode *listNode // ponter to the last node of the list
	pnode *listNode // Internal pointer. It is moved using the struct functions.
	length int      // List size
}

// see(List.addAfter)
func (il *internalList) addAfter(it Item) {
	node := listNode{}
	node.item = it

	if il.fnode == nil {
		// internalList is empty. Add the first node
		il.fnode = &node
		il.lnode = &node
		il.pnode = &node
		il.length++
		return
	}

	node.next = il.pnode.next
	node.prev = il.pnode
	il.pnode.next = &node

	if node.next != nil {
		node.next.prev = &node
	}
	il.length++

	if node.next == nil {
		// the node has been inserted in the last position.
		il.lnode = &node
	}

	il.pnode =  &node
}

// see(List.AddBefore)
func (il *internalList) addBefore(it Item) {
	node := listNode{}
	node.item = it

	if il.fnode == nil {
		// internalList is empty. Add the first node
		il.fnode = &node
		il.lnode = &node
		il.pnode = &node
		il.length++
		return
	}

	node.next = il.pnode
	node.prev = il.pnode.prev
	il.pnode.prev = &node

	if node.prev != nil {
		node.prev.next = &node
	}
	il.length++

	if node.prev == nil {
		// the value inserted is the first.
		il.fnode = &node
	}

	il.pnode = &node
}

// see(List.Next)
func (il *internalList) next() bool{
	if il.pnode.next != nil {
		il.pnode = il.pnode.next
		return true
	}

	return false
}

// see(List.Prev)
func (il *internalList) prev() bool {
	if il.pnode.prev != nil {
		il.pnode = il.pnode.prev
		return true
	}

	return false
}

// see(List.First)
func (il *internalList) first() {
	il.pnode = il.fnode
}

// see(List.Last)
func (il *internalList) last() {
	il.pnode = il.lnode
}


// see(List.Get)
func (il *internalList) get() (Item, bool) {
	if il.pnode != nil {
		return il.pnode.item, true
	}

	return nil, false
}

// see(List.Delete)
func (il *internalList) delete() (Item, bool) {
	var item Item

	if il.length == 0 {
		return item, false
	}

	if il.length == 1 {
		item = il.pnode.item

		// create empty list.
		il.fnode = nil
		il.pnode = nil
		il.lnode = nil
		il.length = 0

		return item, true
	}

	tmpPnode := il.pnode
	item = tmpPnode.item
	il.length--

	if tmpPnode.prev != nil {
		il.pnode = tmpPnode.prev
		tmpPnode.prev.next = tmpPnode.next
	} else {
		//pnode is the first node in the list
		il.fnode = tmpPnode.next
		il.pnode = tmpPnode.next
	}

	if tmpPnode.next != nil {
		tmpPnode.next.prev = tmpPnode.prev
	} else {
		il.lnode = tmpPnode.prev
	}

	return item, true
}

// see(List.Search)
func (il *internalList) search(it Item) (Item, bool) {
	if (il.fnode == nil) {
		return nil, false // list is empty.
	}

	var tmpn *listNode

	for tmpn = il.fnode; tmpn != nil && !tmpn.item.Eq(it); tmpn = tmpn.next {
	}

	if tmpn != nil {
		il.pnode = tmpn
		return tmpn.item, true
	}

	return nil, false
}
