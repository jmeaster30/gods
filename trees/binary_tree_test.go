package trees

import (
	"testing"
)

func TestBinaryTreeCreateLeaf(t *testing.T) {
	tree := NewBinaryTree[int](12)

	if !tree.IsLeaf() {
		t.Error("Expected the created tree node to be a leaf but it wasn't")
	}
}
