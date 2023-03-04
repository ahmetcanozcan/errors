package errors

func Code(err error) int {
	stack := Stack(err)
	for _, err := range stack {
		if c, ok := err.(coder); ok {
			return c.Code()
		}
	}
	return DefaultValues.Code
}

func Status(err error) int {
	stack := Stack(err)
	for _, err := range stack {
		if c, ok := err.(statuser); ok {
			return c.Status()
		}
	}
	return DefaultValues.Status
}
