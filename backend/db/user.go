package db

import (
	"context"

	models "Logger.Fitness/go-libs/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB_NAME         = "logger_fitness_db"
	USER_COLLECTION = "users"
)

// CheckUser function for checking if user exists based on email.
func (db *DbClient) CheckUserBasedOnEmail(lookupEmail string) (bool, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Checking if user exists
	var u models.User
	findErr := collection.FindOne(context.TODO(), bson.M{"email": lookupEmail}).Decode(&u)
	if findErr != nil && findErr != mongo.ErrNoDocuments {
		return false, findErr
	}
	if u.Email == lookupEmail {
		return true, nil
	}

	return false, nil
}

// AddUser function to add user for the DbClient class.
func (db *DbClient) AddUser(user models.User) error {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Input validation
	if inputErr := user.IsValid(); inputErr != nil {
		return inputErr
	}

	// Manually generating an _id
	user.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// GetUsersAll function to get all of the users for the DbClient class.
func (db *DbClient) GetUsersAll() ([]models.User, error) {
	dbc := db.Conn
	var users []models.User

	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)
	curr, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return users, err
	}

	err = curr.All(context.TODO(), &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetUserByEmail function to get a document by the email.
func (db *DbClient) GetUserByEmail(lookupEmail string) (models.User, error) {
	dbc := db.Conn
	collection := dbc.Database(DB_NAME).Collection(USER_COLLECTION)

	// Query db for user
	var u models.User
	filter := bson.M{"email": lookupEmail}
	findErr := collection.FindOne(context.TODO(), filter).Decode(&u)

	if findErr != nil {
		return u, findErr
	}

	return u, nil
}
