package db

import (
	"context"
	"errors"

	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Need to validate all id params in here for empty

// ExerciseTypeCollection name of collection
const ExerciseTypeCollection = "exercise_types"

// InsertNewExerciseType for inserting new type
func (db *DbClient) InsertNewExerciseType(exerciseType types.ExerciseType) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)

	err := exerciseType.IsValid()
	if err != nil {
		return err
	}

	exerciseType.ID = primitive.NewObjectID()

	_, err = collection.InsertOne(context.TODO(), exerciseType)
	if err != nil {
		return err
	}

	return nil
}

// UpdateExerciseType is for updating an existing exercise type
func (db *DbClient) UpdateExerciseType(userID primitive.ObjectID, exerciseType types.ExerciseType) error {

	if userID.Hex() == "" {
		return errors.New("need to provide a userID")
	}

	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)

	filter := bson.M{"_id": exerciseType.ID, "user_id": userID}
	update := bson.M{"$set": exerciseType}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	log.Info(result)
	return nil
}

// DeleteExerciseTypeByID is for deleting an exercise type by ID
func (db *DbClient) DeleteExerciseTypeByID(exerciseTypeID, userID primitive.ObjectID) error {
	if exerciseTypeID.Hex() == "" {
		return errors.New("need to provide a exerciseTypeID")
	}
	if userID.Hex() == "" {
		return errors.New("need to provide a userID")
	}
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)

	filter := bson.M{
		"_id":     exerciseTypeID,
		"user_id": userID,
	}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	log.Info(result)
	return nil
}

// GetExerciseTypesByUserID getting all exercise types belonging to a user
func (db *DbClient) GetExerciseTypesByUserID(userID primitive.ObjectID) ([]types.ExerciseType, error) {
	if userID.Hex() == "" {
		return nil, errors.New("need to provide a userID")
	}
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)

	var exerciseTypes []types.ExerciseType
	filter := bson.M{"user_id": userID}

	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return exerciseTypes, nil
	}

	err = curr.All(context.TODO(), &exerciseTypes)
	if err != nil {
		return exerciseTypes, err
	}

	return exerciseTypes, nil
}

// GetExerciseTypesByID find a specific exercise type
func (db *DbClient) GetExerciseTypesByID(exerciseTypeID, userID primitive.ObjectID) (types.ExerciseType, error) {
	var exerciseType types.ExerciseType
	if userID.Hex() == "" {
		return exerciseType, errors.New("need to provide a userID")
	}
	if exerciseTypeID.Hex() == "" {
		return exerciseType, errors.New("need to provide a exerciseTypeID")
	}

	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)

	filter := bson.M{"_id": exerciseTypeID, "user_id": userID}

	err := collection.FindOne(context.TODO(), filter).Decode(&exerciseType)
	if err != nil {
		return exerciseType, err
	}

	return exerciseType, nil
}

// GetExerciseTypesBySearch search for exercise types
// TODO: Need to implement, add userID also for sec validation
// func (db *DbClient) GetExerciseTypesBySearch(query string) ([]types.ExerciseType, error) {
// dbc := db.Conn
// collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)
// var exerciseTypes []types.ExerciseType

// filter := bson.M{"name": query, "description": query}
// cur, err := collection.Find(context.Background(), filter)
// if err != nil {
// return exerciseTypes, err
// }

// for cur.Next(context.TODO()) {
// var exerciseType types.ExerciseType
// if err := cur.Decode(exerciseType); err != nil {
// return exerciseTypes, err
// }
// exerciseTypes = append(exerciseTypes, exerciseType)
// }

// return exerciseTypes, nil
// }
