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
	so := NewSortedList(false)
	data := []int {8, 2, 9, 10, 3, 5, 6, 7, 1, 4};

	min, max, l := 999, 0, 1
	for _, value := range data {
		inserted := so.Add(It(value))


		if value < min {
			min = value
		}

		if value > max {
			max = value
		}

		as.True(inserted)
		as.Equal(so.list.avl.length, l)
		as.Equalf(so.list.pnode.item.(IntItem).value, value, "value: %d", value)
		as.Equalf(so.list.fnode.item.(IntItem).value, min, "value: %d", value)
		as.Equalf(so.list.lnode.item.(IntItem).value, max, "value: %d", value)

		l++
	}


	node := so.list.fnode
	for a := 1; a <= 10; a++ {
		as.Equal(node.item.(IntItem).value, a)
		node = node.next
	}

	node = so.list.lnode
	for a := 10; a >= 1; a-- {
		as.Equal(node.item.(IntItem).value, a)
		node = node.prev
	}
}

func Test_SortedList_Next_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		so.Add(It(i))
	}

	so.list.pnode = so.list.fnode // move the pointer to first item.

	for i := 1; i <= so.list.avl.length; i++ {
		as.Equal(so.list.pnode.item.(IntItem).value, i)
		if i < so.list.avl.length {
			as.True(so.Next())
			continue
		}

		// last item
		as.False(so.Next()) // In the end of the list. No continue
		as.Equal(so.list.pnode, so.list.lnode)
	}

	as.False(so.Next())
	as.False(so.Next())
}

func Test_SortedList_Prev_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)
	values := []int{1, 2, 3, 4, 5}

	// insert the value
	for _, i := range values {
		so.Add(It(i))
	}

	so.list.pnode = so.list.lnode //move the pointer to the last item.

	for i := so.list.avl.length; i >= 1; i-- {
		as.Equal(so.list.pnode.item.(IntItem).value, i)
		if i > 1 {
			as.True(so.Prev())
			continue
		}

		// last item
		as.False(so.Prev()) // In the end of the list. No continue
		as.Equal(so.list.pnode, so.list.fnode)
	}

	as.False(so.Prev())
	as.False(so.Prev())
}

func Test_SortedList_First_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	as.Equal(so.list.pnode, so.list.lnode)
	so.First()
	as.Equal(so.list.pnode, so.list.fnode)
}

func Test_SortedList_Last_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.list.pnode = so.list.fnode
	so.Last()
	as.Equal(so.list.pnode, so.list.lnode)
}

func Test_SortedList_Advance_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.First()
	i := 1
	for item, cont := so.Get(); cont; item, cont = so.Advance() {
		as.Equal(item.(IntItem).value, i)
		i++
	}
}

func Test_SortedList_Rewind_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	so.Last()
	i := 5
	for item, cont := so.Get(); cont; item, cont = so.Rewind() {
		as.Equal(item.(IntItem).value, i)
		i--
	}
}

func Test_SortedList_Get_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	// Get item in an empty list
	v, exists := so.Get()
	as.Nil(v)
	as.False(exists)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	as.Equal(so.list.avl.length, 5)

	so.First()
	for i := 1; i <= 5; i++ {
		v, exists = so.Get()
		as.True(exists)
		as.Equal(v.(IntItem).value, i)
		so.Next()
	}
}

