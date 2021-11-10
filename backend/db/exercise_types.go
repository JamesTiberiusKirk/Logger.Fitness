package db

import (
	"context"

	"Logger.Fitness/go-libs/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
// TODO: add userID also for sec validation
func (db *DbClient) UpdateExerciseType(userID primitive.ObjectID, exerciseType types.ExerciseType) error {
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
// TODO: add userID also for sec validation
func (db *DbClient) DeleteExerciseTypeByID(exerciseTypeID, userID primitive.ObjectID) error {
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
// TODO: Havent finished using userID
func (db *DbClient) GetExerciseTypesByUserID(userID primitive.ObjectID) ([]types.ExerciseType, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)
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
func (db *DbClient) GetExerciseTypesByID(id string) (types.ExerciseType, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)
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
	collection := dbc.Database(DB_NAME).Collection(ExerciseTypeCollection)
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
