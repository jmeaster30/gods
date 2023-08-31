package aviary

// https://www.angelfire.com/tx4/cus/combinator/birds.html

func Kestrel[T any, U any](a T, b U) T {
	return a
}

func Idiot[T any](a T) T {
	return a
}

func Identity[T any](a T) T {
	return Starling(Kestrel[T, T], CurryTwo(Kestrel[T, T], a), a)
}

func Starling[A any, B any, C any](f func(A, B) C, g func(A) B, h A) C {
	return f(h, g(h))
}
