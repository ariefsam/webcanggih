package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

}

type Config struct {
	Appname string
	Code    string
}

func LoadConfig() (data Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.Appname = os.Getenv("APPNAME")
	data.Code = os.Getenv("CODE")
	return data
}
