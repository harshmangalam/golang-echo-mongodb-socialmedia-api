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

type PostHandler struct {
	postColl *mongo.Collection
	userColl *mongo.Collection
}

func NewPostHandler(db *mongo.Database) *PostHandler {
	postColl := db.Collection("post")
	userColl := db.Collection("user")
	return &PostHandler{postColl, userColl}
}

func (h *PostHandler) GetPosts(c echo.Context) error {
	cursor, err := h.postColl.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	defer cursor.Close(context.TODO())
	var posts []models.Post

	cursor.All(context.TODO(), &posts)

	return c.JSON(http.StatusOK, posts)

}

func (h *PostHandler) GetPost(c echo.Context) error {

	postId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// match := bson.D{{"$match", bson.D{{"_id", postId}}}}
	// lookup := bson.D{{"$lookup", bson.D{{"from", "user"}, {"localField", "userId"}, {"foreignField", "_id"}, {"as", "User"}}}}
	// cursor, err := h.postColl.Aggregate(context.TODO(), mongo.Pipeline{lookup, match})

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// res := bson.M{}
	// _ = cursor.All(context.TODO(), res)

	var post models.Post
	var user models.User

	h.postColl.FindOne(context.TODO(), bson.M{"_id": postId}).Decode(&post)
	h.userColl.FindOne(context.TODO(), bson.M{"_id": post.UserId}).Decode(&user)

	post.User = &user

	post.User.Password = ""
	return c.JSON(http.StatusOK, post)

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
	inseredPost, err := h.postColl.InsertOne(context.TODO(), doc)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	var createdPost models.Post

	if err := h.postColl.FindOne(context.TODO(), bson.M{"_id": inseredPost.InsertedID}).Decode(&createdPost); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, createdPost)

}
