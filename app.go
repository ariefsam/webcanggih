package main

import (
	"log"
	"os"

	"github.com/ariefsam/webcanggih/webservice"

	"github.com/joho/godotenv"
)

func main() {
	webservice.StartService()
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
