package server

import (
	"net/http"
	"strconv"
)

const StatusCodeDefault = http.StatusOK

type HandlerFunctions struct {
	Get    func(request *Request, response *Response)
	Post   func(request *Request, response *Response)
	Patch  func(request *Request, response *Response)
	Put    func(request *Request, response *Response)
	Delete func(request *Request, response *Response)
}

type Request struct {
	Body       []byte
	Method     string
	Parameters map[string][]string
	Headers    map[string][]string
}

type Response struct {
	Body       interface{}
	StatusCode int
	Error      error
}

func NewResponse() *Response {
	return &Response{
		StatusCode: StatusCodeDefault,
		Body: struct {
		}{},
	}
}

func newErrorResponse(message string, statusCode int, code int) *Response {
	return &Response{
		StatusCode: statusCode,
		Body: map[string]string{
			"msg":  message,
			"code": strconv.Itoa(code),
		},
	}
}
