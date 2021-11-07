package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataType string

const (
	Sets        DataType = "sets"
	SingleValue DataType = "single-value"
)

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
