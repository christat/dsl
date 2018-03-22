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
	head *gost.Node
	size int
}

// Enqueue a new node containing data (interface{}) into the stack.
func (stack *NodeStack) Push(data interface{}) {
	head := &gost.Node{Data: data, Next: stack.head}
	stack.head = head
	stack.size++
}

// Dequeue the head node from the stack. Returns the data or nil if empty.
func (stack *NodeStack) Pop() interface{} {
	if stack.size > 0 {
		data := stack.head.Data
		stack.head = stack.head.Next
		stack.size--
		return data
	}
	return nil
}

// Peek at the content of the stack head (nil if empty) without removing it afterwards.
func (stack *NodeStack) Peek() interface{} {
	if stack.size > 0 {
		return stack.head.Data
	}
	return nil
}

// Size returns the depth of the current NodeStack.
func (stack *NodeStack) Size() int {
	return stack.size
}
