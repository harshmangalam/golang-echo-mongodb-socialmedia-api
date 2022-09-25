package models

import "time"

type Post struct {
	Id        string    `json:"_id" bson:"_id"`
	Content   string    `json:"content" bson:"content"`
	Image     string    `json:"image" bson:"image"`
	UserId    string    `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
