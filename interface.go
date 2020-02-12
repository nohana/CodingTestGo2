package stack

// Stackable is a string stack.
type Stackable interface {
	Push(string)
	Pop() string
}
