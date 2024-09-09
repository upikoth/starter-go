// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// V1CheckCurrentSession implements V1CheckCurrentSession operation.
	//
	// Получить информацию валидна ли текущая сессия.
	//
	// GET /api/v1/session
	V1CheckCurrentSession(ctx context.Context, params V1CheckCurrentSessionParams) (*SuccessResponse, error)
	// V1CheckHealth implements V1CheckHealth operation.
	//
	// Получить информацию о работоспособности приложения.
	//
	// GET /api/v1/health
	V1CheckHealth(ctx context.Context) (*SuccessResponse, error)
	// V1ConfirmPasswordRecoveryRequest implements V1ConfirmPasswordRecoveryRequest operation.
	//
	// Подтверждение заявки на восстановление пароля.
	//
	// PATCH /api/v1/passwordRecoveryRequests
	V1ConfirmPasswordRecoveryRequest(ctx context.Context, req *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, error)
	// V1ConfirmRegistration implements V1ConfirmRegistration operation.
	//
	// Подтверждение заявки на регистрацию.
	//
	// PATCH /api/v1/registrations
	V1ConfirmRegistration(ctx context.Context, req *V1RegistrationsConfirmRegistrationRequestBody) (*V1RegistrationsConfirmRegistrationResponse, error)
	// V1CreatePasswordRecoveryRequest implements V1CreatePasswordRecoveryRequest operation.
	//
	// Создать заявку на восстановление пароля.
	//
	// POST /api/v1/passwordRecoveryRequests
	V1CreatePasswordRecoveryRequest(ctx context.Context, req *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, error)
	// V1CreateRegistration implements V1CreateRegistration operation.
	//
	// Создать заявку на регистрацию пользователя.
	//
	// POST /api/v1/registrations
	V1CreateRegistration(ctx context.Context, req *V1RegistrationsCreateRegistrationRequestBody) (*V1RegistrationsCreateRegistrationResponse, error)
	// V1CreateSession implements V1CreateSession operation.
	//
	// Создание сессии пользователя.
	//
	// POST /api/v1/sessions
	V1CreateSession(ctx context.Context, req *V1SessionsCreateSessionRequestBody) (*V1SessionsCreateSessionResponse, error)
	// V1DeleteSession implements V1DeleteSession operation.
	//
	// Удаление сессии пользователя.
	//
	// DELETE /api/v1/sessions/{id}
	V1DeleteSession(ctx context.Context, params V1DeleteSessionParams) (*SuccessResponse, error)
	// V1GetUsers implements V1GetUsers operation.
	//
	// Получение информации обо всех пользователях.
	//
	// GET /api/v1/users
	V1GetUsers(ctx context.Context, params V1GetUsersParams) (*V1UsersGetUsersResponse, error)
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