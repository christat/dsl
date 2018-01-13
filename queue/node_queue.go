package gost

import (
	"errors"
	"github.com/christat/gost/list"
)

/*
NodeQueue is a single-linked list backed implementation of queues. It takes any interface{} and
allows:

- Enqueuing: inserting an item into the last position of the queue.

- De-queuing: retrieving the first item in the queue.

Note that the implementation is NOT thread-safe.
*/
type NodeQueue gost.NodeList

// Enqueue a new Node containing data (interface{}) to the tail of the queue.
func (queue *NodeQueue) Enqueue(data interface{}) {
	node := &gost.Node{Data: data, Next: nil}
	if queue.Size > 1 {
		queue.Tail.Next = node
	} else if queue.Size == 0 {
		queue.Head = node
	} else {
		queue.Head.Next = node
	}
	queue.Tail = node
	queue.Size++
}

// Dequeue the head node of the queue. Returns the data or an error value if failed.
func (queue *NodeQueue) Dequeue() (interface{}, error) {
	if queue.Size > 0 {
		data := queue.Head.Data
		next := queue.Head.Next
		queue.Head.Next = nil
		queue.Head = next
		queue.Size--
		return data, nil
	}
	return nil, errors.New("cannot Dequeue() from an empty NodeQueue")
}
