package models

import "net/http"

type ErrorCode string

const (
	ErrorCodeValidationByOpenapi ErrorCode = "1"
	ErrorCodeUserUnauthorized    ErrorCode = "2"

	ErrorCodeRegistrationSMTPSendEmail                 ErrorCode = "100"
	ErrorCodeRegistrationYdbCreateRegistration         ErrorCode = "101"
	ErrorCodeRegistrationYdbFindUser                   ErrorCode = "102"
	ErrorCodeRegistrationUserWithThisEmailAlreadyExist ErrorCode = "103"

	ErrorCodeRegistrationYdbCheckConfirmationToken ErrorCode = "200"
	ErrorCodeRegistrationRegistrationNotFound      ErrorCode = "201"
	ErrorCodeRegistrationGeneratePasswordHash      ErrorCode = "202"
	ErrorCodeRegistrationCreateSession             ErrorCode = "203"
	ErrorCodeRegistrationDBError                   ErrorCode = "204"

	ErrorCodeSessionsCreateSessionDBError              ErrorCode = "300"
	ErrorCodeSessionsCreateSessionWrongEmailOrPassword ErrorCode = "301"

	ErrorCodeSessionsDeleteSessionDBError  ErrorCode = "400"
	ErrorCodeSessionsDeleteSessionNotFound ErrorCode = "401"

	ErrorCodePasswordRecoveryRequestYdbFindUser                      ErrorCode = "500"
	ErrorCodePasswordRecoveryRequestYdbCreatePasswordRecoveryRequest ErrorCode = "501"
	ErrorCodePasswordRecoveryRequestSMTPSendEmail                    ErrorCode = "502"

	ErrorCodePasswordRecoveryRequestYdbCheckConfirmationToken       ErrorCode = "600"
	ErrorCodePasswordRecoveryRequestPasswordRecoveryRequestNotFound ErrorCode = "601"
	ErrorCodePasswordRecoveryRequestGeneratePasswordHash            ErrorCode = "602"
	ErrorCodePasswordRecoveryRequestFindUserByEmail                 ErrorCode = "603"
	ErrorCodePasswordRecoveryRequestCreateSession                   ErrorCode = "604"
	ErrorCodePasswordRecoveryRequestUpdateUserPassword              ErrorCode = "605"

	ErrorCodeUsersGetListForbidden ErrorCode = "700"
	ErrorCodeUsersGetListDBError   ErrorCode = "701"

	ErrorCodeSessionsCheckTokenDBError ErrorCode = "702"
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
