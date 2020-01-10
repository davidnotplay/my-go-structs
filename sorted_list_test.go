package mygostructs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSortedList_func(t *testing.T) {
	for _, dup := range []bool{true, false} {
		sl := NewSortedList(dup)

		assert.Nil(t, sl.list.fnode)
		assert.Nil(t, sl.list.pnode)
		assert.Nil(t, sl.list.lnode)

		assert.Nil(t, sl.list.avl.root)
		assert.Equal(t, sl.list.avl.length, 0)
		assert.Equal(t, sl.list.avl.rebalance, true)
		assert.Equal(t, sl.list.avl.duplicated, dup)
	}
}

func Test_SortedList_Add_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(false)
	data := []int{8, 2, 9, 3, 0, 5, 6, 7, 4, 1}
	size := len(data)
	min, max := 999, -1

	for _, value := range data {
		if  value < min {
			min = value
		}

		if value > max {
			max = value
		}

		item := It(value)
		inserted := list.Add(item)
		as.True(inserted, "item %s wasn't inserted", item)
		as.Equal(
			list.list.pnode.item.(IntItem).value,
			value,
			"item pointed by current pointer is invalid: item pointed: %s",
			item)
		as.Equal(
			list.list.fnode.item.(IntItem).value,
			min,
			"item pointed by first pointer is invalid: item pointed: %s",
			item)
		as.Equal(
			list.list.lnode.item.(IntItem).value,
			max,
			"item pointed by last pointer is invalid: item pointed: %s",
			item)
	}

	// Loop the list to check the items order.
	counter := 0
	for node := list.list.fnode; node != nil; counter, node = counter+1, node.next {
		value := node.item.(IntItem).value
		as.Equal(value, counter, "value of item is invalid")
	}

	as.Equal(counter, size, "number of iterations in loop is invalid")


	// Loop backward the list to check the items order.
	counter = 0
	for node := list.list.lnode; node != nil; counter, node = counter+1, node.prev {
		value := node.item.(IntItem).value
		as.Equal(value, size-1-counter, "value of item is invalid")
	}

	// Avoid duplicated items
	list = NewSortedList(false)
	as.True(list.Add(It(1)), "item wasn't inserted")
	as.False(list.Add(It(1)), "duplicated item was inserted")


	// List with duplicated items.
	list = NewSortedList(true)
	data = []int{3, 2, 0, 0, 1, 2, 4, 3, 1, 4}
	min, max = 999, -1
	for _, value := range data {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}

		item := It(value)
		inserted := list.Add(item)
		as.True(inserted, "item %s wasn't inserted", item)
		as.Equal(
			list.list.pnode.item.(IntItem).value,
			value,
			"item pointed by current pointer is invalid: item pointed: %s",
			item)
		as.Equal(
			list.list.fnode.item.(IntItem).value,
			min,
			"item pointed by first pointer is invalid: item pointed: %s",
			item)
		as.Equal(
			list.list.lnode.item.(IntItem).value,
			max,
			"item pointed by last pointer is invalid: item pointed: %s",
			item)
	}

	// Loop the list to check the items order.
	counter = 0
	for node := list.list.fnode; node != nil; counter, node = counter+1, node.next {
		value := node.item.(IntItem).value
		as.Equal(value, (counter/2), "value of item is invalid")
	}

	as.Equal(counter, size, "number of iterations in loop is invalid")

	// Loop backward the list to check the items order.
	counter = 0
	for node := list.list.lnode; node != nil; counter, node = counter+1, node.prev {
		value := node.item.(IntItem).value
		as.Equal(value, (size-1-counter)/2, "value of item is invalid")
	}
}

