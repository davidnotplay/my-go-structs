package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// ip returns a pointer of the integer.
func ip(i int) *int {
	return &i
}

// checkln checks the node properties
func checkln(t *testing.T, node *listNode, value int, prev, next *int) {
	assert.Equal(t, node.item.(IntItem).value, value, "item doesn't match")

	if prev == nil {
		assert.Nil(t, node.prev, "prev item isn't nil")
	} else {
		if node.prev == nil {
			assert.Fail(t, "item is nil", "prev item is nil")
		} else {
			value := node.prev.item.(IntItem).value
			assert.Equal(t, value, *prev, "prev item doesn't match")
		}
	}

	if next == nil {
		assert.Nil(t, node.next, "next item isnt't nil")
	} else {
		if node.next == nil {
			assert.Fail(t, "item is nil", "next item is nil")
		} else {
			value := node.next.item.(IntItem).value
			assert.Equal(t, value, *next, "next item doesn't match")
		}
	}
}

// moveNextPrev changes the pointer and after restore it
func changeListProperties(l *List, max int, done chan bool) {
	for i := 0; i < max; i++ {
		l.mutex.Lock()
		fnode, pnode, lnode, length := l.fnode, l.pnode, l.lnode, l.avl.length
		l.fnode, l.pnode, l.lnode, l.avl.length = nil, nil, nil, 0
		l.fnode, l.pnode, l.lnode, l.avl.length = pnode, lnode, fnode, -1
		time.Sleep(time.Nanosecond)
		l.fnode, l.pnode, l.lnode, l.avl.length = fnode, pnode, lnode, length
		l.mutex.Unlock()
	}

	done <- true
}

func Test_listNode_Less_func(t *testing.T) {
	var n1, n2 listNode
	as := assert.New(t)
	n1 = listNode{nil, nil, It(1)}

	n2 = listNode{nil, nil, It(2)}
	as.True(n1.Less(&n2), "%s isnt't less than %s", n1, n2)
	n2 = listNode{nil, nil, It(1)}
	as.False(n1.Less(&n2), "%s is less than %s", n1, n2)
	n2 = listNode{nil, nil, It(0)}
	as.False(n1.Less(&n2), "%s is less than %s", n1, n2)
}

func Test_listNode_Eq_func(t *testing.T) {
	var n1, n2 listNode
	as := assert.New(t)
	n1 = listNode{nil, nil, It(1)}

	n2 = listNode{nil, nil, It(1)}
	as.True(n1.Eq(&n2), "%s isn't equal to %s", n1, n2)
	n2 = listNode{nil, nil, It(0)}
	as.False(n1.Eq(&n2), "%s is equal to %s", n1, n2)
}

func Test_listNode_String_func(t *testing.T) {
	assert.Equalf(t, listNode{nil, nil, It(1)}.String(), "1", "item stringify is invalid")
}

func Test_NewList_func(t *testing.T) {
	as := assert.New(t)

	for _, duplicated := range []bool{true, false} {
		list := NewList(duplicated)

		as.Nil(list.fnode, "pointer to first node isn't nil in empty list")
		as.Nil(list.pnode, "pointer to current node isn't nil in empty list")
		as.Nil(list.lnode, "pointer to last node isn't nil in empty list")
		as.Nil(list.avl.root, "avl root isn't nil in empty tree")
		as.Equal(list.avl.length, 0, "avl length isn't 0 in empty tree")
		as.True(list.avl.rebalance, "avl rebalance flag, is false")
		as.Equal(list.avl.duplicated, duplicated, "avl duplicated flag is invalid")
	}
}

