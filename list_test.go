package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkln(t *testing.T, node *listNode, value int, prev, next *int) {
	assert.Equal(t, node.item.(IntItem).value, value, "node item is invalid")

	if prev == nil {
		assert.Nil(t, node.prev, "prev isn't nil")
	} else {
		assert.Equal(t, node.prev.item.(IntItem).value, *prev, "prev item is invalid")
	}

	if next == nil {
		assert.Nil(t, node.next, "next isn't nil")
	} else {
		assert.Equal(t, node.next.item.(IntItem).value, *next, "next item is invalid")
	}
}

func Test_NewList_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	as.Nil(l.fnode)
	as.Nil(l.pnode)
	as.Nil(l.lnode)
	as.Equal(l.length, 0)
}


func Test_List_AddAfter_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	// insert the first element
	l.AddAfter(It(1))

	checkln(t, l.fnode, 1, nil, nil)
	checkln(t, l.pnode, 1, nil, nil)
	checkln(t, l.lnode, 1, nil, nil)
	as.Equal(l.length, 1)

	// Insert more elements
	for _, a := range []int{2, 3, 4, 5} {
		l.addAfter(It(a))

		// first item always is the same
		checkln(t, l.fnode, 1, nil, i(2))

		// the item pointed and the last element is the last inserted.
		checkln(t, l.lnode, a, i(a-1), nil)
		checkln(t, l.pnode, a, i(a-1), nil)
	}

	as.Equal(l.length, 5)

	// check the next and prev pointers.
	for tmpNode, a := l.fnode, 1; a <= l.length; a++ {
		if a == 1 {
			checkln(t, tmpNode, 1, nil, i(2))
		} else if a == l.length {
			checkln(t, tmpNode, 5, i(4), nil)
		} else {
			checkln(t, tmpNode, a, i(a-1), i(a+1))
		}
		tmpNode = tmpNode.next
	}

	l.first()
	l.addAfter(It(12))
	checkln(t, l.fnode, 1, nil, i(12))
	checkln(t, l.pnode, 12, i(1), i(2))
	checkln(t, l.lnode, 5, i(4), nil)
}

func Test_List_AddBefore_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}

	// insert the first element
	il.addBefore(It(5))

	checkln(t, il.fnode, 5, nil, nil)
	checkln(t, il.pnode, 5, nil, nil)
	checkln(t, il.lnode, 5, nil, nil)
	as.Equal(il.length, 1)


	// Insert more elements
	for _, a := range []int{4, 3, 2, 1} {
		il.addBefore(It(a))

		// last item always is the same
		checkln(t, il.lnode, 5, i(4), nil)

		// the item pointed and the first element is the last inserted.
		checkln(t, il.fnode, a, nil, i(a+1))
		checkln(t, il.pnode, a, nil, i(a+1))
	}

	as.Equal(il.length, 5)

	// check the next and prev pointers.
	for tmpNode, a := il.fnode, 1; a <= il.length; a++ {
		if a == 1 {
			checkln(t, tmpNode, 1, nil, i(2))
		} else if a == il.length {
			checkln(t, tmpNode, 5, i(4), nil)
		} else {
			checkln(t, tmpNode, a, i(a-1), i(a+1))
		}
		tmpNode = tmpNode.next
	}

	il.last()
	il.addBefore(It(45))
	checkln(t, il.fnode, 1, nil, i(2))
	checkln(t, il.pnode, 45, i(4), i(5))
	checkln(t, il.lnode, 5, i(45), nil)
}

func Test_List_Next_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		il.addAfter(It(i))
	}

	il.pnode = il.fnode // move the pointer to first item.

	for i := 1; i <= il.length; i++ {
		as.Equal(il.pnode.item.(IntItem).value, i)
		if i < il.length {
			as.True(il.next())
			continue
		}

		// last item
		as.False(il.next()) // In the end of the list. No continue
		as.Equal(il.pnode, il.lnode)
	}

	as.False(il.next())
	as.False(il.next())
}

func Test_List_Prev_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		il.addAfter(It(i))
	}

	il.pnode = il.lnode //move the pointer to the last item.

	for i := il.length; i >= 1 ; i-- {
		as.Equal(il.pnode.item.(IntItem).value, i)
		if i > 1 {
			as.True(il.prev())
			continue
		}

		// last item
		as.False(il.prev()) // In the end of the list. No continue
		as.Equal(il.pnode, il.fnode)
	}

	as.False(il.prev())
	as.False(il.prev())
}

func Test_List_First_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}

	for i := 1; i <= 5; i++ {
		il.addAfter(It(i))
	}

	as.Equal(il.pnode, il.lnode)
	il.first()
	as.Equal(il.pnode, il.fnode)
}

func Test_List_Last_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}

	for i := 1; i <= 5; i++ {
		il.addBefore(It(i))
	}

	as.Equal(il.pnode, il.fnode)
	il.last()
	as.Equal(il.pnode, il.lnode)
}

