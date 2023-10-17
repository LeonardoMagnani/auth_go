package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required, min=2, max=100"`
	Last_name     *string            `json:"last_name" validate:"required, min=2, max=100"`
	Password      *string            `json:"password" validate:"required, min=8"`
	Email         *string            `json:"email" validate:"required, min=2, max=100"`
	Phone         *string            `json:"phone" validate:"required, min=2, max=100"`
	Token         *string            `json:"token" validate:"required, min=2, max=100"`
	User_type     *string            `json:"user_type" validate:"required, min=2, max=100"`
	Refresh_token *string            `json:"refresh_token" validate:"required, min=2, max=100"`
	Created_at    time.Time
	Updated_at    time.Time
	User_id       *string `json:"user_id" validate:"required, min=2, max=100"`
}
