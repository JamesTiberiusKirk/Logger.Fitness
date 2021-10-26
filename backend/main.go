package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	log "github.com/sirupsen/logrus"

	"Logger.Fitness/backend/db"
	"Logger.Fitness/backend/server"
)

type Config struct {
	HttpPort string `envconfig:"HTTP_PORT" default:":5000"`
	DbUrl    string `envconfig:"MONGO_URL" default:"localhost"`
	DbUser   string `envconfig:"MONGO_USER"`
	DbPass   string `envconfig:"MONGO_USER_PASS"`
	DbPort   string `envconfig:"MONGO_PORT" default:"27017"`
	DbName   string `envconfig:"MONGO_NAME"`
}

func main() {
	godotenv.Load()

	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatalf("Failed to proccess evn: %s", err)
	}
	log.Info(conf)

	dbClient, err := db.Connect(db.DbConfigOpts{
		Url:  conf.DbUrl,
		User: conf.DbUser,
		Pass: conf.DbPass,
		Port: conf.DbPort,
		Name: conf.DbName,
	})
	if err != nil {
		log.Fatalf("Error connecting to the db: %s", err)
	}

	srvParams := &server.ContextParams{
		DbClient: *dbClient,
		Port:     conf.HttpPort,
	}

	server.Run(srvParams)
}
