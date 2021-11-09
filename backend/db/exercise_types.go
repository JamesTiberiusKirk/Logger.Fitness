package db

import (
	"context"

	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EXERCISE_TYPE_COLLECTION name of collection
const EXERCISE_TYPE_COLLECTION = "exercise_types"

// InsertNewExerciseType for inserting new type
func (db *DbClient) InsertNewExerciseType(exerciseType types.ExerciseType) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)

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
// TODO: add userID also for sec validation
func (db *DbClient) UpdateExerciseType(userID primitive.ObjectID, exerciseType types.ExerciseType) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": exerciseType.ID, "user_id": userID}, exerciseType)
	if err != nil {
		return err
	}
	log.Info(result)
	return nil
}

// DeleteExerciseTypeById is for deleting an exercise type by ID
// TODO: add userID also for sec validation
func (db *DbClient) DeleteExerciseTypeById(id string) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	log.Info(result)
	return nil
}

// GetExerciseTypesByUserID getting all exercise types belonging to a user
func (db *DbClient) GetExerciseTypesByUserID(userid primitive.ObjectID) ([]types.ExerciseType, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)
	var exerciseTypes []types.ExerciseType

	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return exerciseTypes, nil
	}

	err = curr.All(context.TODO(), &exerciseTypes)
	if err != nil {
		return exerciseTypes, err
	}

	return exerciseTypes, nil
}

// GetExerciseTypesById find a specific exercise type
// TODO: add userID also for sec validation
func (db *DbClient) GetExerciseTypesById(id string) (types.ExerciseType, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)
	var exerciseType types.ExerciseType

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&exerciseType)
	if err != nil {
		return exerciseType, err
	}

	return exerciseType, nil
}

// GetExerciseTypesBySearch search for exercise types
// TODO: add userID also for sec validation
func (db *DbClient) GetExerciseTypesBySearch(query string) ([]types.ExerciseType, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(EXERCISE_TYPE_COLLECTION)
	var exerciseTypes []types.ExerciseType

	filter := bson.M{"name": query, "description": query}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return exerciseTypes, err
	}

	for cur.Next(context.TODO()) {
		var exerciseType types.ExerciseType
		if err := cur.Decode(exerciseType); err != nil {
			return exerciseTypes, err
		}
		exerciseTypes = append(exerciseTypes, exerciseType)
	}

	return exerciseTypes, nil
}
