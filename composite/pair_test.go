package composite

import (
	"testing"
)

func TestPairCreateFirstSecond(t *testing.T) {
	pair_value := NewPair[int, string](5, "test")

	if pair_value.First() != 5 {
		t.Errorf("Expected first of the pair to be 5 but got %d", pair_value.First())
	}

	if pair_value.Second() != "test" {
		t.Errorf("Expected second of the pair to be 'test' but got '%s'", pair_value.Second())
	}
}

func TestPairFlip(t *testing.T) {
	pair_value := NewPair[int, string](5, "test")

	flipped := pair_value.Flip()

	if flipped.First() != "test" {
		t.Errorf("Expected first of the pair to be 'test' but got '%s'", flipped.First())
	}

	if flipped.Second() != 5 {
		t.Errorf("Expected second of the pair to be 5 but got %d", flipped.Second())
	}
}
