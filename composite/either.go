package composite

type Either[T any, U any] struct {
	leftExists  bool
	left        T
	rightExists bool
	right       U
}

func Left[T any, U any](value T) Either[T, U] {
	return Either[T, U]{
		leftExists:  true,
		left:        value,
		rightExists: false,
	}
}

func Right[T any, U any](value U) Either[T, U] {
	return Either[T, U]{
		leftExists:  false,
		right:       value,
		rightExists: true,
	}
}

func (e Either[T, U]) IsLeft() bool {
	return e.leftExists
}

func (e Either[T, U]) IsRight() bool {
	return e.rightExists
}

func (e Either[T, U]) Left() T {
	if !e.leftExists {
		panic("Expected the Either to have a right result")
	}

	return e.left
}

func (e Either[T, U]) Right() U {
	if !e.rightExists {
		panic("Expected the Either to have a right value")
	}

	return e.right
}

func (e Either[T, U]) LeftOr(defaultValue T) T {
	if !e.leftExists {
		return defaultValue
	}

	return e.left
}

func (e Either[T, U]) RightOr(defaultValue U) U {
	if !e.rightExists {
		return defaultValue
	}

	return e.right
}
