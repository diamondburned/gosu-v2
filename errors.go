package osu

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type ResponseError struct {
	HTTPCode int
	HTTPBody []byte
}

func NewResponseError(r *http.Response, body []byte) error {
	if body == nil {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		body = b
	}

	return &ResponseError{
		HTTPCode: r.StatusCode,
		HTTPBody: body,
	}
}

func (err *ResponseError) Error() string {
	return "Invalid status code: " + strconv.Itoa(err.HTTPCode)
}
