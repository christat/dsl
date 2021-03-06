package gost_test

import (
	"testing"

	"github.com/christat/gost/queue"
)

// test helper function; initializes and adds size elements to the structure.
func generateEqualInverseQueue(size int) (queue gost.MinPriorityQueue) {
	for i := 0; i < size; i++ {
		queue.Enqueue(newVector(i), 1)
	}
	return
}

func TestMinPriorityQueue_Len(t *testing.T) {
	queue := gost.MinPriorityQueue{}
	if queue.Size() != 0 {
		t.Error("len() failed to return empty queue length")
	}

	queue = generateEqualInverseQueue(10)
	if queue.Size() != 10 {
		t.Errorf("len() length: %v, expected: %v", queue.Size(), 10)
	}
}

func TestMinPriorityQueue_Push(t *testing.T) {
	queue := gost.NewMinPriorityQueue()
	queue.Enqueue("last", 0)
	queue.Enqueue("middle", 1)
	queue.Enqueue("first", 2)

	if queue.Size() != 3 {
		t.Error("Enqueue() failed to add items")
	}
}

func TestMinPriorityQueue_Pop(t *testing.T) {
	pq := gost.NewMinPriorityQueue()
	pq.Enqueue("a", 0)
	pq.Enqueue("b", 5)
	pq.Enqueue("c", 10)
	pq.Enqueue("d", 5)

	value := pq.Dequeue().(string)
	if value != "a" {
		t.Errorf("Dequeue() failed: returned: %v, expected: %v", value, "a")
	}

	value = pq.Dequeue().(string)
	if value != "b" {
		t.Errorf("Dequeue() failed: returned: %v, expected: %v", value, "b")
	}

	value = pq.Dequeue().(string)
	if value != "d" {
		t.Errorf("Dequeue() failed: returned: %v, expected: %v", value, "d")
	}

	value = pq.Dequeue().(string)
	if value != "c" {
		t.Errorf("Dequeue() failed: returned: %v, expected: %v", value, "c")
	}

	if pq.Size() != 0 {
		t.Error("Dequeue() failed: MinPriorityQueue should be empty")
	}

	null := pq.Dequeue()
	if null != nil {
		t.Error("Dequeue() failed: MinPriorityQueue returned non-nil value when empty")
	}
}
