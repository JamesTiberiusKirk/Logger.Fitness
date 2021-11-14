package db

import (
	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WorkoutsCollection const for the mongodb collection
const WorkoutsCollection = "workouts"

// InsertNewWorkout mongodb insert operation
func (db *DbClient) InsertNewWorkout(newWorkout types.Workout) error {
	log.Info("Test ")
	return nil
}

// GetUserWorkout mongodb get operation
func (db *DbClient) GetUserWorkout(userID primitive.ObjectID) ([]types.Workout, error) {
	return nil, nil
}

// EditWorkout re-write operation
func (db *DbClient) EditWorkout(types.Workout) error {
	return nil
}

// DeleteWorkout delete operation
func (db *DbClient) DeleteWorkout(workoutID, userID primitive.ObjectID) error {
	return nil
}
