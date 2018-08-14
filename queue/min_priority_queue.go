package gost

import "container/heap"

// MinPriorityQueue implements a heap-based priority queue, only exposing methods Enqueue() and Dequeue() for simplicity.
// Inverse priority means that items with lower priority are dequeued faster than higher priority ones.
// This implementation uses FIFO order as tiebreaker when elements have the same priority.
type MinPriorityQueue struct {
	contents MinHeapContents
	counter  int // counter ensures FIFO when priority between elements is equal
}

// NewMinPriorityQueue initializes the heap-based priority queue and returns the instance.
func NewMinPriorityQueue() (pq *MinPriorityQueue) {
	pq = new(MinPriorityQueue)
	heap.Init(&pq.contents)
	return
}

// Enqueue adds an interface item and its priority into the MinPriorityQueue.
func (pq *MinPriorityQueue) Enqueue(item interface{}, priority float64) {
	heap.Push(&pq.contents, newPriorityItem(item, priority, pq.counter))
	pq.counter++
}

// Dequeue removes the item in the MinPriorityQueue with the lowest priority, or insertion order when there's no lower priority contents.
// If the queue is empty, returns nil.
func (pq *MinPriorityQueue) Dequeue() interface{} {
	if pq.Size() == 0 {
		pq.counter = 0 // reset FIFO ordering counter (opportunistic)
		return nil
	}
	return heap.Pop(&pq.contents)
}

// Size returns the size of the MinPriorityQueue.
func (pq *MinPriorityQueue) Size() int {
	return pq.contents.Len()
}

/*
	The types defined below implement heap.Interface.
	MinPriorityQueue is intended to abstract the underlying implementation details.
*/

// heapContents implements heap.Interface and holds priorityItems.
type MinHeapContents []*priorityItem

// len returns the length of heapContents.
func (mhc MinHeapContents) Len() int { return len(mhc) }

// Less responds whether item in index i should be sorted before j (or will take "Less" time to dequeue).
// If two contents have the same priority, the response will be false as it strictly checks for higher priority.
func (mhc MinHeapContents) Less(i, j int) bool {
	iPriority, jPriority := mhc[i].priority, mhc[j].priority
	if iPriority == jPriority {
		return mhc[i].counter < mhc[j].counter
	}
	return iPriority < jPriority
}

// Swap switches places between both priorityItems in the designated indices.
func (mhc MinHeapContents) Swap(i, j int) {
	mhc[i], mhc[j] = mhc[j], mhc[i]
	mhc[i].index = i
	mhc[j].index = j
}

// Enqueue expects an element x of type *priorityItem and appends it to heapContents.
func (mhc *MinHeapContents) Push(x interface{}) {
	item := x.(*priorityItem)
	item.index = len(*mhc)
	*mhc = append(*mhc, item)
}

// Dequeue removes the first value to be dequeued from heapContents.
func (mhc *MinHeapContents) Pop() interface{} {
	old := *mhc
	item := old[len(old)-1]
	*mhc = old[0 : len(old)-1]
	return item.value
}
