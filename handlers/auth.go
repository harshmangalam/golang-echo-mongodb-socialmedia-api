package handlers

import (
	"context"
	"net/http"
	"socialmedia/models"
	"socialmedia/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	coll *mongo.Collection
}

func NewAuthHandler(db *mongo.Database) *AuthHandler {
	coll := db.Collection("user")
	return &AuthHandler{coll}
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

	// check duplicate email address

	// user := new(models.User)

	var user models.User
	filter := bson.D{{"email", signupData.Email}}
	if err := h.coll.FindOne(context.TODO(), filter).Decode(&user); err != nil {

		// create new user if user not exists
		if err == mongo.ErrNoDocuments {
			// hash plain password
			hashedPassword, err := utils.HashPassword(signupData.Password)

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// insert data
			result, err := h.coll.InsertOne(context.TODO(), bson.D{{"name", signupData.Name}, {"email", signupData.Email}, {"password", hashedPassword}})
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			// find new inserted data
			if err := h.coll.FindOne(context.TODO(), bson.D{{"_id", result.InsertedID}}).Decode(&user); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// remove password fiemd in response
			user.Password = ""

			return c.JSON(http.StatusCreated, user)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// user already exists
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "User already exists"})
}

func (h *AuthHandler) AuthLogout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

func (h *AuthHandler) AuthMe(c echo.Context) error {
	return c.String(http.StatusOK, "Me")
}
