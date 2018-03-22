package gost_test

import (
	"testing"

	"github.com/christat/gost/stack"
)

// test helper function; initializes and adds size elements to the structure.
func generateStack(size int) *gost.Stack {
	stack := gost.NewStack(10)
	for i := 0; i < size; i++ {
		stack.Push(newVector(i))
	}
	return stack
}

func TestStack_Push(t *testing.T) {
	stack := generateStack(num)
	if stack.Size() != num {
		t.Error("Enqueue did not grow stack size properly")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := gost.NewStack(10)
	stack = generateStack(num)
	value := stack.Pop()
	if value == nil {
		t.Error("Dequeue() failed on non-empty stack")
	}
	if *(value.(*vector)) != *newVector(num - 1) {
		t.Errorf("Dequeue() error: expected: %v, got: %v", newVector(num-1), value)
	}
	for i := 0; i < num-2; i++ {
		_ = stack.Pop()
	}
	value = stack.Pop()
	if value == nil {
		t.Error("Dequeue() failed on stack of size 1")
	}
	value = stack.Pop()
	if value != nil {
		t.Error("Dequeue() did not return nil on empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := gost.NewStack(10)
	value := stack.Peek()
	if value != nil {
		t.Error("Peek() did not return nil on empty stack")
	}
	stack = generateStack(num)
	value = stack.Peek()
	if value == nil {
		t.Error("Peek() failed on non-empty stack")
	}
	if *(value.(*vector)) != *newVector(num - 1) {
		t.Errorf("Peek() error: expected: %v, got: %v", newVector(num-1), value)
	}
}

/*
Stack Benchmark:
The following methods are meant to put this implementation to the test against the linked list backed version.

	- Basic tests fill and subsequently empty the structure with N vector elements (check testUtils.go).
	- Growth tests aim to analyze the performance hit taken due to the array copies made
	by this version against the linked list one.
*/

// benchmark helper function to add num items to the stack.
func fillStack(stack *gost.Stack, num int) {
	for i := 0; i < num; i++ {
		stack.Push(newVector(i))
	}
}

// benchmark helper function to remove num items from the stack.
func emptyStack(stack *gost.Stack, num int) {
	for i := 0; i < num; i++ {
		stack.Pop()
	}
}

func benchmarkStackBasicTest(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := gost.NewStack(10)
		fillStack(stack, len)
		emptyStack(stack, len)
	}
}

func BenchmarkStack_BasicTest10(b *testing.B) {
	benchmarkStackBasicTest(10, b)
}

func BenchmarkStack_BasicTest20(b *testing.B) {
	benchmarkStackBasicTest(20, b)
}

func BenchmarkStack_BasicTest40(b *testing.B) {
	benchmarkStackBasicTest(40, b)
}

func BenchmarkStack_BasicTest80(b *testing.B) {
	benchmarkStackBasicTest(80, b)
}

func BenchmarkStack_BasicTest160(b *testing.B) {
	benchmarkStackBasicTest(160, b)
}

func BenchmarkStack_GrowthDecay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := generateStack(bigNum)
		for {
			emptyStack(stack, bigNum/2)
			if stack.Size() > 0 {
				fillStack(stack, stack.Size()/2)
			} else {
				break
			}
		}
	}
}

func BenchmarkStack_GrowthIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := generateStack(bigNum)
		for stack.Size() <= bigNum {
			emptyStack(stack, num/4)
			fillStack(stack, num/2)
		}
	}

}
