// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// V1CheckHealth implements V1CheckHealth operation.
	//
	// Получить информацию о работоспособности приложения.
	//
	// GET /api/v1/health
	V1CheckHealth(ctx context.Context) (*SuccessResponse, error)
	// V1CreateRegistration implements V1CreateRegistration operation.
	//
	// Создать заявку на регистрацию пользователя.
	//
	// POST /api/v1/registrations
	V1CreateRegistration(ctx context.Context, req *V1RegistrationsCreateRegistrationRequestBody) (*SuccessResponse, error)
	// NewError creates *ErrorResponseStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorResponseStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
