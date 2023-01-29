package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Email        string             `json:"email"`
	HashPassword string             `json:"hash_password"`
}
