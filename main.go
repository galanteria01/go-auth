package main

import (
	"example/go-auth/configs"
	"example/go-auth/routes"
)

func main() {
	router := configs.SetupRouter()
	configs.SetupMongo()
	routes.UserRoute(router)
	routes.AuthRoute(router)

	router.Run("localhost:8080")
}
