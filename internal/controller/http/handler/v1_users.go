package handler

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/models"
)

func (h *Handler) V1GetUsers(
	inputCtx context.Context,
	params app.V1GetUsersParams,
) (*app.V1UsersGetUsersResponse, error) {
	span := sentry.StartSpan(inputCtx, "Controller: V1GetUsers")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session, err := h.service.Sessions.CheckToken(ctx, params.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	if !session.UserRole.CheckAccessToAction(models.UserActionGetAnyUserInfo) {
		return nil, &models.Error{
			Code:        models.ErrorCodeUsersGetListForbidden,
			Description: "Insufficient rights",
			StatusCode:  http.StatusForbidden,
		}
	}

	usersGetListParams := models.UsersGetListParams{
		Limit:  params.Limit.Value,
		Offset: params.Offset.Value,
	}

	userList, err := h.service.Users.GetList(ctx, usersGetListParams)

	if err != nil {
		return nil, err
	}

	var usersResult []app.User
	for _, user := range userList.Users {
		usersResult = append(usersResult, app.User{
			ID:    user.ID,
			Email: user.Email,
			Role:  app.UserRole(user.Role),
		})
	}

	return &app.V1UsersGetUsersResponse{
		Success: true,
		Data: app.V1UsersGetUsersResponseData{
			Users:  usersResult,
			Limit:  usersGetListParams.Limit,
			Offset: usersGetListParams.Offset,
			Total:  userList.Total,
		},
	}, nil
}
