package errors

import (
	"strings"
)

type ErrorStack []error

func Stack(err error) ErrorStack {
	var stack ErrorStack

	for {
		if err == nil {
			break
		}
		stack = append(stack, err)
		err = Unwrap(err)
	}
	return stack
}

func (stack ErrorStack) String() string {
	return strings.Join(stack.Messages(), "\n")
}

func (stack ErrorStack) Messages() []string {
	var s []string
	for _, err := range stack {
		s = append(s, err.Error())
	}

	return s
}

func (stack ErrorStack) Has(target error) bool {
	for _, err := range stack {
		if Is(err, target) {
			return true
		}
	}
	return false

}
