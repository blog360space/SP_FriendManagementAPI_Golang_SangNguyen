package controllers

import (
	"github.com/gin-gonic/gin"

	"spapp/src/commands/friend"
)

func FriendController(router *gin.RouterGroup)  {
	apis := router.Group("/friend")
	{
		apis.POST("/make-friend", friend.MakeFriendCommand)
		apis.POST("/get-friends", friend.GetFriendsCommand)
		apis.POST("/get-common-friends", friend.GetCommonFriendsCommand)
	}
}