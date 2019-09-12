package mygostructs

import "fmt"

/*
	Tree
	====
*/

// Basic usage
func ExampleTree_Insert() {
	tree := Tree{}

	// insert numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
	}

	fmt.Printf("Number of items is %d.", tree.Length())

	/*
	Output:

Number of items is 10.
	*/
}

func ExampleTree_Length() {
	tree := Tree{}

	// insert numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
	}

	fmt.Printf("Number of items is %d.", tree.Length())

	/*
	Output:

Number of items is 10.
	*/
}

// Basic usage
func ExampleTree_Search() {
	tree := Tree{}

	// insert numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
	}

	if it, found := tree.Search(It(1)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}

	if it, found := tree.Search(It(11)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}
	/*
	Output:

Item 1 found.
Item not found.
	*/
}

// Basic usage
func ExampleTree_Delete() {
	tree := Tree{}

	// insert numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
	}

	// Search the 1.
	if it, found := tree.Search(It(1)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}

	// Delete the 1
	if it, found := tree.Delete(It(1)); found {
		fmt.Printf("Item %s deleted.\n", it.String())
	} else {
		fmt.Printf("Item not deleted.\n")
	}

	// Search the 1 again.
	if it, found := tree.Search(It(1)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}

	// Delete the 12. This number isn't in the tree.
	if it, found := tree.Delete(It(11)); found {
		fmt.Printf("Item %s deleted.\n", it.String())
	} else {
		fmt.Printf("Item not deleted.\n")
	}
	/*
	Output:

Item 1 found.
Item 1 deleted.
Item not found.
Item not deleted.
	*/
}

func ExampleTree_Clear() {
	tree := Tree{}

	// insert numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		tree.Insert(It(i))
	}

	fmt.Printf("Number of items is %d.\n", tree.Length())

	// Clear the tree
	tree.Clear()
	fmt.Printf("Number of items is %d.\n", tree.Length())

	/*
	Output:

Number of items is 10.
Number of items is 0.
	*/
}

// Basic usage
func ExampleAvl() {
	// create new avl
	avl := NewAvl()

	for i := 1; i <= 5; i++ {
		avl.Insert(It(i))
	}

	// Search the item 3
	if item, found := avl.Search(It(3)); found {
		fmt.Printf("Item %s found.\n", item.String())
	}

	// Delete the item 2
	if itemDeleted, deleted := avl.Delete(It(2)); deleted {
		fmt.Printf("Item %s deleted.\n", itemDeleted.String())
	}

	/*
	Output:

Item 3 found.
Item 2 deleted.
	*/
}

// Basic usage
func ExampleBst() {
	// create new Bst
	bst := NewBst()

	for i := 1; i <= 5; i++ {
		bst.Insert(It(i))
	}

	// Search the item 3
	if item, found := bst.Search(It(3)); found {
		fmt.Printf("Item %s found.\n", item.String())
	}

	// Delete the item 2
	if itemDeleted, deleted := bst.Delete(It(2)); deleted {
		fmt.Printf("Item %s deleted.\n", itemDeleted.String())
	}

	/*
	Output:

Item 3 found.
Item 2 deleted.
	*/
}

/*
	List
	====
*/

// Basic usage
func ExampleList() {
	// Create en empty List, that accepts duplicated items.
	list := NewList(true)

	// Insert in the list, the numbers: 1, 2, 3, 4, 5
	for _, i := range []int{1, 2, 3, 4, 5} {
		list.AddAfter(It(i))
	}

	// Iterate the List
	fmt.Printf("\nIterate the list:\n")
	list.First()
	for item := true; item; item = list.Next() {
		item, _ := list.Get()
		fmt.Printf("List item: %s\n", item.String())
	}

	// Iterate the list in reverse.
	fmt.Printf("\nIterate the list in reverse:\n")
	list.Last()
	for found := true; found; found = list.Prev() {
		item, _ := list.Get()
		fmt.Printf("List item: %s\n", item.String())
	}

	// Search an item
	fmt.Printf("\nSearch items:\n")
	if item, found := list.Search(It(3)); found {
		fmt.Printf("Item %s found.\n", item)
	} else {
		fmt.Printf("Item not found.\n")
	}


	// Delete an item
	fmt.Printf("\nDelete items:\n")
	if _, found := list.Search(It(4)); found {
		itemDeleted, _ := list.Delete()
		fmt.Printf("Item %s deleted.\n", itemDeleted)
	}

	/*
	Output:

Iterate the list:
List item: 1
List item: 2
List item: 3
List item: 4
List item: 5

Iterate the list in reverse:
List item: 5
List item: 4
List item: 3
List item: 2
List item: 1

Search items:
Item 3 found.

Delete items:
Item 4 deleted.
	*/
}

