package linear

type LinkedList[T any] struct {
	data T
	next *LinkedList[T]
}

func (list *LinkedList[T]) Next() *LinkedList[T] {
	return list.next
}

func (list *LinkedList[T]) Size() uint {
	var result uint = 0
	current := list
	for current != nil {
		result += 1
		current = current.next
	}

	return result
}

func (list *LinkedList[T]) ForEach(process func(T)) {
	current := list
	for current != nil {
		process(current.data)
		current = current.next
	}
}

func (list *LinkedList[T]) Prepend(data T) *LinkedList[T] {
	return &LinkedList[T]{
		data: data,
		next: list,
	}
}

func (list *LinkedList[T]) Append(data T) *LinkedList[T] {
	final_node := list
	for final_node.next != nil {
		final_node = final_node.next
	}

	final_node.next = &LinkedList[T]{
		data: data,
		next: nil,
	}

	return final_node.next
}
