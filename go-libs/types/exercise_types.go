package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DataTypeLabel is the data type of the exercise type
type DataTypeLabel string

const (
	// SetLabel is for sets
	SetLabel DataTypeLabel = "sets"
	// SingleValueLabel is for just single values
	SingleValueLabel DataTypeLabel = "single-value"
)

// ExerciseType is for custom types of exercises
type ExerciseType struct {
	ID          primitive.ObjectID `json:"exercise_type_id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description"`
	DataType    DataTypeLabel      `json:"data_type" bson:"data_type"`
}

// IsValid validates the struct
func (exerciseType ExerciseType) IsValid() error {
	validate := validator.New()
	return validate.Struct(exerciseType)
}
