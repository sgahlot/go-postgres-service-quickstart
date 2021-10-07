package db

import (
	"log"
	"os"
)

const (
	DB_NAME_KEY      = "DB_NAME"
	DB_URL_KEY       = "DB_URL"
	RESPOSNE_SUCCESS = "Success"

	POST = "POST"
	PUT  = "PUT"
	GET  = "GET"
)

func CheckErrorWithPanic(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func GetEnvOrDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
