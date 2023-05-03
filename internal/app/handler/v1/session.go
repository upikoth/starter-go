package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/upikoth/starter-go/internal/app/constants"
	"github.com/upikoth/starter-go/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

type createSessionRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createSessionResponseData struct {
	User model.User `json:"user"`
}

// CreateSession godoc
// @Summary      Создание сессии пользователя
// @Accept       json
// @Produce      json
// @Param        body body  createSessionRequestBody true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=createSessionResponseData}
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1700, 1701, 1702, 1703, 1704]"
// @Router       /api/v1/session [post]
func (h *HandlerV1) CreateSession(c *gin.Context) {
	requestBody := createSessionRequestBody{}
	err := c.BindJSON(&requestBody)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrSessionPostNoValidJson)
	}

	if requestBody.Email == "" {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrSessionPostNoEmail)
		return
	}

	if requestBody.Password == "" {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrSessionPostNoPassword)
		return
	}

	user, err := h.store.GetUserByEmail(requestBody.Email)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password))

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
	})

	jwtToken, err := token.SignedString(h.jwtSecret)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostCreateJwtToken)
		return
	}

	responseData := createSessionResponseData{User: user}
	c.Set("responseData", responseData)
	c.SetCookie("Authorization", jwtToken, int(constants.Month/time.Second), "", "", true, true)
}

// DeleteSession godoc
// @Summary      Удаление сессии
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/session [delete]
func (h *HandlerV1) DeleteSession(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", true, true)
}

// GetSession godoc
// @Summary      Получение сессии
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/session [get]
func (h *HandlerV1) GetSession(c *gin.Context) {}
