package controllers

import (
	"github.com/gin-gonic/gin"

	"spapp/src/handlers/friend"
)

func FriendController(router *gin.RouterGroup)  {
	apis := router.Group("/friend")
	{
		apis.POST("/make-friend", friend.MakeFriendHandle)
		apis.POST("/get-friends", friend.GetFriendsHandle)
		apis.POST("/get-common-friends", friend.GetCommonFriendsHandle)

		apis.POST("/subscribe-user", friend.SubscribeUserHandle)
		apis.POST("/block-user", friend.BlockUserHandle)
		apis.POST("/get-recipients", friend.GetRecipientsHandle)


	}
}