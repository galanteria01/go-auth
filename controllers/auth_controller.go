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
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				},
			)
		}

		var dbUser models.Auth

		err := authCollection.FindOne(ctx, bson.M{"email": user.Email})

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": "The user has not been found"},
				},
			)
		}

		if utils.CheckPassword(user.Password, dbUser.HashPassword) {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    map[string]interface{}{"data": "The password doesn't match"},
				},
			)
			return
		}

		tokenString, _ := utils.GenerateJWT(dbUser.Email)

		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "success",
				Data:    map[string]interface{}{"token": tokenString},
			},
		)
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.Auth
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				},
			)
			return
		}

		result, err := authCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "success",
				Data:    map[string]interface{}{"data": result},
			},
		)
	}
}
