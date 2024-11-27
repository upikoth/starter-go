// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// V1AuthorizeUsingOauth implements V1AuthorizeUsingOauth operation.
//
// Авторизация в приложении с помощью oauth.
//
// POST /api/v1/oauth
func (UnimplementedHandler) V1AuthorizeUsingOauth(ctx context.Context, req *V1AuthorizeUsingOauthRequestBody) (r *V1AuthorizeUsingOauthResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AuthorizeUsingOauthHandleMailRedirect implements V1AuthorizeUsingOauthHandleMailRedirect operation.
//
// Обработка редиректа после авторизации в mail.ru.
//
// GET /api/v1/oauthRedirect/mail
func (UnimplementedHandler) V1AuthorizeUsingOauthHandleMailRedirect(ctx context.Context, params V1AuthorizeUsingOauthHandleMailRedirectParams) (r *V1AuthorizeUsingOauthHandleMailRedirectFound, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AuthorizeUsingOauthHandleVkRedirect implements V1AuthorizeUsingOauthHandleVkRedirect operation.
//
// Обработка редиректа после авторизации в vk.
//
// GET /api/v1/oauthRedirect/vk
func (UnimplementedHandler) V1AuthorizeUsingOauthHandleVkRedirect(ctx context.Context, params V1AuthorizeUsingOauthHandleVkRedirectParams) (r *V1AuthorizeUsingOauthHandleVkRedirectFound, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AuthorizeUsingOauthHandleYandexRedirect implements V1AuthorizeUsingOauthHandleYandexRedirect operation.
//
// Обработка редиректа после авторизации в yandex.
//
// GET /api/v1/oauthRedirect/yandex
func (UnimplementedHandler) V1AuthorizeUsingOauthHandleYandexRedirect(ctx context.Context, params V1AuthorizeUsingOauthHandleYandexRedirectParams) (r *V1AuthorizeUsingOauthHandleYandexRedirectFound, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CheckCurrentSession implements V1CheckCurrentSession operation.
//
// Получить информацию валидна ли текущая сессия.
//
// GET /api/v1/session
func (UnimplementedHandler) V1CheckCurrentSession(ctx context.Context, params V1CheckCurrentSessionParams) (r *SuccessResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CheckHealth implements V1CheckHealth operation.
//
// Получить информацию о работоспособности приложения.
//
// GET /api/v1/health
func (UnimplementedHandler) V1CheckHealth(ctx context.Context) (r *SuccessResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1ConfirmPasswordRecoveryRequest implements V1ConfirmPasswordRecoveryRequest operation.
//
// Подтверждение заявки на восстановление пароля.
//
// PATCH /api/v1/passwordRecoveryRequests
func (UnimplementedHandler) V1ConfirmPasswordRecoveryRequest(ctx context.Context, req *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) (r *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1ConfirmRegistration implements V1ConfirmRegistration operation.
//
// Подтверждение заявки на регистрацию.
//
// PATCH /api/v1/registrations
func (UnimplementedHandler) V1ConfirmRegistration(ctx context.Context, req *V1RegistrationsConfirmRegistrationRequestBody) (r *V1RegistrationsConfirmRegistrationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CreatePasswordRecoveryRequest implements V1CreatePasswordRecoveryRequest operation.
//
// Создать заявку на восстановление пароля.
//
// POST /api/v1/passwordRecoveryRequests
func (UnimplementedHandler) V1CreatePasswordRecoveryRequest(ctx context.Context, req *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) (r *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CreateRegistration implements V1CreateRegistration operation.
//
// Создать заявку на регистрацию пользователя.
//
// POST /api/v1/registrations
func (UnimplementedHandler) V1CreateRegistration(ctx context.Context, req *V1RegistrationsCreateRegistrationRequestBody) (r *V1RegistrationsCreateRegistrationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CreateSession implements V1CreateSession operation.
//
// Создание сессии пользователя.
//
// POST /api/v1/sessions
func (UnimplementedHandler) V1CreateSession(ctx context.Context, req *V1SessionsCreateSessionRequestBody) (r *V1SessionsCreateSessionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1DeleteSession implements V1DeleteSession operation.
//
// Удаление сессии пользователя.
//
// DELETE /api/v1/sessions/{id}
func (UnimplementedHandler) V1DeleteSession(ctx context.Context, params V1DeleteSessionParams) (r *SuccessResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// V1GetUsers implements V1GetUsers operation.
//
// Получение информации обо всех пользователях.
//
// GET /api/v1/users
func (UnimplementedHandler) V1GetUsers(ctx context.Context, params V1GetUsersParams) (r *V1UsersGetUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorResponseStatusCode) {
	r = new(ErrorResponseStatusCode)
	return r
}
