package composite

type Equatable interface {
	IsEqual(Equatable) bool
	IsNotEqual(Equatable) bool
}
