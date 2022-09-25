package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name      string             `json:"name,omitempty" bson:"name"`
	Email     string             `json:"email,omitempty" bson:"email"`
	IsActive  bool               `json:"isActive,omitempty" bson:"isActive"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}
