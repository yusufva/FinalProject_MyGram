package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("error while getting .env file : %s", err.Error())
	}

	return os.Getenv(key)
}
