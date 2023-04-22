package env

import (
	"github.com/joho/godotenv"
	"log"
)

var Env map[string]string

// Load used to load environment variables from a .env file.
func Load() {
	var err error
	Env, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
