package db

import (
	"context"
	"errors"

	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WorkoutsCollection const for the mongodb collection
const WorkoutsCollection = "workouts"

// InsertNewWorkout insert operation
func (db *DbClient) InsertNewWorkout(newWorkout types.Workout) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)

	err := newWorkout.IsValid()
	if err != nil {
		return err
	}

	newWorkout.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(context.TODO(), newWorkout)
	if err != nil {
		return err
	}

	return nil
}

// EditWorkout re-write operation
// TODO: maybe I need to find something todo with the result?
func (db *DbClient) EditWorkout(workout types.Workout) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)

	filter := bson.M{"_id": workout.ID, "user_id": workout.UserID}
	update := bson.M{"$set": workout}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Info(result)

	return nil
}

// DeleteWorkout delete operation
func (db *DbClient) DeleteWorkout(workoutID, userID primitive.ObjectID) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)

	filter := bson.M{"_id": workoutID, "user_id": userID}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// GetUserWorkout get operation
func (db *DbClient) GetUserWorkout(userID primitive.ObjectID) ([]types.Workout, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)

	if userID.Hex() == "" {
		return nil, errors.New("need to provide a userID")
	}

	var result []types.Workout
	filter := bson.M{"user_id": userID}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// GetUserWorkoutByID get operation
func (db *DbClient) GetUserWorkoutByID(workoutID, userID primitive.ObjectID) (types.Workout, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)
	var result types.Workout

	if userID.Hex() == "" {
		return result, errors.New("need to provide a userID")
	}

	if workoutID.Hex() == "" {
		return result, errors.New("need to provide a workoutID")
	}

	filter := bson.M{"_id": workoutID, "user_id": userID}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetUserWorkoutByQuery get operation
// TODO: Need to implement
// func (db *DbClient) GetUserWorkoutByDateRange(begin, end string, userID primitive.ObjectID) (types.Workout, error) {
// dbc := db.Conn
// collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)
// var result types.Workout

// if userID.Hex() == "" {
// return result, fmt.Errorf("need to provide a userID")
// }

// filter := bson.M{"user_id": userID}
// err := collection.FindOne(context.TODO(), filter).Decode(&result)
// if err != nil {
// return result, err
// }
// return result, nil
// }

// GetUserWorkoutByQuery get operation
// TODO: Need to implement
// func (db *DbClient) GetUserWorkoutByQuery(qery string, userID primitive.ObjectID) (types.Workout, error) {
// dbc := db.Conn
// collection := dbc.Database(DB_NAME).Collection(WorkoutsCollection)
// var result types.Workout

// if userID.Hex() == "" {
// return result, fmt.Errorf("need to provide a userID")
// }

// filter := bson.M{"user_id": userID}
// err := collection.FindOne(context.TODO(), filter).Decode(&result)
// if err != nil {
// return result, err
// }
// return result, nil
// }
