package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostHandler struct {
	coll *mongo.Collection
}

func NewPostHandler(db *mongo.Database) *PostHandler {
	coll := db.Collection("user")
	return &PostHandler{coll}
}

func (h *PostHandler) GetPosts(c echo.Context) error {
	user := c.Get("user")

	return c.JSON(http.StatusOK, user)
}
