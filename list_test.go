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

	l := NewList(true)
	as.Nil(l.fnode)
	as.Nil(l.pnode)
	as.Nil(l.lnode)
	as.Equal(l.length, 0)
	as.Equal(l.avl.length, 0)
	as.True(l.avl.duplicated)
	as.True(l.avl.rebalance)

	l = NewList(false)
	as.Nil(l.fnode)
	as.Nil(l.pnode)
	as.Nil(l.lnode)
	as.Equal(l.length, 0)
	as.Equal(l.avl.length, 0)
	as.False(l.avl.duplicated)
	as.True(l.avl.rebalance)
}


func Test_List_AddAfter_func(t *testing.T) {
	var (
		inserted bool
		found    bool
		node     *listNode
		values   = []int{1, 2, 2, 3, 4, 5, 5}
		ndvalues = []int{1, 2, 3, 4, 5}
		as	 = assert.New(t)
		l	 = NewList(true)
	)

	for indx, value := range values {
		// Test value inserted
		inserted = l.AddAfter(It(value))
		as.True(inserted)

		// Test the internal pointers are corrects.
		if indx == 0 {
			checkln(t, l.fnode, values[0], nil, nil)
			checkln(t, l.pnode, values[0], nil, nil)
			checkln(t, l.lnode, values[0], nil, nil)

		} else {
			checkln(t, l.fnode, values[0], nil, i(values[1]))
			checkln(t, l.lnode, values[indx], i(values[indx - 1]), nil)
			checkln(t, l.pnode, values[indx], i(values[indx - 1]), nil)
		}

		// node saved in the avl
		_, found = l.avl.Search(l.pnode)
		as.True(found)
	}

	// next pointers.
	node = l.fnode
	for i := 0; i < len(values); i++ {
		// as.Equal(node.item.(IntItem).value, values[i])
		as.Equal(node.item.(IntItem).value, values[i])
		node = node.next
	}
	as.Nil(node)

	// prev pointers.
	node = l.lnode
	for i := len(values) - 1; i >= 0; i-- {
		as.Equal(node.item.(IntItem).value, values[i])
		node = node.prev
	}
	as.Nil(node)


	// No duplicates items.
	l = NewList(false)

	for indx, value := range values {
		// Test value inserted
		inserted = l.AddAfter(It(value))
		// 2 and 6 index with items duplicated.
		as.Equalf(inserted, indx != 2 && indx != 6, "index %d", indx)

		// node saved in the avl
		_, found = l.avl.Search(l.pnode)
		as.True(found)
	}

	checkln(t, l.fnode, 1, nil, i(2))
	checkln(t, l.pnode, 5, i(4), nil)
	checkln(t, l.lnode, 5, i(4), nil)

	// next pointers.
	node = l.fnode
	for i := 0; i < len(ndvalues); i++ {
		// as.Equal(node.item.(IntItem).value, values[i])
		as.Equal(node.item.(IntItem).value, ndvalues[i])
		node = node.next
	}
	as.Nil(node)

	// prev pointers.
	node = l.lnode
	for i := len(ndvalues) - 1; i >= 0; i-- {
		as.Equal(node.item.(IntItem).value, ndvalues[i])
		node = node.prev
	}
	as.Nil(node)
}

func Test_List_AddBefore_func(t *testing.T) {
	var (
		inserted bool
		found    bool
		node     *listNode
		values   = []int{1, 2, 2, 3, 4, 5, 5}
		ndvalues = []int{1, 2, 3, 4, 5}
		as	 = assert.New(t)
		l	 = NewList(true)
	)

	for indx, value := range values {
		// Test value inserted
		inserted = l.AddBefore(It(value))
		as.True(inserted)

		// Test the internal pointers are corrects.
		if indx == 0 {
			checkln(t, l.fnode, values[0], nil, nil)
			checkln(t, l.pnode, values[0], nil, nil)
			checkln(t, l.lnode, values[0], nil, nil)

		} else {
			checkln(t, l.fnode, values[indx], nil, i(values[indx-1]))
			checkln(t, l.pnode, values[indx], nil, i(values[indx-1]))
			checkln(t, l.lnode, values[0], i(values[1]), nil)
		}

		// node saved in the avl
		_, found = l.avl.Search(l.pnode)
		as.True(found)
	}

	// next pointers.
	node = l.lnode
	for i := 0; i < len(values); i++ {
		// as.Equal(node.item.(IntItem).value, values[i])
		as.Equal(node.item.(IntItem).value, values[i])
		node = node.prev
	}
	as.Nil(node)

	// prev pointers.
	node = l.fnode
	for i := len(values) - 1; i >= 0; i-- {
		as.Equal(node.item.(IntItem).value, values[i])
		node = node.next
	}
	as.Nil(node)


	// No duplicates items.
	l = NewList(false)

	for indx, value := range values {
		// Test value inserted
		inserted = l.AddBefore(It(value))
		// 2 and 6 index with items duplicated.
		as.Equalf(inserted, indx != 2 && indx != 6, "index %d", indx)

		// node saved in the avl
		_, found = l.avl.Search(l.pnode)
		as.True(found)
	}

	checkln(t, l.fnode, 5, nil, i(4))
	checkln(t, l.pnode, 5, nil, i(4))
	checkln(t, l.lnode, 1, i(2), nil)

	// next pointers.
	node = l.lnode
	for i := 0; i < len(ndvalues); i++ {
		// as.Equal(node.item.(IntItem).value, values[i])
		as.Equal(node.item.(IntItem).value, ndvalues[i])
		node = node.prev
	}
	as.Nil(node)

	// prev pointers.
	node = l.fnode
	for i := len(ndvalues) - 1; i >= 0; i-- {
		as.Equal(node.item.(IntItem).value, ndvalues[i])
		node = node.next
	}
	as.Nil(node)
}

func Test_List_Next_func(t *testing.T) {
	as := assert.New(t)
	l := NewList(true)
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
	l := NewList(true)
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
	l := NewList(true)

	for i := 1; i <= 5; i++ {
		l.AddAfter(It(i))
	}

	as.Equal(l.pnode, l.lnode)
	l.First()
	as.Equal(l.pnode, l.fnode)
}

func Test_List_Last_func(t *testing.T) {
	as := assert.New(t)
	l := NewList(true)

	for i := 1; i <= 5; i++ {
		l.AddBefore(It(i))
	}

	as.Equal(l.pnode, l.fnode)
	l.Last()
	as.Equal(l.pnode, l.lnode)
}

func Test_List_Get_func(t *testing.T) {
	as := assert.New(t)
	l := NewList(true)

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
	l := NewList(true)

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

	// insert items duplicated
	l.AddAfter(It(2))
	l.AddAfter(It(4))
	l.AddAfter(It(7))

	maxLength := l.Length()
	for i, a := range []int{3, 2, 4, 6, 9, 1, 2, 7, 5, 10, 8, 7, 4} {
		_, found := l.Search(It(a))
		as.True(found)

		vdeleted, deleted := l.Delete()
		as.True(deleted)
		as.Equal(vdeleted.(IntItem).value, a)
		as.Equal(l.pnode, l.fnode)
		as.Equal(l.Length(), maxLength - (i+1))

		_, found = l.Search(It(a))
		as.Equal(found, i == 1 || i == 2 || i == 7)
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
	l := NewList(true)

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
	l := NewList(true)

	as.Equal(l.Length(), 0)

	for i, a := range []int{1, 2, 3, 4, 5, 6} {
		l.AddAfter(It(a))
		as.Equal(l.Length(), i + 1)
	}
}
