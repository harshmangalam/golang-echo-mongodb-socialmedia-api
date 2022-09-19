package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"socialmedia/handlers"
	"socialmedia/models"
	"socialmedia/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
		ParseTokenFunc: utils.ParseTokenFunc,
		SuccessHandler: func(c echo.Context) {
			userColl := db.Collection("user")
			user := new(models.User)
			userId, err := primitive.ObjectIDFromHex(c.Get("userId").(string))
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			if err := userColl.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(user); err != nil {

				c.JSON(http.StatusNotFound, err.Error())
				return
			}

			user.Password = ""
			c.Set("user", user)
		},
		ContextKey: "userId",
	}))
	post.GET("/", postHandler.GetPosts)

	e.Logger.Fatal(e.Start(":4000"))
}