func Test_List_AddAfter_func(t *testing.T) {
	var (
		inserted bool
		counter  int
	)

	size := 100
	as := assert.New(t)
	list := NewList(false)

	// Insert items.
	for i := 0; i < size; i++ {
		item := It(i)
		inserted = list.AddAfter(item)
		as.True(inserted, "item %s no inserted", item)
		_, found := list.avl.Search(&listNode{nil, nil, item})
		as.True(found, "item %s not found in the AVL tree", item)
	}
	as.Equal(list.Length(), size, "list length is invalid")

	// check the pointers visiting all elements in the list from the start
	counter = 0
	list.First()
	for item, advance := list.Get(); advance; item, advance = list.Advance() {
		as.Equal(item.(IntItem).value, counter, "position of %s item is incorrect", item)
		counter++
	}
	as.Equal(list.Length(), counter, "number of item visited is inavlid")

	// check the pointers visiting all elements in the list from the end
	counter = 0
	list.Last()
	for item, cont := list.Get(); cont; item, cont = list.Rewind() {
		value := item.(IntItem).value
		as.Equal(value, list.Length()-counter-1, "position of %s item is incorrect", item)
		counter++
	}
	as.Equal(list.Length(), counter, "number of item visited is inavlid")

	// List with duplicated items.
	list = NewList(true)
	as.True(list.AddAfter(It(1)), "item wasn't inserted")
	as.True(list.AddAfter(It(1)), "item wasn't inserted")
	as.Equal(list.Length(), 2, "the tree length doesn't match")

	// list without duplicated items
	list = NewList(false)
	as.True(list.AddAfter(It(1)), "item wasn't inserted")
	as.False(list.AddAfter(It(1)), "duplicated item was inserted")
	as.Equal(list.avl.Length(), 1, "the tree length doesn't match")
}

func Test_List_AddAfter_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 1000
	done := make(chan bool)
	insert := func(min, max int) {
		for i := min; i < max; i++ {
			list.AddAfter(It(i))
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go changeListProperties(&list, size, done)
		go insert(i*size, (i+1)*size)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(list.avl.Length(), concurrence*size, "list length doesn't match")
	for i := 0; i < concurrence*size; i++ {
		item := It(i)
		_, found := list.avl.Search(&listNode{nil, nil, item})
		as.Truef(found, "item %s not found", item)
	}
}

func Test_List_AddBefore_func(t *testing.T) {
	var (
		inserted bool
		counter  int
	)

	size := 100
	as := assert.New(t)
	list := NewList(false)

	// Insert items.
	for i := 0; i < size; i++ {
		item := It(i)
		inserted = list.AddBefore(item)
		as.True(inserted, "item %s no inserted", item)
		_, found := list.avl.Search(&listNode{nil, nil, item})
		as.True(found, "item %s not found in the AVL tree", item)
	}
	as.Equal(list.Length(), size, "list length is invalid")

	// check the pointers visiting all elements in the list from the start
	counter = 0
	list.First()
	for item, advance := list.Get(); advance; item, advance = list.Advance() {
		value := item.(IntItem).value
		as.Equal(value, list.Length()-counter-1, "position of %s item is incorrect", item)
		counter++
	}
	as.Equal(list.Length(), counter, "number of item visited is inavlid")

	// check the pointers visiting all elements in the list from the end
	counter = 0
	list.Last()
	for item, cont := list.Get(); cont; item, cont = list.Rewind() {
		value := item.(IntItem).value
		as.Equal(value, counter, "position of %s item is incorrect", item)
		counter++
	}
	as.Equal(list.Length(), counter, "number of item visited is inavlid")

	// List with duplicated items.
	list = NewList(true)
	as.True(list.AddBefore(It(1)), "item wasn't inserted")
	as.True(list.AddBefore(It(1)), "item wasn't inserted")
	as.Equal(list.Length(), 2, "the tree length doesn't match")

	// list without duplicated items
	list = NewList(false)
	as.True(list.AddBefore(It(1)), "item wasn't inserted")
	as.False(list.AddBefore(It(1)), "duplicated item was inserted")
	as.Equal(list.avl.Length(), 1, "the tree length doesn't match")
}

func Test_List_AddBefore_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 1000
	done := make(chan bool)
	insert := func(min, max int) {
		for i := min; i < max; i++ {
			list.AddBefore(It(i))
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go changeListProperties(&list, size, done)
		go insert(i*size, (i+1)*size)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(list.avl.Length(), concurrence*size, "list length doesn't match")
	for i := 0; i < concurrence*size; i++ {
		item := It(i)
		_, found := list.avl.Search(&listNode{nil, nil, item})
		as.True(found, "item %s not found", item)
	}
}

func Test_List_Next_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)

	// empty list
	as.False(list.Next(), "there is a next item in the empty list")

	for i := 0; i < 5; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < 5; i++ {
		item, _ := list.Get()
		value := item.(IntItem).value
		as.Equal(value, i, "item is incorrect")

		if i < 4 {
			as.True(list.Next(), "cannot access to next item")
		} else {
			as.False(list.Next(), "last item has next element")
		}
	}
}

