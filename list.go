package structs

// List is a structure for saves items sequentially, The struct allows: insert an item in any
// position, get, search or delete any item and iterate for get all elements.
type List struct {
	internalList
}

// NewList creates and returns a new empty List.
func NewList() List{
	return List{}
}

// AddAfter adds the item after the item pointed by internal pointer and moves the internal
//pointer to the new item inserted.
func (l *List) AddAfter(it Item) {
	l.addAfter(it)
}

// AddBefore adds the item before the item pointed by internal pointer and moves the internal
// pointer to the new item inserted.
func (l *List) AddBefore(it Item) {
	l.addBefore(it)
}

// Next moves the internal pointer to the next item, if it possible. Returns true if the pointer
// was moved correctly or false if the internal pointer is in the last item and cannot advance
// more.
func (l *List) Next() bool {
	return l.Next()
}

// Prev moves the internal pointer to the previous item, if it possible. Returns true if the
// pointer was moved correctly or false if the internal pointer is in the first item and cannot
// backwards more.
func (l *List) Prev() bool {
	return l.Prev()
}

// First moves the internal pointer to the first item of the list.
func (l *List) First() {
	l.first()
}

// Last moves the internal pointer to the last item of the list.
func (l *List) Last() {
	l.last()
}

// Get gets the item pointed by the internal pointer. Returns the item and a boolean flag with
// the true value if the item was getted or false if the list is empty.
func (l *List) Get() (Item, bool) {
	return l.get()
}

// Delete deletes the item pointed by the internal pointer. After the pointer backwards
// or advances one position, depending of the situation.
//
// The function returns 2 values: The value deleted and a flag indicating if the value was deleted.
// The second value returned only will be false if it try delete an item in an empty list.
func (l *List) Delete() (Item, bool) {
	return l.delete()
}

// Search searchs the item in the list. It returns the item found and a flag indicating if the item
// exists in the list. This function also move the internal pointer to the item found.
func (l *List) Search(it Item) (Item, bool) {
	return l.search(it)
}

// Length returns the number of items in the list.
func (l *List)Length() int {
	return l.length
}
