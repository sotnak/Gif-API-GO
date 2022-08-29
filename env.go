package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Environment struct {
	GinMode       string
	MongoUrl      string
	Secret        string
	TagsCacheTime time.Duration
	GifsCacheTime time.Duration
}

var Env Environment = Environment{}

func initEnv() {
	if Env.GinMode = os.Getenv("GIN_MODE"); Env.GinMode != "release" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	Env.MongoUrl = os.Getenv("MONGO_URL")
	Env.Secret = os.Getenv("SECRET")

	log.Println("using mongo on: " + Env.MongoUrl)
	log.Println("secret set to: " + Env.Secret)

	tagsCacheTimeInt, err := strconv.Atoi(os.Getenv(""))
	if err != nil {
		tagsCacheTimeInt = 3
	}

	Env.TagsCacheTime = time.Duration(int(time.Second) * tagsCacheTimeInt)

	gifsCacheTimeInt, err := strconv.Atoi(os.Getenv(""))
	if err != nil {
		gifsCacheTimeInt = 3
	}

	Env.GifsCacheTime = time.Duration(int(time.Second) * gifsCacheTimeInt)

	log.Println("tagsCacheTime: " + Env.TagsCacheTime.String())
	log.Println("gifsCacheTime: " + Env.GifsCacheTime.String())
}
