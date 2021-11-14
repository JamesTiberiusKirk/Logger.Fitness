package db

import (
	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExercisesCollection const for mongodb collection
const ExerciseCollection = "workouts"

// InsertNewExercise insert operation
func (db *DbClient) InsertNewExercise(newExercise types.Exercise) error {
	log.Info("Test ")
	return nil
}

// GetUserExercise get operation
func (db *DbClient) GetUserExercise(userID primitive.ObjectID) ([]types.Exercise, error) {
	return nil, nil
}

// EditExercise re-write operation
func (db *DbClient) EditExercise(types.Exercise) error {
	return nil
}

// DeleteExercise delete operation
func (db *DbClient) DeleteExercise(workoutID, userID primitive.ObjectID) error {
	return nil
}
