package gost

/*
Queue is a slice-backed implementation of queues. It takes any type implementing interface{} and
allows:

- Enqueuing: inserting an item into the last position of the queue.

- De-queuing: retrieving the first item in the queue.

Note that the implementation is NOT thread-safe.
*/
type Queue struct {
	slice []interface{}
}

// NewQueue creates a new queue with initial len() zero and capacity cap.
func NewQueue(cap int) *Queue {
	return &Queue{slice: make([]interface{}, 0, cap)}
}

// Internal function meant to replace the current slice, copying its contents and resizing it to cap.
func (queue *Queue) resize(cap int) {
	resize := make([]interface{}, len(queue.slice), cap)
	copy(resize, queue.slice)
	queue.slice = resize
}

// Enqueue a new node containing data (interface{}) to the tail of the queue.
func (queue *Queue) Enqueue(data interface{}) {
	queue.slice = append(queue.slice, data)
}

// Dequeue the head node of the queue. Returns the data or nil if empty.
func (queue *Queue) Dequeue() interface{} {
	if len(queue.slice) > 0 {
		value := queue.slice[0]
		queue.slice = queue.slice[1:]
		// Shrink Slice if 10+ elements but less than half the capacity used
		if length := len(queue.slice); length > 10 && length < cap(queue.slice)/2 {
			queue.resize(length)
		}
		return value
	}
	return nil
}

// Size returns the length of the queue's underlying slice.
func (queue *Queue) Size() int {
	return len(queue.slice)
}
