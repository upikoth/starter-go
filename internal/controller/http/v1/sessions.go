package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
)

type sessionsResponseData struct {
	Sessions []model.Session `json:"sessions"`
}

// GetSessions godoc
// @Summary      Получение списка сессий
// @Tags         sessions
// @Produce      json
// @Success      200  {object}  http.ResponseSuccess{data=sessionsResponseData}
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/sessions [get].
func (h *HandlerV1) GetSessions(c *gin.Context) {
	sessions, err := h.service.Sessions.GetAll()

	if err != nil {
		c.Set("ResponseCode", http.StatusInternalServerError)
		c.Set("responseErrorCode", constants.ErrDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	c.Set("responseData", sessionsResponseData{
		Sessions: sessions,
	})
}

type sessionCreationRequestData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=5,lte=72"`
}

// CreateSession godoc
// @Summary      Создание сессии
// @Tags         sessions
// @Produce      json
// @Param        body body  sessionCreationRequestData true "Параметры запроса"
// @Success      200  {object}  http.ResponseSuccess
// @Failure      400  {object}  http.ResponseError
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/sessions [post].
func (h *HandlerV1) CreateSession(c *gin.Context) {
	data := sessionCreationRequestData{}
	err := c.BindJSON(&data)

	if err != nil {
		c.Set("ResponseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrRequestValidation)
		c.Set("responseErrorDetails", err)
		return
	}

	session, serviceErr := h.service.Sessions.Create(data.Email, data.Password, c.Request.UserAgent())

	if serviceErr != nil {
		c.Set("ResponseCode", serviceErr.ResponseCode)
		c.Set("responseErrorCode", serviceErr.ResponseErrorCode)
		c.Set("responseErrorDetails", serviceErr.ResponseErrorDetails)
		return
	}

	expiredAt, _ := time.Parse(time.RFC3339Nano, session.ExpiredAt)
	maxAge := expiredAt.UnixMilli() - time.Now().UnixMilli()

	c.SetCookie("AuthToken", session.Token, int(maxAge), "/", h.config.SiteURL, true, true)
}
