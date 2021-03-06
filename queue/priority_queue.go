package gost

import "container/heap"

// PriorityQueue implements a heap-based priority queue, only exposing methods Enqueue() and Dequeue() for simplicity.
// This implementation uses FIFO order as tiebreaker when elements have the same priority.
type PriorityQueue struct {
	contents heapContents
	counter  int // counter ensures FIFO when priority between elements is equal
}

// NewPriorityQueue initializes the heap-based priority queue and returns the instance.
func NewPriorityQueue() (pq *PriorityQueue) {
	pq = new(PriorityQueue)
	heap.Init(&pq.contents)
	return
}

// Enqueue adds an interface item and its priority into the PriorityQueue.
func (pq *PriorityQueue) Enqueue(item interface{}, priority float64) {
	heap.Push(&pq.contents, newPriorityItem(item, priority, pq.counter))
	pq.counter++
}

// Dequeue removes the item in the PriorityQueue with the highest priority, or insertion order when there's no higher priority contents.
// If the queue is empty, returns nil.
func (pq *PriorityQueue) Dequeue() interface{} {
	if pq.Size() == 0 {
		pq.counter = 0 // reset FIFO ordering counter (opportunistic)
		return nil
	}
	return heap.Pop(&pq.contents)
}

// Size returns the size of the PriorityQueue.
func (pq *PriorityQueue) Size() int {
	return pq.contents.Len()
}

/*
	The types defined below implement heap.Interface.
	PriorityQueue is intended to abstract the underlying implementation details.
*/

// priorityItem wraps a value with a priority and an index, required for heap.Interface.
type priorityItem struct {
	value    interface{}
	priority float64
	counter  int // Counter ensures FIFO when priority between elements is equal
	index    int // The index is needed by update and is maintained by the heap.Interface methods.
}

// newPriorityItem is a queue data wrapper, used as item container in heapContents.
// It must be added via heapContents.Enqueue() operator so that it gets an index.
func newPriorityItem(value interface{}, priority float64, counter int) *priorityItem {
	item := priorityItem{}
	item.value = value
	item.counter = counter
	item.priority = priority
	return &item
}

// heapContents implements heap.Interface and holds priorityItems.
type heapContents []*priorityItem

// len returns the length of heapContents.
func (hc heapContents) Len() int { return len(hc) }

// Less responds whether item in index i should be sorted before j (or will take "Less" time to dequeue).
// If two contents have the same priority, the response will be false as it strictly checks for higher priority.
func (hc heapContents) Less(i, j int) bool {
	iPriority, jPriority := hc[i].priority, hc[j].priority
	if iPriority == jPriority {
		return hc[i].counter < hc[j].counter
	}
	return iPriority > jPriority
}

// Swap switches places between both priorityItems in the designated indices.
func (hc heapContents) Swap(i, j int) {
	hc[i], hc[j] = hc[j], hc[i]
	hc[i].index = i
	hc[j].index = j
}

// Enqueue expects an element x of type *priorityItem and appends it to heapContents.
func (hc *heapContents) Push(x interface{}) {
	item := x.(*priorityItem)
	item.index = len(*hc)
	*hc = append(*hc, item)
}

// Dequeue removes the first value to be dequeued from heapContents.
func (hc *heapContents) Pop() interface{} {
	old := *hc
	item := old[len(old)-1]
	*hc = old[0 : len(old)-1]
	return item.value
}
