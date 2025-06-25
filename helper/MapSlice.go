package helper

func MapSlice[TIn any, TOut any](in []TIn, callback func(TIn) TOut) []TOut {
	result := make([]TOut, 0, len(in))
	for _, item := range in {
		result = append(result, callback(item))
	}

	return result
}
