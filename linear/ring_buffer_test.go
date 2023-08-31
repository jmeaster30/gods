package linear

import (
	"testing"

	"github.com/jmeaster30/gods/composite"
)

func TestNewRingBuffer(t *testing.T) {
	ring_buffer := NewRingBuffer[string](10)

	if ring_buffer.capacity != 10 {
		t.Errorf("Expected capacity of ring buffer to be 10 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected start of ring buffer to be 1 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.None[uint64]() {
		t.Errorf("Expected end to be a None but got %v", ring_buffer.end)
	}
}

func TestRingBufferPushFront(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(7)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected start of ring buffer to be 0 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](0) {
		t.Errorf("Expected end to be a None but got %v", ring_buffer.end)
	}
}

func TestRingBufferPushFrontTwice(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(7)
	ring_buffer.PushFront(8)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 3 {
		t.Errorf("Expected start of ring buffer to be 3 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](0) {
		t.Errorf("Expected end to be some 0 but got %v", ring_buffer.end)
	}
}

func TestRingBufferPushFrontWrap(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(1)
	ring_buffer.PushFront(2)
	ring_buffer.PushFront(3)
	ring_buffer.PushFront(4)
	ring_buffer.PushFront(5)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected start of ring buffer to be 0 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](3) {
		t.Errorf("Expected end to be some 3 but got %v", ring_buffer.end)
	}

	expected := []int{5, 4, 3, 2}
	for i := 0; i < 4; i++ {
		if ring_buffer.data[i] != expected[i] {
			t.Errorf("Expected data at index %d to be %d but got %d", i, expected[i], ring_buffer.data[i])
		}
	}
}

func TestRingBufferSizeWithSomething(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(1)
	ring_buffer.PushFront(2)
	ring_buffer.PushFront(3)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.Size() != 3 {
		t.Errorf("Expected size of ring buffer to be 3 but got %d", ring_buffer.Size())
	}
}

func TestRingBufferSizeWithNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.Size() != 0 {
		t.Errorf("Expected size of ring buffer to be 0 but got %d", ring_buffer.Size())
	}
}

func TestRingBufferSizeWithSomething2(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushBack(3)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.Size() != 1 {
		t.Errorf("Expected size of ring buffer to be 1 but got %d", ring_buffer.Size())
	}
}

func TestRingBufferIsEmpty(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if !ring_buffer.IsEmpty() {
		t.Error("Expected size of ring buffer to be empty but it wasn't")
	}
}

func TestRingBufferIsEmptyWithSomething(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(1)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.IsEmpty() {
		t.Error("Expected size of ring buffer to be non-empty but it wasn't")
	}
}

func TestRingBufferIsFullWithNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.IsFull() {
		t.Error("Expected size of ring buffer to not be full but it was")
	}
}

func TestRingBufferIsFullWithSomething(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(1)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.IsFull() {
		t.Error("Expected size of ring buffer to not be full but it was")
	}
}

func TestRingBufferIsFull(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushFront(1)
	ring_buffer.PushFront(1)
	ring_buffer.PushFront(1)
	ring_buffer.PushFront(1)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if !ring_buffer.IsFull() {
		t.Error("Expected size of ring buffer to be full but it wasn't")
	}
}

func TestRingBufferPopFrontPopNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	value := ring_buffer.PopFront()
	if value != composite.None[string]() {
		t.Errorf("Expected popped value to be nothing but got %v", value)
	}

	if ring_buffer.capacity != 5 {
		t.Errorf("Expected capacity of ring buffer to be 5 but got %d", ring_buffer.capacity)
	}
}

func TestRingBufferPopFrontPopSomethingToNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	ring_buffer.PushFront("Hey")

	value := ring_buffer.PopFront()
	if value != composite.Some[string]("Hey") {
		t.Errorf("Expected popped value to be some 'Hey' but got %v", value)
	}

	if !ring_buffer.IsEmpty() {
		t.Errorf("Expected ring_buffer to be empty but it wasn't")
	}
}

func TestRingBufferPopFrontPopSomethingToNotEmpty(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	ring_buffer.PushFront("Hey")
	ring_buffer.PushFront("World")

	value := ring_buffer.PopFront()
	if value != composite.Some[string]("World") {
		t.Errorf("Expected popped value to be some 'World' but got %v", value)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected ring_buffer.start to equal 0 but it was %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](0) {
		t.Errorf("Expected ting_buffer.end to equal 0 but it was %v", ring_buffer.end)
	}

	if ring_buffer.Size() != 1 {
		t.Errorf("Expected ring_buffer to have size of 1 but it was %d", ring_buffer.Size())
	}
}