func ExampleList_AddAfter() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for it, cont := list.Get(); cont; it, cont = list.Advance() {
		fmt.Printf("List item %s\n", it)
	}

	/*
	Output:

List item 1
List item 2
List item 3
List item 4
List item 5
	*/
}

func ExampleList_AddBefore() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddBefore(It(i))
	}

	list.First()
	for it, cont := list.Get(); cont; it, cont = list.Advance() {
		fmt.Printf("List item %s\n", it)
	}

	/*
	Output:

List item 5
List item 4
List item 3
List item 2
List item 1
	*/
}

func ExampleList_Next() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	for cont := true; cont; cont = list.Next() {
		it, _ := list.Get()
		fmt.Printf("List item %s\n", it)
	}
	/*
	Output:

List item 1
List item 2
List item 3
List item 4
List item 5
	*/
}

func ExampleList_Prev() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	for cont := true; cont; cont = list.Prev() {
		it, _ := list.Get()
		fmt.Printf("List item %s\n", it)
	}
	/*
	Output:

List item 5
List item 4
List item 3
List item 2
List item 1
	*/
}

func ExampleList_First() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	list.First()
	it, _ := list.Get() // get first item
	fmt.Printf("List item %s\n", it)

	/*
	Output:

List item 1
	*/
}

func ExampleList_Last() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	list.Last()
	it, _ := list.Get() // get first item
	fmt.Printf("List item %s\n", it)

	/*
	Output:

List item 5
	*/
}

func ExampleList_Advance() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	// Iterate the List
	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		fmt.Printf("List item %s\n", item.String())
	}

	/*
	Output:

List item 1
List item 2
List item 3
List item 4
List item 5
	*/
}

func ExampleList_Rewind() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	// Iterate the list reversed.
	list.Last()
	for item, cont := list.Get(); cont; item, cont = list.Rewind() {
		fmt.Printf("List item %s\n", item.String())
	}

	/*
	Output:

List item 5
List item 4
List item 3
List item 2
List item 1
	*/
}

func ExampleList_Get() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	// Iterate the List
	list.First()

	// Get the items pointed by internal pointer.
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		fmt.Printf("List item %s\n", item.String())
	}

	/*
	Output:

List item 1
List item 2
List item 3
List item 4
List item 5
	*/
}

func ExampleList_Replace() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	// Replace the items: New item is the previous item + 10
	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		list.Replace(It(item.(IntItem).value + 10))
	}

	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		fmt.Printf("List item %s.\n", item)
	}
	/*
	Output:

List item 11.
List item 12.
List item 13.
List item 14.
List item 15.
	*/
}

func ExampleList_Delete() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	// Delete even numbers
	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		if item.(IntItem).value % 2 == 0 {
			list.Delete()
		}
	}

	list.First()
	for item, cont := list.Get(); cont; item, cont = list.Advance() {
		fmt.Printf("List item %s.\n", item)
	}
	/*
	Output:

List item 1.
List item 3.
List item 5.
	*/
}

// Basic usage
func ExampleList_Clear() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	fmt.Printf("Number of items in the list is %d.\n", list.Length())

	// Clear the list
	list.Clear()
	fmt.Printf("Number of items in the list is %d.\n", list.Length())

	/*
	Output:

Number of items in the list is 5.
Number of items in the list is 0.
	*/
}

// Basic usage
func ExampleList_Search() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
	}

	if it, found := list.Search(It(3)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}

	if it, found := list.Search(It(11)); found {
		fmt.Printf("Item %s found.\n", it.String())
	} else {
		fmt.Printf("Item not found.\n")
	}

	/*
	Output:

Item 3 found.
Item not found.
	*/
}

// Basic usage
func ExampleList_Length() {
	list := NewList(true)

	// insert the sequence: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		list.AddAfter(It(i))
		fmt.Printf("Number of items: %d\n", list.Length())
	}
	/*
	Output:

Number of items: 1
Number of items: 2
Number of items: 3
Number of items: 4
Number of items: 5
	*/

}

// Basic usage
func ExampleList_ForEach() {
	// Create en empty List, that accepts duplicated items.
	list := NewList(true)

	// Insert in the list, the numbers: 1, 2, 3, 4, 5
	for _, i := range []int{1, 2, 3, 4, 5} {
		list.AddAfter(It(i))
	}

	list.ForEach(func (it Item) {
		fmt.Printf("List item %s\n", it.String())
	})

	/*
	Output:

List item 1
List item 2
List item 3
List item 4
List item 5
	*/
}

// Basic usage
func ExampleList_Map() {
	// Create en empty List, that accepts duplicated items.
	list := NewList(true)

	// Insert in the list, the numbers: 1, 2, 3, 4, 5
	for _, i := range []int{1, 2, 3, 4, 5} {
		list.AddAfter(It(i))
	}

	pow2List := list.Map(func (it Item) Item {

		if number, valid := it.(IntItem); valid {
			return It(number.value * number.value)
		}

		return nil
	})

	// Move the internal pointer to the first of the list.
	pow2List.First()
	for it, found := pow2List.Get(); found; it, found = pow2List.Advance() {
		fmt.Printf("List item %s\n", it.String())
	}

	/*
	Output:

List item 1
List item 4
List item 9
List item 16
List item 25
	*/
}

