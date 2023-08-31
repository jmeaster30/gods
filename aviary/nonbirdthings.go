package aviary

func CurryOne[T any, U any](f func(T) U, a T) func() U {
	return func() U {
		return f(a)
	}
}

func CurryTwo[T any, U any, V any](f func(T, U) V, a T) func(U) V {
	return func(b U) V {
		return f(a, b)
	}
}

func CurryThree[T any, U any, V any, W any](f func(T, U, V) W, a T) func(U, V) W {
	return func(b U, c V) W {
		return f(a, b, c)
	}
}
