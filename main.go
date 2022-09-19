package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	auth := e.Group("/auth")

	auth.POST("/login", func(c echo.Context) error {})
	auth.POST("/signup", func(c echo.Context) error {})
	auth.GET("/logout", func(c echo.Context) error {})
	auth.GET("/me", func(c echo.Context) error {})

	e.Logger.Fatal(e.Start(":4000"))
}
