package handlers

import (
	"context"
	"net/http"
	"socialmedia/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

func (h *PostHandler) CreatePost(c echo.Context) error {
	userId := c.Get("user").(models.User).Id

	type CreatePostData struct {
		Content string `json:"content"`
		Image   string `json:"image"`
	}

	var newPost CreatePostData

	if err := c.Bind(&newPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	doc := bson.M{"content": newPost.Content, "image": newPost.Image, "userId": userId}
	inseredPost, err := h.coll.InsertOne(context.TODO(), doc)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	var createdPost models.Post

	if err := h.coll.FindOne(context.TODO(), bson.M{"_id": inseredPost.InsertedID}).Decode(&createdPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, createdPost)

}
