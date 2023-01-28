package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Name	string
	Email string
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Connect with mongo db
	atlas_uri := "mongodb+srv://shoury:shoury@cluster0.aomtf.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(atlas_uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	dbs, _ := client.ListDatabaseNames(ctx, bson.D{{}})

	fmt.Println(dbs)

	coll := client.Database("User").Collection("Users")

	// Routes

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is up and running")
	})

	r.GET("/users", func(c *gin.Context) {
		users, err := coll.Find(ctx, bson.D{{}})
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, users)
	})

	r.GET("/create", func(c *gin.Context) {
		doc := User{Name: "Shoury Sharma", Email: "shanuu12e@gmail.com"}
		result, err := coll.InsertOne(ctx, doc)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, result)
	})

	return r
}

func main() {
	router := setupRouter()

	// Serve on port 8080
	router.Run("localhost:8080")
}