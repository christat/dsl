package dsl

import (
	"errors"
)

/*
NodeStack is a single-linked list backed implementation of stacks. It takes any interface{} and
allows:

- Pushing: adding a new element on top of the stack.

- Popping: retrieving the element on top of the stack.

- Peeking: obtaining the element on top of the stack without removing it.

Note that the implementation is NOT thread-safe.
*/
type NodeStack struct {
	Head *Node
	Size uint
}

// Push a new node containing data (interface{}) into the stack.
func (stack *NodeStack) Push(data interface{}) {
	head := &Node{Data: data, Next: stack.Head}
	stack.Head = head
	stack.Size++
}

// Pop the head node from the stack. Returns the data or an error value if failed.
func (stack *NodeStack) Pop() (interface{}, error) {
	if stack.Size > 0 {
		data := stack.Head.Data
		stack.Head = stack.Head.Next
		stack.Size--
		return data, nil
	}
	return nil, errors.New("cannot Pop() from an empty NodeStack")
}

// Peek at the content of the stack head ((nil, error) if empty) without removing it afterwards.
func (stack *NodeStack) Peek() (interface{}, error) {
	if stack.Size > 0 {
		return stack.Head.Data, nil
	}
	return nil, errors.New("cannot Peek() into an empty NodeStack")
}
