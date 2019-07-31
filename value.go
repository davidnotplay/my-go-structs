package structs

type Value interface {
	Less(Value) bool
	Eq(Value)   bool
	String()    string
}
