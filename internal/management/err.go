package management

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrHttp struct {
	Method     string
	Url        string
	StatusCode int
	Response   ErrorResponse
}

func NewErrHttp(method, url string, statusCode int, body []byte) *ErrHttp {
	var response ErrorResponse
	json.Unmarshal(body, &response)
	return &ErrHttp{Method: method, Url: url, StatusCode: statusCode, Response: response}
}

func (e *ErrHttp) Error() string {
	return fmt.Sprintf("%s: %s - %d %s %s %s",
		e.Method, e.Url, e.StatusCode, http.StatusText(e.StatusCode), e.Response.Error, e.Response.ErrorDescription)
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
