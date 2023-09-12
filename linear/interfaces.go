package linear

type List[T any] interface {
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() uint
	Data() []T
}

type IndexableList[T any] interface {
	Add(T)
	Remove(T)
	Contains(T)
	Size() uint
	Data() []T

	Push(T)
	Pop() T

	Get(uint) T
	Insert(T, uint)
	RemoveAt(uint) T
}
