package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewStack_func(t *testing.T) {
	st := NewStack()

	assert.Nil(t, st.list.fnode)
	assert.Nil(t, st.list.pnode)
	assert.Nil(t, st.list.lnode)
	assert.Equal(t, st.list.length, 0)
}

func Test_Push_Stack_func(t *testing.T) {
	st := NewStack()

	// insert one item
	st.Push(It(1))
	checkln(t, st.list.fnode, 1, nil, nil)
	checkln(t, st.list.pnode, 1, nil, nil)
	checkln(t, st.list.lnode, 1, nil, nil)
	assert.Equal(t, st.list.length, 1)

	// insert more items
	for _, a := range []int{2, 3, 4, 5, 6} {
		st.Push(It(a))
		checkln(t, st.list.fnode, 1, nil, i(2))
		checkln(t, st.list.pnode, a, i(a-1), nil)
		checkln(t, st.list.lnode, a, i(a-1), nil)
		assert.Equal(t, st.list.length, a)
	}
}

func Test_Pop_Stack_func(t *testing.T) {
	var (
		it Item
		popped bool
	)
	st := NewStack()

	// empty list
	it, popped = st.Pop()
	assert.Nil(t, it)
	assert.False(t, popped)

	// insert items
	for a := 1; a<=5; a++ {
		st.Push(It(a))
	}

	for _, a := range []int{5, 4, 3, 2, 1} {
		// first check pointers
		assert.Equal(t, st.list.pnode, st.list.lnode)

		it, popped = st.Pop()
		assert.True(t, popped)
		assert.Equal(t, it.(IntItem).value, a)

	}

	it, popped = st.Pop()
	assert.Nil(t, it)
	assert.False(t, popped)

	assert.Nil(t, st.list.fnode)
	assert.Nil(t, st.list.pnode)
	assert.Nil(t, st.list.lnode)
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
