package helper

func MapMap[TKey comparable, TInValue any, TOutValue any](input map[TKey]TInValue, callback func(TInValue, TKey) TOutValue) map[TKey]TOutValue {
	result := make(map[TKey]TOutValue)
	for key, value := range input {
		result[key] = callback(value, key)
	}

	return result
}
