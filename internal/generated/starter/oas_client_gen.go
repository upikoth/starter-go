// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// V1CheckCurrentSession invokes V1CheckCurrentSession operation.
	//
	// Получить информацию валидна ли текущая сессия.
	//
	// GET /api/v1/session
	V1CheckCurrentSession(ctx context.Context, params V1CheckCurrentSessionParams) (*SuccessResponse, error)
	// V1CheckHealth invokes V1CheckHealth operation.
	//
	// Получить информацию о работоспособности приложения.
	//
	// GET /api/v1/health
	V1CheckHealth(ctx context.Context) (*SuccessResponse, error)
	// V1ConfirmPasswordRecoveryRequest invokes V1ConfirmPasswordRecoveryRequest operation.
	//
	// Подтверждение заявки на восстановление пароля.
	//
	// PATCH /api/v1/passwordRecoveryRequests
	V1ConfirmPasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, error)
	// V1ConfirmRegistration invokes V1ConfirmRegistration operation.
	//
	// Подтверждение заявки на регистрацию.
	//
	// PATCH /api/v1/registrations
	V1ConfirmRegistration(ctx context.Context, request *V1RegistrationsConfirmRegistrationRequestBody) (*V1RegistrationsConfirmRegistrationResponse, error)
	// V1CreatePasswordRecoveryRequest invokes V1CreatePasswordRecoveryRequest operation.
	//
	// Создать заявку на восстановление пароля.
	//
	// POST /api/v1/passwordRecoveryRequests
	V1CreatePasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, error)
	// V1CreateRegistration invokes V1CreateRegistration operation.
	//
	// Создать заявку на регистрацию пользователя.
	//
	// POST /api/v1/registrations
	V1CreateRegistration(ctx context.Context, request *V1RegistrationsCreateRegistrationRequestBody) (*V1RegistrationsCreateRegistrationResponse, error)
	// V1CreateSession invokes V1CreateSession operation.
	//
	// Создание сессии пользователя.
	//
	// POST /api/v1/sessions
	V1CreateSession(ctx context.Context, request *V1SessionsCreateSessionRequestBody) (*V1SessionsCreateSessionResponse, error)
	// V1DeleteSession invokes V1DeleteSession operation.
	//
	// Удаление сессии пользователя.
	//
	// DELETE /api/v1/sessions/{id}
	V1DeleteSession(ctx context.Context, params V1DeleteSessionParams) (*SuccessResponse, error)
	// V1GetUsers invokes V1GetUsers operation.
	//
	// Получение информации обо всех пользователях.
	//
	// GET /api/v1/users
	V1GetUsers(ctx context.Context, params V1GetUsersParams) (*V1UsersGetUsersResponse, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}
type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrorResponseStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// V1CheckCurrentSession invokes V1CheckCurrentSession operation.
//
// Получить информацию валидна ли текущая сессия.
//
// GET /api/v1/session
func (c *Client) V1CheckCurrentSession(ctx context.Context, params V1CheckCurrentSessionParams) (*SuccessResponse, error) {
	res, err := c.sendV1CheckCurrentSession(ctx, params)
	return res, err
}