func TestRingBufferPeekFrontSomething(t *testing.T) {
	ring_buffer := NewRingBuffer[int](5)

	ring_buffer.PushFront(123)
	ring_buffer.PushFront(456)

	value := ring_buffer.PeekFront()
	if value != composite.Some[int](456) {
		t.Errorf("Expected peeked value to be some 456 but it was %v", value)
	}

	if ring_buffer.Size() != 2 {
		t.Errorf("Expected ring buffer size to be 2 but it was %d", ring_buffer.Size())
	}
}

func TestRingBufferPeekFrontNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[int](5)

	value := ring_buffer.PeekFront()
	if value != composite.None[int]() {
		t.Errorf("Expected peeked value to be None but it was %v", value)
	}
}

func TestRingBufferPushBack(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushBack(7)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected start of ring buffer to be 0 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](0) {
		t.Errorf("Expected end to be a None but got %v", ring_buffer.end)
	}
}

func TestRingBufferPushBackTwice(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushBack(7)
	ring_buffer.PushBack(8)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 0 {
		t.Errorf("Expected start of ring buffer to be 3 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](1) {
		t.Errorf("Expected end to be some 0 but got %v", ring_buffer.end)
	}
}

func TestRingBufferPushBackWrap(t *testing.T) {
	ring_buffer := NewRingBuffer[int](4)

	ring_buffer.PushBack(1)
	ring_buffer.PushBack(2)
	ring_buffer.PushBack(3)
	ring_buffer.PushBack(4)
	ring_buffer.PushBack(5)

	if ring_buffer.capacity != 4 {
		t.Errorf("Expected capacity of ring buffer to be 4 but got %d", ring_buffer.capacity)
	}

	if ring_buffer.start != 1 {
		t.Errorf("Expected start of ring buffer to be 0 but got %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](0) {
		t.Errorf("Expected end to be some 3 but got %v", ring_buffer.end)
	}

	expected := []int{5, 2, 3, 4}
	for i := 0; i < 4; i++ {
		if ring_buffer.data[i] != expected[i] {
			t.Errorf("Expected data at index %d to be %d but got %d", i, expected[i], ring_buffer.data[i])
		}
	}
}

func TestRingBufferPopBackPopNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	value := ring_buffer.PopBack()
	if value != composite.None[string]() {
		t.Errorf("Expected popped value to be nothing but got %v", value)
	}

	if ring_buffer.capacity != 5 {
		t.Errorf("Expected capacity of ring buffer to be 5 but got %d", ring_buffer.capacity)
	}
}

func TestRingBufferPopBackPopSomethingToNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	ring_buffer.PushBack("Hey")

	value := ring_buffer.PopBack()
	if value != composite.Some[string]("Hey") {
		t.Errorf("Expected popped value to be some 'Hey' but got %v", value)
	}

	if !ring_buffer.IsEmpty() {
		t.Errorf("Expected ring_buffer to be empty but it wasn't")
	}
}

func TestRingBufferPopBackPopSomethingToNotEmpty(t *testing.T) {
	ring_buffer := NewRingBuffer[string](5)

	ring_buffer.PushFront("Hey")
	ring_buffer.PushFront("World")

	value := ring_buffer.PopBack()
	if value != composite.Some[string]("Hey") {
		t.Errorf("Expected popped value to be some 'World' but got %v", value)
	}

	if ring_buffer.start != 4 {
		t.Errorf("Expected ring_buffer.start to equal 4 but it was %d", ring_buffer.start)
	}

	if ring_buffer.end != composite.Some[uint64](4) {
		t.Errorf("Expected ting_buffer.end to equal 4 but it was %v", ring_buffer.end)
	}

	if ring_buffer.Size() != 1 {
		t.Errorf("Expected ring_buffer to have size of 1 but it was %d", ring_buffer.Size())
	}
}

func TestRingBufferPeekBackSomething(t *testing.T) {
	ring_buffer := NewRingBuffer[int](5)

	ring_buffer.PushBack(123)
	ring_buffer.PushBack(456)

	value := ring_buffer.PeekBack()
	if value != composite.Some[int](456) {
		t.Errorf("Expected peeked value to be some 456 but it was %v", value)
	}

	if ring_buffer.Size() != 2 {
		t.Errorf("Expected ring buffer size to be 2 but it was %d", ring_buffer.Size())
	}
}

func TestRingBufferPeekBackNothing(t *testing.T) {
	ring_buffer := NewRingBuffer[int](5)

	value := ring_buffer.PeekBack()
	if value != composite.None[int]() {
		t.Errorf("Expected peeked value to be None but it was %v", value)
	}
}
