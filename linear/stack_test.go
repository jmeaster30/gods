package linear

import (
	"testing"

	"github.com/jmeaster30/gods/composite"
)

func TestPush(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Errorf("The stack was expected to be 3 but actually was %d :(", stack.Size())
	}

	peeked := stack.Peek()
	if peeked != composite.Some(3) {
		t.Errorf("The top of the stack was expected to be 3 but actually was %v :(", peeked)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val := stack.Pop()
	if val != composite.Some(3) {
		t.Errorf("The stack was expected to pop 3 but actually was %v :(", val)
	}

	val = stack.Pop()
	if val != composite.Some(2) {
		t.Errorf("The stack was expected to pop 2 but actually was %v :(", val)
	}

	val = stack.Pop()
	if val != composite.Some(1) {
		t.Errorf("The stack was expected to pop 1 but actually was %v :(", val)
	}

	last := stack.Pop()
	if last != composite.None[int]() {
		t.Errorf("The stack was expected to pop none but actually was %v :(", last)
	}
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack[int]()
	if !stack.IsEmpty() {
		t.Errorf("The stack was expected to be empty after creating new stack :(")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("The stack was supposed to have an element in it after adding an element :(")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("The stack was expected to be empty after removing an element :(")
	}
}
