package linear

import "github.com/jmeaster30/gods/composite"

type Queue[T any] struct {
	store []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (s *Queue[T]) Peek() composite.Optional[T] {
	if s.IsEmpty() {
		return composite.None[T]()
	}
	return composite.Some(s.store[0])
}

func (s *Queue[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Queue[T]) Pop() composite.Optional[T] {
	if s.IsEmpty() {
		return composite.None[T]()
	}
	result := s.store[0]
	s.store = s.store[1:len(s.store)]
	return composite.Some(result)
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *Queue[T]) Size() uint64 {
	return uint64(len(s.store))
}
