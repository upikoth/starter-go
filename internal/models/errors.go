package models

import "net/http"

type ErrorCode string

const (
	ErrorCodeValidationByOpenapi ErrorCode = "1"

	ErrorCodeRegistrationSMTPSendEmail                 ErrorCode = "100"
	ErrorCodeRegistrationYdbStarterCreateRegistration  ErrorCode = "101"
	ErrorCodeRegistrationYdbStarterFindUser            ErrorCode = "102"
	ErrorCodeRegistrationUserWithThisEmailAlreadyExist ErrorCode = "103"

	ErrorCodeRegistrationYdbStarterCheckConfirmationToken ErrorCode = "200"
	ErrorCodeRegistrationYdbStarterRegistrationNotFound   ErrorCode = "201"
	ErrorCodeRegistrationGeneratePasswordHash             ErrorCode = "202"
	ErrorCodeRegistrationCreateSession                    ErrorCode = "203"
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
