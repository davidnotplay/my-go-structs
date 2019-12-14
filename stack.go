package mygostructs

import "sync"

//stackNode is the node for Stack struct.
type stackNode struct {
	item Item
	prev *stackNode
}

// Stack is a struct it implements a stack type abstract data structure, where the items are
// inserted linearly and the last item in enter is the first in out. (LIFO)
type Stack struct {
	top    *stackNode
	length int
	mutex sync.Mutex
}

// NewStack creates and returns a new empty stack
func NewStack() Stack {
	return Stack{}
}

// Push inserts a the item to top of the stack
func (st *Stack) Push(it Item) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	node := &stackNode{item: it}
	node.prev = st.top
	st.top = node
	st.length++
}

// Pop deletes and returns the item in the top of the stack. If the second argment returned is
// false then the stack is empty.
func (st *Stack) Pop() (Item, bool) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	if st.length == 0 {
		return nil, false
	}

	node := st.top
	st.top = node.prev
	st.length--

	return node.item, true
}

// Top reads the top item in the stack. The second value returned is false if the stack is
// empty.
func (st *Stack) Top() (Item, bool) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	if st.length > 0 {
		return st.top.item, true
	}

	return nil, false
}

// Length returns the number of items in the stack.
func (st *Stack) Length() int {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	return st.length
}

// Clear clears the stack.
func (st *Stack) Clear() {
	st.mutex.Lock()
	defer st.mutex.Unlock()
	st.top, st.length = nil, 0
}
