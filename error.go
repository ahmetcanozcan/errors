package errors

import "fmt"

type unwrapper interface {
	Unwrap() error
}

type coder interface {
	Code() int
}

type statuser interface {
	Status() int
}

type Error interface {
	error
	unwrapper

	coder
	statuser

	Message() string
	WithCause(cause error) Error
	WithDesc(desc string, level ...int) Error
	Desc() []Description
}

type _error struct {
	cause   error
	code    *int
	status  *int
	message string
	desc    []Description
}

func newError(message string, code, status *int, cause error) Error {

	err := &_error{
		message: message,
		code:    code,
		status:  status,
		cause:   cause,
	}

	return err
}

func New(message string, codestatus ...int) Error {
	var (
		code   *int
		status *int
	)

	if len(codestatus) > 0 {
		code = &codestatus[0]
	}

	if len(codestatus) > 1 {
		status = &codestatus[1]
	}

	return newError(message, code, status, nil)
}

func Parse(err error) Error {
	r, ok := err.(Error)
	if ok {
		return r
	}
	return New(err.Error())
}

func (err *_error) Message() string {
	return err.message
}

func (err *_error) Error() string {
	if err.isTextError() {
		return err.message
	}

	var s, c string

	if err.status != nil {
		s = fmt.Sprintf("status=%d; ", *err.status)
	}

	if err.code != nil {
		c = fmt.Sprintf("code=%d; ", *err.code)
	}

	return fmt.Sprintf("%s%smessage=%s", s, c, err.message)
}

func (err *_error) Unwrap() error {
	return err.cause
}

func (err *_error) WithCause(cause error) Error {
	return err.copy().wrap(cause)
}

func (err *_error) Code() int {
	if err.code != nil {
		return *err.code
	}

	if err.cause != nil {
		return Code(err.cause)
	}

	return DefaultValues.Code
}

func (err *_error) Status() int {
	if err.status != nil {
		return *err.status
	}

	if err.cause != nil {
		return Status(err.cause)
	}

	return DefaultValues.Status
}

func (err *_error) WithDesc(desc string, level ...int) Error {
	err.desc = append(err.desc, newDesc(desc, level...))
	return err
}

func (err *_error) Desc() []Description {
	return err.desc
}

func (err *_error) isTextError() bool {
	return err.code == nil && err.status == nil
}

func (err *_error) wrap(cause error) Error {
	err.cause = cause
	return err
}

func (err *_error) copy() *_error {
	return &_error{
		cause:   err.cause,
		code:    err.code,
		status:  err.status,
		message: err.message,
		desc:    err.desc,
	}
}
