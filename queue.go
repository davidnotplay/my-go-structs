package mygostructs

import "sync"

//queueNode is the node for Queue struct.
type queueNode struct {
	item Item
	next *queueNode
}

// Queue is a struct it implements a queue type abstract data structure, where the items are
// inserted linearly and the first element in enter is the first element in out. (FIFO).
//
// The struct is adapted to run in multithread code.
type Queue struct {
	length int
	fnode  *queueNode
	lnode  *queueNode
	mutex  sync.Mutex
}

// NewQueue creates and returns a new empty queue.
func NewQueue() Queue {
	return Queue{}
}

// Enqueue adds the item of the paramter in the end of the queue.
func (qu *Queue) Enqueue(it Item) {
	qu.mutex.Lock()
	defer qu.mutex.Unlock()

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

// Dequeue returns and delete the first item of the queue. The second value returned is flag
// indicating the operation was success.
func (qu *Queue) Dequeue() (Item, bool) {
	qu.mutex.Lock()
	defer qu.mutex.Unlock()

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
	qu.mutex.Lock()
	defer qu.mutex.Unlock()

	if qu.length > 0 {
		return qu.fnode.item, true
	}

	return nil, false
}

// Length returns the number of items in the queue.
func (qu *Queue) Length() int {
	qu.mutex.Lock()
	defer qu.mutex.Unlock()
	return qu.length
}

// Clear clears the queue.
func (qu *Queue) Clear() {
	qu.mutex.Lock()
	defer qu.mutex.Unlock()
	qu.fnode, qu.lnode, qu.length = nil, nil, 0
}
