package go_structures

import (
	"errors"
)

// Stack which is exposed to be used by other packages
type Stack struct {
	top  *element
	size int
}

// Adds any value to the stack
func (stack *Stack) Push(value interface{}) {
	stack.top = &element{value, stack.top}
	stack.size++
}

// Removes the top element from the stack and returns it
// Returns an error if the stack is empty
func (stack *Stack) Pop() (interface{}, error) {
	if stack.size == 0 {
		return nil, errors.New("Stack empty")
	}

	value := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return value, nil
}

// Returns the current size of the stack
func (stack *Stack) Size() int {
	return stack.size
}
