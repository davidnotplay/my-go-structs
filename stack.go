package mygostructs

//stackNode is the node for Stack struct.
type stackNode struct {
	item Item
	prev *stackNode
}

// Stack is the classic stack type data structure (LIFO). More info:
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack struct {
	top    *stackNode
	length int
}

// NewStack creates and returns a new empty stack
func NewStack() Stack {
	return Stack{}
}

// Push inserts a the item to top of the stack
func (st *Stack) Push(it Item) {
	node := &stackNode{item:it}
	node.prev = st.top
	st.top = node
	st.length++
}

// Pop deletes and returns the item in the top of the stack. If the second argment returned is
// false then the stack is empty.
func (st *Stack) Pop() (Item, bool) {
	if st.length == 0 {
		return nil, false
	}

	node := st.top
	st.top = node.prev
	st.length--

	return node.item, true
}

// Length returns the number of items in the stack.
func (st *Stack) Length() int {
	return st.length
}
