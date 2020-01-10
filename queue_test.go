package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func checkQueueNode(t *testing.T, node *queueNode, item int, next *int) {
	as := assert.New(t)
	as.Equal(node.item.(IntItem).value, item, "node value is incorrect")

	if next != nil {
		nvalue := node.next.item.(IntItem).value
		as.Equal(nvalue, *next, "value of next item is incorrect")
	} else {
		if node.next != nil {
			as.Fail("next item isn't nil", "next item isn't nil. item: %s", node.next)
		}
	}
}

func moveQueueProperties(qu *Queue, size int, done chan bool) {
	for i := 0; i < size; i++ {
		qu.mutex.Lock()
		length, fnode, lnode := qu.length, qu.fnode, qu.lnode
		qu.length, qu.fnode, qu.lnode = -1, lnode, fnode
		time.Sleep(time.Nanosecond)
		qu.length, qu.fnode, qu.lnode = length, fnode, lnode
		qu.mutex.Unlock()
	}

	done <- true
}

func Test_NewQueue_func(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()

	as.Nil(queue.fnode, "in empty queue, pointer to first item isn't nil")
	as.Nil(queue.lnode, "in empty queue pointer to last item isn't nil")
	as.Equal(queue.length, 0, "in empty queue, the length isn't 0")
}

func Test_Queue_Enqueue_func(t *testing.T) {
	queue := NewQueue()

	// insert one item
	queue.Enqueue(It(1))
	checkQueueNode(t, queue.fnode, 1, nil)
	checkQueueNode(t, queue.lnode, 1, nil)

	// insert more items
	for _, a := range []int{2, 3, 4, 5, 6} {
		queue.Enqueue(It(a))
		checkQueueNode(t, queue.fnode, 1, i(2))
		checkQueueNode(t, queue.lnode, a, nil)
	}
}

func Test_Queue_Enqueue_func_sync(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()
	size := 2000
	concurrence := 8
	done := make(chan bool)
	enqueue := func() {
		for i := 0; i < size; i++ {
			queue.Enqueue(It(i))
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go enqueue()
		go moveQueueProperties(&queue, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}

	as.Equal(queue.length, size*concurrence, "the length doesn't match")
}

func Test_Queue_Dequeue_func(t *testing.T) {
	var (
		it       Item
		dequeued bool
	)
	queue := NewQueue()
	as := assert.New(t)

	// first an empty queue
	it, dequeued = queue.Dequeue()
	as.Nil(it, "item returned in an empty queue isn't nil")
	as.False(dequeued, "Dequeued flag is true in empty queue")

	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		queue.Enqueue(It(i))
	}

	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		it, dequeued = queue.Dequeue()
		as.Equal(it.(IntItem).value, i, "value returned is incorrect")
		as.True(dequeued, "Dequeued flag is incorrect")
	}

	// Now the queue must be empty.
	it, dequeued = queue.Dequeue()
	as.Nil(it, "item returned in an empty queue isn't nil")
	as.False(dequeued, "Dequeued flag is true in empty queue")
}

func Test_Queue_Dequeue_func_sync(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()
	size := 2000
	concurrence := 8
	done := make(chan bool)
	dequeue := func() {
		for i := 0; i < size; i++ {
			_, dequeued := queue.Dequeue()
			as.True(dequeued, "queue is empty")
		}

		done <- true
	}

	for i := 0; i < size*concurrence+1; i++ {
		queue.Enqueue(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go dequeue()
		go moveQueueProperties(&queue, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}

	as.Equal(queue.Length(), 1, "length is invalid")
	item, dequeued := queue.Dequeue()
	as.Equal(item.(IntItem).value, size*concurrence, "item returned is incorrect")
	as.True(dequeued, "Dequeued flag is incorrect")
	as.Equal(queue.Length(), 0, "length is invalid")
}

func Test_Queue_Front_func(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()

	// Queue empty
	item, exists := queue.Front()
	assert.Nil(t, item, "item returned in empty queue")
	assert.False(t, exists, "exists flag is true in empty queue")

	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	for i := 1; queue.Length() > 0; i++ {
		item, exists = queue.Front()
		as.Equal(item.(IntItem).value, i, "value returned is incorrect")
		as.True(exists, "exists flag is false and the queue isn't empty")

		queue.Dequeue()
	}

	// the queue is empty atagin
	item, exists = queue.Front()
	assert.Nil(t, item, "item returned in empty queue")
	assert.False(t, exists, "exists flag is true in empty queue")
}

func Test_Queue_Front_func_sync(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()
	size := 2000
	concurrence := 8
	done := make(chan bool)
	from := func() {
		for i := 0; i < size; i++ {
			item, exists := queue.Front()
			value := item.(IntItem).value
			as.Equal(value, 0, "item returned is invalid")
			as.True(exists, "exists flag is false and the queue isn't empty")
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		queue.Enqueue(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go from()
		go moveQueueProperties(&queue, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}
}

func Test_Queue_Length_func(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()
	size := 5
	for i := 0; i < size; i++ {
		as.Equal(queue.Length(), i, "length is incorrect")
		queue.Enqueue(It(i))
	}

	for i := size; i >= 0; i-- {
		as.Equal(queue.Length(), i, "length is incorrect")
		queue.Dequeue()
	}

	as.Equal(queue.Length(), 0, "length is incorrect")
}

func Test_Queue_Length_func_sync(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()
	size := 1000
	concurrence := 8
	done := make(chan bool)
	length := func() {
		for i := 0; i < size; i++ {
			as.Equal(queue.Length(), size, "the length is invalid")
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		queue.Enqueue(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go length()
		go moveQueueProperties(&queue, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<-done
		<-done
	}
}

func Test_Queue_Clear_func(t *testing.T) {
	as := assert.New(t)
	queue := NewQueue()

	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	queue.Clear()

	as.Nil(queue.fnode, "pointer to first node isn't nil in empty queue")
	as.Nil(queue.lnode, "pointer to last node isn't nil in empty queue")
	as.Equal(queue.length, 0, "length isn't 0 in empty queue")
}
