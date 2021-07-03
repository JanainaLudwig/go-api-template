package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var App AppConfig

type AppConfig struct {
	AppEnv string
	ApiPort string
}

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env")
	}

	App.AppEnv = loadString("APP_ENV")
	App.ApiPort = loadString("API_PORT")
}

func RootPath() string {
	wd, _ := os.Getwd()

	return wd
}

func loadString(code string) string {
	return os.Getenv(code)
}

func loadInt(code string) int {
	str := os.Getenv(code)

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Unable to convert", code, "to int")
	}

	return val
}
