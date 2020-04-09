package controllers

import (
	"github.com/gin-gonic/gin"
	"spapp/src/handlers/user"
)

func UserController(router *gin.RouterGroup)  {
	apis := router.Group("/user")
	{
		apis.POST("/register-user", user.RegisterUserHandle)

	}
}