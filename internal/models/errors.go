package models

import "net/http"

type ErrorCode string

const (
	ErrCodeValidationByOpenapi ErrorCode = "1"
	ErrCodeUserUnauthorized    ErrorCode = "2"
	ErrCodeInterval            ErrorCode = "3"

	ErrCodeRegistrationUserWithThisEmailAlreadyExist ErrorCode = "100"
	ErrCodeRegistrationNotFound                      ErrorCode = "101"

	ErrorCodeCreateSessionWrongEmailOrPassword ErrorCode = "200"
	ErrorCodeSessionNotFound                   ErrorCode = "201"

	ErrorCodePasswordRecoveryRequestNotFound ErrorCode = "300"

	ErrorCodeUsersGetListForbidden ErrorCode = "400"

	ErrorCodeOauthSourceNotExist ErrorCode = "500"
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
		return http.StatusInternalServerError
	}

	return e.StatusCode
}
