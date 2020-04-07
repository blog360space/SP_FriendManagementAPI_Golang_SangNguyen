package controllers

import (
	"github.com/gin-gonic/gin"
)


func Initialize(router *gin.Engine){
	api := router.Group("/api")
	{
		UserController(api)
		FriendController(api)
	}
}