func Test_List_Next_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	next := func() {
		for i := 0; i < size; i++ {
			as.True(list.Next(), "cannot access next item")
		}

		done <- true
	}

	for i := 0; i < concurrence*size+1; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go next()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.False(list.Next(), "current item isn't the last item")
}

func Test_List_Prev_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)

	// empty list
	as.False(list.Prev(), "there is a previous item in the empty list")

	for i := 0; i < 5; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	for i := 4; i >= 0; i-- {
		if i > 0 {
			as.True(list.Prev(), "cannot access to prev item")
		} else {
			as.False(list.Prev(), "first item has a previous element")
		}
	}
}

func Test_List_Prev_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	prev := func() {
		for i := 0; i < size; i++ {
			as.True(list.Prev(), "cannot access next item")
		}
		done <- true
	}

	for i := 0; i < size*concurrence+1; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	for i := 0; i < concurrence; i++ {
		go prev()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equalf(list.pnode, list.fnode, "pointed node isn't the first node")
}

func Test_List_First_func(t *testing.T) {
	list := NewList(true)

	for i := 0; i < 5; i++ {
		list.AddAfter(It(i))
	}

	list.pnode = list.lnode
	list.First()
	assert.Equalf(t, list.pnode, list.fnode, "pointed node isn't the first node")
}

