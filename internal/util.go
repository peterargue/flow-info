package internal

func Extract[T any](container map[string]interface{}, name string) T {
	if value, found := container[name]; found {
		return value.(T)
	}

	var result T
	return result
}
