package models

import "net/http"

type ErrorCode string

const (
	ErrorCodeValidationByOpenapi ErrorCode = "1"
)

type Error struct {
	Code        ErrorCode
	StatusCode  int
	Description string
}

func (e *Error) Error() string {
	return e.Description
}

func (e *Error) GetStatusCode() int {
	if e.StatusCode == 0 {
		return http.StatusBadRequest
	}

	return e.StatusCode
}
