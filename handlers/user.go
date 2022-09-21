package handlers

import (
	"context"
	"net/http"
	"socialmedia/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	coll *mongo.Collection
}

func NewUserHandler(db *mongo.Database) *UserHandler {
	coll := db.Collection("user")
	return &UserHandler{coll}
}

func (h *UserHandler) GetUsers(c echo.Context) error {

	cursor, err := h.coll.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	var users []models.User

	cursor.All(context.TODO(), &users)
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(models.User)
	userId, _ := primitive.ObjectIDFromHex(user.Id)
	if err := h.coll.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) LogoutUser(c echo.Context) error {
	user := c.Get("user").(models.User)
	userId, _ := primitive.ObjectIDFromHex(user.Id)
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"isActive": false}}
	h.coll.UpdateOne(context.TODO(), filter, update)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User log out successfully",
	})
}
