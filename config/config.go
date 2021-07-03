package config

import (
	"path"
	"runtime"
)

var App AppConfig

type AppConfig struct {
	AppEnv string
	ApiPort string
}

func LoadEnv(path string) {
	load(path)

	App.AppEnv = loadString("APP_ENV", envStr("development"))
	App.ApiPort = loadString("API_PORT", envStr("8080"))
}

func RootPath() string {
	_, file, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(file))

	return root
}