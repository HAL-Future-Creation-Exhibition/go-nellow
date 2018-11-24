package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBConfig = NewDBConfig()

type db struct {
	Host     string
	User     string
	Pass     string
	Database string
}

func NewDBConfig() db {
	Env_load()
	return db{
		Host:     os.Getenv("DBHOST"),
		User:     os.Getenv("DBUSER"),
		Pass:     os.Getenv("DBPASSWORD"),
		Database: os.Getenv("DATABASE"),
	}
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
