package httperror

import (
	"fmt"
	"net/http"
)

// HTTPError http error
type HTTPError struct {
	Status int
	Name   string
	Msg    string
}

func (err HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", err.Name, err.Msg)
}

// WithName return copy error with given name
func (err HTTPError) WithName(name string) *HTTPError {
	err.Name = name
	return &err
}

// WithMsg return copy error with given msg
func (err HTTPError) WithMsg(msg string) *HTTPError {
	err.Msg = msg
	return &err
}

// WithStatus return copy error with given status
func (err HTTPError) WithStatus(status int) *HTTPError {
	err.Status = status
	return &err
}

// WithMsg return copy error with given error
func (err HTTPError) WithErr(e error) *HTTPError {
	err.Msg = e.Error()
	return &err
}

var err = HTTPError{}

// common http error
var (
	InvalidParams       = err.WithName("InvalidParams").WithStatus(http.StatusBadRequest)
	Forbidden           = err.WithName("Forbidden").WithStatus(http.StatusForbidden)
	Unauthorized        = err.WithName("Unauthorized").WithStatus(http.StatusUnauthorized)
	InternalServerError = err.WithName("InternalServerError").WithStatus(http.StatusInternalServerError)
)

// business error
var (
	AccessTokenRequired = Unauthorized.WithName("AccessTokenRequired").WithMsg("access token required")
	AccessTokenInvalid  = Unauthorized.WithName("AccessTokenInvalid").WithMsg("access token invalid")

	EmailPasswordNotMatch = Unauthorized.WithName("EmailPasswordNotMatch").WithMsg("email/password not match")
)
