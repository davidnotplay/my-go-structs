package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkQueueNode(t *testing.T, node *queueNode, item int, next *int) {
	assert.Equal(t, node.item.(IntItem).value, item)

	if next != nil {
		assert.Equal(t, node.next.item.(IntItem).value, *next)
	} else {
		assert.Nil(t, node.next)
	}
}

func Test_NewQueue_func(t *testing.T) {
	qu := NewQueue()

	assert.Nil(t, qu.fnode)
	assert.Nil(t, qu.lnode)
	assert.Equal(t, qu.length, 0)
}

func Test_Enqueue_Queue_func(t *testing.T) {
	qu := NewQueue()

	// insert one item
	qu.Enqueue(It(1))
	checkQueueNode(t, qu.fnode, 1, nil)
	checkQueueNode(t, qu.lnode, 1, nil)

	// insert more items
	for _, a := range []int{2, 3, 4, 5, 6} {
		qu.Enqueue(It(a))
		checkQueueNode(t, qu.fnode, 1, i(2))
		checkQueueNode(t, qu.lnode, a, nil)
	}
}

func Test_Dequeue_Queue_func(t *testing.T) {
	var (
		it       Item
		dequeued bool
	)
	qu := NewQueue()

	// first an empty queue
	it, dequeued = qu.Dequeue()
	assert.Nil(t, it)
	assert.False(t, dequeued)

	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		qu.Enqueue(It(i))
	}

	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		it, dequeued = qu.Dequeue()
		assert.True(t, dequeued)
		assert.Equal(t, it.(IntItem).value, i)
	}

	// Now the queue must be empty.
	it, dequeued = qu.Dequeue()
	assert.Nil(t, it)
	assert.False(t, dequeued)
}

func Test_Front_Queue_func(t *testing.T) {
	qu := NewQueue()

	// Queue empty
	item, cont := qu.Front()
	assert.Nil(t, item)
	assert.False(t, cont)

	for i := 1; i <= 5; i++ {
		qu.Enqueue(It(i))
	}

	for i := 1; qu.Length() > 0; i++ {
		item, cont := qu.Front()
		assert.Equal(t, item.(IntItem).value, i)
		assert.True(t, cont)

		qu.Dequeue()
	}

	assert.Nil(t, item)
	assert.False(t, cont)
}

func Test_Length_Queue_func(t *testing.T) {
	qu := NewQueue()

	for indx, i := range []int{1, 2, 3, 4, 5} {
		assert.Equal(t, qu.Length(), indx)
		qu.Enqueue(It(i))
	}

	for i := qu.Length(); i > 0; i-- {
		assert.Equal(t, qu.Length(), i)
		qu.Dequeue()
	}

	assert.Equal(t, qu.Length(), 0)
	_, popped := qu.Dequeue()
	assert.False(t, popped)
}

func Test_Clear_Queue_func(t *testing.T) {
	qu := NewQueue()

	for i := 1; i <= 5; i++ {
		qu.Enqueue(It(i))
	}

	qu.Clear()

	assert.Nil(t, qu.fnode)
	assert.Nil(t, qu.lnode)
	assert.Equal(t, qu.length, 0)
}
