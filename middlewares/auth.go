package middlewares

import (
	"os"
	"socialmedia/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CheckAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
		ParseTokenFunc: utils.ParseTokenFunc,
	})
}
