// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorResponseStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	Success ErrorResponseSuccess `json:"success"`
	Data    ErrorResponseData    `json:"data"`
	Error   ErrorResponseError   `json:"error"`
}

// GetSuccess returns the value of Success.
func (s *ErrorResponse) GetSuccess() ErrorResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *ErrorResponse) GetData() ErrorResponseData {
	return s.Data
}

// GetError returns the value of Error.
func (s *ErrorResponse) GetError() ErrorResponseError {
	return s.Error
}

// SetSuccess sets the value of Success.
func (s *ErrorResponse) SetSuccess(val ErrorResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *ErrorResponse) SetData(val ErrorResponseData) {
	s.Data = val
}

// SetError sets the value of Error.
func (s *ErrorResponse) SetError(val ErrorResponseError) {
	s.Error = val
}

type ErrorResponseData struct{}

type ErrorResponseError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// GetCode returns the value of Code.
func (s *ErrorResponseError) GetCode() string {
	return s.Code
}

// GetDescription returns the value of Description.
func (s *ErrorResponseError) GetDescription() string {
	return s.Description
}

// SetCode sets the value of Code.
func (s *ErrorResponseError) SetCode(val string) {
	s.Code = val
}

// SetDescription sets the value of Description.
func (s *ErrorResponseError) SetDescription(val string) {
	s.Description = val
}

// ErrorResponseStatusCode wraps ErrorResponse with StatusCode.
type ErrorResponseStatusCode struct {
	StatusCode int
	Response   ErrorResponse
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorResponseStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorResponseStatusCode) GetResponse() ErrorResponse {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorResponseStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorResponseStatusCode) SetResponse(val ErrorResponse) {
	s.Response = val
}

type ErrorResponseSuccess bool

const (
	ErrorResponseSuccessFalse ErrorResponseSuccess = false
)

// AllValues returns all ErrorResponseSuccess values.
func (ErrorResponseSuccess) AllValues() []ErrorResponseSuccess {
	return []ErrorResponseSuccess{
		ErrorResponseSuccessFalse,
	}
}

// Ref: #/components/schemas/SuccessResponse
type SuccessResponse struct {
	Success SuccessResponseSuccess `json:"success"`
	Data    SuccessResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *SuccessResponse) GetSuccess() SuccessResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *SuccessResponse) GetData() SuccessResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *SuccessResponse) SetSuccess(val SuccessResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *SuccessResponse) SetData(val SuccessResponseData) {
	s.Data = val
}

type SuccessResponseData struct{}

type SuccessResponseSuccess bool

const (
	SuccessResponseSuccessTrue SuccessResponseSuccess = true
)

// AllValues returns all SuccessResponseSuccess values.
func (SuccessResponseSuccess) AllValues() []SuccessResponseSuccess {
	return []SuccessResponseSuccess{
		SuccessResponseSuccessTrue,
	}
}

// Ref: #/components/schemas/V1RegistrationsCreateRegistrationRequestBody
type V1RegistrationsCreateRegistrationRequestBody struct {
	Email string `json:"email"`
}

// GetEmail returns the value of Email.
func (s *V1RegistrationsCreateRegistrationRequestBody) GetEmail() string {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *V1RegistrationsCreateRegistrationRequestBody) SetEmail(val string) {
	s.Email = val
}

// Ref: #/components/schemas/V1RegistrationsCreateRegistrationResponse
type V1RegistrationsCreateRegistrationResponse struct {
	Success V1RegistrationsCreateRegistrationResponseSuccess `json:"success"`
	Data    V1RegistrationsCreateRegistrationResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1RegistrationsCreateRegistrationResponse) GetSuccess() V1RegistrationsCreateRegistrationResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1RegistrationsCreateRegistrationResponse) GetData() V1RegistrationsCreateRegistrationResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1RegistrationsCreateRegistrationResponse) SetSuccess(val V1RegistrationsCreateRegistrationResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1RegistrationsCreateRegistrationResponse) SetData(val V1RegistrationsCreateRegistrationResponseData) {
	s.Data = val
}

type V1RegistrationsCreateRegistrationResponseData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// GetID returns the value of ID.
func (s *V1RegistrationsCreateRegistrationResponseData) GetID() string {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *V1RegistrationsCreateRegistrationResponseData) GetEmail() string {
	return s.Email
}

// SetID sets the value of ID.
func (s *V1RegistrationsCreateRegistrationResponseData) SetID(val string) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *V1RegistrationsCreateRegistrationResponseData) SetEmail(val string) {
	s.Email = val
}

type V1RegistrationsCreateRegistrationResponseSuccess bool

const (
	V1RegistrationsCreateRegistrationResponseSuccessTrue V1RegistrationsCreateRegistrationResponseSuccess = true
)

// AllValues returns all V1RegistrationsCreateRegistrationResponseSuccess values.
func (V1RegistrationsCreateRegistrationResponseSuccess) AllValues() []V1RegistrationsCreateRegistrationResponseSuccess {
	return []V1RegistrationsCreateRegistrationResponseSuccess{
		V1RegistrationsCreateRegistrationResponseSuccessTrue,
	}
}
