package gost_test

import (
	"testing"

	"github.com/christat/gost/stack"
)

// test helper function; initializes and adds size elements to the structure.
func generateNodeStack(size int) *gost.NodeStack {
	stack := &gost.NodeStack{}
	for i := 0; i < size; i++ {
		stack.Push(newVector(i))
	}
	return stack
}

func TestNodeStack_Push(t *testing.T) {
	stack := new(gost.NodeStack)
	elem := newVector(0)
	stack.Push(elem)
	stack.Push(elem)
	if stack.Size != 2 {
		t.Error("Push() on stack failed, change undetected")
	}
}

func TestNodeStack_Pop(t *testing.T) {
	stack := new(gost.NodeStack)
	value := stack.Pop()
	if value != nil {
		t.Error("Pop() did not return nil on empty stack")
	}
	elem := newVector(0)
	stack.Push(elem)
	value = stack.Pop()
	if value == nil {
		t.Error("Pop() failed")
	}
	if *(value.(*vector)) != *elem {
		t.Errorf("Pop() error: expected %v, got %v", elem, value)
	}
	stack.Push(elem)
	stack.Push(elem)
	stack.Pop()
	stack.Push(elem)
	if stack.Size != 2 {
		t.Error("Stack size not updated properly")
	}
}

func TestNodeStack_Peek(t *testing.T) {
	stack := new(gost.NodeStack)
	value := stack.Peek()
	if value != nil {
		t.Error("Peek() did not return nil on empty stack")
	}
	elem := newVector(0)
	stack.Push(elem)
	value = stack.Peek()
	if value == nil {
		t.Error("Peek() failed")
	}
	if *(value.(*vector)) != *elem {
		t.Errorf("Peek() error: expected %v, got %v", elem, value)
	}
}

/*
NodeStack Benchmark:
The following methods are meant to put this implementation to the test against the slice-backed version.

	- Basic tests fill and subsequently empty the structure with N vector elements (check testUtils.go).
	- Growth tests aim to analyze the performance hit taken due to the array copies made
	by the slice-backed version against this one.
*/

// benchmark helper function to add num items to the stack.
func fillNodeStack(stack *gost.NodeStack, num int) {
	for i := 0; i < num; i++ {
		stack.Push(newVector(i))
	}
}

// benchmark helper function to remove num items from the stack.
func emptyNodeStack(stack *gost.NodeStack, num int) {
	for i := 0; i < num; i++ {
		stack.Pop()
	}
}

func benchmarkNodeStackBasicTest(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := new(gost.NodeStack)
		fillNodeStack(stack, len)
		emptyNodeStack(stack, len)
	}
}

func BenchmarkNodeStack_BasicTest10(b *testing.B) {
	benchmarkNodeStackBasicTest(10, b)
}

func BenchmarkNodeStack_BasicTest20(b *testing.B) {
	benchmarkNodeStackBasicTest(20, b)
}

func BenchmarkNodeStack_BasicTest40(b *testing.B) {
	benchmarkNodeStackBasicTest(40, b)
}

func BenchmarkNodeStack_BasicTest80(b *testing.B) {
	benchmarkNodeStackBasicTest(80, b)
}

func BenchmarkNodeStack_BasicTest160(b *testing.B) {
	benchmarkNodeStackBasicTest(160, b)
}

func BenchmarkNodeStack_GrowthDecay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := generateNodeStack(bigNum)
		for {
			emptyNodeStack(stack, bigNum/2)
			if stack.Size > 0 {
				fillNodeStack(stack, int(stack.Size/2))
			} else {
				break
			}
		}
	}
}

func BenchmarkNodeStack_GrowthIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := generateNodeStack(bigNum)
		for stack.Size <= bigNum {
			emptyNodeStack(stack, num/4)
			fillNodeStack(stack, num/2)
		}
	}

}