func Test_SortedList_Add_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
	size := 1000
	concurrence := 8
	done := make(chan bool)
	add := func (min, max int) {
		for i := min; i < max; i++ {
			item := It(i)
			as.True(list.Add(item), "item %s wasn't inserted", item)
		}

		done <- true
	}

	for i := 0; i < concurrence; i++ {
		go add (size*i, (i+1)*size)
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(list.Length(), size*concurrence, "list length is invalid")
	list.First()
	for i := 0; i < concurrence*size; i++ {
		item := It(i)
		_, found := list.Search(item)
		as.True(found, "item %s wasn't found", item)
	}
}

func Test_SortedList_Next_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(false)


	// test in empty list
	as.False(list.Next(), "empty list has a next item")

	for _, value := range []int{3, 2, 4, 1, 5, 8, 0, 6, 7, 9} {
		list.Add(It(value))
	}

	list.First()
	i := 0
	for cont := true; cont; cont, i = list.Next(), i+1 {
		item, _ := list.Get()
		value := item.(IntItem).value
		as.Equal(value, i, "item %s is in an invalid position in the list", item)
	}

	as.Equal(list.Length(), i, "list didn't run the loop")
	as.False(list.Next(), "item found after last item")
}

func Test_SortedList_Next_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
	size := 1000
	concurrence := 8
	done := make(chan bool)
	next := func() {
		for i := 0; i < size; i++ {
			as.True(list.Next())
		}

		done <- true
	}

	for i := 0; i < (size*concurrence)+1; i++ {
		list.Add(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go next()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.False(list.Next(), "item found after last item")
}

func Test_SortedList_Prev_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(false)

	// test in empty list
	as.False(list.Prev(), "empty list has a previous item")

	for _, value := range []int{3, 2, 4, 1, 5, 8, 0, 6, 7, 9} {
		list.Add(It(value))
	}

	list.Last()
	i := 0
	for cont := true; cont; cont, i = list.Prev(), i+1 {
		item, _ := list.Get()
		value := item.(IntItem).value
		as.Equal(value, 9-i, "item %s is in an invalid position in the list", item)
	}

	as.Equal(list.Length(), i, "list didn't run the loop")
	as.False(list.Prev(), "item found before first item")
}

func Test_SortedList_Prev_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
	size := 1000
	concurrence := 8
	done := make(chan bool)
	prev := func() {
		for i := 0; i < size; i++ {
			as.True(list.Prev())
		}

		done <- true
	}

	for i := 0; i < (size*concurrence)+1; i++ {
		list.Add(It(i))
	}

	list.Last()
	for i := 0; i < concurrence; i++ {
		go prev()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.False(list.Prev(), "item found before first item")
}

func Test_SortedList_First_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		list.Add(It(i))
	}

	list.list.pnode = list.list.lnode
	list.First()
	as.Equal(list.list.pnode, list.list.fnode, "item pointed isn't the first item")
}

func Test_SortedList_First_func_sync(t *testing.T) {
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go first()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	assert.Equal(t, list.list.pnode, list.list.fnode, "pointed node isn't the first node")
}

func Test_SortedList_Last_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.list.pnode = so.list.fnode
	so.Last()
	as.Equal(so.list.pnode, so.list.lnode, "item pointed isn't the last item")
}

func Test_SortedList_Last_func_sync(t *testing.T) {
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go last()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	assert.Equalf(t, list.list.pnode, list.list.lnode, "node pointed isn't the last node")
}

func Test_SortedList_Advance_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	// In empty list
	item, cont := so.Advance()
	if item != nil {
		as.Fail(
			"item isn't nil",
			"the item returned when it advances in an empty list isn't nil: item: %s",
			item,
		)
	}

	as.False(cont, "item returned when it advances in an empty list")

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.First()
	i := 1
	for item, cont := so.Get(); cont; item, cont = so.Advance() {
		as.Equal(item.(IntItem).value, i, "value is incorrect")
		i++
	}
}


func Test_SortedList_Advance_func_sync(t *testing.T) {
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go advance()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	item, cont := list.Advance()
	assert.Nil(t, item, "item isn't nil when advance from last position")
	assert.False(t, cont, "item can advance from the last position")
}

