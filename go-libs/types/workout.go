package types

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Workout is for a collection of exercises
type Workout struct {
	ID        primitive.ObjectID `json:"workout_id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	StartTime time.Time          `json:"start_time" bson:"start_time"`
	EndTime   *time.Time         `json:"end_time" bson:"end_time"`
	Notes     string             `json:"notes" bson:"notes"`
	Title     string             `json:"title" bson:"title" validation:"required"`
}

// IsValid checks if instance of User is valid using the validator.
func (w Workout) IsValid() error {
	validate := validator.New()
	return validate.Struct(w)
}
