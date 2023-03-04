// package errors implement generic version of the std errors package functionality
package errors

import (
	std "errors"
)

func As(err error, target any) bool {
	return std.As(err, target)
}

func Is(err, target error) bool {
	return std.Is(err, target)
}

func Unwrap(err error) error {
	return std.Unwrap(err)
}
