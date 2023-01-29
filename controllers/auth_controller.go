package controllers

import (
	"context"
	"example/go-auth/configs"
	"example/go-auth/models"
	"example/go-auth/responses"
	"example/go-auth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCollection *mongo.Collection = configs.GetCollection(configs.DB, "auth")

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.LoginAuth
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status: http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{"data": err.Error()},
				},
			)
		}

		var dbUser models.Auth

		err := authCollection.FindOne(ctx, bson.M{"email": user.Email})

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": "The user has not been found"},
				},
			)
		}

		if utils.CheckPassword(user.Password, dbUser.HashPassword) {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status: http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{"data": "The password doesn't match"},
				},
			)
			return
		}

		
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