func Test_SortedList_Search_func(t *testing.T) {
	var (
		item  Item
		found bool
	)
	as := assert.New(t)
	so := NewSortedList(true)

	// First checks an empty list
	item, found = so.Search(It(1))
	as.Nil(item)
	as.False(found)

	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		so.Add(It(a))
	}

	// Check the pointers are correct
	checkln(t, so.list.fnode, 1, nil, i(2))
	checkln(t, so.list.pnode, 6, i(5), nil)
	checkln(t, so.list.lnode, 6, i(5), nil)

	// Search items they are in the list
	for _, a := range []int{1, 2, 3, 4, 5, 6} {
		item, found := so.Search(It(a))
		as.Equal(item.(IntItem).value, a)
		as.True(found)

		// Check the pointers are correct
		checkln(t, so.list.fnode, 1, nil, i(2))
		checkln(t, so.list.lnode, 6, i(5), nil)

		switch {
		case a == 1:
			checkln(t, so.list.pnode, 1, nil, i(2))
		case a == 6:
			checkln(t, so.list.pnode, 6, i(5), nil)
		default:
			checkln(t, so.list.pnode, a, i(a-1), i(a+1))
		}

	}

	// Search items they arent' in the list
	for _, a := range []int{7, 8, 9, 10, -1, -2} {
		item, found := so.Search(It(a))
		as.Nil(item)
		as.False(found)

		// Check the pointers are correct
		checkln(t, so.list.fnode, 1, nil, i(2))
		checkln(t, so.list.pnode, 6, i(5), nil)
		checkln(t, so.list.lnode, 6, i(5), nil)
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
	as.Nil(item)
	as.False(deleted)

	// only one item
	so.Add(It(3))
	item, deleted = so.Delete()

	as.Equal(item.(IntItem).value, 3)
	as.True(deleted)
	as.Nil(so.list.pnode)
	as.Nil(so.list.fnode)
	as.Nil(so.list.lnode)
	as.Equal(so.list.avl.length, 0)

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
		as.True(deleted)
		as.Equal(vdeleted.(IntItem).value, a)
		as.Equal(so.list.pnode, so.list.fnode)
		as.Equal(so.list.Length(), maxLength-(i+1))

		_, found = so.Search(It(a))
		as.Equal(found, i == 1 || i == 2 || i == 7)
	}

	as.Equal(so.list.avl.length, 0)
	as.Nil(so.list.fnode)
	as.Nil(so.list.pnode)
	as.Nil(so.list.lnode)
}

func Test_SortedList_Clear_func(t *testing.T) {
	as := assert.New(t)

	for _, duplicated := range []bool{true, false} {
		so := NewSortedList(duplicated)
		as.Nil(so.list.fnode)
		as.Nil(so.list.pnode)
		as.Nil(so.list.lnode)
		as.Equal(so.list.avl.length, 0)

		as.Equal(so.list.avl.duplicated, duplicated)
		as.Equal(so.list.avl.length, 0)
		as.True(so.list.avl.rebalance)
		as.Nil(so.list.avl.root)

		for i := 1; i <= 10; i++ {
			so.list.AddAfter(It(i))
		}

		// clear the list
		so.Clear()

		as.Nil(so.list.fnode)
		as.Nil(so.list.pnode)
		as.Nil(so.list.lnode)
		as.Equal(so.list.avl.length, 0)

		as.Equal(so.list.avl.duplicated, duplicated)
		as.Equal(so.list.avl.length, 0)
		as.True(so.list.avl.rebalance)
		as.Nil(so.list.avl.root)
	}
}

func Test_SortedList_Length_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	as.Equal(so.Length(), 0)

	for i, a := range []int{1, 2, 3, 4, 5, 6} {
		so.Add(It(a))
		as.Equal(so.Length(), i+1)
	}
}

func Test_SortedList_ForEach_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	// Move the internal pointer.
	so.First()
	so.Next()
	so.Next()
	itPrev, _ := so.Get()

	i := 1
	so.ForEach(func(it Item) {
		as.Equal(it.(IntItem).value, i)
		i++
	})
	as.Equal(i, 6)

	// test if the internal pointer is pointed the same node that before of execute the
	// ForEach function.
	itNext, _ := so.Get()
	as.Equal(itNext, itPrev)
	as.Equal(itNext.(IntItem).value, 3)
}

func Test_SortedList_Map_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	for i := 1; i <= 5; i++ {
		so.Add(It(i))
	}

	// Move the internal pointer.
	so.First()
	so.Next()
	so.Next()
	itPrev, _ := so.Get()

	pow2List := so.Map(func(it Item) Item {
		num := it.(IntItem).value
		return It(num * num)
	})

	as.Equal(pow2List.Length(), 5)
	pow2List.First()
	for _, i := range []int{1, 4, 9, 16, 25} {
		it, _ := pow2List.Get()
		as.Equal(it.(IntItem).value, i)
		pow2List.Next()
	}

	as.False(pow2List.Next())

	// test if the internal pointer is pointed the same node that before of execute the
	// ForEach function.
	itNext, _ := so.Get()
	as.Equal(itNext, itPrev)
	as.Equal(itNext.(IntItem).value, 3)
}

func Test_SortedList_Filter_func(t *testing.T) {
	as := assert.New(t)
	so := NewSortedList(true)

	filter := func(it Item) bool {
		return it.(IntItem).value%2 == 1
	}

	for i := 1; i <= 10; i++ {
		so.Add(It(i))
	}

	newList := so.Filter(filter)

	i := 0
	newList.First()
	for it, cont := newList.Get(); cont; it, cont = newList.Advance() {
		as.Equal(it.(IntItem).value, i*2+1)
		i++
	}
}
