package mygostructs

import "fmt"

// Basic usage
func ExampleAvl() {
	// create new avl
	avl := NewAvl()

	for i := 1; i <= 5; i++ {
		avl.Insert(It(1))
	}

	// Search the item 3
	if item, found := avl.Search(It(3)); found {
		fmt.Printf("Item %s found", item.String())
	}
	/*
	Output:

		Item 3 found
	*/

	// Delete the item 2
	if itemDeleted, deleted := avl.Delete(It(2)); deleted {
		fmt.Printf("Item %s deleted", itemDeleted.String())
	}

	/**
	Output:

		Item 2 deleted
	*/
}

// Basic usage
func ExampleBst() {
	// create new bst
	bst := NewBst()

	for i := 1; i <= 5; i++ {
		bst.Insert(It(1))
	}

	// Search the item 3
	if item, found := bst.Search(It(3)); found {
		fmt.Printf("Item %s found", item.String())
	}
	/*
	Output:

		Item 3 found
	*/

	// Delete the item 2
	if itemDeleted, deleted := bst.Delete(It(2)); deleted {
		fmt.Printf("Item %s deleted", itemDeleted.String())
	}
	/**
	Output:

		Item 2 deleted
	*/
}

// Basic usage
func ExampleList() {
	// Create en empty List, that accepts duplicated items.
	list := NewList(true)

	// Insert in the list, the numbers: 1, 2, 3, 4, 5
	for _, i := range []int{1, 2, 3, 4, 5} {
		list.AddAfter(It(i))
	}


	// Iterate the List

	// Move the internal pointer to the first of the list.
	list.First()
	for it, found := list.Get(); found; found = list.Next() {
		fmt.Printf("List item %s", it.String())
	}

	/**
	Output:

		 List item 1
		 List item 2
		 List item 3
		 List item 4
		 List item 5
	*/

	// Iterate the list reversed.

	// Move the internal pointer to the end of the list.
	list.Last()
	for it, found := list.Get(); found; found = list.Prev() {
		fmt.Printf("List item %s", it.String())
	}

	/**
	Output:

		 List item 5
		 List item 4
		 List item 3
		 List item 2
		 List item 1
	*/

	// Search an item
	if item, found := list.Search(It(3)); found {
		println("Item found ", item.String())
	} else {
		println("Item not found")
	}

	// Delete the item 4
	if _, found := list.Search(It(4)); found {
		itemDeleted, _ := list.Delete();
		println("Item %s deleted", itemDeleted.String())
	}

	/**
	Output:

		Item 4 deleted.
	*/
}

// Basic usage
func ExampleQueue() {
	queue := NewQueue()

	// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		queue.Enqueue(It(5))
	}


	// Get and remove the first item of the queue.
	for item, found := queue.Dequeue(); found; item, found = queue.Dequeue() {
		fmt.Printf("Item is number is: %s", item.String())
	}
	/**
	Output:

		 Item is number: 1
		 Item is number: 2
		 Item is number: 3
		 Item is number: 4
		 Item is number: 5
	*/
}

// Basic usage
func ExampleStack() {
	stack := NewStack()

	// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
	for i := 1; i <= 5; i++ {
		stack.Push(It(5))
	}


	// Get and remove the first item of the stack.
	for item, found := stack.Pop(); found; item, found = stack.Pop() {
		fmt.Printf("Item is number is: %s", item.String())
	}
	/**
	Output:

		Item is number: 5
		Item is number: 4
		Item is number: 3
		Item is number: 2
		Item is number: 1
	*/
}