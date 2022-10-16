package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Logger.Fitness/go-libs/types"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	HTTPPort string `envconfig:"HTTP_PORT" default:":5000"`
	DBUrl    string `envconfig:"MONGO_URL" default:"localhost"`
	DbUser   string `envconfig:"MONGO_USER"`
	DbPass   string `envconfig:"MONGO_USER_PASS"`
	DbPort   string `envconfig:"MONGO_PORT" default:"27017"`
	DbName   string `envconfig:"MONGO_NAME"`
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

func main() {
	config := buildConfig()
	postData := types.User{
		Username: "test",
		Email:    "test@mail.com",
		Password: "test123",
	}
	postDataRaw, err := json.Marshal(postData)
	if err != nil {
		logrus.Fatalf("Error mashaling: %w", err)
	}

	url := fmt.Sprintf("http://localhost%s/api/v2/auth/register", config.HTTPPort)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postDataRaw))
	if err != nil {
		logrus.Fatalf("Error posting: %d", err.Error())
	}
	if resp.Status != "200 OK" {
		logrus.Fatalf("Non 200 response: %d", resp)
	}
	logrus.Infof("Success: %w", resp)
}
