package structs

import "fmt"

// Item the interface used as main item in the different structs. Any data that you
// want store in a struct must be implements this interface.
type Item interface {
	// Less checks if the item is more less than the item of the parameter.
	Less(Item) bool

	// Eq checks if the item is Eq to the item of the parameter.
	Eq(Item) bool

	// String transforms the item to string.
	String() string
}

// IntItem structs is an implementation of the Item interface specific for storing int numbers.
type IntItem struct {
	value int // number stored
}

// Less checks if the iit item is more less than the item of the parameter.
// The function also returns false if it paramater isn't type IntItem.
func (iit IntItem) Less(it Item) bool {
	iitp, valid := it.(IntItem)
	return valid &&  iit.value < iitp.value
}

// Eq checks if the iit item is equal to the item of the paramater.
// The function also returns false if it paramater isn't type IntItem.
func (iit IntItem) Eq(it Item) bool {
	iitp, valid := it.(IntItem)
	return valid &&  iit.value == iitp.value
}

// String returns the number as string.
func (iit IntItem) String() string {
	return fmt.Sprintf("%d", iit.value)
}

// Value returns the number stored in iit.
func (iit IntItem) Value() int {
	return iit.value
}

// It creates an IntItem object with the number of the param.
func It(num int) IntItem{
	return IntItem{num}
}
