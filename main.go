package main

import (
	"example/go-auth/configs"
	"example/go-auth/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	configs.SetupMongo()

	// Routes
	routes.UserRoute(r)

	return r
}

func main() {
	router := setupRouter()

	router.Run("localhost:8080")
}