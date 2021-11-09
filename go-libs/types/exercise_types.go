package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DataType is the data type of the exercise type
type DataType string

const (
	// Sets is for sets
	Sets DataType = "sets"
	// SingleValue is for just single values
	SingleValue DataType = "single-value"
)

// ExerciseType is for vustom types of exercises
type ExerciseType struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"userId"`
	Name        string             `json:"name" bson:"name" validate:""`
	Description string             `json:"description" bson:"description"`
	DataType    DataType           `json:"data_type" bson:"data_type"`
}

func (exerciseType ExerciseType) IsValid() error {
	validate := validator.New()
	return validate.Struct(exerciseType)
}
