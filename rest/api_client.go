package rest

type ApiClient interface {
	Get(string) (*Response, error)
}

type Response struct {
	getBody       func() []byte
	getStatusCode func() int
}

func NewResponse(getBody func() []byte, getStatusCode func() int) *Response {
	return &Response{getBody: getBody, getStatusCode: getStatusCode}
}

func (r Response) Body() []byte {
	return r.getBody()
}

func (r Response) StatusCode() int {
	return r.getStatusCode()
}
