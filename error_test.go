package errors_test

import (
	"testing"

	"github.com/ahmetcanozcan/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name        string
		err         errors.Error
		wantMessage string
		wantCode    int
		wantStatus  int
	}{
		{
			"error with only message",
			errors.New("test"), "test",
			errors.DefaultValues.Code, errors.DefaultValues.Status,
		},
		{
			"error with message and code",
			errors.New("test", 123), "test",
			123, errors.DefaultValues.Status,
		},
		{
			"error with message, code, and status",
			errors.New("test", 123, 401), "test",
			123, 401,
		},
	}

	for _, tt := range tests {
		qt.Equal(tt.err.Message(), tt.wantMessage)
		qt.Equal(tt.err.Code(), tt.wantCode)
		qt.Equal(tt.err.Status(), tt.wantStatus)
	}

}

func Test_error_WithDesc(t *testing.T) {
	t.Run("error with multiple desc", func(t *testing.T) {
		qt := assert.New(t)

		err := errors.New("test").WithDesc("desc1").WithDesc("desc2")

		d := err.Desc()

		var res []string
		for _, dd := range d {
			res = append(res, dd.Content)
		}
		want := []string{"desc1", "desc2"}
		qt.Equal(want, res)
	})

}

func Test_error_Error(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name string
		err  error
		want string
	}{
		{"Error with only message", errors.New("test"), "test"},
		{"Error with message and code", errors.New("test", 123), "code=123; message=test"},
		{"Error with message, code, and status", errors.New("test", 123, 401), "status=401; code=123; message=test"},
		{"Error with empty message", errors.New(""), ""},
	}
	for _, tt := range tests {
		qt.Equal(tt.want, tt.err.Error(), tt.name)
	}
}

func Test_error_WithCause(t *testing.T) {

	t.Run("with cause error", func(t *testing.T) {
		qt := assert.New(t)
		var (
			cause error = errors.New("this error cause something")
			err   error = errors.New("ops!").WithCause(cause)
		)

		qt.True(errors.Is(err, cause))
		qt.Equal(cause, errors.Unwrap(err))
	})

	t.Run("without error", func(t *testing.T) {
		qt := assert.New(t)
		var err error = errors.New("ops!").WithCause(nil)
		qt.Nil(errors.Unwrap(err))
	})
}

func Test_error_Code(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name string
		err  errors.Error
		want int
	}{
		{"Error without code", errors.New("test"), errors.DefaultValues.Code},
		{"Error with code", errors.New("test", 123), 123},
		{"wrapped error with code", errors.New("test").WithCause(errors.New("test", 123)), 123},
		{"wrapped error without code", errors.New("test").WithCause(errors.New("test")), errors.DefaultValues.Code},
	}
	for _, tt := range tests {
		qt.Equal(tt.want, tt.err.Code(), tt.name)
	}
}

func Test_error_Status(t *testing.T) {
	qt := assert.New(t)

	tests := []struct {
		name string
		err  errors.Error
		want int
	}{
		{"Error without status", errors.New("test"), errors.DefaultValues.Status},
		{"Error with status", errors.New("test", 123, 401), 401},
		{"wrapped error with status", errors.New("test").WithCause(errors.New("test", 123, 401)), 401},
		{"wrapped error without status", errors.New("test").WithCause(errors.New("test")), errors.DefaultValues.Status},
	}
	for _, tt := range tests {
		qt.Equal(tt.want, tt.err.Status(), tt.name)
	}
}
