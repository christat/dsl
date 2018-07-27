package gost

import "container/heap"

// InversePriorityQueue implements a heap-based priority queue, only exposing methods Enqueue() and Dequeue() for simplicity.
// Inverse priority means that items with lower priority are dequeued faster than higher priority ones.
// This implementation uses FIFO order as tiebreaker when elements have the same priority.
type InversePriorityQueue struct {
	contents InverseHeapContents
	counter  int // counter ensures FIFO when priority between elements is equal
}

// NewInversePriorityQueue initializes the heap-based priority queue and returns the instance.
func NewInversePriorityQueue() (pq *InversePriorityQueue) {
	pq = new(InversePriorityQueue)
	heap.Init(&pq.contents)
	return
}

// Enqueue adds an interface item and its priority into the InversePriorityQueue.
func (pq *InversePriorityQueue) Enqueue(item interface{}, priority float64) {
	heap.Push(&pq.contents, newPriorityItem(item, priority, pq.counter))
	pq.counter++
}

// Dequeue removes the item in the InversePriorityQueue with the lowest priority, or insertion order when there's no lower priority contents.
// If the queue is empty, returns nil.
func (pq *InversePriorityQueue) Dequeue() interface{} {
	if pq.Size() == 0 {
		pq.counter = 0 // reset FIFO ordering counter (opportunistic)
		return nil
	}
	return heap.Pop(&pq.contents)
}

// Size returns the size of the InversePriorityQueue.
func (pq *InversePriorityQueue) Size() int {
	return pq.contents.Len()
}

/*
	The types defined below implement heap.Interface.
	InversePriorityQueue is intended to abstract the underlying implementation details.
*/

// heapContents implements heap.Interface and holds priorityItems.
type InverseHeapContents []*priorityItem

// len returns the length of heapContents.
func (ihc InverseHeapContents) Len() int { return len(ihc) }

// Less responds whether item in index i should be sorted before j (or will take "Less" time to dequeue).
// If two contents have the same priority, the response will be false as it strictly checks for higher priority.
func (ihc InverseHeapContents) Less(i, j int) bool {
	iPriority, jPriority := ihc[i].priority, ihc[j].priority
	if iPriority == jPriority {
		return ihc[i].counter < ihc[j].counter
	}
	return iPriority < jPriority
}

// Swap switches places between both priorityItems in the designated indices.
func (ihc InverseHeapContents) Swap(i, j int) {
	ihc[i], ihc[j] = ihc[j], ihc[i]
	ihc[i].index = i
	ihc[j].index = j
}

// Enqueue expects an element x of type *priorityItem and appends it to heapContents.
func (ihc *InverseHeapContents) Push(x interface{}) {
	item := x.(*priorityItem)
	item.index = len(*ihc)
	*ihc = append(*ihc, item)
}

// Dequeue removes the first value to be dequeued from heapContents.
func (ihc *InverseHeapContents) Pop() interface{} {
	old := *ihc
	item := old[len(old)-1]
	*ihc = old[0 : len(old)-1]
	return item.value
}
