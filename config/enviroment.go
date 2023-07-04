package config

import (
	"log"

	"github.com/joho/godotenv"
)

func initEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error ao carregar o arquivo .env")
	}
}
