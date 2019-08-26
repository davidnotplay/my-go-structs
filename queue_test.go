package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewQueue_func(t *testing.T) {
	qu := NewQueue()

	assert.Nil(t, qu.fnode)
	assert.Nil(t, qu.pnode)
	assert.Nil(t, qu.lnode)
	assert.Equal(t, qu.length, 0)
}

func Test_Enqueue_Queue_func(t *testing.T) {
	qu := NewQueue()

	// insert one element
	qu.Enqueue(It(1))
	checkln(t, qu.fnode, 1, nil, nil)
	checkln(t, qu.pnode, 1, nil, nil)
	checkln(t, qu.lnode, 1, nil, nil)
	assert.Equal(t, qu.length, 1)

	// insert more elements
	for _, a := range []int{2, 3, 4, 5, 6} {
		qu.Enqueue(It(a))
		checkln(t, qu.fnode, 1, nil, i(2))
		checkln(t, qu.pnode, a, i(a-1), nil)
		checkln(t, qu.lnode, a, i(a-1), nil)
		assert.Equal(t, qu.length, a)
	}
}

func Test_Dequeue_Queue_func(t *testing.T) {
	var (
		it	 Item
		dequeued bool
	)
	qu := NewQueue()

	// first an empty queue
	it, dequeued = qu.Dequeue()
	assert.Nil(t, it)
	assert.False(t, dequeued)

	// insert one element
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
