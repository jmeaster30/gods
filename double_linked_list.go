package gods

type DoubleLinkedList[T any] struct {
	Data     T
	previous *DoubleLinkedList[T]
	next     *DoubleLinkedList[T]
}

func NewDoubleLinkedList[T any](data T) *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		Data:     data,
		previous: nil,
		next:     nil,
	}
}

func (list *DoubleLinkedList[T]) First() *DoubleLinkedList[T] {
	current := list
	for current.previous != nil {
		current = current.previous
	}
	return current
}

func (list *DoubleLinkedList[T]) Last() *DoubleLinkedList[T] {
	current := list
	for current.next != nil {
		current = current.next
	}
	return current
}

func (list *DoubleLinkedList[T]) Get(index int) *DoubleLinkedList[T] {
	current := list.First()
	currentIndex := 0
	for currentIndex < index {
		current = list.next
	}
	return current
}

func (list *DoubleLinkedList[T]) Next() *DoubleLinkedList[T] {
	return list.next
}

func (list *DoubleLinkedList[T]) Previous() *DoubleLinkedList[T] {
	return list.previous
}

func (list *DoubleLinkedList[T]) IsFirst() bool {
	return list.previous == nil
}

func (list *DoubleLinkedList[T]) IsLast() bool {
	return list.next == nil
}

func (list *DoubleLinkedList[T]) Size() int {
	result := 0
	current := list.First()
	for current != nil {
		result += 1
		current = current.next
	}

	return result
}

func (list *DoubleLinkedList[T]) ForEach(process func(T)) {
	current := list.First()
	for current != nil {
		process(current.Data)
		current = current.next
	}
}

func (list *DoubleLinkedList[T]) Prepend(data T) *DoubleLinkedList[T] {
	first := list.First()
	newNode := &DoubleLinkedList[T]{
		Data:     data,
		previous: nil,
		next:     first,
	}
	first.previous = newNode
	return newNode
}

func (list *DoubleLinkedList[T]) Append(data T) *DoubleLinkedList[T] {
	final_node := list
	for final_node.next != nil {
		final_node = final_node.next
	}

	final_node.next = &DoubleLinkedList[T]{
		Data:     data,
		previous: final_node,
		next:     nil,
	}

	return final_node.next
}
