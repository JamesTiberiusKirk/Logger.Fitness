package types

import "gorm.io/gorm"

// TODO: figure out how to store exercise type

type WorkoutTemplate struct {
	gorm.Model
	ID        string           `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name"`
	Exercises [][]ExerciseType `json:"exercises"`
	Notes     string           `json:"notes"`
	Tags      []string         `json:"tags"`
}
