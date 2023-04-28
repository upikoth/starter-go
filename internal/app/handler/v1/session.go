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

type createSessionRequestBodyV1 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createSessionResponseDataV1 struct {
	User model.User `json:"user"`
}

func (h *HandlerV1) CreateSession(c *gin.Context) {
	requestBody := createSessionRequestBodyV1{}
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

	responseData := createSessionResponseDataV1{User: user}
	c.Set("responseData", responseData)
	c.SetCookie("Authorization", jwtToken, int(constants.Month/time.Second), "", "", true, true)
}

func (h *HandlerV1) DeleteSession(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", true, true)
}

func (h *HandlerV1) GetSession(c *gin.Context) {}
