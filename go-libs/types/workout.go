package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Workout is for a collection of exercises
type Workout struct {
	ID        primitive.ObjectID `json:"workout_id" bson:"_id"`
	StartTime time.Time          `json:"start_time" bson:"start_time"`
	EndTime   time.Time          `json:"end_time" bson:"end_time"`
	Notes     string             `json:"notes" bson:"notes"`
}
