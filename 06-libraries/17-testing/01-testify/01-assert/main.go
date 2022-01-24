package main

import "errors"

type Request struct {
	Page int
}

type Response struct {
	Status	bool
	Data 	[]string
	Error	error
}

func Success(data []string) *Response {
	r := Response{}
	r.Status = true
	r.Data = data
	r.Error = nil
	return &r
}

func Fail(err error) *Response {
	r := Response{}
	r.Status = false
	r.Data = make([]string, 0)
	r.Error = err
	return &r
}

var (
	RequestCouldNotBeNil = errors.New("the request could not be nil")
	RequestPageCouldNotBeLessThanOrEqualZero = errors.New("the request page could not be less than or equal zero")
)

func Get(request *Request) *Response {
	if request == nil {
		return Fail(RequestCouldNotBeNil)
	}

	if request.Page <= 0 {
		return Fail(RequestPageCouldNotBeLessThanOrEqualZero)
	}

	return Success([]string{
		"Item 1",
		"Item 2",
		"Item 3",
		"Item 4",
	})
}