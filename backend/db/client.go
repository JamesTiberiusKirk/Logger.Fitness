package db

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client contains the mongo connection and custom functions.
type Client struct {
	Conn *mongo.Client
}

type DbConfigOpts struct {
	Url    string
	User   string
	Pass   string
	Port   string
	DbName string
}

// Connect for connecting to the mongo serer.
func Connect(creds DbConfigOpts) (*Client, error) {
	mongoURI := "mongodb://" + creds.Url + ":" + creds.Port + "/"
	log.Infof("Connecting to %s", mongoURI)

	credential := options.Credential{
		AuthSource: "logger_fitness_db",
		Username:   "logger_fitness",
		Password:   "lfPass",
	}
	clientOpts := options.Client().ApplyURI(mongoURI).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		panic(err)
	}

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping, mongoDB:", err)
	}

	log.Printf("Database connected!")
	return &Client{Conn: client}, nil
}
