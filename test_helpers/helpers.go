package helpers

import "testing"

func ShouldPanic(t *testing.T, f func(), errorMessage string) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Error(errorMessage)
}

func AssertEqual[T comparable](t *testing.T, expectedValue T, actualValue T) {
	if actualValue != expectedValue {
		t.Errorf("Expected (%v) but got (%v)", expectedValue, actualValue)
	}
}
