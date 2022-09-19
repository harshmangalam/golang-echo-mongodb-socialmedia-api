package main

import (
	"socialmedia/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	authHandler := handlers.NewAuthHandler()

	// auth group routes
	auth := e.Group("/auth")
	auth.POST("/login", authHandler.AuthLogin)
	auth.POST("/signup", authHandler.AuthSignup)
	auth.GET("/me", authHandler.AuthMe)
	auth.GET("/logout", authHandler.AuthLogout)

	e.Logger.Fatal(e.Start(":4000"))
}
