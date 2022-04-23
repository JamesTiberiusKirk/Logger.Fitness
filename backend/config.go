package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type googleOAuth struct {
	ClientID     string `envconfig:"GOOGLE_OAUTH_CLIENT_ID"`
	ClientSecret string `envconfig:"GOOGLE_OAUTH_CLIENT_SECRET"`
}

type config struct {
	HTTPPort string `envconfig:"HTTP_PORT" default:":5000"`
	DBUrl    string `envconfig:"MONGO_URL" default:"localhost"`
	DbUser   string `envconfig:"MONGO_USER"`
	DbPass   string `envconfig:"MONGO_USER_PASS"`
	DbPort   string `envconfig:"MONGO_PORT" default:"27017"`
	DbName   string `envconfig:"MONGO_NAME"`
	Google   googleOAuth
}

func buildConfig() config {
	godotenv.Load()
	var conf config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatalf("Error while building config: %v", err)
	}

	return conf
}
