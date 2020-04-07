package controllers

import (
	"github.com/gin-gonic/gin"
	"spapp/src/commands/user"
)

func UserController(router *gin.RouterGroup)  {
	apis := router.Group("/user")
	{
		apis.POST("/register-user", user.RegisterUserCommand)

	}
}