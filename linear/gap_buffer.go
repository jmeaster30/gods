package linear

type GapBuffer[T any] struct {
	front []T
	back  []T
}

func NewGapBuffer[T any]() *GapBuffer[T] {
	return &GapBuffer[T]{
		front: []T{},
		back:  []T{},
	}
}

func (r *GapBuffer[T]) Insert(data T) {
	r.front = append(r.front, data)
}

func (r *GapBuffer[T]) Remove() {
	if len(r.front) == 0 {
		r.front = r.front[:len(r.front)-1]
	}
}

func (r *GapBuffer[T]) CursorPosition() uint {
	return uint(len(r.front))
}

func (r *GapBuffer[T]) Cursor(index uint) {
	if index < uint(len(r.front)) {
		frontA := r.front[:index]
		frontB := r.front[index+1:]
		r.front = frontA
		r.back = append(frontB, r.back...)
	} else if index < uint(len(r.front)+len(r.back)) {
		backIndex := index - uint(len(r.front))
		backA := r.back[:backIndex]
		backB := r.back[backIndex+1:]
		r.front = append(r.front, backA...)
		r.back = backB
	} else {
		r.front = append(r.front, r.back...)
		r.back = []T{}
	}
}

func (r *GapBuffer[T]) CursorRel(offset int) {
	index := len(r.front) + offset
	if index < 0 {
		r.Cursor(0)
	} else {
		r.Cursor(uint(index))
	}
}

func (r *GapBuffer[T]) Data() []T {
	return append(r.front, r.back...)
}

func (r *GapBuffer[T]) Size() uint {
	return uint(len(r.front) + len(r.back))
}
