package composite

import (
	"testing"

	helpers "github.com/jmeaster30/gods/test_helpers"
)

func TestOptionalSome(t *testing.T) {
	optional_number := Some[int](10)

	if !optional_number.HasValue() {
		t.Error("Expected the optional number to have a value but it didn't :(")
	}

	value := optional_number.Value()
	if optional_number.Value() != 10 {
		t.Errorf("Expected the optional number to have a value of 10 but it had %d", value)
	}
}

func TestOptionalNone(t *testing.T) {
	optional_number := None[int]()

	if optional_number.HasValue() {
		t.Error("Expected the optional number to not have a value but it did :(")
	}

	helpers.ShouldPanic(t, func() {
		_ = optional_number.Value()
	}, "Expected getting the value from a none to panic!")
}
