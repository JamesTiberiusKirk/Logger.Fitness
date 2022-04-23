package main

import (
	"Logger.Fitness/backend/server"
	"Logger.Fitness/go-libs/auth/providers"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := buildConfig()
	log.Printf("%+v", conf)

	dbClient := initMongoDBClinet(conf)
	googleOauthConfig := providers.SetupGogoleConfig(conf.Google.ClientID, conf.Google.ClientSecret)

	server.Run(dbClient, conf.HTTPPort, googleOauthConfig)
}
