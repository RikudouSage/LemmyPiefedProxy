package helper

func SafeDereference[TIn any, TOut any](in *TIn, callable func(TIn) *TOut) *TOut {
	if in == nil {
		return nil
	}

	return callable(*in)
}

func DefaultOnNil[T any](in *T, defaultValue T) T {
	if in == nil {
		return defaultValue
	}

	return *in
}
