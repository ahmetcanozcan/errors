package errors

type defaultValues struct {
	Code    int
	Status  int
	DescRef int
}

var DefaultValues = defaultValues{
	Code:    0,
	Status:  500,
	DescRef: 0,
}

func fallbackDefault[T any](t *T, d T) T {
	if t == nil {
		return d
	}
	return *t
}

func optionalParam[T any](t []T, d T) T {
	if len(t) == 1 {
		return t[0]
	}
	return d
}
