package gost

import (
	"github.com/christat/gost/list"
)

/*
NodeQueue is a single-linked contents backed implementation of queues. It takes any interface{} and
allows:

- Enqueuing: inserting an item into the last position of the queue.

- De-queuing: retrieving the first item in the queue.

Note that the implementation is NOT thread-safe.
*/
type NodeQueue struct {
	head *gost.Node
	tail *gost.Node
	size int
}

// Enqueue a new Node containing data (interface{}) to the tail of the queue.
func (queue *NodeQueue) Enqueue(data interface{}) {
	node := &gost.Node{Data: data, Next: nil}
	if queue.size > 1 {
		queue.tail.Next = node
	} else if queue.size == 0 {
		queue.head = node
	} else {
		queue.head.Next = node
	}
	queue.tail = node
	queue.size++
}

// Dequeue the head node of the queue. Returns the data or nil if empty.
func (queue *NodeQueue) Dequeue() interface{} {
	if queue.size > 0 {
		data := queue.head.Data
		next := queue.head.Next
		queue.head.Next = nil
		queue.head = next
		queue.size--
		return data
	}
	return nil
}

// Size returns the length of the NodeQueue.
func (queue *NodeQueue) Size() int {
	return queue.size
}
