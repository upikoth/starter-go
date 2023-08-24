package model

type ExtendedError struct {
	ResponseCode         int
	ResponseErrorCode    error
	ResponseErrorDetails error
}
