package config

import (
	"fmt"
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
		fmt.Println("Error loading .env file")
	}
}
