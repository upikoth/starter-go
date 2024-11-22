package oauth

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-go/internal/pkg/tracing"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/oauth2"
)

func (o *Oauth) HandleVkRedirect(
	inputCtx context.Context,
	code string,
) (string, error) {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
	defer span.End()

	token, err := o.vkConfig.Exchange(ctx, code)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return "", &models.Error{
			Code:        models.ErrorCodeOauthVkTokenCreating,
			Description: err.Error(),
		}
	}

	email, err := getEmail(token)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return "", &models.Error{
			Code:        models.ErrorCodeOauthVkEmailInvalid,
			Description: err.Error(),
		}
	}

	userVkID, err := getUserVkID(token)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return "", &models.Error{
			Code:        models.ErrorCodeOauthVkUserIDInvalid,
			Description: err.Error(),
		}
	}

	span.SetAttributes(attribute.String("email", email))
	span.SetAttributes(attribute.String("userVkID", userVkID))

	_, err = o.repository.YDB.Users.GetByVkID(ctx, userVkID)

	if err != nil && !errors.Is(err, constants.ErrDBEntityNotFound) {
		span.RecordError(err)
		sentry.CaptureException(err)

		return "", &models.Error{
			Code:        models.ErrorCodeOauthVkGetUserByVkID,
			Description: err.Error(),
		}
	}

	if err != nil && errors.Is(err, constants.ErrDBEntityNotFound) {
		_, errByVkEmail := o.repository.YDB.Users.GetByEmail(ctx, email)

		if errByVkEmail != nil && !errors.Is(errByVkEmail, constants.ErrDBEntityNotFound) {
			span.RecordError(errByVkEmail)
			sentry.CaptureException(errByVkEmail)

			return "", &models.Error{
				Code:        models.ErrorCodeOauthVkGetUserByVkEmail,
				Description: errByVkEmail.Error(),
			}
		}

		if errByVkEmail != nil && errors.Is(errByVkEmail, constants.ErrDBEntityNotFound) {
			// create user
			// create session
			// return session
		}

		// update vk_id
		// create session
		// return session
	}

	// update email if needed
	// create session
	// return session

	return "", nil
}

func getEmail(token *oauth2.Token) (string, error) {
	email := token.Extra("email")

	parsedEmail, emailOk := email.(string)

	if !emailOk || parsedEmail == "" {
		return "", errors.New("Invalid email")
	}

	return parsedEmail, nil
}

func getUserVkID(token *oauth2.Token) (string, error) {
	vkUserID := token.Extra("user_id")

	userIDInt, userIDOk := vkUserID.(float64)

	if !userIDOk || userIDInt == 0 {
		return "", errors.New("Invalid user ID")
	}

	parsedUserID := fmt.Sprintf("%.0f", userIDInt)

	return parsedUserID, nil
}
