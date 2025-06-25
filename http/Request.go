package http

type Headers map[string]string

type Request struct {
	Body        []byte
	Method      string
	RouteParams map[string]string
	QueryParams map[string]string
	Headers     Headers
}

func NewRequest(body []byte, method string, headers Headers) *Request {
	return &Request{
		Body:        body,
		Method:      method,
		RouteParams: make(map[string]string),
		QueryParams: make(map[string]string),
		Headers:     headers,
	}
}
