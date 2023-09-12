package linear

import "github.com/jmeaster30/gods/composite"

type Stack[T any] struct {
	store []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() composite.Optional[T] {
	if s.IsEmpty() {
		return composite.None[T]()
	}
	return composite.Some(s.store[len(s.store)-1])
}

func (s *Stack[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Stack[T]) Pop() composite.Optional[T] {
	if s.IsEmpty() {
		return composite.None[T]()
	}
	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return composite.Some(result)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *Stack[T]) Size() uint64 {
	return uint64(len(s.store))
}
