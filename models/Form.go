package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}