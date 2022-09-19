package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	mongoClient *mongo.Client
}

func NewAuthHandler(mongoClient *mongo.Client) *AuthHandler {
	return &AuthHandler{mongoClient}
}

func (h *AuthHandler) AuthLogin(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

func (h *AuthHandler) AuthSignup(c echo.Context) error {
	// parse request body in form of struct
	type SignupData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	signupData := new(SignupData)

	// bind incomming request body to  struct
	if err := c.Bind(signupData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//
	return c.JSON(http.StatusCreated, signupData)
}

func (h *AuthHandler) AuthLogout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

func (h *AuthHandler) AuthMe(c echo.Context) error {
	return c.String(http.StatusOK, "Me")
}
