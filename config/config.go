package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"runtime"
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
		log.Fatal("Error loading .env from ", path)
	}

	App.AppEnv = loadString("APP_ENV")
	App.ApiPort = loadString("API_PORT")
}

func RootPath() string {
	_, file, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(file))

	log.Println("root", root)
	return root
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
