package structs

import "fmt"

// Value is the interface used as main value in the different structs. Any data that you
// want store in a struct must be implements this interface.
type Value interface {
	// Less checks if the value stored in the interface is equal that the value of
	// the parameter.
	Less(Value) bool

	// Eq checks if the value stored in the interface is equal that the value of
	// the parameter.
	Eq(Value) bool

	// String transforms the value to string.
	String() string
}

// IntValue structs is an implementation of the Value interface specific for storing int numbers.
type IntValue struct {
	value int // number stored
}

// Less checks if the iv number is less than the v value. v value must be type IntValue, if not
// the function returns false.
func (iv IntValue) Less(v Value) bool {
	intvalue, valid := v.(IntValue)
	return valid &&  iv.value < intvalue.value
}

// Eq checks if the iv number is equal than the number in v value. v value must be type IntValue,
// if not the function returns false.
func (iv IntValue) Eq(v Value) bool {
	intvalue, valid := v.(IntValue)
	return valid &&  iv.value == intvalue.value
}

// String returns the number as string.
func (iv IntValue) String() string {
	return fmt.Sprintf("%d", iv.value)
}

// Value returns the number stored in iv.
func (iv IntValue) Value() int {
	return iv.value
}

// Iv creates an IntValue object with the number of the param.
func Iv(num int) IntValue{
	return IntValue{num}
}
