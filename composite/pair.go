package composite

type Pair[T any, U any] struct {
	first  T
	second U
}

func NewPair[T any, U any](first T, second U) *Pair[T, U] {
	return &Pair[T, U]{
		first:  first,
		second: second,
	}
}

func (value *Pair[T, U]) First() T {
	return value.first
}

func (value *Pair[T, U]) Second() U {
	return value.second
}

func (value *Pair[T, U]) Flip() *Pair[U, T] {
	return &Pair[U, T]{
		first:  value.second,
		second: value.first,
	}
}
