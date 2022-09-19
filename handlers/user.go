package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

	return c.JSON(http.StatusOK, "users")
}