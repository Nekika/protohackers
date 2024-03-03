package lib

func Map[S any, D any](src []S, fn func(value S, index int) D) []D {
	dst := make([]D, len(src))

	for index, value := range src {
		dst[index] = fn(value, index)
	}

	return dst
}