func Test_List_Get_func(t *testing.T) {
	as := assert.New(t)
	il := internalList{}

	// Get item in an empty list
	v, exists := il.get()
	as.Nil(v)
	as.False(exists)

	for i := 1; i <= 5; i++ {
		il.addAfter(It(i))
	}

	as.Equal(il.length, 5)

	il.first()
	for i := 1; i <= 5; i++ {
		v, exists = il.get()
		as.True(exists)
		as.Equal(v.(IntItem).value, i)
		il.next()
	}
}

func Test_List_Delete_func(t *testing.T) {
	var (
		item   Item
		deleted bool
	)
	as := assert.New(t)
	il := internalList{}

	// test empty list
	item, deleted = il.delete()
	as.Nil(item)
	as.False(deleted)

	// only one element
	il.addAfter(It(3))
	item, deleted = il.delete()

	as.Equal(item.(IntItem).value, 3)
	as.True(deleted)
	as.Nil(il.pnode)
	as.Nil(il.fnode)
	as.Nil(il.lnode)
	as.Equal(il.length, 0)

	il.addAfter(It(1))
	il.addAfter(It(2))
	il.addAfter(It(3))
	il.addAfter(It(4))

	//  remove the element 2
	il.first()
	il.next()
	item, _ = il.get()
	as.Equal(item.(IntItem).value, 2)

	item, deleted = il.delete()

	as.Equal(item.(IntItem).value, 2)
	as.True(deleted)
	as.Equal(il.length, 3)
	checkln(t, il.pnode, 1, nil, i(3))
	checkln(t, il.fnode, 1, nil, i(3))
	checkln(t, il.lnode, 4, i(3), nil)

	// remove the element 3
	il.next()
	as.Equal(il.pnode.item.(IntItem).value, 3)
	item, deleted = il.delete()

	as.Equal(item.(IntItem).value, 3)
	as.True(deleted)
	as.Equal(il.length, 2)

	checkln(t, il.pnode, 1, nil, i(4))
	checkln(t, il.fnode, 1, nil, i(4))
	checkln(t, il.lnode, 4, i(1), nil)

	// remove the element 4
	il.next()
	as.Equal(il.pnode.item.(IntItem).value, 4)
	item, deleted = il.delete()
	as.Equal(item.(IntItem).value, 4)
	as.True(deleted)
	as.Equal(il.length, 1)

	checkln(t, il.pnode, 1, nil, nil)
	checkln(t, il.fnode, 1, nil, nil)
	checkln(t, il.lnode, 1, nil, nil)

	// insert 2 and remove element 1
	il.addAfter(It(2))
	as.Equal(il.pnode.item.(IntItem).value, 2)
	as.Equal(il.lnode.item.(IntItem).value, 2)

	il.first()
	item, deleted = il.delete()
	as.Equal(item.(IntItem).value, 1)
	as.True(deleted)
	as.Equal(il.length, 1)

	checkln(t, il.pnode, 2, nil, nil)
	checkln(t, il.fnode, 2, nil, nil)
	checkln(t, il.lnode, 2, nil, nil)
}

func Test_List_Search_func(t *testing.T) {
	var (
		item  Item
		found bool
	)
	as := assert.New(t)
	il := internalList{}

	// First checks an empty list
	item, found = il.search(It(1))
	as.Nil(item)
	as.False(found)

	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		il.addAfter(It(a))
	}


	// Check the pointers are correct
	checkln(t, il.fnode, 1, nil, i(2))
	checkln(t, il.pnode, 6, i(5), nil)
	checkln(t, il.lnode, 6, i(5), nil)

	// Search elements they are in the list
	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		item, found := il.search(It(a))
		as.Equal(item.(IntItem).value, a)
		as.True(found)

		// Check the pointers are correct
		checkln(t, il.fnode, 1, nil, i(2))
		checkln(t, il.lnode, 6, i(5), nil)

		switch {
		case a == 1:
			checkln(t, il.pnode, 1, nil, i(2))
		case a == 6:
			checkln(t, il.pnode, 6, i(5), nil)
		default:
			checkln(t, il.pnode, a, i(a-1), i(a+1))
		}

	}

	// Search items they arent' in the list
	for _, a := range []int{7, 8, 9, 10, -1, -2} {
		item, found := il.search(It(a))
		as.Nil(item)
		as.False(found)

		// Check the pointers are correct
		checkln(t, il.fnode, 1, nil, i(2))
		checkln(t, il.pnode, 6, i(5), nil)
		checkln(t, il.lnode, 6, i(5), nil)
	}
}

func Test_List_Length_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	as.Equal(l.Length(), 0)

	for i, a := range []int{1, 2, 3, 4, 5, 6} {
		l.AddAfter(It(a))
		as.Equal(l.Length(), i + 1)
	}
}
