package src

import (
	persistence "spapp/src/persistence"
	routers "spapp/src/routers"

	"github.com/joho/godotenv"
	"log"
)

func Bootstrap(){
	// Config
	err := godotenv.Load("env/env.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	persistence.DbContext = persistence.UseMySql()

	// Routes
	server := routers.GetRoutes()

	// Run
	server.Run()
}