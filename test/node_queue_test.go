package gost_test

import (
	"testing"

	"github.com/christat/gost/queue"
)

// test helper function; initializes and adds size elements to the structure.
func generateNodeQueue(size uint) *gost.NodeQueue {
	queue := new(gost.NodeQueue)
	for i := 0; uint(i) < size; i++ {
		queue.Enqueue(newVector(i))
	}
	return queue
}

func TestNodeQueue_Enqueue(t *testing.T) {
	queue := generateNodeQueue(num)
	if queue.Size != num {
		t.Errorf("Enqueue() error; queue size expected: %v, got: %v", num, queue.Size)
	}
}

func TestNodeQueue_Dequeue(t *testing.T) {
	queue := generateNodeQueue(num)
	// Dequeue half the contents
	for i := 0; i < num/2; i++ {
		val := queue.Dequeue()
		if val == nil {
			t.Error("Dequeue() failed")
		}
		if i == num/4 {
			if *(val.(*vector)) != *newVector(num / 4) {
				t.Errorf("Dequeue() error; expected to get: %v, got: %v", newVector(num/4), val)
			}
		}
	}
	if queue.Size != num/2 {
		t.Errorf("Enqueue() error; queue size expected: %v, got: %v", num/2, queue.Size)
	}
	// Empty the queue
	for i := 0; i < num/2-1; i++ {
		val := queue.Dequeue()
		if val == nil {
			t.Error("Dequeue() failed")
		}
	}
	val := queue.Dequeue()
	if queue.Size != 0 {
		t.Errorf("Dequeue() error; queue size expected: %v, got: %v", 0, queue.Size)
	}
	if *(val.(*vector)) != *newVector(num - 1) {
		t.Errorf("Dequeue() error; expected to get: %v, got: %v", newVector(num-1), val)
	}
	val = queue.Dequeue()
	if val != nil {
		t.Error("Dequeue() error: On empty queue, no error was returned")
	}
}

/*
NodeQueue Benchmark:
The following methods are meant to put this implementation to the test against the slice-backed version.

	- Basic tests fill and subsequently empty the structure with N vector elements (check testUtils.go).
	- Growth tests aim to analyze the performance hit taken due to the array copies made
	by the slice-backed version against this one.
*/

// benchmark helper function to add num items to the queue.
func fillNodeQueue(queue *gost.NodeQueue, num int) {
	for i := 0; i < num; i++ {
		queue.Enqueue(newVector(i))
	}
}

// benchmark helper function to remove num items from the queue.
func emptyNodeQueue(queue *gost.NodeQueue, num int) {
	for i := 0; i < num; i++ {
		queue.Dequeue()
	}
}

func benchmarkNodeQueueBasicTest(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := new(gost.NodeQueue)
		fillNodeQueue(queue, len)
		emptyNodeQueue(queue, len)
	}
}

func BenchmarkNodeQueue_BasicTest10(b *testing.B) {
	benchmarkNodeQueueBasicTest(10, b)
}

func BenchmarkNodeQueue_BasicTest20(b *testing.B) {
	benchmarkNodeQueueBasicTest(20, b)
}

func BenchmarkNodeQueue_BasicTest40(b *testing.B) {
	benchmarkNodeQueueBasicTest(40, b)
}

func BenchmarkNodeQueue_BasicTest80(b *testing.B) {
	benchmarkNodeQueueBasicTest(80, b)
}

func BenchmarkNodeQueue_BasicTest160(b *testing.B) {
	benchmarkNodeQueueBasicTest(160, b)
}

func BenchmarkNodeQueue_GrowthDecay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := generateNodeQueue(bigNum)
		for {
			emptyNodeQueue(queue, bigNum/2)
			if queue.Size > 0 {
				fillNodeQueue(queue, int(queue.Size/2))
			} else {
				break
			}
		}
	}
}

func BenchmarkNodeQueue_GrowthIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := generateNodeQueue(bigNum)
		for queue.Size <= bigNum {
			emptyNodeQueue(queue, num/4)
			fillNodeQueue(queue, num/2)
		}
	}

}
