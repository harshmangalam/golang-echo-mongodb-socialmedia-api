package main

import (
	"context"
	"log"
	"os"
	"socialmedia/handlers"
	"socialmedia/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	// mongodb connection

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		panic(err)
	}
	// disconnect mongodb

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := mongoClient.Database(os.Getenv("MONGO_DB_NAME"))

	// auth group routes
	authHandler := handlers.NewAuthHandler(db)
	auth := e.Group("/api/auth")
	auth.POST("/login", authHandler.AuthLogin)
	auth.POST("/signup", authHandler.AuthSignup)

	// users group
	userHandler := handlers.NewUserHandler(db)
	users := e.Group("/api/users")
	users.Use(middlewares.CheckAuth())
	users.GET("/", userHandler.GetUsers)
	users.GET("/me", userHandler.GetCurrentUser)
	users.GET("/logout", userHandler.LogoutUser)
	users.GET("/:id", userHandler.GetUser)

	// posts group
	postHandler := handlers.NewPostHandler(db)
	post := e.Group("/api/posts")
	post.Use(middlewares.CheckAuth())
	post.GET("/", postHandler.GetPosts)
	post.POST("/", postHandler.CreatePost)

	e.Logger.Fatal(e.Start(":4000"))
}