func Test_List_First_func_sync(t *testing.T) {
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	first := func() {
		for i := 0; i < size; i++ {
			list.First()
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go first()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	assert.Equal(t, list.pnode, list.fnode, "pointed node isn't the first node")
}

func Test_List_Last_func(t *testing.T) {
	list := NewList(true)

	for i := 0; i < 5; i++ {
		list.AddAfter(It(i))
	}

	list.pnode = list.fnode
	list.Last()
	assert.Equalf(t, list.pnode, list.lnode, "pointed node isn't the first node")
}

func Test_List_Last_func_sync(t *testing.T) {
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	last := func() {
		for i := 0; i < size; i++ {
			list.Last()
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go last()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	assert.Equalf(t, list.pnode, list.lnode, "node pointed isn't the last node")
}

func Test_List_Advance_func(t *testing.T) {
	list := NewList(true)
	size := 5
	as := assert.New(t)

	// In empty list
	item, cont := list.Advance()
	if item != nil {
		as.Fail(
			"item isn't nil",
			"the item returned when it advances in an empty list isn't nil: item: %s",
			item,
		)
	}

	as.False(cont, "item returned when it advances in an empty list")

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	i := 0
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		value := item.(IntItem).value
		as.Equal(value, i, "item value is invalid, value: %d", value)
		i++
	}

	as.Equal(i, size, "list length is invalid")
}

func Test_List_Advance_func_sync(t *testing.T) {
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	as := assert.New(t)
	advance := func() {
		for i := 0; i < size; i++ {
			_, cont := list.Advance()
			as.True(cont, "list cannot advance")
		}

		done <- true
	}

	for i := 0; i < size*concurrence+1; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go advance()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	item, cont := list.Advance()
	assert.Nil(t, item, "item isn't nil when advance from last position")
	assert.False(t, cont, "item can advance from the last position")
}

func Test_List_Rewind_func(t *testing.T) {
	list := NewList(true)
	as := assert.New(t)
	size := 5

	// In empty list
	item, cont := list.Rewind()
	if item != nil {
		as.Fail(
			"item isn't nil",
			"the item returned when it rewinds in an empty list isn't nil: item: %s",
			item,
		)
	}

	as.False(cont, "item returned when it rewinds in an empty list")

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	i := size
	for item, cont := list.Get(); cont; item, cont = list.Rewind() {
		value := item.(IntItem).value
		as.Equal(value, i-1, "item value is invalid, value: %d", value)
		i--
	}

	as.Equalf(i, 0, "list length is invalid")
}

func Test_List_Rewind_func_sync(t *testing.T) {
	list := NewList(true)
	as := assert.New(t)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	rewind := func() {
		for i := 0; i < size; i++ {
			_, cont := list.Rewind()
			as.True(cont, "the list cannot Rewind")
		}

		done <- true
	}

	for i := 0; i < size*concurrence+1; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	for i := 0; i < concurrence; i++ {
		go rewind()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	item, cont := list.Rewind()
	as.Nil(item, "item isn't nil, when the list rewind from the first position")
	as.False(cont, "can continue, when the list rewind from the first position")
}

func Test_List_Get_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)

	// empty list
	item, getted := list.Get()
	as.Nil(item, "list is empty but the item returned isn't nil")
	as.False(getted, "list is empty but the item returned isn't nil")


	for i := 0; i < 5; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < 5; i++ {
		item, getted := list.Get()
		list.Next()

		value := item.(IntItem).value
		as.Equal(value, i, "value is invalid")
		as.True(getted, "item wasn't getted")
	}
}

func Test_List_Get_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	get := func() {
		for i := 0; i < 2000; i++ {
			item, getted := list.Get()
			value := item.(IntItem).value
			as.Equalf(value, 0, "value isn't 0")
			as.True(getted, "item wasn't got")
		}

		done <- true
	}

	for i := 0; i < 10; i++ {
		list.AddAfter(It(i))
	}

	list.First()

	for i := 0; i < concurrence; i++ {
		go get()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}
}

func Test_List_Replace_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	size := 5

	// empty list
	as.Falsef(list.Replace(It(-1)), "item was replaced in an empty list")

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		value := item.(IntItem).value
		as.Truef(list.Replace(It(value+10)), "value %d wasn't replaced", value)
	}


	list.First()
	for i := 0; i < size; i++ {
		item, _ := list.Get()
		value := item.(IntItem).value
		as.Equalf(value, i+10, "value is invalid")
		list.Next()
	}
}

func Test_List_Replace_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 1000
	done := make(chan bool)
	replace := func() {
		for i := 0; i < size; i++ {
			as.True(list.Replace(It(-1)), "item wasn't replaced")
		}

		done <- true
	}

	for i := 0; i < size*concurrence; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go replace()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	for i := 1; i < size*concurrence; i++ {
		_, exist := list.Search(It(i))
		as.True(exist, "item %d doesn't exist", i)
	}
}

func Test_List_Search_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	size := 100

	// In empty list
	item, found := list.Search(It(3))
	as.Nil(item, "item returned in an empty list isn't nil")
	as.Falsef(found, "item found in an empty list")

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	for i := 0; i < size; i++ {
		item, found = list.Search(It(i))
		value := item.(IntItem).value
		as.Equal(value, i, "item returned in the search doesn't match")
		as.True(found, "item %d wasn't found", value)
	}

	// Search an item it doesn't exist in the list
	item, found = list.Search(It(-1))
	if item != nil {
		err := "item was returned"
		msg :=  "item -1 was returned in the list: value: %d"
		as.Fail(err, msg, item.(IntItem).value)
	}
	as.False(found, "item -1 was found in the list")
}

