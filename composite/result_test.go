package composite

import (
	"errors"
	"testing"

	helpers "github.com/jmeaster30/gods/test_helpers"
)

func TestResultOkIsOkAndNotIsError(t *testing.T) {
	result := Ok(123)

	helpers.AssertEqual(t, true, result.IsOk())
	helpers.AssertEqual(t, false, result.IsError())
	helpers.AssertEqual(t, 123, result.Value())
}

func TestResultErrorIsNotOkAndIsError(t *testing.T) {
	result := Error[int](errors.New("Test"))

	helpers.AssertEqual(t, false, result.IsOk())
	helpers.AssertEqual(t, true, result.IsError())
	helpers.AssertEqual(t, "Test", result.Error())
}

func TestResultOkErrorPanics(t *testing.T) {
	result := Ok(123)

	helpers.ShouldPanic(t, func() {
		_ = result.Error()
	}, "Expected error from ok result to panic")

	helpers.AssertEqual(t, 123, result.Value())
}

func TestResultErrorValuePanics(t *testing.T) {
	result := Error[int](errors.New("Test"))

	helpers.ShouldPanic(t, func() {
		_ = result.Value()
	}, "Expected value from error result to panic")

	helpers.AssertEqual(t, "Test", result.Error())
}

func TestResultPanicGuard(t *testing.T) {
	result := (func() (result Result[int]) {
		defer PanicGuard(&result)

		panic("Oh no!")
	})()

	helpers.AssertEqual(t, false, result.IsOk())
	helpers.AssertEqual(t, true, result.IsError())
	helpers.AssertEqual(t, "Oh no!", result.Error())
}

func TestResultPanicGuardError(t *testing.T) {
	result := (func() (result Result[int]) {
		defer PanicGuard(&result)

		panic(errors.New("Oh no!"))
	})()

	helpers.AssertEqual(t, false, result.IsOk())
	helpers.AssertEqual(t, true, result.IsError())
	helpers.AssertEqual(t, "Oh no!", result.Error())
}

func TestResultOkValueOrDefault(t *testing.T) {
	result := Ok(123)

	helpers.AssertEqual(t, 123, result.ValueOrDefault(1))
}

func TestResultErrorValueOrDefault(t *testing.T) {
	result := Errorf[int]("Oh no!")

	helpers.AssertEqual(t, 12, result.ValueOrDefault(12))
}
