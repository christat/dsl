package gost_test

import (
	"testing"

	"github.com/christat/gost/queue"
)

// test helper function; initializes and adds size elements to the structure.
func generateEqualQueue(size int) (queue gost.PriorityQueue) {
	for i := 0; i < size; i++ {
		queue.Push(newVector(i), 1)
	}
	return
}

func TestPriorityQueue_Len(t *testing.T) {
	queue := gost.PriorityQueue{}
	if queue.Len() != 0 {
		t.Error("Len() failed to return empty queue length")
	}

	queue = generateEqualQueue(10)
	if queue.Len() != 10 {
		t.Errorf("Len() length: %v, expected: %v", queue.Len(), 10)
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	queue := gost.NewPriorityQueue()
	queue.Push("last", 0)
	queue.Push("middle", 1)
	queue.Push("first", 2)

	if queue.Len() != 3 {
		t.Error("Push() failed to add items")
	}
}

func TestPriorityQueue_Pop(t *testing.T) { //TODO
	pq := gost.NewPriorityQueue()
	pq.Push("a", 0)
	pq.Push("b", 5)
	pq.Push("c", 10)
	pq.Push("d", 5)

	value := pq.Pop().(string)
	if value != "c" {
		t.Errorf("Pop() failed: returned: %v, expected: %v", value, "c")
	}

	value = pq.Pop().(string)
	if value != "b" {
		t.Errorf("Pop() failed: returned: %v, expected: %v", value, "b")
	}

	value = pq.Pop().(string)
	if value != "d" {
		t.Errorf("Pop() failed: returned: %v, expected: %v", value, "d")
	}

	value = pq.Pop().(string)
	if value != "a" {
		t.Errorf("Pop() failed: returned: %v, expected: %v", value, "a")
	}

	if pq.Len() != 0 {
		t.Error("Pop() failed: PriorityQueue should be empty")
	}

	null := pq.Pop()
	if null != nil {
		t.Error("Pop() failed: PriorityQueue returned non-nil value when empty")
	}
}
