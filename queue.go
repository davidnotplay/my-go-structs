package mygostructs

//queueNode is the node for Queue struct.
type queueNode struct {
	item Item
	next *queueNode
}

// Queue is a struct it implements a queue type abstract data structure, where the items are
// inserted linearly and the first element in enter is the first element in out. (FIFO).
type Queue struct {
	length int
	fnode  *queueNode
	lnode  *queueNode
}

// NewQueue returns a an empty queue.
func NewQueue() Queue {
	return Queue{}
}

// Enqueue adds a new item in the end of the queue.
func (qu *Queue) Enqueue(it Item) {
	node := &queueNode{item: it}

	if qu.length == 0 {
		qu.fnode = node
		qu.lnode = node
	} else {
		qu.lnode.next = node
	}

	qu.lnode = node
	qu.length++
}

// Dequeue returns and delete teh first item of the queue. The second value returned is flag
// indicating the operation was success.
func (qu *Queue) Dequeue() (Item, bool) {
	if qu.length == 0 {
		return nil, false
	}

	node := qu.fnode
	qu.fnode = qu.fnode.next
	qu.length--

	return node.item, true
}

// Front reads the first item in the queue. The second value is a flag indicating if the item
// was read successlly.
func (qu *Queue) Front() (Item, bool) {
	if qu.Length() > 0 {
		return qu.fnode.item, true
	}

	return nil, false
}

// Length returns the number of items in the queue.
func (qu *Queue) Length() int {
	return qu.length
}

// Clear clears the queue.
func (qu *Queue) Clear() {
	*qu = NewQueue()
}
