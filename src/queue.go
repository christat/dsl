package dsl

import (
	"errors"
)

/*
Queue is a slice-backed implementation of queues. It takes any type implementing interface{} and
allows:

- Enqueuing: inserting an item into the last position of the queue.

- De-queuing: retrieving the first item in the queue.

Note that the implementation is NOT thread-safe.
*/
type Queue struct {
	Slice []interface{}
}

// NewQueue creates a new queue with initial len() zero and capacity cap.
func NewQueue(cap int) *Queue {
	return &Queue{Slice: make([]interface{}, 0, cap)}
}

// Internal function meant to replace the current slice, copying its contents and resizing it to cap.
func (queue *Queue) resize(cap int) {
	resize := make([]interface{}, len(queue.Slice), cap)
	copy(resize, queue.Slice)
	queue.Slice = resize
}

// Enqueue a new node containing data (interface{}) to the tail of the queue.
func (queue *Queue) Enqueue(data interface{}) {
	queue.Slice = append(queue.Slice, data)
}

// Dequeue the head node of the queue. Returns the data or an error if failed.
func (queue *Queue) Dequeue() (interface{}, error) {
	if len(queue.Slice) > 0 {
		value := queue.Slice[0]
		queue.Slice = queue.Slice[1:]
		//Shrink Slice if 10+ elements but less than half the capacity used
		if length := len(queue.Slice); length > 10 && length < cap(queue.Slice)/2 {
			queue.resize(length)
		}
		return value, nil
	}
	return nil, errors.New("cannot Dequeue() from an empty Queue")
}
