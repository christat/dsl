package gost_test

import (
	"testing"

	"github.com/christat/gost/queue"
)

// test helper function; initializes and adds size elements to the structure.
func generateQueue(size int) *gost.Queue {
	queue := gost.NewQueue(10)
	for i := 0; i < size; i++ {
		queue.Enqueue(newVector(i))
	}
	return queue
}

func TestQueue_Enqueue(t *testing.T) {
	queue := generateQueue(num)
	if queue.Size() != num {
		t.Error("Enqueue() did not grow queue size properly")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := generateQueue(num)
	value := queue.Dequeue()
	if value == nil {
		t.Error("Dequeue() failed on non-empty queue")
	}
	if *(value.(*vector)) != *newVector(0) {
		t.Errorf("Dequeue() error: expected %v, got %v", value, newVector(0))
	}
	for i := 0; i < num-2; i++ {
		queue.Dequeue()
	}
	value = queue.Dequeue()
	if value == nil {
		t.Error("Dequeue() failed on queue of size 1")
	}
	if *(value.(*vector)) != *newVector(num - 1) {
		t.Errorf("Dequeue() error: expected %v, got %v", value, newVector(num-1))
	}
	value = queue.Dequeue()
	if value != nil {
		t.Error("Dequeue() did not return nil on empty queue")
	}
}

/*
Queue Benchmark:
The following methods are meant to put this implementation to the test against the linked list backed version.

	- Basic tests fill and subsequently empty the structure with N vector elements (check testUtils.go).
	- Growth tests aim to analyze the performance hit taken due to the array copies made
	by this version against the linked list one.
*/

// benchmark helper function to add num items to the queue.
func fillQueue(queue *gost.Queue, num int) {
	for i := 0; i < num; i++ {
		queue.Enqueue(newVector(i))
	}
}

// benchmark helper function to remove num items from the queue.
func emptyQueue(queue *gost.Queue, num int) {
	for i := 0; i < num; i++ {
		queue.Dequeue()
	}
}

func benchmarkQueueBasicTest(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := gost.NewQueue(10)
		fillQueue(queue, len)
		emptyQueue(queue, len)
	}
}

func BenchmarkQueue_BasicTest10(b *testing.B) {
	benchmarkQueueBasicTest(10, b)
}

func BenchmarkQueue_BasicTest20(b *testing.B) {
	benchmarkQueueBasicTest(20, b)
}

func BenchmarkQueue_BasicTest40(b *testing.B) {
	benchmarkQueueBasicTest(40, b)
}

func BenchmarkQueue_BasicTest80(b *testing.B) {
	benchmarkQueueBasicTest(80, b)
}

func BenchmarkQueue_BasicTest160(b *testing.B) {
	benchmarkQueueBasicTest(160, b)
}

func BenchmarkQueue_GrowthDecay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := generateQueue(bigNum)
		for {
			emptyQueue(queue, bigNum/2)
			if queue.Size() > 0 {
				fillQueue(queue, queue.Size()/2)
			} else {
				break
			}
		}
	}
}

func BenchmarkQueue_GrowthIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := generateQueue(bigNum)
		for queue.Size() <= bigNum {
			emptyQueue(queue, num/4)
			fillQueue(queue, num/2)
		}
	}

}
