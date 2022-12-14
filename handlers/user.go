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

	defer cursor.Close(context.TODO())
	var users []models.User

	cursor.All(context.TODO(), &users)
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var user models.User

	if err := h.coll.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(models.User)
	if err := h.coll.FindOne(context.TODO(), bson.M{"_id": user.Id}).Decode(&user); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) LogoutUser(c echo.Context) error {
	user := c.Get("user").(models.User)
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{"isActive": false}}
	h.coll.UpdateOne(context.TODO(), filter, update)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User log out successfully",
	})
}
