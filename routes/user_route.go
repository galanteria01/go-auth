package routes

import (
	"example/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/users", controllers.GetAllUsers())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.GET("/user/:userId", controllers.GetUser())
}
