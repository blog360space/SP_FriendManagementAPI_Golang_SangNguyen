package src

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	persistence "spapp/src/persistence"
	routers "spapp/src/routers"
	docs "spapp/docs"

	"github.com/joho/godotenv"
	"log"


)

func Bootstrap(){
	// Config
	err := godotenv.Load("env/env.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	persistence.UseMySql()

	// Routes
	server := routers.GetRoutes()

	// Swagger
	docs.SwaggerInfo.Title = "Friend Management APIs"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*.html", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run
	server.Run()
}