func Test_List_Search_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	size := 1000
	concurrence := 8
	done := make(chan bool)
	search := func(min, max int) {
		for i := min; i < max; i++ {
			item, found := list.Search(It(i))
			value := item.(IntItem).value
			as.Equal(value, i, "values don't match")
			as.True(found, "value %d not found", value)
		}
		done <- true
	}

	for i := 0; i < concurrence*size; i++ {
		list.AddAfter(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go search(i*size, (i+1)*size)
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}
}

func Test_List_Delete_func(t *testing.T) {
	var (
		list    List
		item    Item
		deleted bool
	)
	list = NewList(false)
	as := assert.New(t)
	size := 100

	// Delete item in empty list
	item, deleted = list.Delete()
	as.Nil(item, "item returned isn't nil when it deletes an item in empty list")
	as.False(deleted, "item was deleted in an empty list")

	// Delete item in a list without duplicated items
	list = NewList(false)
	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < size; i++ {
		item, deleted = list.Delete()
		if item == nil {
			msg := "item returned is nil when delete the item in position %d"
			as.Fail("item is nil", msg, i)
			continue

		}
		as.True(deleted, "item wasn't deleted: position: %d", i)

		_, found := list.avl.Search(item)
		as.False(found, "item deleted was found in list avl tree")
	}

	as.Equal(list.Length(), 0, "list isn't empty")


	// Delete item in a list with duplicated items
	list = NewList(true)
	for i := 0; i < size; i++ {
		list.AddAfter(It(1))
	}

	list.First()
	for i := 0; i < size; i++ {
		item, deleted = list.Delete()
		if item == nil {
			msg := "item returned is nil when delete the item in position %d"
			as.Fail("item is nil", msg, i)
			continue

		}
		as.True(deleted, "item wasn't deleted: position: %d", i)

		_, found := list.avl.Search(item)
		as.False(found, "item deleted was found in list avl tree")
	}

	as.Equal(list.Length(), 0, "list isn't empty")
}

func Test_List_Delete_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 1000
	done := make(chan bool)
	deletef := func() {
		for i := 0; i < size; i++ {
			_, deleted := list.Delete()
			as.True(deleted, "item wasn't deleted")
		}

		done <- true
	}

	for i := 0; i < (concurrence+1)*size; i++ {
		list.AddAfter(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go deletef()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(list.Length(), 1000, "list length is invalid")
}

func Test_List_Length_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(false)
	size := 10

	/*
		No duplicate items
	*/
	for i := 1; i <= size; i++ {
		list.AddAfter(It(i))
		as.Equal(list.Length(), i, "list length is invalid")
	}

	list.AddAfter(It(1))
	as.Equal(list.Length(), size, "list length is invalid")

	// remove items
	for i := 1; i <= size; i++ {
		list.Delete()
		as.Equal(list.Length(), size-i, "list length is invalid")
	}

	as.Equal(list.Length(), 0, "list length isn't 0 in empty list")

	/*
		Duplicate items
	*/
	list = NewList(true)
	for i := 1; i <= size; i++ {
		list.AddAfter(It(i))
		as.Equal(list.Length(), i, "list length is invalid")
	}

	list.AddAfter(It(1))
	as.Equal(list.Length(), size+1, "list length is invalid")

	// remove items
	for i := 1; i <= size+1; i++ {
		list.Delete()
		as.Equal(list.Length(), size+1-i, "list length is invalid")
	}

	as.Equal(list.Length(), 0, "list length isn't 0 in empty list")
}

func Test_List_Length_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 1000
	done := make(chan bool)
	length := func () {
		for i := 0; i < size; i++ {
			as.Equal(list.Length(), size, "list length is invalid")
		}

		done <- true
	}

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}


	for i := 0; i < concurrence; i++ {
		go length()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence*2; i++ {
		<- done
	}
}