func (c *Client) sendV1CheckCurrentSession(ctx context.Context, params V1CheckCurrentSessionParams) (res *SuccessResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1CheckCurrentSession"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/session"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1CheckCurrentSession",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/session"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Authorization-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AuthorizationToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1CheckCurrentSessionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1CheckHealth invokes V1CheckHealth operation.
//
// Получить информацию о работоспособности приложения.
//
// GET /api/v1/health
func (c *Client) V1CheckHealth(ctx context.Context) (*SuccessResponse, error) {
	res, err := c.sendV1CheckHealth(ctx)
	return res, err
}

func (c *Client) sendV1CheckHealth(ctx context.Context) (res *SuccessResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1CheckHealth"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/health"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1CheckHealth",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/health"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1CheckHealthResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1ConfirmPasswordRecoveryRequest invokes V1ConfirmPasswordRecoveryRequest operation.
//
// Подтверждение заявки на восстановление пароля.
//
// PATCH /api/v1/passwordRecoveryRequests
func (c *Client) V1ConfirmPasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, error) {
	res, err := c.sendV1ConfirmPasswordRecoveryRequest(ctx, request)
	return res, err
}

func (c *Client) sendV1ConfirmPasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestRequestBody) (res *V1PasswordRecoveryRequestsConfirmPasswordRecoveryRequestResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1ConfirmPasswordRecoveryRequest"),
		semconv.HTTPMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/api/v1/passwordRecoveryRequests"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1ConfirmPasswordRecoveryRequest",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/passwordRecoveryRequests"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1ConfirmPasswordRecoveryRequestRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1ConfirmPasswordRecoveryRequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1ConfirmRegistration invokes V1ConfirmRegistration operation.
//
// Подтверждение заявки на регистрацию.
//
// PATCH /api/v1/registrations
func (c *Client) V1ConfirmRegistration(ctx context.Context, request *V1RegistrationsConfirmRegistrationRequestBody) (*V1RegistrationsConfirmRegistrationResponse, error) {
	res, err := c.sendV1ConfirmRegistration(ctx, request)
	return res, err
}

func (c *Client) sendV1ConfirmRegistration(ctx context.Context, request *V1RegistrationsConfirmRegistrationRequestBody) (res *V1RegistrationsConfirmRegistrationResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1ConfirmRegistration"),
		semconv.HTTPMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/api/v1/registrations"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1ConfirmRegistration",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/registrations"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1ConfirmRegistrationRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1ConfirmRegistrationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1CreatePasswordRecoveryRequest invokes V1CreatePasswordRecoveryRequest operation.
//
// Создать заявку на восстановление пароля.
//
// POST /api/v1/passwordRecoveryRequests
func (c *Client) V1CreatePasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) (*V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, error) {
	res, err := c.sendV1CreatePasswordRecoveryRequest(ctx, request)
	return res, err
}

func (c *Client) sendV1CreatePasswordRecoveryRequest(ctx context.Context, request *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestRequestBody) (res *V1PasswordRecoveryRequestsCreatePasswordRecoveryRequestResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1CreatePasswordRecoveryRequest"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/passwordRecoveryRequests"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1CreatePasswordRecoveryRequest",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/passwordRecoveryRequests"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1CreatePasswordRecoveryRequestRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1CreatePasswordRecoveryRequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1CreateRegistration invokes V1CreateRegistration operation.
//
// Создать заявку на регистрацию пользователя.
//
// POST /api/v1/registrations
func (c *Client) V1CreateRegistration(ctx context.Context, request *V1RegistrationsCreateRegistrationRequestBody) (*V1RegistrationsCreateRegistrationResponse, error) {
	res, err := c.sendV1CreateRegistration(ctx, request)
	return res, err
}

func (c *Client) sendV1CreateRegistration(ctx context.Context, request *V1RegistrationsCreateRegistrationRequestBody) (res *V1RegistrationsCreateRegistrationResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1CreateRegistration"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/registrations"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1CreateRegistration",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/registrations"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1CreateRegistrationRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1CreateRegistrationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1CreateSession invokes V1CreateSession operation.
//
// Создание сессии пользователя.
//
// POST /api/v1/sessions
func (c *Client) V1CreateSession(ctx context.Context, request *V1SessionsCreateSessionRequestBody) (*V1SessionsCreateSessionResponse, error) {
	res, err := c.sendV1CreateSession(ctx, request)
	return res, err
}

func (c *Client) sendV1CreateSession(ctx context.Context, request *V1SessionsCreateSessionRequestBody) (res *V1SessionsCreateSessionResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1CreateSession"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/sessions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1CreateSession",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/sessions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeV1CreateSessionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1CreateSessionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1DeleteSession invokes V1DeleteSession operation.
//
// Удаление сессии пользователя.
//
// DELETE /api/v1/sessions/{id}
func (c *Client) V1DeleteSession(ctx context.Context, params V1DeleteSessionParams) (*SuccessResponse, error) {
	res, err := c.sendV1DeleteSession(ctx, params)
	return res, err
}

func (c *Client) sendV1DeleteSession(ctx context.Context, params V1DeleteSessionParams) (res *SuccessResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1DeleteSession"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/api/v1/sessions/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1DeleteSession",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/api/v1/sessions/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1DeleteSessionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// V1GetUsers invokes V1GetUsers operation.
//
// Получение информации обо всех пользователях.
//
// GET /api/v1/users
func (c *Client) V1GetUsers(ctx context.Context, params V1GetUsersParams) (*V1UsersGetUsersResponse, error) {
	res, err := c.sendV1GetUsers(ctx, params)
	return res, err
}

func (c *Client) sendV1GetUsers(ctx context.Context, params V1GetUsersParams) (res *V1UsersGetUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("V1GetUsers"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/users"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "V1GetUsers",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "limit" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Limit.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "offset" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Offset.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Authorization-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AuthorizationToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeV1GetUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
