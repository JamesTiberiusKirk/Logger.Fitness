package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Workout is for a collection of exercises
type Workout struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"user_id"`
	StartTime int64              `json:"startTime" bson:"start_time"`
	EndTime   int64              `json:"endTime" bson:"end_time"`
	Notes     string             `json:"notes" bson:"notes"`
	Title     string             `json:"title" bson:"title" validation:"required"`
}

// IsValid checks if instance of User is valid using the validator.
func (w Workout) IsValid() error {
	validate := validator.New()
	return validate.Struct(w)
}
