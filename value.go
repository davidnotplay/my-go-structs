package structs

type Key interface{}

type Value interface {
	Less(Value) bool
	Eq(Value) bool
	String() string

	Key() Key
	EqKey(Key) bool
	LessKey(Key) bool
}
