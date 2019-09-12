package mygostructs

//queueNode is the node for Queue struct.
type queueNode struct {
	item Item
	next *queueNode
}

// Queue is the classic queue type data structure (FIFO). More info:
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Queue struct {
	length int
	fnode  *queueNode
	lnode   *queueNode
}

// NewQueue returns a an empty queue.
func NewQueue() Queue {
	return Queue{}
}

// Enqueue adds a new item in the queue.
func (qu *Queue) Enqueue(it Item) {
	node := &queueNode{item:it}

	if qu.length == 0 {
		qu.fnode = node
		qu.lnode = node
	} else {
		qu.lnode.next = node
	}

	qu.lnode = node
	qu.length++
}

// Dequeue gets the first item of the queue. The second value returned, is a flag indicating if
// was possible fetch the first element or if the Queue was empty.
func (qu *Queue) Dequeue() (Item, bool) {
	if qu.length == 0 {
		return nil, false
	}

	node := qu.fnode
	qu.fnode = qu.fnode.next
	qu.length--

	return node.item, true
}

// Front reads the first element in the queue, if the que isn't empty. The second value returned,
// is a flag indicating if element was read, true, or if the Queue is empty, false.
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
