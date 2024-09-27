package composite

import (
	"fmt"
)

type Result[T any] struct {
	data *T
	err  error
}

func (result Result[T]) IsOk() bool {
	return result.data != nil
}

func (result Result[T]) IsError() bool {
	return result.err != nil
}

func (result Result[T]) ErrorIs(err error) bool {
	return result.err == err
}

func (result Result[T]) Value() T {
	if result.data == nil {
		panic(result.err)
	}
	return *result.data
}

func (result Result[T]) Error() string {
	if result.err == nil {
		panic("Ok result doesn't have an error")
	}
	return result.err.Error()
}

func (result Result[T]) ValueOrDefault(defaultValue T) T {
	if result.data != nil {
		return *result.data
	}
	return defaultValue
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		data: &value,
		err:  nil,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		data: nil,
		err:  err,
	}
}

func Errorf[T any](message string, values ...any) Result[T] {
	return Error[T](fmt.Errorf(message, values...))
}

func PanicGuard[T any](result *Result[T]) {
	if err := recover(); err != nil {
		if e, ok := err.(error); ok {
			*result = Error[T](e)
		} else {
			*result = Errorf[T]("%v", err)
		}
	}
}
