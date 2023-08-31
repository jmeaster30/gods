package helpers

import "testing"

func ShouldPanic(t *testing.T, f func(), errorMessage string) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Errorf(errorMessage)
}
