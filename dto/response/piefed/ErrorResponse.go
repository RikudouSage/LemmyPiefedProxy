package piefed

type ErrorResponse struct {
	ErrorCode  string `json:"error"`
	StatusCode int    `json:"-"`
}

func (receiver *ErrorResponse) Error() string {
	return receiver.ErrorCode
}
