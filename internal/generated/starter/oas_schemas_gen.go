// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *DefaultErrorResponseStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/DefaultErrorResponse
type DefaultErrorResponse struct {
	Success DefaultErrorResponseSuccess `json:"success"`
	Data    DefaultErrorResponseData    `json:"data"`
	Error   DefaultErrorResponseError   `json:"error"`
}

// GetSuccess returns the value of Success.
func (s *DefaultErrorResponse) GetSuccess() DefaultErrorResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *DefaultErrorResponse) GetData() DefaultErrorResponseData {
	return s.Data
}

// GetError returns the value of Error.
func (s *DefaultErrorResponse) GetError() DefaultErrorResponseError {
	return s.Error
}

// SetSuccess sets the value of Success.
func (s *DefaultErrorResponse) SetSuccess(val DefaultErrorResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *DefaultErrorResponse) SetData(val DefaultErrorResponseData) {
	s.Data = val
}

// SetError sets the value of Error.
func (s *DefaultErrorResponse) SetError(val DefaultErrorResponseError) {
	s.Error = val
}

type DefaultErrorResponseData struct{}

type DefaultErrorResponseError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// GetCode returns the value of Code.
func (s *DefaultErrorResponseError) GetCode() string {
	return s.Code
}

// GetDescription returns the value of Description.
func (s *DefaultErrorResponseError) GetDescription() string {
	return s.Description
}

// SetCode sets the value of Code.
func (s *DefaultErrorResponseError) SetCode(val string) {
	s.Code = val
}

// SetDescription sets the value of Description.
func (s *DefaultErrorResponseError) SetDescription(val string) {
	s.Description = val
}

// DefaultErrorResponseStatusCode wraps DefaultErrorResponse with StatusCode.
type DefaultErrorResponseStatusCode struct {
	StatusCode int
	Response   DefaultErrorResponse
}

// GetStatusCode returns the value of StatusCode.
func (s *DefaultErrorResponseStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *DefaultErrorResponseStatusCode) GetResponse() DefaultErrorResponse {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *DefaultErrorResponseStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *DefaultErrorResponseStatusCode) SetResponse(val DefaultErrorResponse) {
	s.Response = val
}

type DefaultErrorResponseSuccess bool

const (
	DefaultErrorResponseSuccessFalse DefaultErrorResponseSuccess = false
)

// AllValues returns all DefaultErrorResponseSuccess values.
func (DefaultErrorResponseSuccess) AllValues() []DefaultErrorResponseSuccess {
	return []DefaultErrorResponseSuccess{
		DefaultErrorResponseSuccessFalse,
	}
}

// Ref: #/components/schemas/DefaultSuccessResponse
type DefaultSuccessResponse struct {
	Success DefaultSuccessResponseSuccess `json:"success"`
	Data    DefaultSuccessResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *DefaultSuccessResponse) GetSuccess() DefaultSuccessResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *DefaultSuccessResponse) GetData() DefaultSuccessResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *DefaultSuccessResponse) SetSuccess(val DefaultSuccessResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *DefaultSuccessResponse) SetData(val DefaultSuccessResponseData) {
	s.Data = val
}

type DefaultSuccessResponseData struct{}

type DefaultSuccessResponseSuccess bool

const (
	DefaultSuccessResponseSuccessTrue DefaultSuccessResponseSuccess = true
)

// AllValues returns all DefaultSuccessResponseSuccess values.
func (DefaultSuccessResponseSuccess) AllValues() []DefaultSuccessResponseSuccess {
	return []DefaultSuccessResponseSuccess{
		DefaultSuccessResponseSuccessTrue,
	}
}
