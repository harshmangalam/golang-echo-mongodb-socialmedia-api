package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Content   string             `json:"content,omitempty" bson:"content"`
	Image     string             `json:"image,omitempty" bson:"image"`
	User      User               `json:"user,omitempty" bson:"user"`
	UserId    primitive.ObjectID `json:"-" bson:"userId"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}
