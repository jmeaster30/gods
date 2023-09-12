package trees

import "github.com/jmeaster30/gods/linear"

type TraversalOrder int

const (
	Prefix TraversalOrder = iota
	Postfix
	Infix
)

type BinaryTree[T any] struct {
	data  T
	left  *BinaryTree[T]
	right *BinaryTree[T]
}

func NewBinaryTree[T any](value T) *BinaryTree[T] {
	return &BinaryTree[T]{
		data:  value,
		left:  nil,
		right: nil,
	}
}

func (bt *BinaryTree[T]) IsLeaf() bool {
	return bt.left == nil && bt.right == nil
}

func (bt *BinaryTree[T]) Data() T {
	return bt.data
}

func (bt *BinaryTree[T]) Left() *BinaryTree[T] {
	return bt.left
}

func (bt *BinaryTree[T]) AddLeft(left *BinaryTree[T]) {
	bt.left = left
}

func (bt *BinaryTree[T]) Right() *BinaryTree[T] {
	return bt.right
}

func (bt *BinaryTree[T]) AddRight(right *BinaryTree[T]) {
	bt.right = right
}

func (bt *BinaryTree[T]) BreadthFirst() []T {
	searchQueue := linear.NewQueue[*BinaryTree[T]]()

	searchQueue.Push(bt)
	results := []T{}

	for !searchQueue.IsEmpty() {
		current := searchQueue.Pop()
		if !current.HasValue() {
			break
		}
		results = append(results, current.Value().Data())
		if current.Value().left != nil {
			searchQueue.Push(current.Value().left)
		}

		if current.Value().right != nil {
			searchQueue.Push(current.Value().right)
		}
	}

	return results
}

func (bt *BinaryTree[T]) DepthFirst(order TraversalOrder) []T {
	left := []T{}
	if bt.left != nil {
		left = bt.left.DepthFirst(order)
	}

	right := []T{}
	if bt.right != nil {
		right = bt.right.DepthFirst(order)
	}

	results := []T{}
	if order == Prefix {
		results = append(results, bt.data)
		results = append(results, left...)
		results = append(results, right...)
	} else if order == Postfix {
		results = append(results, left...)
		results = append(results, right...)
		results = append(results, bt.data)
	} else if order == Infix {
		results = append(results, left...)
		results = append(results, bt.data)
		results = append(results, right...)
	}
	return results
}
