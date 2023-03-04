package errors

type Description struct {
	Content string
	Level   int
}

func newDesc(content string, level ...int) Description {
	return Description{content, optionalParam(level, DefaultValues.DescRef)}
}
