package lemmy

type ErrorCode string

const (
	ErrorCodeUnknown        ErrorCode = "unknown"
	ErrorCodeIncorrectLogin ErrorCode = "incorrect_login"
)
