package controllers

import (
	"github.com/gin-gonic/gin"

	"spapp/src/commands/friend"
)

func FriendController(router *gin.RouterGroup)  {
	apis := router.Group("/friend")
	{
		apis.POST("/make-friend", friend.MakeFriendCommand)

	}
}