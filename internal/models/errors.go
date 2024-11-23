package models

import "net/http"

type ErrorCode string

const (
	ErrCodeValidationByOpenapi ErrorCode = "1"
	ErrCodeUserUnauthorized    ErrorCode = "2"
	ErrCodeInterval            ErrorCode = "3"

	ErrorCodeRegistrationUserWithThisEmailAlreadyExist ErrorCode = "103"

	ErrorCodeRegistrationRegistrationNotFound ErrorCode = "201"

	ErrorCodeSessionsCreateSessionWrongEmailOrPassword ErrorCode = "301"

	ErrorCodeSessionsDeleteSessionNotFound ErrorCode = "401"

	ErrorCodePasswordRecoveryRequestPasswordRecoveryRequestNotFound ErrorCode = "601"

	ErrorCodeUsersGetListForbidden ErrorCode = "700"

	ErrorCodeOauthSourceNotExist     ErrorCode = "900"
	ErrorCodeOauthVkTokenCreating    ErrorCode = "901"
	ErrorCodeOauthVkEmailInvalid     ErrorCode = "902"
	ErrorCodeOauthVkUserIDInvalid    ErrorCode = "903"
	ErrorCodeOauthVkGetUserByVkID    ErrorCode = "904"
	ErrorCodeOauthVkGetUserByVkEmail ErrorCode = "905"
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
