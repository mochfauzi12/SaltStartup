package handler

import (
	"SaltStartup/helper"
	"SaltStartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// menangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"salah": errors}

		response := helper.APIResponse("Register Account Fail", http.StatusUnprocessableEntity, "Error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Fail", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token
	formatter := user.FormatUser(newUser, "tokenTokenToket")

	response := helper.APIResponse("Account has been Registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)

}