func Test_SortedList_Rewind_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	// In empty list
	item, cont := so.Rewind()
	if item != nil {
		as.Fail(
			"item isn't nil",
			"the item returned when it rewinds in an empty list isn't nil: item: %s",
			item,
		)
	}

	as.False(cont, "item returned when it rewinds in an empty list")

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.Last()
	i := 5
	for item, cont := so.Get(); cont; item, cont = so.Rewind() {
		as.Equal(item.(IntItem).value, i, "item value is incorrect")
		i--
	}
}

func Test_SortedList_Rewind_func_sync(t *testing.T) {
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	list.Last()
	for i := 0; i < concurrence; i++ {
		go rewind()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	item, cont := list.Rewind()
	as.Nil(item, "item isn't nil, when the list rewind from the first position")
	as.False(cont, "can continue, when the list rewind from the first position")
}

func Test_SortedList_Get_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	// Get item in an empty list
	v, exists := so.Get()
	as.Nil(v, "item getted in an empty list")
	as.False(exists, "item getted in an empty list")

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	as.Equal(so.list.avl.length, 5)

	so.First()
	for i := 1; i <= 5; i++ {
		v, exists = so.Get()
		as.True(exists)
		as.Equal(v.(IntItem).value, i, "item getted is invalid")
		so.Next()
	}
}

func Test_SortedList_Get_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	list.First()

	for i := 0; i < concurrence; i++ {
		go get()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}
}

func Test_SortedList_Search_func(t *testing.T) {
	var (
		item  Item
		found bool
	)
	as := assert.New(t)
	so := NewSortedList(true)
	size := 10

	// First checks an empty list
	item, found = so.Search(It(1))
	as.Nil(item, "item found in an empty list")
	as.False(found, "item found in an empty list")

	for i := 0; i < size; i++ {
		so.Add(It(i))
	}

	for i := 0; i < size; i++ {
		item, found := so.Search(It(i))
		as.Equal(item.(IntItem).value, i, "value of item found is invalid")
		as.True(found, "item not found")
	}

	// search item it isn't in the list
	item, found = so.Search(It(11))

	if item != nil {
		as.Fail(
			"item isn't nil",
			"item no inserted was found in the list: item: %s",
			item,
		)
	}

	as.False(found, "item no inserted was found in the list")
}

func Test_SortedList_Search_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go search(i*size, (i+1)*size)
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}
}

func Test_SortedList_Delete_func(t *testing.T) {
	var (
		item    Item
		deleted bool
	)
	as := assert.New(t)
	so := NewSortedList(true)

	// test empty list
	item, deleted = so.Delete()
	if item != nil {
		as.Fail("item isn't nil", "item %s was deleted in an empty list", item)
	}
	as.False(deleted, "item was deleted in an empty list")

	// only one item
	so.Add(It(3))
	item, deleted = so.Delete()

	as.Equal(item.(IntItem).value, 3, "item deleted is invalid")
	as.True(deleted, "item wasn't deleted")
	as.Nil(so.list.pnode, "internal pointer isn't nil when the last item was deleted")
	as.Nil(so.list.fnode, "pointer to first item isn't nil when the last item was deleted")
	as.Nil(so.list.lnode, "pointer to last item isn't nil when the last item was deleted")
	as.Equal(so.Length(), 0, "list isn't empty, when the last item was deleted")

	for i := 1; i <= 10; i++ {
		so.Add(It(i))
	}

	// insert items duplicated
	so.Add(It(2))
	so.Add(It(4))
	so.Add(It(7))

	maxLength := so.list.avl.length
	for i, a := range []int{3, 2, 4, 6, 9, 1, 2, 7, 5, 10, 8, 7, 4} {
		_, found := so.Search(It(a))
		as.True(found)

		vdeleted, deleted := so.Delete()
		as.True(deleted, "item wasn't deleted")
		as.Equal(vdeleted.(IntItem).value, a, "the value of item deleted is invalid")
		as.Equal(
			so.list.pnode,
			so.list.fnode,
			"internal pointer isn't pointed to firs item",
		)

		as.Equal(
			so.list.Length(),
			maxLength-(i+1),
			"list length is invalid after delete an item",
		)

		_, found = so.Search(It(a))
		as.Equal(
			found,
			i == 1 || i == 2 || i == 7,
			"item wasn't found when it has got an duplicated item in the list",
		)
	}

	as.Nil(so.list.pnode, "internal pointer isn't nil when the last item was deleted")
	as.Nil(so.list.fnode, "pointer to first item isn't nil when the last item was deleted")
	as.Nil(so.list.lnode, "pointer to last item isn't nil when the last item was deleted")
	as.Equal(so.Length(), 0, "list isn't empty, when the last item was deleted")
}

