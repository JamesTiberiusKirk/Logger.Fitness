package db

import (
	"context"

	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const exerciseCollectionName = "exercises"

// InsertNewExercise mongodb insert operation
func (db *DbClient) InsertNewExercise(newExercise types.Exercise) (types.Exercise, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)

	err := newExercise.IsValid()
	if err != nil {
		log.Error(err.Error())
		return types.Exercise{}, err
	}

	newExercise.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(context.TODO(), newExercise)
	if err != nil {
		log.Error(err.Error())
		return types.Exercise{}, err
	}

	return newExercise, nil
}

// GetUserExercises mongodb get operation
func (db *DbClient) GetUserExercises(userID primitive.ObjectID) ([]types.Exercise, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)
	var results []types.Exercise

	filter := bson.M{"user_id": userID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err.Error())
		return results, err
	}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Error(err.Error())
		return results, err
	}

	return results, nil
}

// GetUserExercisesByWorkout mongodb get operation
func (db *DbClient) GetUserExercisesByWorkout(userID, workoutID primitive.ObjectID) ([]types.Exercise, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)
	var results []types.Exercise

	filter := bson.M{"user_id": userID, "workout_id": workoutID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err.Error())
		return results, err
	}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Error(err.Error())
		return results, err
	}

	return results, nil
}

// EditExercise re-write operation
func (db *DbClient) EditExercise(exercise types.Exercise) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)

	filter := bson.M{"_id": exercise.ID, "user_id": exercise.UserID}
	update := bson.M{"$set": exercise}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

// DeleteExerciseByID delete operation
func (db *DbClient) DeleteExerciseByID(userID, exerciseID primitive.ObjectID) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)

	filter := bson.M{"_id": exerciseID, "user_id": userID}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

// DeleteExerciseByWorkoutID delete operation
func (db *DbClient) DeleteExerciseByWorkoutID(userID, workoutID primitive.ObjectID) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(exerciseCollectionName)

	filter := bson.M{"workout_id": workoutID, "user_id": userID}

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
