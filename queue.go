package mygostructs

// Queue is the classic queue type data structure (FIFO). More info:
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Queue struct {
	internalList
}

// NewQueue returns a an empty queue.
func NewQueue() Queue{
	return Queue{}
}

// Enqueue adds a new item in the queue.
func (qu *Queue) Enqueue (it Item) {
	qu.last()
	qu.addAfter(it)
}

// Dequeue gets the first item of the queue.
func (qu *Queue) Dequeue() (Item, bool) {
	qu.first()
	return qu.delete()
}

// Length returns the number of items in the queue.
func (qu *Queue) Length() int {
	return qu.length
}
