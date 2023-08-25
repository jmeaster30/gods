package gods

import "testing"

func shouldPanic(t *testing.T, f func(), errorMessage string) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Errorf(errorMessage)
}
