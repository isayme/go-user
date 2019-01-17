package httperror

import (
	"fmt"
	"net/http"
)

// HTTPError http error
type HTTPError struct {
	Status int
	Msg    string
}

func (err HTTPError) Error() string {
	return fmt.Sprintf("stauts: %d, msg: %s", err.Status, err.Msg)
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

var err = HTTPError{}

// common http error
var (
	InvalidParams       = err.WithStatus(http.StatusBadRequest)
	Forbidden           = err.WithStatus(http.StatusForbidden)
	Unauthorized        = err.WithStatus(http.StatusUnauthorized)
	InternalServerError = err.WithStatus(http.StatusInternalServerError)
)

// business error
var (
	AccessTokenRequired = Unauthorized.WithMsg("access token required")
	AccessTokenInvalid  = Unauthorized.WithMsg("access token invalid")

	UsernamePasswordNotMatch = Unauthorized.WithMsg("username/password not match")
)
