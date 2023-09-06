package linear

import "github.com/jmeaster30/gods/composite"

type RingBuffer[T any] struct {
	capacity uint64
	start    uint64
	end      composite.Optional[uint64]
	data     []T
}

func NewRingBuffer[T any](capacity uint64) *RingBuffer[T] {
	return &RingBuffer[T]{
		capacity: capacity,
		end:      composite.None[uint64](),
		data:     make([]T, capacity),
	}
}

func (r *RingBuffer[T]) PushFront(value T) {
	if !r.end.HasValue() {
		r.end = composite.Some[uint64](r.start)
	} else {
		r.start = ((r.start - 1) + r.capacity) % r.capacity
		if r.start == r.end.Value() {
			r.end = composite.Some[uint64](((r.end.Value() - 1) + r.capacity) % r.capacity)
		}
	}

	r.data[r.start] = value
}

func (r *RingBuffer[T]) PopFront() composite.Optional[T] {
	if !r.end.HasValue() {
		return composite.None[T]()
	}

	value := r.data[r.start]

	if r.end.Value() == r.start {
		r.end = composite.None[uint64]()
	}
	r.start = (r.start + 1) % r.capacity

	return composite.Some[T](value)
}

func (r *RingBuffer[T]) PeekFront() composite.Optional[T] {
	if !r.end.HasValue() {
		return composite.None[T]()
	}
	return composite.Some[T](r.data[r.start])
}

func (r *RingBuffer[T]) PushBack(value T) {
	if !r.end.HasValue() {
		r.end = composite.Some[uint64](r.start)
	} else {
		r.end = composite.Some[uint64]((r.end.Value() + 1) % r.capacity)
		if r.start == r.end.Value() {
			r.start = (r.start + 1) % r.capacity
		}
	}

	// r.end will always have a value here!
	r.data[r.end.Value()] = value
}

func (r *RingBuffer[T]) PopBack() composite.Optional[T] {
	if !r.end.HasValue() {
		return composite.None[T]()
	}

	value := r.data[r.end.Value()]

	if r.end.Value() == r.start {
		r.end = composite.None[uint64]()
	} else {
		r.end = composite.Some[uint64](((r.end.Value() - 1) + r.capacity) % r.capacity)
	}

	return composite.Some[T](value)
}

func (r *RingBuffer[T]) PeekBack() composite.Optional[T] {
	if !r.end.HasValue() {
		return composite.None[T]()
	}
	return composite.Some[T](r.data[r.end.Value()])
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return !r.end.HasValue()
}

func (r *RingBuffer[T]) IsFull() bool {
	return r.Size() == r.capacity
}

func (r *RingBuffer[T]) Size() uint64 {
	if !r.end.HasValue() {
		return 0
	}

	end := r.end.Value()
	if r.start < end {
		return end - r.start + 1
	} else if r.start == end {
		return 1
	} else {
		return (r.capacity - r.start) + end + 1
	}
}
