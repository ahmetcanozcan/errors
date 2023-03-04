package errors_test

import (
	std "errors"
	"fmt"
	"testing"

	"github.com/ahmetcanozcan/errors"
	"github.com/stretchr/testify/assert"
)

var (
	errNotWrapped = errors.New("this error is not wrapped")
	err1          = errors.New("error 1")
	err2          = errors.New("error 2").WithCause(err1)
	err3          = errors.New("error 2").WithCause(err2)
)

func TestStack(t *testing.T) {
	t.Run("error stack from wrapped errors", func(t *testing.T) {
		qt := assert.New(t)
		want := errors.ErrorStack([]error{err3, err2, err1})
		stack := errors.Stack(err3)
		qt.Equal(want, stack)
	})

	t.Run("error stack from nil error", func(t *testing.T) {
		qt := assert.New(t)
		stack := errors.Stack(nil)
		qt.Nil(stack)
	})

	t.Run("error stack from std errors", func(t *testing.T) {
		qt := assert.New(t)

		stderr1 := std.New("test1")
		stderr2 := fmt.Errorf("test2: %w", stderr1)
		stderr3 := fmt.Errorf("test3: %w", stderr2)

		want := errors.ErrorStack([]error{stderr3, stderr2, stderr1})
		stack := errors.Stack(stderr3)
		qt.Equal(stack, want)
	})
}

func TestStack_String(t *testing.T) {
	t.Run("error stack messages from wrapped errors", func(t *testing.T) {
		qt := assert.New(t)
		want := []string{err3.Error(), err2.Error(), err1.Error()}
		stack := errors.Stack(err3)
		qt.Equal(want, stack.Messages())
	})

	t.Run("error stack messages from nil error", func(t *testing.T) {
		qt := assert.New(t)
		stack := errors.Stack(nil)
		qt.Nil(stack)
	})

}

func TestStack_Has(t *testing.T) {
	qt := assert.New(t)
	tests := []struct {
		stack errors.ErrorStack
		arg   error
		want  bool
	}{
		{errors.Stack(err3), err1, true},
		{errors.Stack(err3), nil, false},
		{errors.Stack(nil), err1, false},
	}

	for _, tt := range tests {
		qt.Equal(tt.want, tt.stack.Has(tt.arg))
	}
}
