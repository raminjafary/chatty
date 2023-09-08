package handler

import (
	"io"
	"net/http"
	"server/core/user"
	"server/logger"
	jsonS "server/serializer/json"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	GetUserById(*gin.Context)
}

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) serializer(contentType string) user.UserSerializer {
	return &jsonS.User{}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	requestBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := h.serializer(contentType).Decode(requestBody)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = h.userService.CreateUser(user)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(user)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	contextLogger, _ := logger.LogWithFields(user)

	contextLogger.Info("creating user ------> ")

	setupResponse(c, contentType, responseBody, http.StatusCreated)

}

func (h *userHandler) GetUserById(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	id := c.Param("id")

	user, err := h.userService.GetUserById(id)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusBadRequest)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(user)

	setupResponse(c, contentType, responseBody, http.StatusCreated)
}
