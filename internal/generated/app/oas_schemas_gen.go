// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"
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

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Session
type Session struct {
	ID       string   `json:"id"`
	Token    string   `json:"token"`
	UserRole UserRole `json:"userRole"`
}

// GetID returns the value of ID.
func (s *Session) GetID() string {
	return s.ID
}

// GetToken returns the value of Token.
func (s *Session) GetToken() string {
	return s.Token
}

// GetUserRole returns the value of UserRole.
func (s *Session) GetUserRole() UserRole {
	return s.UserRole
}

// SetID sets the value of ID.
func (s *Session) SetID(val string) {
	s.ID = val
}

// SetToken sets the value of Token.
func (s *Session) SetToken(val string) {
	s.Token = val
}

// SetUserRole sets the value of UserRole.
func (s *Session) SetUserRole(val UserRole) {
	s.UserRole = val
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

// Ref: #/components/schemas/User
type User struct {
	ID    string   `json:"id"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

// GetID returns the value of ID.
func (s *User) GetID() string {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetRole returns the value of Role.
func (s *User) GetRole() UserRole {
	return s.Role
}

// SetID sets the value of ID.
func (s *User) SetID(val string) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val UserRole) {
	s.Role = val
}

type UserPassword string

// Ref: #/components/schemas/UserRole
type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

// AllValues returns all UserRole values.
func (UserRole) AllValues() []UserRole {
	return []UserRole{
		UserRoleAdmin,
		UserRoleUser,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UserRole) MarshalText() ([]byte, error) {
	switch s {
	case UserRoleAdmin:
		return []byte(s), nil
	case UserRoleUser:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserRole) UnmarshalText(data []byte) error {
	switch UserRole(data) {
	case UserRoleAdmin:
		*s = UserRoleAdmin
		return nil
	case UserRoleUser:
		*s = UserRoleUser
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/V1AuthorizeUsingOauthRequestBody
type V1AuthorizeUsingOauthRequestBody struct {
	OauthSource V1AuthorizeUsingOauthRequestBodyOauthSource `json:"oauthSource"`
}

// GetOauthSource returns the value of OauthSource.
func (s *V1AuthorizeUsingOauthRequestBody) GetOauthSource() V1AuthorizeUsingOauthRequestBodyOauthSource {
	return s.OauthSource
}

// SetOauthSource sets the value of OauthSource.
func (s *V1AuthorizeUsingOauthRequestBody) SetOauthSource(val V1AuthorizeUsingOauthRequestBodyOauthSource) {
	s.OauthSource = val
}

type V1AuthorizeUsingOauthRequestBodyOauthSource string

const (
	V1AuthorizeUsingOauthRequestBodyOauthSourceVk     V1AuthorizeUsingOauthRequestBodyOauthSource = "vk"
	V1AuthorizeUsingOauthRequestBodyOauthSourceOk     V1AuthorizeUsingOauthRequestBodyOauthSource = "ok"
	V1AuthorizeUsingOauthRequestBodyOauthSourceMail   V1AuthorizeUsingOauthRequestBodyOauthSource = "mail"
	V1AuthorizeUsingOauthRequestBodyOauthSourceYandex V1AuthorizeUsingOauthRequestBodyOauthSource = "yandex"
)

// AllValues returns all V1AuthorizeUsingOauthRequestBodyOauthSource values.
func (V1AuthorizeUsingOauthRequestBodyOauthSource) AllValues() []V1AuthorizeUsingOauthRequestBodyOauthSource {
	return []V1AuthorizeUsingOauthRequestBodyOauthSource{
		V1AuthorizeUsingOauthRequestBodyOauthSourceVk,
		V1AuthorizeUsingOauthRequestBodyOauthSourceOk,
		V1AuthorizeUsingOauthRequestBodyOauthSourceMail,
		V1AuthorizeUsingOauthRequestBodyOauthSourceYandex,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s V1AuthorizeUsingOauthRequestBodyOauthSource) MarshalText() ([]byte, error) {
	switch s {
	case V1AuthorizeUsingOauthRequestBodyOauthSourceVk:
		return []byte(s), nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceOk:
		return []byte(s), nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceMail:
		return []byte(s), nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceYandex:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *V1AuthorizeUsingOauthRequestBodyOauthSource) UnmarshalText(data []byte) error {
	switch V1AuthorizeUsingOauthRequestBodyOauthSource(data) {
	case V1AuthorizeUsingOauthRequestBodyOauthSourceVk:
		*s = V1AuthorizeUsingOauthRequestBodyOauthSourceVk
		return nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceOk:
		*s = V1AuthorizeUsingOauthRequestBodyOauthSourceOk
		return nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceMail:
		*s = V1AuthorizeUsingOauthRequestBodyOauthSourceMail
		return nil
	case V1AuthorizeUsingOauthRequestBodyOauthSourceYandex:
		*s = V1AuthorizeUsingOauthRequestBodyOauthSourceYandex
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/V1AuthorizeUsingOauthResponse
type V1AuthorizeUsingOauthResponse struct {
	URL string `json:"url"`
}

// GetURL returns the value of URL.
func (s *V1AuthorizeUsingOauthResponse) GetURL() string {
	return s.URL
}

// SetURL sets the value of URL.
func (s *V1AuthorizeUsingOauthResponse) SetURL(val string) {
	s.URL = val
}

// Ref: #/components/schemas/V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody
type V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody struct {
	ConfirmationToken string       `json:"confirmationToken"`
	NewPassword       UserPassword `json:"newPassword"`
}

// GetConfirmationToken returns the value of ConfirmationToken.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) GetConfirmationToken() string {
	return s.ConfirmationToken
}

// GetNewPassword returns the value of NewPassword.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) GetNewPassword() UserPassword {
	return s.NewPassword
}

// SetConfirmationToken sets the value of ConfirmationToken.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) SetConfirmationToken(val string) {
	s.ConfirmationToken = val
}

// SetNewPassword sets the value of NewPassword.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) SetNewPassword(val UserPassword) {
	s.NewPassword = val
}

// Ref: #/components/schemas/V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse
type V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse struct {
	Success V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess `json:"success"`
	Data    V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse) GetSuccess() V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse) GetData() V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse) SetSuccess(val V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse) SetData(val V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData) {
	s.Data = val
}

type V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData struct {
	Session Session `json:"session"`
}

// GetSession returns the value of Session.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData) GetSession() Session {
	return s.Session
}

// SetSession sets the value of Session.
func (s *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseData) SetSession(val Session) {
	s.Session = val
}

type V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess bool

const (
	V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccessTrue V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess = true
)

// AllValues returns all V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess values.
func (V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess) AllValues() []V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess {
	return []V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccess{
		V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponseSuccessTrue,
	}
}

// Ref: #/components/schemas/V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody
type V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody struct {
	Email string `json:"email"`
}

// GetEmail returns the value of Email.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) GetEmail() string {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) SetEmail(val string) {
	s.Email = val
}

// Ref: #/components/schemas/V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse
type V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse struct {
	Success V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess `json:"success"`
	Data    V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse) GetSuccess() V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse) GetData() V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse) SetSuccess(val V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse) SetData(val V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData) {
	s.Data = val
}

type V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// GetID returns the value of ID.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData) GetID() string {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData) GetEmail() string {
	return s.Email
}

// SetID sets the value of ID.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData) SetID(val string) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseData) SetEmail(val string) {
	s.Email = val
}

type V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess bool

const (
	V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccessTrue V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess = true
)

// AllValues returns all V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess values.
func (V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess) AllValues() []V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess {
	return []V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccess{
		V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponseSuccessTrue,
	}
}

// Ref: #/components/schemas/V1RegistrationsConfirmRegistrationRequestBody
type V1RegistrationsConfirmRegistrationRequestBody struct {
	ConfirmationToken string       `json:"confirmationToken"`
	Password          UserPassword `json:"password"`
}

// GetConfirmationToken returns the value of ConfirmationToken.
func (s *V1RegistrationsConfirmRegistrationRequestBody) GetConfirmationToken() string {
	return s.ConfirmationToken
}

// GetPassword returns the value of Password.
func (s *V1RegistrationsConfirmRegistrationRequestBody) GetPassword() UserPassword {
	return s.Password
}

// SetConfirmationToken sets the value of ConfirmationToken.
func (s *V1RegistrationsConfirmRegistrationRequestBody) SetConfirmationToken(val string) {
	s.ConfirmationToken = val
}

// SetPassword sets the value of Password.
func (s *V1RegistrationsConfirmRegistrationRequestBody) SetPassword(val UserPassword) {
	s.Password = val
}

// Ref: #/components/schemas/V1RegistrationsConfirmRegistrationResponse
type V1RegistrationsConfirmRegistrationResponse struct {
	Success V1RegistrationsConfirmRegistrationResponseSuccess `json:"success"`
	Data    V1RegistrationsConfirmRegistrationResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1RegistrationsConfirmRegistrationResponse) GetSuccess() V1RegistrationsConfirmRegistrationResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1RegistrationsConfirmRegistrationResponse) GetData() V1RegistrationsConfirmRegistrationResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1RegistrationsConfirmRegistrationResponse) SetSuccess(val V1RegistrationsConfirmRegistrationResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1RegistrationsConfirmRegistrationResponse) SetData(val V1RegistrationsConfirmRegistrationResponseData) {
	s.Data = val
}

type V1RegistrationsConfirmRegistrationResponseData struct {
	Session Session `json:"session"`
}

// GetSession returns the value of Session.
func (s *V1RegistrationsConfirmRegistrationResponseData) GetSession() Session {
	return s.Session
}

// SetSession sets the value of Session.
func (s *V1RegistrationsConfirmRegistrationResponseData) SetSession(val Session) {
	s.Session = val
}

type V1RegistrationsConfirmRegistrationResponseSuccess bool

const (
	V1RegistrationsConfirmRegistrationResponseSuccessTrue V1RegistrationsConfirmRegistrationResponseSuccess = true
)

// AllValues returns all V1RegistrationsConfirmRegistrationResponseSuccess values.
func (V1RegistrationsConfirmRegistrationResponseSuccess) AllValues() []V1RegistrationsConfirmRegistrationResponseSuccess {
	return []V1RegistrationsConfirmRegistrationResponseSuccess{
		V1RegistrationsConfirmRegistrationResponseSuccessTrue,
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

// Ref: #/components/schemas/V1SessionsCreateSessionRequestBody
type V1SessionsCreateSessionRequestBody struct {
	Email    string       `json:"email"`
	Password UserPassword `json:"password"`
}

// GetEmail returns the value of Email.
func (s *V1SessionsCreateSessionRequestBody) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *V1SessionsCreateSessionRequestBody) GetPassword() UserPassword {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *V1SessionsCreateSessionRequestBody) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *V1SessionsCreateSessionRequestBody) SetPassword(val UserPassword) {
	s.Password = val
}

// Ref: #/components/schemas/V1SessionsCreateSessionResponse
type V1SessionsCreateSessionResponse struct {
	Success V1SessionsCreateSessionResponseSuccess `json:"success"`
	Data    V1SessionsCreateSessionResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1SessionsCreateSessionResponse) GetSuccess() V1SessionsCreateSessionResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1SessionsCreateSessionResponse) GetData() V1SessionsCreateSessionResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1SessionsCreateSessionResponse) SetSuccess(val V1SessionsCreateSessionResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1SessionsCreateSessionResponse) SetData(val V1SessionsCreateSessionResponseData) {
	s.Data = val
}

type V1SessionsCreateSessionResponseData struct {
	Session Session `json:"session"`
}

// GetSession returns the value of Session.
func (s *V1SessionsCreateSessionResponseData) GetSession() Session {
	return s.Session
}

// SetSession sets the value of Session.
func (s *V1SessionsCreateSessionResponseData) SetSession(val Session) {
	s.Session = val
}

type V1SessionsCreateSessionResponseSuccess bool

const (
	V1SessionsCreateSessionResponseSuccessTrue V1SessionsCreateSessionResponseSuccess = true
)

// AllValues returns all V1SessionsCreateSessionResponseSuccess values.
func (V1SessionsCreateSessionResponseSuccess) AllValues() []V1SessionsCreateSessionResponseSuccess {
	return []V1SessionsCreateSessionResponseSuccess{
		V1SessionsCreateSessionResponseSuccessTrue,
	}
}

// Ref: #/components/schemas/V1UsersGetUsersResponse
type V1UsersGetUsersResponse struct {
	Success V1UsersGetUsersResponseSuccess `json:"success"`
	Data    V1UsersGetUsersResponseData    `json:"data"`
}

// GetSuccess returns the value of Success.
func (s *V1UsersGetUsersResponse) GetSuccess() V1UsersGetUsersResponseSuccess {
	return s.Success
}

// GetData returns the value of Data.
func (s *V1UsersGetUsersResponse) GetData() V1UsersGetUsersResponseData {
	return s.Data
}

// SetSuccess sets the value of Success.
func (s *V1UsersGetUsersResponse) SetSuccess(val V1UsersGetUsersResponseSuccess) {
	s.Success = val
}

// SetData sets the value of Data.
func (s *V1UsersGetUsersResponse) SetData(val V1UsersGetUsersResponseData) {
	s.Data = val
}

type V1UsersGetUsersResponseData struct {
	Users  []User `json:"users"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
}

// GetUsers returns the value of Users.
func (s *V1UsersGetUsersResponseData) GetUsers() []User {
	return s.Users
}

// GetLimit returns the value of Limit.
func (s *V1UsersGetUsersResponseData) GetLimit() int {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *V1UsersGetUsersResponseData) GetOffset() int {
	return s.Offset
}

// GetTotal returns the value of Total.
func (s *V1UsersGetUsersResponseData) GetTotal() int {
	return s.Total
}

// SetUsers sets the value of Users.
func (s *V1UsersGetUsersResponseData) SetUsers(val []User) {
	s.Users = val
}

// SetLimit sets the value of Limit.
func (s *V1UsersGetUsersResponseData) SetLimit(val int) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *V1UsersGetUsersResponseData) SetOffset(val int) {
	s.Offset = val
}

// SetTotal sets the value of Total.
func (s *V1UsersGetUsersResponseData) SetTotal(val int) {
	s.Total = val
}

type V1UsersGetUsersResponseSuccess bool

const (
	V1UsersGetUsersResponseSuccessTrue V1UsersGetUsersResponseSuccess = true
)

// AllValues returns all V1UsersGetUsersResponseSuccess values.
func (V1UsersGetUsersResponseSuccess) AllValues() []V1UsersGetUsersResponseSuccess {
	return []V1UsersGetUsersResponseSuccess{
		V1UsersGetUsersResponseSuccessTrue,
	}
}
