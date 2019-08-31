package mygostructs

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

	// insert the first item
	l.AddAfter(It(1))

	checkln(t, l.fnode, 1, nil, nil)
	checkln(t, l.pnode, 1, nil, nil)
	checkln(t, l.lnode, 1, nil, nil)
	as.Equal(l.length, 1)

	// Insert more items
	for _, a := range []int{2, 3, 4, 5} {
		l.AddAfter(It(a))

		// first item always is the same
		checkln(t, l.fnode, 1, nil, i(2))

		// the item pointed and the last item is the last inserted.
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

	l.First()
	l.AddAfter(It(12))
	checkln(t, l.fnode, 1, nil, i(12))
	checkln(t, l.pnode, 12, i(1), i(2))
	checkln(t, l.lnode, 5, i(4), nil)
}

func Test_List_AddBefore_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	// insert the First item
	l.AddBefore(It(5))

	checkln(t, l.fnode, 5, nil, nil)
	checkln(t, l.pnode, 5, nil, nil)
	checkln(t, l.lnode, 5, nil, nil)
	as.Equal(l.length, 1)


	// Insert more items
	for _, a := range []int{4, 3, 2, 1} {
		l.AddBefore(It(a))

		// last item always is the same
		checkln(t, l.lnode, 5, i(4), nil)

		// the item pointed and the First item is the last inserted.
		checkln(t, l.fnode, a, nil, i(a+1))
		checkln(t, l.pnode, a, nil, i(a+1))
	}

	as.Equal(l.length, 5)

	// check the Next and prev pointers.
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

	l.Last()
	l.AddBefore(It(45))
	checkln(t, l.fnode, 1, nil, i(2))
	checkln(t, l.pnode, 45, i(4), i(5))
	checkln(t, l.lnode, 5, i(45), nil)
}

func Test_List_Next_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		l.AddAfter(It(i))
	}

	l.pnode = l.fnode // move the pointer to first item.

	for i := 1; i <= l.length; i++ {
		as.Equal(l.pnode.item.(IntItem).value, i)
		if i < l.length {
			as.True(l.Next())
			continue
		}

		// last item
		as.False(l.Next()) // In the end of the list. No continue
		as.Equal(l.pnode, l.lnode)
	}

	as.False(l.Next())
	as.False(l.Next())
}

func Test_List_Prev_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		l.AddAfter(It(i))
	}

	l.pnode = l.lnode //move the pointer to the last item.

	for i := l.length; i >= 1 ; i-- {
		as.Equal(l.pnode.item.(IntItem).value, i)
		if i > 1 {
			as.True(l.Prev())
			continue
		}

		// last item
		as.False(l.Prev()) // In the end of the list. No continue
		as.Equal(l.pnode, l.fnode)
	}

	as.False(l.Prev())
	as.False(l.Prev())
}

func Test_List_First_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	for i := 1; i <= 5; i++ {
		l.AddAfter(It(i))
	}

	as.Equal(l.pnode, l.lnode)
	l.First()
	as.Equal(l.pnode, l.fnode)
}

func Test_List_Last_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	for i := 1; i <= 5; i++ {
		l.AddBefore(It(i))
	}

	as.Equal(l.pnode, l.fnode)
	l.Last()
	as.Equal(l.pnode, l.lnode)
}

func Test_List_Get_func(t *testing.T) {
	as := assert.New(t)
	l := NewList()

	// Get item in an empty list
	v, exists := l.Get()
	as.Nil(v)
	as.False(exists)

	for i := 1; i <= 5; i++ {
		l.AddAfter(It(i))
	}

	as.Equal(l.length, 5)

	l.First()
	for i := 1; i <= 5; i++ {
		v, exists = l.Get()
		as.True(exists)
		as.Equal(v.(IntItem).value, i)
		l.Next()
	}
}

func Test_List_Delete_func(t *testing.T) {
	var (
		item   Item
		deleted bool
	)
	as := assert.New(t)
	l := NewList()

	// test empty list
	item, deleted = l.Delete()
	as.Nil(item)
	as.False(deleted)

	// only one item
	l.AddAfter(It(3))
	item, deleted = l.Delete()

	as.Equal(item.(IntItem).value, 3)
	as.True(deleted)
	as.Nil(l.pnode)
	as.Nil(l.fnode)
	as.Nil(l.lnode)
	as.Equal(l.length, 0)


	for i := 1; i <= 10; i++ {
		l.AddAfter(It(i))
	}

	maxLength := l.Length()
	for i, a := range []int{3, 4, 6, 9, 1, 2, 7, 5, 10, 8} {
		_, found := l.Search(It(a))
		as.True(found)

		vdeleted, deleted := l.Delete()
		as.True(deleted)
		as.Equal(vdeleted.(IntItem).value, a)
		as.Equal(l.pnode, l.fnode)
		as.Equal(l.Length(), maxLength - (i+1))

		_, found = l.Search(It(a))
		as.False(found)
	}

	as.Equal(l.Length(), 0)
	as.Nil(l.fnode)
	as.Nil(l.pnode)
	as.Nil(l.lnode)
}

func Test_List_Search_func(t *testing.T) {
	var (
		item  Item
		found bool
	)
	as := assert.New(t)
	l := NewList()

	// First checks an empty list
	item, found = l.Search(It(1))
	as.Nil(item)
	as.False(found)

	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		l.AddAfter(It(a))
	}


	// Check the pointers are correct
	checkln(t, l.fnode, 1, nil, i(2))
	checkln(t, l.pnode, 6, i(5), nil)
	checkln(t, l.lnode, 6, i(5), nil)

	// Search items they are in the list
	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		item, found := l.Search(It(a))
		as.Equal(item.(IntItem).value, a)
		as.True(found)

		// Check the pointers are correct
		checkln(t, l.fnode, 1, nil, i(2))
		checkln(t, l.lnode, 6, i(5), nil)

		switch {
		case a == 1:
			checkln(t, l.pnode, 1, nil, i(2))
		case a == 6:
			checkln(t, l.pnode, 6, i(5), nil)
		default:
			checkln(t, l.pnode, a, i(a-1), i(a+1))
		}

	}

	// Search items they arent' in the list
	for _, a := range []int{7, 8, 9, 10, -1, -2} {
		item, found := l.Search(It(a))
		as.Nil(item)
		as.False(found)

		// Check the pointers are correct
		checkln(t, l.fnode, 1, nil, i(2))
		checkln(t, l.pnode, 6, i(5), nil)
		checkln(t, l.lnode, 6, i(5), nil)
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
