package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/entity"
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Register User
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	err := h.Service.RegisterUser(c, &user)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{Status: http.StatusCreated, Message: "Success Register User"})
}

// Get User by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(c, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, response.Response[entity.GetUserByIdResponse]{Status: http.StatusOK, Message: "Success", Data: *user})
}
