package errors_test

import (
	std "errors"
	"testing"

	"github.com/ahmetcanozcan/errors"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name string
		err  error
		want int
	}{
		{"error with message", errors.New("test"), errors.DefaultValues.Code},
		{"error with code", errors.New("test", 123), 123},
		{"wrapped error with code", errors.New("test").WithCause(errors.New("test", 123)), 123},
		{"std error", std.New("std error"), errors.DefaultValues.Code},
	}
	for _, tt := range tests {
		qt.Equal(tt.want, errors.Code(tt.err), tt.name)
	}
}

func TestStatus(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name string
		err  error
		want int
	}{
		{"error with message", errors.New("test"), errors.DefaultValues.Status},
		{"error with status", errors.New("test", 123, 401), 401},
		{"wrapped error with status", errors.New("test").WithCause(errors.New("test", 123, 401)), 401},
		{"std error", std.New("std error"), errors.DefaultValues.Status},
	}
	for _, tt := range tests {
		qt.Equal(tt.want, errors.Status(tt.err), tt.name)
	}
}
