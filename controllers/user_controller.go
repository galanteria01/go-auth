package controllers

import (
	"context"
	"example/go-auth/configs"
	"example/go-auth/models"
	"example/go-auth/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status:  http.StatusBadRequest,
					Message: "Error during request",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		newUser := models.User{
			Id:      primitive.NewObjectID(),
			Name:    user.Name,
			Email:   user.Email,
			JobRole: user.JobRole,
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			responses.Response{
				Status:  http.StatusCreated,
				Message: "Request Success",
				Data: map[string]interface{}{
					"data": result,
				},
			},
		)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{},
			)
		}

		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "success",
				Data:    map[string]interface{}{"data": user},
			},
		)
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//edit a user code goes here
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, err := primitive.ObjectIDFromHex(userId)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.Response{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
		}

		update := models.User{
			Name:    user.Name,
			Email:   user.Email,
			JobRole: user.JobRole,
		}

		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, update)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					responses.Response{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data: map[string]interface{}{
							"data": err.Error(),
						},
					},
				)
				return
			}
		}

		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{
					"data": updatedUser,
				},
			},
		)

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//delete a user code goes here
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, err := primitive.ObjectIDFromHex(userId)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "Failed to find user",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(
				http.StatusNotFound,
				responses.Response{
					Status:  http.StatusNotFound,
					Message: "error",
					Data: map[string]interface{}{
						"data": "User not found with specified ID",
					},
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "error",
				Data: map[string]interface{}{
					"data": result,
				},
			},
		)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "Error fetching data",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)
			return
		}

		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					responses.Response{
						Status:  http.StatusInternalServerError,
						Message: "Error fetching data",
						Data: map[string]interface{}{
							"data": err.Error(),
						},
					},
				)
			}
			users = append(users, singleUser)
		}
		c.JSON(
			http.StatusOK,
			responses.Response{
				Status:  http.StatusOK,
				Message: "Successfully fetched users",
				Data: map[string]interface{}{
					"data": users,
				},
			},
		)
	}
}
