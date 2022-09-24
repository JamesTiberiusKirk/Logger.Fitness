package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SingleValue for the data in exercises
type SingleValue struct {
	DataType string `json:"dataType" bson:"data_type"`
	Value    int    `json:"value" bson:"value"`
}

// Set for the data in exercises
type Set struct {
	Resps      int  `json:"reps" bson:"reps"`
	Resistance int  `json:"resistance" bson:"resistance"`
	IsDropSet  bool `json:"isDropSet" bson:"is_drop_set"`
}

// Exercise is for one individual exercise in a workout
type Exercise struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	WorkoutID      primitive.ObjectID `json:"workoutId" bson:"workout_id"`
	ExerciseTypeID primitive.ObjectID `json:"exerciseTypeId" bson:"exercise_type_id"`
	UserID         primitive.ObjectID `json:"userId" bson:"user_id"`
	Sets           *[]Set             `json:"sets" bson:"sets"`
	SingleValue    *SingleValue       `json:"singleValue" bson:"single_value"`
	Notes          string             `json:"notes" bson:"notes"`
}

// IsValid checks if instance of User is valid using the validator.
func (e Exercise) IsValid() error {
	validate := validator.New()
	return validate.Struct(e)
}
