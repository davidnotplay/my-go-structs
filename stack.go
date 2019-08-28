package mygostructs

// Stack is the classic stack type data structure (LIFO). More info:
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack struct {
	internalList
}

// NewStack creates and returns a new empty stack
func NewStack() Stack {
	return Stack{}
}

// Push inserts a the item to top of the stack
func (st *Stack)Push(it Item) {
	st.addAfter(it)
}

// Pop deletes and returns the item in the top of the stack. If the second argment returned is
// false then the stack is empty.
func (st *Stack)Pop() (Item, bool) {
	defer st.last()
	return st.delete()
}

// Length returns the number of items in the stack.
func (st *Stack)Length() int {
	return st.length
}
