package gost

import (
	"errors"
)

/*
Stack is a slice-backed implementation of stacks. It takes any type implementing interface{} and
allows:

- Pushing: adding a new element on top of the stack.

- Popping: retrieving the element on top of the stack.

- Peeking: obtaining the element on top of the stack without removing it.

Note that the implementation is NOT thread-safe.
*/
type Stack struct {
	slice []interface{}
}

// NewStack creates a new stack with initial len() zero and capacity cap.
func NewStack(cap int) *Stack {
	return &Stack{slice: make([]interface{}, 0, cap)}
}

// Internal function meant to replace the current slice, copying its contents and resizing it to cap.
func (stack *Stack) resize(cap int) {
	resize := make([]interface{}, len(stack.slice), cap)
	copy(resize, stack.slice)
	stack.slice = resize
}

// Push a new node containing data of type interface{} into the stack.
func (stack *Stack) Push(data interface{}) {
	stack.slice = append(stack.slice, data)
}

// Pop the head node from the stack. Returns the data or an error value if failed.
func (stack *Stack) Pop() (interface{}, error) {
	if len(stack.slice) > 0 {
		value := stack.slice[len(stack.slice)-1]
		stack.slice = stack.slice[:len(stack.slice)-1]
		//Shrink Slice if 10+ elements but less than half the capacity used
		if length := len(stack.slice); length > 10 && length <= cap(stack.slice)/2 {
			stack.resize(length)
		}
		return value, nil
	}
	return nil, errors.New("cannot Pop() from an empty Stack")
}

// Peek at the content of the stack Head (nil if empty) without removing it afterwards.
func (stack *Stack) Peek() (interface{}, error) {
	if len(stack.slice) > 0 {
		return stack.slice[len(stack.slice)-1], nil
	}
	return nil, errors.New("cannot Peek() into an empty Stack")
}

// Size returns the current length of the stack's underlying slice.
func (stack *Stack) Size() int {
	return len(stack.slice)
}
