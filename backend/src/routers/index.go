package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"spapp/src/controllers"
)

func ping(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"Time": time.Now(),
		"TimeUTC": time.Now().UTC(),
	})
}

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", ping)

	controllers.Initialize(router)

	return router
}