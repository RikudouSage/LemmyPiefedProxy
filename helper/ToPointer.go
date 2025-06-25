package helper

func ToPointer[TType any](in TType) *TType {
	return &in
}
