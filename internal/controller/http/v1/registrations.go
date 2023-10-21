package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
)

type registrationCreationRequestData struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=5,lte=72"`
}

// CreateRegistration godoc
// @Summary      Создание заявки на регистрацию
// @Tags         registrations
// @Produce      json
// @Param        body body  registrationCreationRequestData true "Параметры запроса"
// @Success      200  {object}  http.ResponseSuccess
// @Failure      400  {object}  http.ResponseError
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/registrations [post].
func (h *HandlerV1) CreateRegistration(c *gin.Context) {
	data := registrationCreationRequestData{}
	err := c.BindJSON(&data)

	if err != nil {
		c.Set("ResponseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrRequestValidation)
		c.Set("responseErrorDetails", err)
		return
	}

	serviceErr := h.service.Registrations.Create(data.Name, data.Email, data.Password)

	if serviceErr != nil {
		c.Set("ResponseCode", serviceErr.ResponseCode)
		c.Set("responseErrorCode", serviceErr.ResponseErrorCode)
		c.Set("responseErrorDetails", serviceErr.ResponseErrorDetails)
		return
	}
}

type registrationConfirmationRequestData struct {
	Token string `json:"token" binding:"required"`
}

// ConfirmRegistration godoc
// @Summary      Подтверждение заявки на регистрацию
// @Tags         registrations
// @Produce      json
// @Param        body body  registrationConfirmationRequestData true "Параметры запроса"
// @Success      200  {object}  http.ResponseSuccess
// @Failure      400  {object}  http.ResponseError
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/registrations [patch].
func (h *HandlerV1) ConfirmRegistration(c *gin.Context) {
	data := registrationConfirmationRequestData{}
	err := c.BindQuery(&data)

	if err != nil {
		c.Set("ResponseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrRequestValidation)
		c.Set("responseErrorDetails", err)
		return
	}

	serviceErr := h.service.Registrations.Confirm(data.Token)

	if serviceErr != nil {
		c.Set("ResponseCode", serviceErr.ResponseCode)
		c.Set("responseErrorCode", serviceErr.ResponseErrorCode)
		c.Set("responseErrorDetails", serviceErr.ResponseErrorDetails)
		return
	}
}
