package schema

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var ErrNotFoundEmail = errors.New("email not found")
var ErrUsernameOrPasswordFailed = errors.New("username or password failed")
var ErrEmailOrPasswordFailed = errors.New("email or password failed")
var ErrEmailAleadyExist = errors.New("email already exist")
var ErrEmailCodeNotEqual = errors.New("email code not equal")
