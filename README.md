My Go structs
=============

My Go structs is package that contains a set of differents abstract data types:
* [Queue](#queue)
* [Stack](#stack)
* [List](#list)
* [Sorted list](#sortedlist)
* [Bst](#bst)
* [Avl](#avl)

Available structs
-----------------

### Queue
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#Queue)

Basic usage:
```go
queue := NewQueue()

// Insert (Enqueue) in the queue, the numbers: 1, 2, 3, 4, 5
for i := 1; i <= 5; i++ {
	queue.Enqueue(It(i))
}

// Dequeue.
for item, cont := queue.Dequeue(); cont; item, cont = queue.Dequeue() {
	fmt.Printf("Item is number: %s\n", item.String())
}
// Output:
// Item is number: 1
// Item is number: 2
// Item is number: 3
// Item is number: 4
// Item is number: 5
```

### Stack
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#Stack)

Basic usage:
```go
stack := NewStack()

// Insert (push) in the stack, the numbers: 1, 2, 3, 4, 5
for i := 1; i <= 5; i++ {
	stack.Push(It(i))
}

// Get and remove the first item of the stack.
for item, found := stack.Pop(); found; item, found = stack.Pop() {
	fmt.Printf("Item is number: %s\n", item.String())
}

// Output:
// Item is number: 5
// Item is number: 4
// Item is number: 3
// Item is number: 2
// Item is number: 1
```

### List
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#List)
- List features:
  * Insert in items in any position of the list.
  * Get any item in the list.
  * Update any item in the list.
  * Delete any item in the list.
  * Iterate list.
  * Optimized searchs. Internaly the items are stored in an AVL tree.
  * Create list without duplicated items.

Basic usage:
```go
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

// Output:
// Iterate the list:
// List item: 1
// List item: 2
// List item: 3
// List item: 4
// List item: 5

// Iterate the list in reverse:
// List item: 5
// List item: 4
// List item: 3
// List item: 2
// List item: 1

// Search items:
// Item 3 found.

// Delete items:
// Item 4 deleted.
```

### Sorted list
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#SortedList)

```go
// Create en empty List, that accepts duplicated items.
slist := NewSortedList(true)

// Insert in the list, the numbers: 3, 2, 4, 5, 1
for _, i := range []int{3, 2, 4, 5, 1} {
	slist.Add(It(i))
}

// Iterate the List
fmt.Printf("\nIterate the list:\n")
slist.First()
for item := true; item; item = slist.Next() {
	item, _ := slist.Get()
	fmt.Printf("List item: %s\n", item.String())
}

// Iterate the list in reverse.
fmt.Printf("\nIterate the list in reverse:\n")
slist.Last()
for found := true; found; found = slist.Prev() {
	item, _ := slist.Get()
	fmt.Printf("List item: %s\n", item.String())
}

// Search an item
fmt.Printf("\nSearch items:\n")
if item, found := slist.Search(It(3)); found {
	fmt.Printf("Item %s found.\n", item)
} else {
	fmt.Printf("Item not found.\n")
}

// Delete an item
fmt.Printf("\nDelete items:\n")
if _, found := slist.Search(It(4)); found {
	itemDeleted, _ := slist.Delete()
	fmt.Printf("Item %s deleted.\n", itemDeleted)
}

// Output:
// Iterate the list:
// List item: 1
// List item: 2
// List item: 3
// List item: 4
// List item: 5

// Iterate the list in reverse:
// List item: 5
// List item: 4
// List item: 3
// List item: 2
// List item: 1

// Search items:
// Item 3 found.

// Delete items:
// Item 4 deleted.
```

### Bst
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#Bst)

Basic usage:
```go
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

// Output:
// Item 3 found.
// Item 2 deleted.
```

### Avl
- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#Avl)

Basic usage:
```go
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

// Output:
// Item 3 found.
// Item 2 deleted.
```

Item iterface
-------------
The `Item` interface is the data type used as item in all structs. Any item you want use in the 
structs must implements the this interface. 

- [Official documentation](https://godoc.org/github.com/davidnotplay/my-go-structs#Item)

Example:
```go
/* 
	Example of an item implementation to store numbers in the structs:
*/

// IntItem structs is an implementation of the Item interface specific for storing int numbers.
type IntItem struct {
	value int // number stored
}

// Less checks if the iit item is more less than the item of the parameter.
// The function also returns false if it paramater isn't type IntItem.
func (iit IntItem) Less(it Item) bool {
	iitp, valid := it.(IntItem)
	return valid && iit.value < iitp.value
}

// Eq checks if the iit item is equal to the item of the paramater.
// The function also returns false if it paramater isn't type IntItem.
func (iit IntItem) Eq(it Item) bool {
	iitp, valid := it.(IntItem)
	return valid && iit.value == iitp.value
}

// String returns the number as string.
func (iit IntItem) String() string {
	return fmt.Sprintf("%d", iit.value)
}
```

Official documentation
----------------------
[Official documentation in godoc](https://godoc.org/github.com/davidnotplay/my-go-structs)
