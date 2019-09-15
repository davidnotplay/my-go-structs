package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkStackNode(t *testing.T, node *stackNode, item int, next *int) {
	assert.Equal(t, node.item.(IntItem).value, item)

	if next != nil {
		assert.Equal(t, node.prev.item.(IntItem).value, *next)
	} else {
		assert.Nil(t, node.prev)
	}
}

func Test_NewStack_func(t *testing.T) {
	st := NewStack()
	assert.Nil(t, st.top)
	assert.Equal(t, st.length, 0)
}

func Test_Push_Stack_func(t *testing.T) {
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

func Test_Pop_Stack_func(t *testing.T) {
	var (
		it     Item
		popped bool
	)
	st := NewStack()

	// empty list
	it, popped = st.Pop()
	assert.Nil(t, it)
	assert.False(t, popped)

	// insert items
	for a := 1; a <= 5; a++ {
		st.Push(It(a))
	}

	for _, a := range []int{5, 4, 3, 2, 1} {
		it, popped = st.Pop()
		assert.True(t, popped)
		assert.Equal(t, it.(IntItem).value, a)
	}

	it, popped = st.Pop()
	assert.Nil(t, it)
	assert.False(t, popped)
	assert.Nil(t, st.top)
}

func Test_Top_Stack_func(t *testing.T) {
	var (
		it     Item
		exists bool
	)
	st := NewStack()

	// empty list
	it, exists = st.Top()
	assert.Nil(t, it)
	assert.False(t, exists)

	// insert items
	for a := 1; a <= 5; a++ {
		st.Push(It(a))
	}

	for _, a := range []int{5, 4, 3, 2, 1} {
		it, exists = st.Top()
		assert.True(t, exists)
		assert.Equal(t, it.(IntItem).value, a)
		st.Pop()
	}

	it, exists = st.Top()
	assert.Nil(t, it)
	assert.False(t, exists)
	assert.Nil(t, st.top)
}

func Test_Length_Stack_func(t *testing.T) {
	st := NewStack()

	for indx, i := range []int{1, 2, 3, 4, 5} {
		assert.Equal(t, st.Length(), indx)
		st.Push(It(i))
	}

	for i := st.Length(); i > 0; i-- {
		assert.Equal(t, st.Length(), i)
		st.Pop()
	}

	assert.Equal(t, st.Length(), 0)
	_, popped := st.Pop()
	assert.False(t, popped)
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
