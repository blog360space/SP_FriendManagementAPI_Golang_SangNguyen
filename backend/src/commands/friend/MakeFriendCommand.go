package friend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeFriendCommand (context * gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"sang": "sang",
	})
}