package linear

type Queue[T any] struct {
	store []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (s *Queue[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}
	return &s.store[0]
}

func (s *Queue[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Queue[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	result := s.store[0]
	s.store = s.store[1:len(s.store)]
	return &result
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *Queue[T]) Size() uint64 {
	return uint64(len(s.store))
}
