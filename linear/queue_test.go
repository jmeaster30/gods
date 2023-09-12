package linear

import (
	"testing"

	"github.com/jmeaster30/gods/composite"
)

func TestQueuePush(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.Size() != 3 {
		t.Errorf("The queue was expected to be 3 but actually was %d :(", queue.Size())
	}

	peeked := queue.Peek()
	if peeked != composite.Some(1) {
		t.Errorf("The front of the queue was expected to be 3 but actually was %v :(", peeked)
	}
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	val := queue.Pop()
	if val != composite.Some(1) {
		t.Errorf("The queue was expected to pop 1 but actually was %v :(", val)
	}

	val = queue.Pop()
	if val != composite.Some(2) {
		t.Errorf("The queue was expected to pop 2 but actually was %v :(", val)
	}

	val = queue.Pop()
	if val != composite.Some(3) {
		t.Errorf("The queue was expected to pop 3 but actually was %v :(", val)
	}

	last := queue.Pop()
	if last != composite.None[int]() {
		t.Errorf("The queue was expected to pop nil but actually was %v :(", last)
	}
}

func TestQueueIsEmpty(t *testing.T) {
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
