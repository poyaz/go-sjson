package pkg

import (
	"fmt"
	"reflect"
)

type NonPointerError struct{}

var _ error = &NonPointerError{}

func NewNonPointerError() *NonPointerError {
	return &NonPointerError{}
}

func (e *NonPointerError) Error() string {
	return fmt.Sprintf("attempt to decode into a non-pointer")
}

type NonTypeError struct {
	t reflect.Type
}

var _ error = &NonTypeError{}

func NewNonTypeError(dataType reflect.Type) *NonTypeError {
	return &NonTypeError{t: dataType}
}

func (e *NonTypeError) Error() string {
	return fmt.Sprintf("unsupported type recived: %s", e.t.String())
}
