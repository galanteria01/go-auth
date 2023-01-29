package controllers

import (
	"example/go-auth/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCollection *mongo.Collection = configs.GetCollection(configs.DB, "auth")

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
