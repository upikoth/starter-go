package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/ogen-go/ogen/middleware"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http/handler"
	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type HTTP struct {
	logger    logger.Logger
	appServer *http.Server
}

func New(
	config *config.ControllerHTTP,
	loggerInstance logger.Logger,
	service *service.Service,
) (*HTTP, error) {
	handlerInstance := handler.New(loggerInstance, service)

	srv, err := app.NewServer(
		handlerInstance,
		app.WithErrorHandler(getAppErrorHandler(handlerInstance)),
		app.WithMiddleware(
			httpSentryMiddleware,
			logger.GetHTTPMiddleware(loggerInstance),
		),
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", corsMiddleware(srv))

	appInstance := http.FileServer(http.Dir("docs/app"))
	mux.Handle("/api/docs/app/", http.StripPrefix("/api/docs/app/", appInstance))

	swaggerUI := http.FileServer(http.Dir("docs/swagger-ui"))
	mux.Handle("/api/docs/swagger-ui/", http.StripPrefix("/api/docs/swagger-ui/", swaggerUI))

	appServer := &http.Server{
		Addr:              ":" + config.Port,
		ReadHeaderTimeout: time.Minute,
		Handler:           mux,
	}

	return &HTTP{
		logger:    loggerInstance,
		appServer: appServer,
	}, nil
}

func (h *HTTP) Start() error {
	return errors.WithStack(h.appServer.ListenAndServe())
}

func (h *HTTP) Stop(ctx context.Context) error {
	return errors.WithStack(h.appServer.Shutdown(ctx))
}

func getAppErrorHandler(handler *handler.Handler) app.ErrorHandler {
	return func(context context.Context, w http.ResponseWriter, _ *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")

		res := handler.NewError(context, err)
		w.WriteHeader(res.StatusCode)

		bytes, _ := res.Response.MarshalJSON()
		_, _ = w.Write(bytes)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Baggage, Sentry-Trace, Authorization-Token")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func httpSentryMiddleware(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	transaction := sentry.StartTransaction(
		req.Context,
		fmt.Sprintf("%s: %s", req.Raw.Method, req.Raw.URL.Path),
		sentry.ContinueFromRequest(req.Raw),
	)
	defer transaction.Finish()

	setParamsToTransactionData(transaction, req.Params)
	bytes, _ := json.Marshal(req.Body)
	transaction.SetData("Request.Body", string(bytes))

	ctx := transaction.Context()
	req.SetContext(ctx)

	res, errorResponse := next(req)

	if errorResponse != nil {
		sentry.CaptureEvent(&sentry.Event{
			Level:   sentry.LevelInfo,
			Message: errorResponse.Error(),
		})
	}

	if errorResponse == nil {
		bytes, _ = json.Marshal(res.Type)
		transaction.SetData("Response.Body", string(bytes))
	} else {
		transaction.SetData("Response.Error", errorResponse.Error())
	}

	return res, errorResponse
}

func setParamsToTransactionData(transaction *sentry.Span, params middleware.Parameters) string {
	res := ""

	for k, v := range params {
		var val string

		switch vt := v.(type) {
		case app.OptInt:
			val = strconv.Itoa(vt.Value)
		default:
			val = fmt.Sprintf("%v", vt)
		}

		transaction.SetData(
			fmt.Sprintf("Request.Params.%s", k.Name),
			val,
		)
	}

	return res
}
