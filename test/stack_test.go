package dsl_test

import (
	"testing"

	"github.com/christat/dsl/src"
)

// test helper function; initializes and adds size elements to the structure.
func generateStack(size int) *dsl.Stack {
	stack := dsl.NewStack(10)
	for i := 0; i < size; i++ {
		stack.Push(newVector(i))
	}
	return stack
}

func TestStack_Push(t *testing.T) {
	stack := generateStack(num)
	if len(stack.Slice) != num {
		t.Error("Push did not grow stack size properly")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := dsl.NewStack(10)
	stack = generateStack(num)
	value, err := stack.Pop()
	if err != nil {
		t.Error("Pop() failed on non-empty stack")
	}
	if *(value.(*vector)) != *newVector(num - 1) {
		t.Errorf("Pop() error: expected: %v, got: %v", newVector(num-1), value)
	}
	for i := 0; i < num-2; i++ {
		_, _ = stack.Pop()
	}
	_, err = stack.Pop()
	if err != nil {
		t.Error("Pop() failed on stack of size 1")
	}
	_, err = stack.Pop()
	if err == nil {
		t.Error("Pop() did not return error on empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := dsl.NewStack(10)
	_, err := stack.Peek()
	if err == nil {
		t.Error("Peek() did not return error on empty stack")
	}
	stack = generateStack(num)
	value, err := stack.Peek()
	if err != nil {
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
func fillStack(stack *dsl.Stack, num int) {
	for i := 0; i < num; i++ {
		stack.Push(newVector(i))
	}
}

// benchmark helper function to remove num items from the stack.
func emptyStack(stack *dsl.Stack, num int) {
	for i := 0; i < num; i++ {
		stack.Pop()
	}
}

func benchmarkStackBasicTest(len int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := dsl.NewStack(10)
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
			if len(stack.Slice) > 0 {
				fillStack(stack, len(stack.Slice)/2)
			} else {
				break
			}
		}
	}
}

func BenchmarkStack_GrowthIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := generateStack(bigNum)
		for len(stack.Slice) <= bigNum {
			emptyStack(stack, num/4)
			fillStack(stack, num/2)
		}
	}

}
