package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:""`
	movie string `json:"movie,omitempty"`
	watched bool `json:"watched,omitempty"`
}

