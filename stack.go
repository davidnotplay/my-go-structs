package mygostructs

// Stack is the classic stack type data structure (LIFO). More info:
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack struct {
	list List
}

// NewStack creates and returns a new empty stack
func NewStack() Stack {
	return Stack{}
}

// Push inserts a the item to top of the stack
func (st *Stack)Push(it Item) {
	st.list.AddAfter(it)
}

// Pop deletes and returns the item in the top of the stack. If the second argment returned is
// false then the stack is empty.
func (st *Stack)Pop() (Item, bool) {
	defer st.list.Last()
	return st.list.Delete()
}

// Length returns the number of items in the stack.
func (st *Stack)Length() int {
	return st.list.Length()
}
