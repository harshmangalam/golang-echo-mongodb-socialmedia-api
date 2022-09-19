package main

import (
	"context"
	"fmt"
	"socialmedia/handlers"
	"socialmedia/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()
	// mongodb connection

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/go-socialmedia"))

	if err != nil {
		panic(err)
	}
	// disconnect mongodb

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := mongoClient.Database("go-socialmedia")

	authHandler := handlers.NewAuthHandler(db)

	// auth group routes
	auth := e.Group("/auth")
	auth.POST("/login", authHandler.AuthLogin)
	auth.POST("/signup", authHandler.AuthSignup)
	auth.GET("/me", authHandler.AuthMe)
	auth.GET("/logout", authHandler.AuthLogout)

	// post group

	postHandler := handlers.NewPostHandler(db)

	post := e.Group("/posts")
	post.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:     []byte("itssecret"),
		ParseTokenFunc: utils.ParseTokenFunc,
		SuccessHandler: func(c echo.Context) {
			fmt.Println(c.Get("userId"))
		},
		ContextKey: "userId",
	}))
	post.GET("/", postHandler.GetPosts)

	e.Logger.Fatal(e.Start(":4000"))
}