// Basic usage
func ExampleList_Filter() {
	list := NewList(true)

	filter := func (it Item) bool {
		return it.(IntItem).value % 2 == 1
	}

	// Create list with the numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		list.AddAfter(It(i))
	}

	newList := list.Filter(filter)

	newList.First()
	for it, cont := newList.Get(); cont; it, cont = newList.Advance() {
		fmt.Printf("List item: %s\n", it)
	}

	/*
	Output:

List item: 1
List item: 3
List item: 5
List item: 7
List item: 9
	*/
}

/*
	Queue
	=====
*/

// Basic usage
func ExampleQueue() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	// Dequeue.
	for item, cont := queue.Dequeue(); cont; item, cont = queue.Dequeue() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 1
Item is number: 2
Item is number: 3
Item is number: 4
Item is number: 5
	*/
}

// Basic usage
func ExampleQueue_Enqueue() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	// Dequeue.
	for item, cont := queue.Dequeue(); cont; item, cont = queue.Dequeue() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 1
Item is number: 2
Item is number: 3
Item is number: 4
Item is number: 5
	*/
}

// Basic usage
func ExampleQueue_Dequeue() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	// Get and remove the first item of the queue.
	for item, cont := queue.Dequeue(); cont; item, cont = queue.Dequeue() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 1
Item is number: 2
Item is number: 3
Item is number: 4
Item is number: 5
	*/
}

// Basic usage
func ExampleQueue_Front() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	// Get and remove the first item of the queue.
	for item, cont := queue.Front(); cont; item, cont = queue.Front() {
		fmt.Printf("Item is number: %s\n", item.String())
		queue.Dequeue()
	}
	/*
	Output:

Item is number: 1
Item is number: 2
Item is number: 3
Item is number: 4
Item is number: 5
	*/
}

// Basic usage
func ExampleQueue_Length() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
		fmt.Printf("Number of items: %d\n", queue.Length())
	}
	/*

	Output:

Number of items: 1
Number of items: 2
Number of items: 3
Number of items: 4
Number of items: 5
	*/
}

// Basic usage
func ExampleQueue_Clear() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(i))
	}

	fmt.Printf("Number of items: %d\n", queue.Length())

	// clear the queue
	queue.Clear()
	fmt.Printf("Number of items: %d\n", queue.Length())

	/*
	Output:

Number of items: 5
Number of items: 0
	*/
}

/*
	Stack
	=====
*/

// Basic usage
func ExampleStack() {
	stack := NewStack()

	// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
	}

	// Get and remove the first item of the stack.
	for item, found := stack.Pop(); found; item, found = stack.Pop() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 5
Item is number: 4
Item is number: 3
Item is number: 2
Item is number: 1
	*/
}

// Basic usage
func ExampleStack_Push() {
	stack := NewStack()

	// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
	}

	// Get and remove the first item of the stack.
	for item, found := stack.Pop(); found; item, found = stack.Pop() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 5
Item is number: 4
Item is number: 3
Item is number: 2
Item is number: 1
	*/
}

func ExampleStack_Pop() {
	stack := NewStack()

	// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
	}

	// Get and remove the first item of the stack.
	for item, found := stack.Pop(); found; item, found = stack.Pop() {
		fmt.Printf("Item is number: %s\n", item.String())
	}
	/*
	Output:

Item is number: 5
Item is number: 4
Item is number: 3
Item is number: 2
Item is number: 1
	*/
}

// Basic usage
func ExampleStack_Top() {
	stack := NewStack()

	// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
	}

	// Get and remove the first item of the stack.
	for item, found := stack.Top(); found; item, found = stack.Top() {
		fmt.Printf("Item is number: %s\n", item.String())
		stack.Pop()
	}
	/*
	Output:

Item is number: 5
Item is number: 4
Item is number: 3
Item is number: 2
Item is number: 1
	*/
}

// Basic usage
func ExampleStack_Length() {
	stack := NewStack()

	// Insert (pus) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
		fmt.Printf("Number of items: %d\n", stack.Length())
	}
	/*

	Output:

Number of items: 1
Number of items: 2
Number of items: 3
Number of items: 4
Number of items: 5
	*/
}

// Basic usage
func ExampleStack_Clear() {
	stack := NewStack()

	// Insert (pus) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(i))
	}

	fmt.Printf("Number of items: %d\n", stack.Length())

	stack.Clear()
	fmt.Printf("Number of items: %d\n", stack.Length())
	/*

	Output:

Number of items: 5
Number of items: 0
	*/
}
