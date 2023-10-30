package models

import (
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	User_id           any                `json:"User_id" bson:"User_id"`
	ID           any                `json:"id" bson:"_id"`
 	Firstname    string `json:"firstname" bson:"firstname" validate:"required" min=2 max=100`
	Lastname     string             `json:"lastname" bson:"lastname" validate:"required" min=2 max=100`
	Email        string             `json:"email" bson:"email" validate:"required" min=2 max=100`
	Password     string             `json:"password" bson:"password" validate:"required" min=6`
	Phone        string             `json:"phone" bson:"phone" validate:"required"`
	Token        string             `json:"token" bson:"token" validate:"required`
	UserType     string             `json:"userType" bson:"userType" validate:"required"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
