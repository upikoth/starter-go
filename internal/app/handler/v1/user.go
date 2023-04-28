package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/app/constants"
	"github.com/upikoth/starter-go/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

type getUsersResponseDataV1 struct {
	Users []model.User `json:"users"`
}

func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.store.GetUsers()

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := getUsersResponseDataV1{users}
	c.Set("responseData", responseData)
}

type getUserResponseDataV1 struct {
	User model.User `json:"user"`
}

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

	responseData := getUserResponseDataV1{user}
	c.Set("responseData", responseData)
}

type createUserRequestBodyV1 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponseDataV1 struct {
	User model.User `json:"user"`
}

func (h *HandlerV1) CreateUser(c *gin.Context) {
	requestBody := createUserRequestBodyV1{}
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
		PasswordHash: string(hashedBytes[:]),
	}

	createdUser, err := h.store.CreateUser(user)

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := createUserResponseDataV1{User: createdUser}
	c.Set("responseData", responseData)
}

type patchUserRequestBodyV1 struct {
	Email string `json:"email"`
}

type patchUserResponseDataV1 struct {
	User model.User `json:"user"`
}

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

	requestBody := patchUserRequestBodyV1{}
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

	responseData := patchUserResponseDataV1{User: updatedUser}
	c.Set("responseData", responseData)
}

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
