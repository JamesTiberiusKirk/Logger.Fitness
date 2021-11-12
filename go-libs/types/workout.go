package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SingleValue for the data in exercises
type SingleValue struct {
	DataType string `json:"data_type" bson:"data_type"`
	Value    int    `json:"value" bson:"value"`
}

// Set for the data in exercises
type Set struct {
	Resps      int  `json:"reps" bson:"reps"`
	Resistance int  `json:"resistance" bson:"resistance"`
	IsDropSet  bool `json:"is_drop_set" bson:"is_drop_set"`
}

// Exercise is for one individual exercise in a workout
type Exercise struct {
	ID             primitive.ObjectID `json:"exercise_id" bson:"_id"`
	WorkoutID      primitive.ObjectID `json:"workout_id" bson:"workout_id"`
	ExerciseTypeID primitive.ObjectID `json:"exercise_type" bson:"exercise_type"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
	Sets           []Set              `json:"sets" bson:"sets"`
	SingleValue    SingleValue        `json:"single_value" bson:"single_value"`
	Notes          string             `json:"notes" bson:"notes"`
}

// Workout is for a collection of exercises
type Workout struct {
	ID        primitive.ObjectID `json:"workout_id" bson:"_id"`
	StartTime time.Time          `json:"start_time" bson:"start_time"`
	EndTime   time.Time          `json:"end_time" bson:"end_time"`
	Notes     string             `json:"notes" bson:"notes"`
}