func Test_List_Clear_func(t *testing.T) {
	as := assert.New(t)

	for _, duplicated := range []bool{true, false} {
		list := NewList(duplicated)

		for i := 0; i < 5; i++ {
			list.AddAfter(It(i))
		}

		list.Clear()
		as.Nil(list.fnode, "pointer to first item isn't nil")
		as.Nil(list.fnode, "pointer to current item isn't nil")
		as.Nil(list.fnode, "pointer to last item isn't nil")

		as.Nil(list.avl.root, "tree root in list isn't nil")
		as.Equal(list.avl.length, 0, "invalid length in the avl tree")
		as.True(list.avl.rebalance, "tree in the list isn't rebalanced")
		as.Equal(list.avl.duplicated, duplicated, "duplicated flag is invalid")
	}
}



func Test_List_ForEach_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	size := 10

	list.ForEach(func (Item) {
		as.FailNow("function was exectued when the list was empty")
	})

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	// Move the internal pointer to check if the `ForEach` function starts from begining
	list.Last()
	itBefore, _ := list.Get()

	i := 0
	list.ForEach(func(it Item) {
		as.Equal(it.(IntItem).value, i, "value is invalid")
		i++
	})
	as.Equal(i, size, "foreach function wasn't executed")

	// test if the internal pointer is pointed the same node that before of execute the
	// ForEach function.
	itAfter, _ := list.Get()
	as.Equal(
		itAfter,
		itBefore,
		"item pointed by internal pointer is diff after execute foreach function",
	)
}

func Test_List_ForEach_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	concurrence := 8
	size := 2000
	done := make(chan bool)
	foreach := func() {
		i := 0;
		list.ForEach(func(item Item) {
			value := item.(IntItem).value
			as.Equal(value, i, "value is incorrect")
			i++

			//  the internal pointers don't change
			itemGetted, _ := list.Get()
			as.Equal(itemGetted.(IntItem).value, 0, "internal pointer changed")
		})

		as.Equal(i, size, "index doesn't match with the list length")
		done <- true
	}

	for i := 0; i < size; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go foreach()
		go changeListProperties(&list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	// The internal pointer is in the first item.
	itemGetted, _ := list.Get()
	as.Equal(itemGetted.(IntItem).value, 0, "item getted isn't 0")
}

func Test_List_Map_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)
	size := 5

	for i := 1; i <= size; i++ {
		list.AddAfter(It(i))
	}

	// Move the internal pointer.
	list.Last()
	itBefore, _ := list.Get()

	pow2List := list.Map(func(it Item) Item {
		num := it.(IntItem).value
		return It(num*num)
	})

	as.Equal(pow2List.Length(), size, "length of new list is invalid")
	pow2List.First()
	for i := 1; i <= size; i++ {
		it, _ := pow2List.Get()
		as.Equal(it.(IntItem).value, i*i, "value is incorrect")
		pow2List.Next()
	}

	// Check if the internal pointer is pointed the same node that before
	// of execute the Map function.
	itAfter, _ := list.Get()
	as.Equal(
		itAfter,
		itBefore,
		"item pointed by internal pointer is diff after execute map function",
	)
}

func Test_List_Filter_func(t *testing.T) {
	as := assert.New(t)
	list := NewList(true)

	filter := func(it Item) bool {
		return it.(IntItem).value%2 == 1
	}

	for i := 1; i <= 10; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	itBefore, _ := list.Get()

	newList := list.Filter(filter)

	i := 0
	newList.First()
	for it, cont := newList.Get(); cont; it, cont = newList.Advance() {
		as.Equal(it.(IntItem).value, i*2+1)
		i++
	}

	as.Equal(newList.Length(), i, "new list length is invalid")

	// Check if the internal pointer is pointed the same node that before of execute the
	// Map function.
	itAfter, _ := list.Get()
	as.Equal(
		itAfter,
		itBefore,
		"item pointed by internal pointer is diff after execute filter function",
	)
}
