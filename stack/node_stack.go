package gost

import (
	"github.com/christat/gost/list"
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
	Head *gost.Node
	Size uint
}

// Push a new node containing data (interface{}) into the stack.
func (stack *NodeStack) Push(data interface{}) {
	head := &gost.Node{Data: data, Next: stack.Head}
	stack.Head = head
	stack.Size++
}

// Pop the head node from the stack. Returns the data or nil if empty.
func (stack *NodeStack) Pop() interface{} {
	if stack.Size > 0 {
		data := stack.Head.Data
		stack.Head = stack.Head.Next
		stack.Size--
		return data
	}
	return nil
}

// Peek at the content of the stack head (nil if empty) without removing it afterwards.
func (stack *NodeStack) Peek() interface{} {
	if stack.Size > 0 {
		return stack.Head.Data
	}
	return nil
}
