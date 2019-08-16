package structs

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

	// EqKey checks if the value stored in the interface is equal that the value with
	// the key of the parameter.
	EqKey(Key) bool

	// LessKey checks if the value store in the interface is less that the value with
	// the key of the parameter.
	LessKey(Key) bool
}
