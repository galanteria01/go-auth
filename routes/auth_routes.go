package routes

import (
	"example/go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signin", controllers.Signin())
	router.POST("/signup", controllers.Signup())
}