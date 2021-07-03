package main

import (
	"go-api-template/api/router"
	"go-api-template/config"
	"log"
)

func main() {

	config.LoadEnv(config.RootPath() + "/config/.env")

	log.Println("Starting the " + config.App.AppEnv + " API...", "Go to http://localhost:" + config.App.ApiPort)

	router.StartApi()
}
