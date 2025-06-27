package helper

func MapFilter[TKey comparable, TValue any](in map[TKey]TValue, filter func(value TValue, key TKey) bool) map[TKey]TValue {
	result := make(map[TKey]TValue)
	for key, value := range in {
		if filter(value, key) {
			result[key] = value
		}
	}

	return result
}
