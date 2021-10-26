package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client contains the mongo connection and custom functions.
type Client struct {
	Conn *mongo.Client
}

type DbConfigOpts struct {
	Url  string
	User string
	Pass string
	Port string
	Name string
}

// Connect for connecting to the mongo serer.
func Connect(creds DbConfigOpts) (*Client, error) {
	mongoURI := "mongodb://" + creds.User + ":" + creds.Pass + "@" + creds.Url + ":" + creds.Port + "/" + creds.Name + "?authSource=admin&retryWrites=true&w=majority"
	c, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	log.Infof("Connecting to %s", mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect(ctx)

	client := &Client{
		Conn: c,
	}

	log.Info("Database connected!")
	return client, nil
}
