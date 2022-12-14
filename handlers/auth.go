package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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

func (h *AuthHandler) AuthLogin(c echo.Context) error {

	type LoginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginData := new(LoginData)

	// get request body data
	if err := c.Bind(loginData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User

	// match user email
	if err := h.coll.FindOne(context.TODO(), bson.M{"email": loginData.Email}).Decode(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"message": "Invalid credentials"})
	}

	// match user password

	if !utils.MatchPassword(loginData.Password, user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"message": "Invalid credentials"})

	}

	// update user to active state

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{"isActive": true}}
	h.coll.UpdateOne(context.TODO(), filter, update)

	// fetch user

	if err := h.coll.FindOne(context.TODO(), bson.M{"_id": user.Id}).Decode(&user); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// do not show  password
	user.Password = ""

	// jwt token

	payload, _ := json.Marshal(user)
	token, err := utils.GenerateJWTToken(string(payload))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// set access_token in cookies

	cookie := utils.CreateCookie("access_token", fmt.Sprintf("Bearer %v", token))

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{"user": user, "token": token})
}
