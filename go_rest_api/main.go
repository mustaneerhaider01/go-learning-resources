package main

import (
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db.InitDB()
	server := gin.Default()

	router := server.Group("/api")
	routes.RegisterRoutes(router)

	server.Run(":4000") // localhost:4000
}
