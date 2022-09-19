package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthLogin(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

func AuthSignup(c echo.Context) error {
	return c.String(http.StatusCreated, "Signup")
}

func AuthLogout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

func AuthMe(c echo.Context) error {
	return c.String(http.StatusOK, "Me")
}
