package main

import (
	"Logger.Fitness/backend/db"
	log "github.com/sirupsen/logrus"
)

func initMongoDBClinet(conf config) *db.DbClient {
	dbClient, err := db.Connect(db.DbConfigOpts{
		Url:    conf.DBUrl,
		User:   conf.DbUser,
		Pass:   conf.DbPass,
		Port:   conf.DbPort,
		DbName: conf.DbName,
	})

	if err != nil {
		log.Fatalf("Error connecting to the db: %s", err)
	}
	return dbClient
}
