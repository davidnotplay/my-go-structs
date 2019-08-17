package structs

import "fmt"

// Key is used to identify values in the structs.
type Key interface{}

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

	// Key returns the value key.
	Key() Key

	// LessKey checks if the value store in the interface is less that the value with
	// the key of the parameter.
	LessKey(Key) bool

	// EqKey checks if the value stored in the interface is equal that the value with
	// the key of the parameter.
	EqKey(Key) bool
}


// IntValue structs is an implementation of the Value interface specific for storing int numbers.
// The value key is the number itself.
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

// Key returns the key value. The key value is the number itself.
func (iv IntValue) Key() Key {
	return iv.value
}

// LessKey checks if the iv key is less than the k key. k key must be type int,
// if not the function returns false.
func (iv IntValue) LessKey(k Key) bool {
	num, valid := k.(int)
	return valid && iv.value < num
}

// EqKey checks if the iv key is equal than the k key. k key must be type int,
// if not the function returns false.
func (iv IntValue) EqKey(k Key) bool {
	num, valid := k.(int)
	return valid && iv.value == num
}

// Value returns the number stored in iv.
func (iv IntValue) Value() int {
	return iv.value
}

// Iv creates an IntValue object with the number of the param.
func Iv(num int) IntValue{
	return IntValue{num}
}
