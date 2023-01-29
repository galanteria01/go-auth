package routes

import (
	"example/go-auth/controllers"
	"example/go-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.Use(middlewares.JWTMiddleware())
	
	router.POST("/user", controllers.CreateUser())
	router.GET("/users", controllers.GetAllUsers())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.GET("/user/:userId", controllers.GetUser())
}
