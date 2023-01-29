package routes

import (
	"example/go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/signin", controllers.Signin())
	router.POST("/signup", controllers.Signup())
}