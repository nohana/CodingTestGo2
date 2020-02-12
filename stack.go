package stack

// Stack implements Stackable interface.
type Stack []string

// Push do push into the top of stack.
func (stack *Stack) Push(s string) {
	*stack = append(*stack, s)
}

// Pop do pop from the top of stack and remove it.
func (stack *Stack) Pop() string {
	if len(*stack) == 0 {
		return ""
	}
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}
