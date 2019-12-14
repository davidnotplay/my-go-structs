package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func checkStackNode(t *testing.T, node *stackNode, item int, prev *int) {
	as := assert.New(t)

	as.Equal(node.item.(IntItem).value, item, "node value is incorrect")

	if prev != nil {
		if node.prev != nil {
			prevValue := node.prev.item.(IntItem).value
			as.Equal(prevValue, *prev, "value of prev pointer is incorrect")
		} else {
			as.Fail(
				"prev pointer is nil",
				"prev pointer is nil when his values should be: %s",
				prev,
			)
		}
	} else {
		as.Nil(node.prev, "prev pointer isn't nil")
	}
}

func changeStackProperties(stack *Stack, size int, done chan bool) {
	for i := 0; i < size; i++ {
		stack.mutex.Lock()
		length, top := stack.length, stack.top
		stack.length, stack.top = -33, nil
		time.Sleep(time.Nanosecond)
		stack.length, stack.top = length, top
		stack.mutex.Unlock()
	}

	done <- true

}

func Test_NewStack_func(t *testing.T) {
	st := NewStack()
	assert.Nil(t, st.top, "stack top pointer isn't nil when stack is empty")
	assert.Equal(t, st.length, 0, "stack length isn't 0 when stack is empty")
}

func Test_Stack_Push_func(t *testing.T) {
	st := NewStack()

	// insert one item
	st.Push(It(1))
	checkStackNode(t, st.top, 1, nil)

	// insert more items
	for _, a := range []int{2, 3, 4, 5, 6} {
		st.Push(It(a))
		checkStackNode(t, st.top, a, i(a-1))
	}
}

func Test_Stack_Push_func_sync(t *testing.T) {
	as := assert.New(t)
	stack := NewStack()
	concurrence := 8
	size := 1000
	done := make(chan bool)
	insert := func() {
		for i := 0; i < size; i++ {
			stack.Push(It(i))
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go insert()
		go changeStackProperties(&stack, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(stack.Length(), size*concurrence, "stack length is incorrect")
}

func Test_Stack_Pop_func(t *testing.T) {
	var (
		it     Item
		popped bool
	)

	as := assert.New(t)
	st := NewStack()

	// empty list
	it, popped = st.Pop()
	as.Nil(it, "item popped isn't nil in an emtpy stack")
	as.False(popped, "Popped flag is true in an empty stack")

	// insert items
	for a := 1; a <= 5; a++ {
		st.Push(It(a))
	}

	for _, a := range []int{5, 4, 3, 2, 1} {
		it, popped = st.Pop()
		as.Equal(it.(IntItem).value, a, "item returned is incorrect")
		as.True(popped, "popped flag is false when the stack isn't empty")
	}

	it, popped = st.Pop()
	as.Nil(it, "item popped isn't nil in an emtpy stack")
	as.False(popped, "Popped flag is true in an empty stack")
	assert.Nil(t, st.top, "top item isn't nil in empty stack")
}

func Test_Stack_Pop_func_sync(t *testing.T) {
	as := assert.New(t)
	stack := NewStack()
	size := 1000
	concurrence := 8
	done := make(chan bool)
	pop := func() {
		for i := 0; i < size; i++ {
			item, popped := stack.Pop()
			as.NotNil(item, "item returned isn't nil")
			as.True(popped, "popped flag is false when the stack isn't empty ")
		}

		done <- true
	}

	for i := 0; i < concurrence*size; i++ {
		stack.Push(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go pop()
		go changeStackProperties(&stack, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(stack.Length(), 0, "stack isn't empty")
}

func Test_Stack_Top_func(t *testing.T) {
	var (
		it     Item
		exists bool
	)
	st := NewStack()
	as := assert.New(t)

	// empty list
	it, exists = st.Top()
	as.Nil(it, "item of top isn't nil in empty stack")
	as.False(exists, "flag indicating there is an item in top is true when stack is empty")

	// insert items
	for a := 1; a <= 5; a++ {
		st.Push(It(a))
	}

	for _, a := range []int{5, 4, 3, 2, 1} {
		it, exists = st.Top()
		assert.Equal(t, it.(IntItem).value, a, "the item is incorrect")
		as.True(exists, "flag indicating the item exists is incorrect")
		st.Pop()
	}

	it, exists = st.Top()
	as.Nil(it, "item of top isn't nil in empty stack")
	as.False(exists, "flag indicating there is an item in top is true when stack is empty")
}

func Test_Stack_Top_func_sync(t *testing.T) {
	as := assert.New(t)
	stack := NewStack()
	concurrence := 8
	size := 1000
	done := make(chan bool)
	top := func() {
		for i := 0; i < size; i++ {
			item, exists := stack.Top()
			as.Equal(item.(IntItem).value, 99, "item is invalid")
			as.True(exists, "item not found")
		}

		done <- true
	}

	for i := 0; i < 100; i++ {
		stack.Push(It(i))
	}


	for i := 0; i < concurrence; i++ {
		go top()
		go changeStackProperties(&stack, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(stack.Length(), 100, "stack length is invalid")
}

func Test_Stack_Length_func(t *testing.T) {
	stack := NewStack()
	as := assert.New(t)

	for i := 0; i < 10; i++ {
		as.Equal(stack.Length(), i, "stack length is invalid")
		stack.Push(It(i))
	}

	for i := stack.Length(); i >= 0; i-- {
		as.Equal(stack.Length(), i, "stack length is invalid")
		stack.Pop()
	}

	as.Equal(stack.Length(), 0, "stack isn't empty")
}

func Test_Stack_Length_func_sync(t *testing.T) {
	stack := NewStack()
	as := assert.New(t)
	size := 1000
	concurrence := 8
	stackLn := 1000
	done := make(chan bool)
	length := func() {
		for i := 0; i < size; i++ {
			as.Equal(stack.Length(), stackLn, "invalid length")
		}
		done <- true
	}

	for i := 0; i < stackLn; i++ {
		stack.Push(It(1))
	}

	for i := 0; i < concurrence; i++ {
		go length()
		go changeStackProperties(&stack, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}
}

func Test_Clear_Stack_func(t *testing.T) {
	st := NewStack()

	for _, i := range []int{1, 2, 3, 4, 5} {
		st.Push(It(i))
	}

	st.Clear()
	assert.Nil(t, st.top)
	assert.Equal(t, st.length, 0)
}
