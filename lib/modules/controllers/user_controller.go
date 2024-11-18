package controllers

import (
	"GoPorject/lib/modules/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (ctrl *UserController) GetWelcomeMessage(ctx *gin.Context) {
	message := ctrl.UserService.GetWelcomeMessage()
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
