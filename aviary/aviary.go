package aviary

// https://www.angelfire.com/tx4/cus/combinator/birds.html

func Bluebird[A any, B any, C any](a func(B) C, b func(A) B, c A) C {
	return a(b(c))
}

func Idiot[A any](a A) A {
	return a
}

func Identity[A any](a A) A {
	return Starling(Kestrel[A, A], CurryTwo(Kestrel[A, A], a), a)
}

func Kestrel[A any, B any](a A, b B) A {
	return a
}

func Starling[A any, B any, C any](a func(A, B) C, b func(A) B, c A) C {
	return a(c, b(c))
}
