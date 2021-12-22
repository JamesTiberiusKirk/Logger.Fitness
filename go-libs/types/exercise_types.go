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
	ID              primitive.ObjectID `json:"exercise_type_id,omitempty" bson:"_id"`
	UserID          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`
	Name            string             `json:"name,omitempty" bson:"name" validate:"required"`
	Description     string             `json:"description,omitempty" bson:"description" validate:"required"`
	DataType        DataTypeLabel      `json:"data_type,omitempty" bson:"data_type" validate:"required"`
	MeasurementType string             `json:"measurement_type,omitempty" bson:"measurement_type" validate:"required"`
}

// IsValid validates the struct
func (exerciseType ExerciseType) IsValid() error {
	validate := validator.New()
	return validate.Struct(exerciseType)
}
