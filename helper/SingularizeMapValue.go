package helper

func SingularizeMapValue[TKey comparable, TValue any](input map[TKey][]TValue) map[TKey]TValue {
	result := make(map[TKey]TValue)
	for key, value := range input {
		if len(value) == 0 {
			continue
		}
		result[key] = value[0]
	}

	return result
}
