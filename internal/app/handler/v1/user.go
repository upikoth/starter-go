package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/app/constants"
	"github.com/upikoth/starter-go/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

type getUsersResponseData struct {
	Users []model.User `json:"users"`
}

// GetUsers godoc
// @Summary      Возвращает список пользователей
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=getUsersResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1200]"
// @Router       /api/v1/users [get].
func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.store.GetUsers()

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := getUsersResponseData{users}
	c.Set("responseData", responseData)
}

type getUserResponseData struct {
	User model.User `json:"user"`
}

// GetUser godoc
// @Summary      Возвращает информацию о пользователе
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Success      200  {object}  model.ResponseSuccess{data=getUserResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1300, 1301, 1302, 1303, 1304, 1305]"
// @Router       /api/v1/user/:id [get].
func (h *HandlerV1) GetUser(c *gin.Context) {
	userId, isIdExist := c.Params.Get("id")

	if !isIdExist {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserGetIdNotExistInRequest)
		return
	}

	userIdParsed, err := strconv.Atoi(userId)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserGetIdNotNumber)
		return
	}

	user, err := h.store.GetUserById(userIdParsed)

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := getUserResponseData{user}
	c.Set("responseData", responseData)
}

type createUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponseData struct {
	User model.User `json:"user"`
}

// CreateUser godoc
// @Summary      Создание пользователя
// @Accept       json
// @Produce      json
// @Param        body body  createUserRequestBody true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=createUserResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1400, 1401, 1402, 1403, 1404, 1405]"
// @Router       /api/v1/user [post].
func (h *HandlerV1) CreateUser(c *gin.Context) {
	requestBody := createUserRequestBody{}
	err := c.BindJSON(&requestBody)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPostNoValidJson)
	}

	if requestBody.Email == "" {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPostNoEmail)
		return
	}

	if requestBody.Password == "" {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPostNoPassword)
		return
	}

	saltedBytes := []byte(requestBody.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserPostCreateHash)
		return
	}

	user := model.User{
		Email:        requestBody.Email,
		PasswordHash: string(hashedBytes),
	}

	createdUser, err := h.store.CreateUser(user)

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := createUserResponseData{User: createdUser}
	c.Set("responseData", responseData)
}

type patchUserRequestBody struct {
	Email string `json:"email"`
}

type patchUserResponseData struct {
	User model.User `json:"user"`
}

// PatchUser godoc
// @Summary      Обновление информации о пользователе
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Param        body body  patchUserRequestBody true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=patchUserResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1600, 1601, 1602, 1603, 1604, 1605]"
// @Router       /api/v1/user/:id [patch].
func (h *HandlerV1) PatchUser(c *gin.Context) {
	userId, isIdExist := c.Params.Get("id")

	if !isIdExist {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPatchIdNotExistInRequest)
		return
	}

	userIdParsed, err := strconv.Atoi(userId)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPatchIdNotNumber)
		return
	}

	requestBody := patchUserRequestBody{}
	err = c.BindJSON(&requestBody)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPatchNoValidJson)
	}

	user := model.User{
		Id: userIdParsed,
	}

	if requestBody.Email != "" {
		user.Email = requestBody.Email
	}

	updatedUser, err := h.store.PatchUser(user)

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := patchUserResponseData{User: updatedUser}
	c.Set("responseData", responseData)
}

// DeleteUser godoc
// @Summary      Удаление информации о пользователе
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Failure      2001 {object}  model.ResponseError "Коды ошибок: [1500, 1501, 1502, 1503]"
// @Router       /api/v1/user/:id [delete].
func (h *HandlerV1) DeleteUser(c *gin.Context) {
	userId, isIdExist := c.Params.Get("id")

	if !isIdExist {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserDeleteIdNotExistInRequest)
		return
	}

	userIdParsed, err := strconv.Atoi(userId)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserDeleteIdNotNumber)
		return
	}

	err = h.store.DeleteUser(userIdParsed)

	if err != nil {
		c.Set("responseErrorCode", err)
	}
}