func Test_SortedList_Delete_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	for i := 0; i < concurrence; i++ {
		go deletef()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	as.Equal(list.Length(), 1000, "list length is invalid")
}

func Test_SortedList_Length_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)
	size := 10

	as.Equal(so.Length(), 0, "length is invalid in an empty list")

	for i := 0; i < size; i++ {
		so.Add(It(i))
		as.Equal(so.Length(), i+1, "list length is invalid")
	}

	for i := 0; i < size; i++ {
		so.Delete()
		as.Equal(so.Length(), size-i-1, "list length is invalid")
	}

	as.Equal(so.Length(), 0, "length is invalid in an empty list")
}

func Test_SortedList_Length_func_sync(t *testing.T) {
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

func Test_SortedList_Clear_func(t *testing.T) {
	as := assert.New(t)

	for _, duplicated := range []bool{true, false} {
		so := NewSortedList(duplicated)
		for i := 1; i <= 10; i++ {
			so.list.AddAfter(It(i))
		}

		// clear the list
		so.Clear()

		as.Nil(so.list.fnode, "pointer to first item isn't nil")
		as.Nil(so.list.pnode, "internal pointer isn't nil")
		as.Nil(so.list.lnode, "pointer to last item isn't nil")
		as.Equal(so.list.avl.length, 0, "length isn't 0 in an empty list")

		as.Equal(so.list.avl.duplicated, duplicated, "duplicated flag in list is invalid")
		as.Equal(so.list.avl.length, 0,"avl length is invalid when the tree is empty")
		as.True(so.list.avl.rebalance, "rebalanced flag is invalid")
		as.Nil(so.list.avl.root, "avl root pointer isn't 0")
	}
}

func Test_SortedList_ForEach_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
	size := 10

	list.ForEach(func (Item) {
		as.FailNow("function was exectued when the list was empty")
	})

	for i := 0; i < size; i++ {
		list.Add(It(i))
	}

	// Move the internal pointer to check if the `ForEach` function starts from begining
	list.Last()
	itBefore, _ := list.Get()

	i := 0
	list.ForEach(func(it Item) {
		as.Equal(it.(IntItem).value, i, "item of the foreach func is invalid")
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
func Test_SortedList_ForEach_func_sync(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
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
		list.Add(It(i))
	}

	list.First()
	for i := 0; i < concurrence; i++ {
		go foreach()
		go changeListProperties(list.list, size, done)
	}

	for i := 0; i < concurrence; i++ {
		<- done
		<- done
	}

	// The internal pointer is in the first item.
	itemGetted, _ := list.Get()
	as.Equal(itemGetted.(IntItem).value, 0, "item getted isn't 0")
}

func Test_SortedList_Map_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)
	size := 5

	for i := 1; i <= size; i++ {
		list.Add(It(i))
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

func Test_SortedList_Filter_func(t *testing.T) {
	as := assert.New(t)
	list := NewSortedList(true)

	filter := func(it Item) bool {
		return it.(IntItem).value%2 == 1
	}

	for i := 1; i <= 10; i++ {
		list.Add(It(i))
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
