package main

import (
	"context"
	"socialmedia/handlers"

	"github.com/labstack/echo/v4"
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

	authHandler := handlers.NewAuthHandler(mongoClient)

	// auth group routes
	auth := e.Group("/auth")
	auth.POST("/login", authHandler.AuthLogin)
	auth.POST("/signup", authHandler.AuthSignup)
	auth.GET("/me", authHandler.AuthMe)
	auth.GET("/logout", authHandler.AuthLogout)

	e.Logger.Fatal(e.Start(":4000"))
}
